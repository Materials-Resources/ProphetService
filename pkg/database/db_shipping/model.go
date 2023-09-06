package db_shipping

import "database/sql"

type pickTicket struct {
	PickTicketId         int32          `gorm:"column:pick_ticket_no"`
	OrderId              string         `gorm:"column:order_no"`
	ShipName             sql.NullString `gorm:"column:ship2_name"`
	ShipLineOne          sql.NullString `gorm:"column:ship2_add1"`
	ShipLineTwo          sql.NullString `gorm:"column:ship2_add2"`
	ShipCity             sql.NullString `gorm:"column:ship2_city"`
	ShipState            sql.NullString `gorm:"column:ship2_state"`
	ShipPostalCode       sql.NullString `gorm:"column:ship2_zip"`
	ShipCountry          sql.NullString `gorm:"column:ship2_country"`
	Po                   sql.NullString `gorm:"column:po_no"`
	DeliveryInstructions sql.NullString `gorm:"column:delivery_instructions"`
	ContactFirstName     string         `gorm:"column:first_name"`
	ContactLastName      string         `gorm:"column:last_name"`
}
