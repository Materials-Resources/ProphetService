package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/materials-resources/s_prophet/infrastructure/data"
	"github.com/materials-resources/s_prophet/internal/order/domain"
	"strconv"
)

func NewOrderService(models data.Models) *OrderService {
	return &OrderService{models: models}
}

type OrderService struct {
	models data.Models
}

func (s *OrderService) ListOrdersByCustomer(
	ctx context.Context, customerId float64,
	filters domain.Filters) (
	[]*domain.Order, domain.Metadata,
	error) {
	oeHdrs, metadata, err := s.models.OeHdr.SelectByCustomerId(
		ctx, customerId, data.Filters{
			Limit:        100,
			Cursor:       filters.Cursor,
			Direction:    data.Direction(filters.Direction),
			Sort:         "",
			SortSafeList: nil,
		})
	if err != nil {
		return nil, domain.Metadata{}, err
	}
	orderList := make([]*domain.Order, 0, len(oeHdrs))
	for _, oeHdr := range oeHdrs {
		orderList = append(
			orderList, &domain.Order{
				Id: oeHdr.OrderNo,
				ShippingAddress: domain.Address{
					Id:         oeHdr.AddressId.Float64,
					Name:       oeHdr.Ship2Name.String,
					LineOne:    oeHdr.Ship2Add1.String,
					LineTwo:    oeHdr.Ship2Add2.String,
					City:       oeHdr.Ship2City.String,
					State:      oeHdr.Ship2State.String,
					PostalCode: oeHdr.Ship2Zip.String,
				},
				Customer: domain.Customer{
					Id: oeHdr.CustomerId,
				},
				Taker:                oeHdr.Taker.String,
				DeliveryInstructions: oeHdr.DeliveryInstructions.String,
				PurchaseOrder:        oeHdr.PoNo.String,
				Status: domain.OrderStatus{
					Approved:  oeHdr.Approved.String == "Y",
					Cancelled: oeHdr.CancelFlag.String == "Y",
				},
			})

	}

	return orderList, domain.Metadata{
		NextCursor:     metadata.NextCursor,
		PreviousCursor: metadata.PreviousCursor,
		TotalRecords:   0,
	}, nil
}

func (s *OrderService) CreateQuote(ctx context.Context, order *domain.Order) error {
	contactRec, err := s.models.Contacts.Get(ctx, order.Contact.Id)
	if err != nil {
		return err
	}

	contactsXShipToRecs, err := s.models.ContactsXShipTo.GetByContactId(ctx, "MRS", order.Contact.Id)
	if err != nil {
		return err
	}

	contactExistsOnShipTo := false
	// check to see if order.ShippingAddress.Id is in the slice of contactsXShipToRecs
	for _, rec := range contactsXShipToRecs {
		if rec.ShipToId == order.ShippingAddress.Id {
			contactExistsOnShipTo = true
			break
		}
	}

	if !contactExistsOnShipTo {
		return errors.New("provided shipping address is not associated with contact")

	}

	shipToRec, err := s.models.ShipTo.Get(ctx, "MRS", order.ShippingAddress.Id)
	if err != nil {
		return err
	}

	customerRec, err := s.models.Customer.Get(ctx, shipToRec.CustomerId, "MRS")
	if err != nil {
		return err
	}

	addressRec, err := s.models.Address.Get(ctx, shipToRec.ShipToId)
	if err != nil {
		return err
	}

	shipToSalesrepRecs, err := s.models.ShipToSalesrep.GetByShipToId(ctx, "MRS", shipToRec.ShipToId, true)
	if err != nil {
		return err
	}

	// create the order
	oeHdrParams := data.CreateOeHdrParams{
		CustomerId:           customerRec.CustomerId,
		AddressId:            addressRec.Id,
		ContactId:            contactRec.Id,
		Ship2Name:            addressRec.Name,
		Ship2Add1:            addressRec.MailAddress1.String,
		Ship2Add2:            addressRec.MailAddress2.String,
		Ship2City:            addressRec.MailCity.String,
		Ship2State:           addressRec.MailState.String,
		Ship2Zip:             addressRec.MailPostalCode.String,
		Ship2Country:         addressRec.MailCountry.String,
		ShipToPhone:          contactRec.DirectPhone.String,
		PoNo:                 order.PurchaseOrder,
		ProjectedOrder:       "Y",
		DeliveryInstructions: addressRec.DeliveryInstructions1.String,
		PackingBasis:         addressRec.ShipToPackingBasis.String,
		PickTicketType:       customerRec.PickTicketType.String,
		Terms:                customerRec.TermsId,
		RequestedDate:        order.RequestedDate,
	}

	oeHdrParams.CarrierId, err = strconv.ParseFloat(addressRec.CarrierId.String, 64)
	if err != nil {
		return err
	}
	oeHdr, err := s.models.OeHdr.Create(ctx, oeHdrParams)
	if err != nil {
		return err
	}

	// add the salesrep to the order
	_, err = s.models.OeHdrSalesrep.Create(
		ctx, data.OeHdrSalesrepParams{
			OrderNumber:     oeHdr.OrderNo,
			SalesrepId:      shipToSalesrepRecs[0].SalesrepId,
			CommissionSplit: shipToSalesrepRecs[0].CommissionPercentage,
			PrimarySalesrep: shipToSalesrepRecs[0].PrimarySalesrep.String,
		})

	// create the quote record
	_, err = s.models.QuoteHdr.Create(
		ctx, data.CreateQuoteHdrParams{
			OeHdrUid: oeHdr.OeHdrUid,
		})

	for _, item := range order.Items {
		invLoc, err := s.models.InvLoc.Get(ctx, 1001, "MRS", item.ProductUid)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				return fmt.Errorf("product with uid %d not found", item.ProductUid)
			default:
				return err
			}

		}
		inventorySupplier, err := s.models.InventorySupplier.GetPrimarySupplierByLocation(
			ctx, invLoc.LocationId,
			invLoc.InvMastUid)
		if err != nil {
			return err
		}

		oeLineRec, err := s.models.OeLine.Create(
			ctx, data.CreateOeLineParams{
				OrderNo:              oeHdr.OrderNo,
				UnitPrice:            0,
				QtyOrdered:           item.QuantityOrdered,
				BaseUtPrice:          0,
				CalcValue:            0,
				SourceLocId:          1001,
				ShipLocId:            1001,
				SupplierId:           inventorySupplier.SupplierId,
				ProductGroupId:       invLoc.ProductGroupId.String,
				ExtendedDesc:         invLoc.InvMast.ExtendedDesc.String,
				CustomerPartNumber:   invLoc.InvMast.ItemId,
				CommissionCost:       0,
				OeHdrUid:             oeHdr.OeHdrUid,
				InvMastUid:           invLoc.InvMastUid,
				SalesDiscountGroupId: invLoc.SalesDiscountGroup.String,
				UnitQuantity:         item.QuantityOrdered,
				UnitSize:             1,
			})
		if err != nil {
			return err
		}
		_, err = s.models.QuoteLine.Create(
			ctx, data.CreateQuoteLineParams{
				OeLineUid:    oeLineRec.OeLineUid,
				QtyConverted: 0,
				UomConverted: oeLineRec.UnitOfMeasure.String,
			})
		if err != nil {
			return err
		}

	}
	order.Id = oeHdr.OrderNo
	return err
}

func (s *OrderService) CreateOrder(ctx context.Context, order *domain.Order) error {
	// TODO implement me
	panic("implement me")
}

func (s *OrderService) UpdateOrder(ctx context.Context, order *domain.Order) error {
	// TODO implement me
	panic("implement me")
}

func (s *OrderService) DeleteOrder(ctx context.Context, s2 string) error {
	// TODO implement me
	panic("implement me")
}

func (s *OrderService) GetPickTicketById(ctx context.Context, id float64) (domain.PickTicket, error) {

	oePickTicket, err := s.models.OePickTicket.Get(ctx, id)
	if err != nil {
		return domain.PickTicket{}, err

	}

	pickTicket := domain.PickTicket{
		Id:                   oePickTicket.PickTicketNo,
		OrderId:              oePickTicket.OrderNo,
		DeliveryInstructions: oePickTicket.OeHdr.DeliveryInstructions.String,
		ShippingAddress: domain.Address{
			Id:         oePickTicket.OeHdr.AddressId.Float64,
			Name:       oePickTicket.OeHdr.Ship2Name.String,
			LineOne:    oePickTicket.OeHdr.Ship2Add1.String,
			LineTwo:    oePickTicket.OeHdr.Ship2Add2.String,
			City:       oePickTicket.OeHdr.Ship2City.String,
			State:      oePickTicket.OeHdr.Ship2State.String,
			PostalCode: oePickTicket.OeHdr.Ship2Zip.String,
		},
	}

	if oePickTicket.OeHdr.ContactId.Valid {
		contact, err := s.models.Contacts.Get(ctx, oePickTicket.OeHdr.ContactId.String)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):

			default:
				return domain.PickTicket{}, err
			}

		}
		pickTicket.OrderContact = domain.Contact{
			Id:    oePickTicket.OeHdr.ContactId.String,
			Name:  fmt.Sprintf("%s %s", contact.FirstName, contact.LastName),
			Title: contact.Title.String,
		}

	}

	return pickTicket, nil
}

func (s *OrderService) GetOrderById(ctx context.Context, id string) (domain.Order, error) {
	oeHdr, err := s.models.OeHdr.Get(ctx, id)
	if err != nil {
		return domain.Order{}, err
	}

	var order domain.Order

	order = domain.Order{
		Id: oeHdr.OrderNo,
		ShippingAddress: domain.Address{
			Id:         oeHdr.AddressId.Float64,
			Name:       oeHdr.Ship2Name.String,
			LineOne:    oeHdr.Ship2Add1.String,
			LineTwo:    oeHdr.Ship2Add2.String,
			City:       oeHdr.Ship2City.String,
			State:      oeHdr.Ship2State.String,
			PostalCode: oeHdr.Ship2Zip.String,
		},

		Customer: domain.Customer{
			Id: oeHdr.CustomerId,
		},
		DeliveryInstructions: oeHdr.DeliveryInstructions.String,
		Items:                nil,
		PurchaseOrder:        oeHdr.PoNo.String,
	}

	if oeHdr.ContactId.Valid {
		contactData, err := s.models.Contacts.Get(ctx, oeHdr.ContactId.String)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):

			default:
				return domain.Order{}, err
			}

		}
		order.Contact = domain.Contact{
			Id:    oeHdr.ContactId.String,
			Name:  fmt.Sprintf("%s %s", contactData.FirstName, contactData.LastName),
			Title: contactData.Title.String,
		}
	}

	return order, nil
}
