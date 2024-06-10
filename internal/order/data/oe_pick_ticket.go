package data

import (
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/internal/order/domain"
	"github.com/materials-resources/s-prophet/pkg/models"
	"github.com/uptrace/bun"
)

type oePickTicket struct {
	models.OePickTicket `bun:",extend"`
	OeHdr               *oeHdr   `bun:"rel:has-one,join:order_no=order_no"`
	Carrier             *address `bun:"rel:has-one,join:carrier_id=id,join_on:carrier_flag='Y'"`
}

type OePickTicketModel struct {
	bun bun.IDB
}

func NewOePickTicketModel(db bun.IDB) OePickTicketModel {
	return OePickTicketModel{
		bun: db,
	}
}

func (m *OePickTicketModel) Get(ctx context.Context, id string) (*domain.Shipment, error) {
	var pickTicket oePickTicket
	err := m.bun.NewSelect().Model(&pickTicket).
		Relation("Carrier").
		Relation("OeHdr", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Relation("Contacts")
		}).
		Where("pick_ticket_no = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return oePickTicketToShipment(&pickTicket), nil

}

func (m *OePickTicketModel) ListForOrderId(ctx context.Context, orderId string) ([]*domain.Shipment, error) {
	var pickTickets []*oePickTicket
	err := m.bun.NewSelect().Model(&pickTickets).Relation("Carrier").Where("order_no = ?", orderId).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return oePickTicketToShipmentSlice(pickTickets), nil

}

func oePickTicketToShipmentSlice(recs []*oePickTicket) []*domain.Shipment {
	shipments := make([]*domain.Shipment, len(recs))
	for i, rec := range recs {
		shipments[i] = oePickTicketToShipment(rec)
	}
	return shipments
}
func oePickTicketToShipment(rec *oePickTicket) *domain.Shipment {
	shipment := &domain.Shipment{
		Id:                   fmt.Sprintf("%.0f", rec.PickTicketNo),
		CarrierTracking:      rec.TrackingNo.String,
		DeliveryInstructions: rec.Instructions.String,
	}

	if rec.Carrier != nil {
		shipment.CarrierName = rec.Carrier.Name
	}

	if rec.OeHdr != nil {
		shipment.OrderId = rec.OeHdr.OrderNo
		shipment.OrderPurchaseOrder = rec.OeHdr.PoNo.String
		shipment.ShippingAddressId = fmt.Sprintf("%.0f", rec.OeHdr.AddressId.Float64)
		shipment.ShippingAddressName = rec.OeHdr.Ship2Name.String
		shipment.ShippingAddressLineOne = rec.OeHdr.Ship2Add1.String
		shipment.ShippingAddressLineTwo = rec.OeHdr.Ship2Add2.String
		shipment.ShippingAddressCity = rec.OeHdr.Ship2City.String
		shipment.ShippingAddressState = rec.OeHdr.Ship2State.String
		shipment.ShippingAddressPostalCode = rec.OeHdr.Ship2Zip.String
		shipment.ShippingAddressCountry = rec.OeHdr.Ship2Country.String
		shipment.ContactId = rec.OeHdr.ContactId.String

		if rec.OeHdr.Contact != nil {
			fmt.Println("Contact: ", rec.OeHdr.Contact.FirstName, rec.OeHdr.Contact.LastName)
			shipment.ContactName = fmt.Sprintf("%s %s", rec.OeHdr.Contact.FirstName, rec.OeHdr.Contact.LastName)
		}
	}

	if rec.Carrier != nil {
		shipment.CarrierName = rec.Carrier.Name
	}
	return shipment
}
