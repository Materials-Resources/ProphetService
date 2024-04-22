package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/materials-resources/s_prophet/infrastructure/data"
	"github.com/materials-resources/s_prophet/internal/order/domain"
)

func NewOrderService(models data.Models) *OrderService {
	return &OrderService{models: models}
}

type OrderService struct {
	models data.Models
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
