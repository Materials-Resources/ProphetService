package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PoScheduleReceipts struct {
	bun.BaseModel         `bun:"table:po_schedule_receipts"`
	PoScheduleReceiptsUid int32     `bun:"po_schedule_receipts_uid,type:int,pk"` // unique identifier for PO schedule receipts records
	ReceiptNumber         float64   `bun:"receipt_number,type:decimal(19,0)"`
	LineNumber            int32     `bun:"line_number,type:int"`                                      // The line number on the purchase order.
	PoLineScheduleUid     int32     `bun:"po_line_schedule_uid,type:int"`                             // unique identifier for PO line schedule receipts records
	QtyReceived           float64   `bun:"qty_received,type:decimal(19,9)"`                           // Amount received for this PO line schedule
	ReceivedDate          time.Time `bun:"received_date,type:datetime"`                               // Date material was received for this PO line schedule
	DateCreated           time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
