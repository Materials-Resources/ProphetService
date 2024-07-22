package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DeliveryPickTicketDetail struct {
	bun.BaseModel               `bun:"table:delivery_pick_ticket_detail"`
	DeliveryPickTicketDetailUid int32     `bun:"delivery_pick_ticket_detail_uid,type:int,pk"` // Unique Identifier
	DeliveryPickTicketUid       int32     `bun:"delivery_pick_ticket_uid,type:int,unique"`    // Pointer to the delivery pick ticket record that this line belongs to.
	Disposition                 *string   `bun:"disposition,type:char(1)"`                    // Stores the disposition of the item in the case of an undershipment or cancellation. For reporting
	ReturnToStock               *string   `bun:"return_to_stock,type:char(1)"`                // When items is undershipped or cancelled, flag on what to do with the material. For reporting.
	ReasonCode                  *int32    `bun:"reason_code,type:int"`                        // When not returned to stock, explains why. For reporting.
	LineNo                      int32     `bun:"line_no,type:int,unique"`                     // Line number of the pick ticket.
	ShipQuantity                float64   `bun:"ship_quantity,type:decimal(19,9)"`            // The quantity that was actually shipped, for reporting.
	DateCreated                 time.Time `bun:"date_created,type:datetime"`                  // Indicates the date/time this record was created.
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`            // Indicates the date/time this record was last modified.
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30)"`         // ID of the user who last maintained this record
}
