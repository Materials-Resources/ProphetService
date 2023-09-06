package db_shipping

import (
	"fmt"
	pb "github.com/materials-resources/s_prophet/proto/shipping/v1alpha0"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewShippingHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetPickTicket(request *pb.GetPickTicketRequest) (*pb.GetPickTicketResponse, error) {
	qr_pickTicket := new(pickTicket)

	h.db.Raw(`SELECT oe_pick_ticket.pick_ticket_no,
       oe_hdr.order_no,
       oe_hdr.ship2_name,
       oe_hdr.ship2_add1,
       oe_hdr.ship2_add2,
       oe_hdr.ship2_city,
       oe_hdr.ship2_state,
       oe_hdr.ship2_zip,
       oe_hdr.ship2_country,
       oe_hdr.delivery_instructions,
       oe_hdr.po_no,
       contacts.first_name,
       contacts.last_name
FROM oe_pick_ticket
INNER JOIN oe_hdr on oe_pick_ticket.order_no = oe_hdr.order_no
INNER JOIN contacts on oe_hdr.contact_id = contacts.id
WHERE oe_pick_ticket.pick_ticket_no = ?
`, request.GetId()).Scan(&qr_pickTicket)

	pbResponse := &pb.GetPickTicketResponse{
		PickTicketId:         qr_pickTicket.PickTicketId,
		ShipName:             qr_pickTicket.ShipName.String,
		ShipLineOne:          qr_pickTicket.ShipLineOne.String,
		ShipLineTwo:          qr_pickTicket.ShipLineTwo.String,
		ShipCity:             qr_pickTicket.ShipCity.String,
		ShipState:            qr_pickTicket.ShipState.String,
		ShipPostalCode:       qr_pickTicket.ShipPostalCode.String,
		ShipCountry:          qr_pickTicket.ShipCountry.String,
		OrderId:              qr_pickTicket.OrderId,
		OrderPo:              qr_pickTicket.Po.String,
		DeliveryInstructions: qr_pickTicket.DeliveryInstructions.String,
		ContactName:          fmt.Sprintf("%s %s", qr_pickTicket.ContactFirstName, qr_pickTicket.ContactLastName),
	}

	return pbResponse, nil
}
