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

func (s *OrderService) CreateQuote(ctx context.Context, order *domain.Order) error {

	customer, err := s.models.Customer.Get(ctx, order.Customer.Id, "MRS")
	if err != nil {
		return err
	}
	address, err := s.models.Address.Get(ctx, order.ShippingAddress.Id)
	if err != nil {
		return err
	}

	if customer.CustomerId != address.CorpAddressId.Float64 {
		return errors.New("customer and address do not match")
	}

	contact, err := s.models.Contacts.Get(ctx, order.Contact.Id)
	if err != nil {
		return err
	}

	if contact.AddressId != address.Id {
		return errors.New("contact and address do not match")

	}

	// create the order
	oeHdrParams := data.CreateOeHdrParams{
		CustomerId:           order.Customer.Id,
		AddressId:            address.Id,
		ContactId:            contact.Id,
		Ship2Name:            address.Name,
		Ship2Add1:            address.MailAddress1.String,
		Ship2Add2:            address.MailAddress2.String,
		Ship2City:            address.MailCity.String,
		Ship2State:           address.MailState.String,
		Ship2Zip:             address.MailPostalCode.String,
		Ship2Country:         address.MailCountry.String,
		ShipToPhone:          contact.DirectPhone.String,
		PoNo:                 order.PurchaseOrder,
		ProjectedOrder:       "Y",
		DeliveryInstructions: address.DeliveryInstructions1.String,
		PackingBasis:         address.ShipToPackingBasis.String,
		PickTicketType:       customer.PickTicketType.String,
		Terms:                customer.TermsId,
	}

	oeHdrParams.CarrierId, err = strconv.ParseFloat(address.CarrierId.String, 64)
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
			SalesrepId:      customer.SalesrepId.String,
			CommissionSplit: 100,
			PrimarySalesrep: "Y",
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

		_, err = s.models.OeLine.Create(
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
