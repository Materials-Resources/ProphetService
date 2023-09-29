package model

import (
	"database/sql"

	"github.com/uptrace/bun"
)

type OeHdr struct {
	bun.BaseModel        `bun:"table:oe_hdr"`
	OrderNo              string         `bun:"order_no,type:varchar(8),pk"`
	Ship2Name            sql.NullString `bun:"ship2_name,type:varchar(50)"`
	Ship2Add1            sql.NullString `bun:"ship2_add1,type:varchar(50)"`
	Ship2Add2            sql.NullString `bun:"ship2_add2,type:varchar(50)"`
	Ship2City            sql.NullString `bun:"ship2_city,type:varchar(50)"`
	Ship2State           sql.NullString `bun:"ship2_state,type:varchar(50)"`
	Ship2Zip             sql.NullString `bun:"ship2_zip,type:varchar(10)"`
	Ship2Country         sql.NullString `bun:"ship2_country,type:varchar(50)"`
	ContactId            sql.NullString `bun:"contact_id,type:varchar(16)"`
	DeliveryInstructions sql.NullString `bun:"delivery_instructions,type:varchar(255)"`
	PoNo                 sql.NullString `bun:"po_no,type:varchar(50)"`

	OeLineItems []OeLine `bun:"rel:has-many,join:order_no=order_no"`
	Contact     Contact  `bun:"rel:has-one,join:contact_id=id"`
}
