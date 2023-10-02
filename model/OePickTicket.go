package model

import (
	"database/sql"

	"github.com/uptrace/bun"
)

type OePickTicket struct {
	bun.BaseModel `bun:"table:oe_pick_ticket"`
	PickTicketNo  float64        `bun:"pick_ticket_no,pk"`
	OrderNo       string         `bun:"order_no"`
	Instructions  sql.NullString `bun:"instructions,type:varchar(255)"`

	OeHdr OeHdr `bun:"rel:has-one,join:order_no=order_no"`
}
