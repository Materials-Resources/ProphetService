package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DeliveryPickTicket struct {
	bun.BaseModel         `bun:"table:delivery_pick_ticket"`
	DeliveryPickTicketUid int32     `bun:"delivery_pick_ticket_uid,type:int,pk"`      // Unique Identifier
	StopUid               int32     `bun:"stop_uid,type:int"`                         // Pointer to the stop record that this pick ticket is associated with.
	PickTicketNo          float64   `bun:"pick_ticket_no,type:decimal(19,0),unique"`  // Pick ticket number that will be delivered.
	ReconciliationFlag    string    `bun:"reconciliation_flag,type:char(1),nullzero"` // The driver has entered an edit in the quantity delivered or acceptance of the delivery.
	Notes                 string    `bun:"notes,type:varchar(255),nullzero"`          // Notes the driver took about this pick ticket during delviery.
	DateCreated           time.Time `bun:"date_created,type:datetime"`                // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`          // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`       // ID of the user who last maintained this record
}
