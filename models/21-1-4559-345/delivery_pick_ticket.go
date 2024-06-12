package model

type DeliveryPickTicket struct {
	bun.BaseModel         `bun:"table:delivery_pick_ticket"`
	DeliveryPickTicketUid int32     `bun:"delivery_pick_ticket_uid,type:int,pk"`
	StopUid               int32     `bun:"stop_uid,type:int"`
	PickTicketNo          float64   `bun:"pick_ticket_no,type:decimal(19,0)"`
	ReconciliationFlag    string    `bun:"reconciliation_flag,type:char,nullzero"`
	Notes                 string    `bun:"notes,type:varchar(255),nullzero"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`
}
