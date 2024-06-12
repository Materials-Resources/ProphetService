package model

type PoScheduleReceipts struct {
	bun.BaseModel         `bun:"table:po_schedule_receipts"`
	PoScheduleReceiptsUid int32     `bun:"po_schedule_receipts_uid,type:int,pk"`
	ReceiptNumber         float64   `bun:"receipt_number,type:decimal(19,0)"`
	LineNumber            int32     `bun:"line_number,type:int"`
	PoLineScheduleUid     int32     `bun:"po_line_schedule_uid,type:int"`
	QtyReceived           float64   `bun:"qty_received,type:decimal(19,9)"`
	ReceivedDate          time.Time `bun:"received_date,type:datetime"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
