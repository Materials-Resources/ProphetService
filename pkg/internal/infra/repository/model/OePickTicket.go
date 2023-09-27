package model

import (
	"database/sql"
)

func (OePickTicket) TableName() string {
	return "oe_pick_ticket"
}

type OePickTicket struct {
	PickTicketNo float64        `gorm:"column:pick_ticket_no"`
	OrderNo      string         `gorm:"column:order_no"`
	Instructions sql.NullString `gorm:"column:instructions;type:varchar(255)"`

	OeHdr OeHdr `gorm:"foreignKey:OrderNo"`
}
