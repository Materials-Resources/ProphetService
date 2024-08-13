package repository

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/order/domain"
	"github.com/uptrace/bun"
	"strconv"
)

type ShipmentRepository struct {
	db bun.IDB
}

// GetShipment returns a Shipment with its details
func (r *ShipmentRepository) GetShipment(ctx context.Context, id string) (*domain.Shipment, error) {
	var pickTicket oePickTicket
	err := r.db.NewSelect().Model(&pickTicket).
		Relation("CarrierAddress").
		Relation("OeHdr").
		Relation("OeHdr.Customer").
		Relation("OeHdr.Contact").
		Where("pick_ticket_no = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return mapDbShipmentToDomainShipment(&pickTicket), nil
}

func (r *ShipmentRepository) ListShipmentsByOrder(ctx context.Context, orderId string) ([]*domain.Shipment, error) {
	var pickTickets []*oePickTicket
	err := r.db.NewSelect().Model(&pickTickets).
		Relation("CarrierAddress").
		Relation("OeHdr").
		Relation("OeHdr.Customer").
		Relation("OeHdr.Contact").
		Where("oe_hdr.order_no = ?", orderId).Scan(ctx)
	if err != nil {
		return nil, err
	}

	var shipments []*domain.Shipment
	for _, pickTicket := range pickTickets {
		shipments = append(shipments, mapDbShipmentToDomainShipment(pickTicket))
	}

	return shipments, nil

}

func mapDbShipmentToDomainShipment(dbShipment *oePickTicket) *domain.Shipment {

	shipment := &domain.Shipment{
		Id:                   strconv.FormatFloat(dbShipment.PickTicketNo, 'f', 0, 64),
		CarrierName:          dbShipment.CarrierAddress.Name,
		CarrierTracking:      dbShipment.TrackingNo,
		DeliveryInstructions: dbShipment.Instructions,
	}

	if dbShipment.OeHdr != nil {
		shipment.OrderId = dbShipment.OeHdr.OrderNo
		shipment.OrderPurchaseOrder = dbShipment.OeHdr.PoNo

		if dbShipment.OeHdr.Customer != nil {
			shipment.Customer = domain.Customer{
				Id:   strconv.FormatFloat(dbShipment.OeHdr.Customer.CustomerId, 'f', 0, 64),
				Name: dbShipment.OeHdr.Customer.CustomerName,
			}
		}
		shipment.ShippingAddress = domain.Address{
			Id: func(addressId *float64) string {
				if addressId == nil {
					return ""
				}
				return strconv.FormatFloat(*addressId, 'f', 0, 64)
			}(dbShipment.OeHdr.AddressId),
			Name:       dbShipment.OeHdr.Ship2Name,
			LineOne:    dbShipment.OeHdr.Ship2Add1,
			LineTwo:    dbShipment.OeHdr.Ship2Add2,
			City:       dbShipment.OeHdr.Ship2City,
			State:      dbShipment.OeHdr.Ship2State,
			PostalCode: dbShipment.OeHdr.Ship2Zip,
			Country:    dbShipment.OeHdr.Ship2Country,
		}
	}

	return shipment
}

func NewShipmentRepository(db bun.IDB) *ShipmentRepository {

	return &ShipmentRepository{db: db}
}
