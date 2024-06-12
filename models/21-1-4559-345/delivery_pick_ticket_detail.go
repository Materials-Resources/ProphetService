package model

type DeliveryPickTicketDetail struct {
	bun.BaseModel               `bun:"table:delivery_pick_ticket_detail"`
	DeliveryPickTicketDetailUid int32     `bun:"delivery_pick_ticket_detail_uid,type:int,pk"`
	DeliveryPickTicketUid       int32     `bun:"delivery_pick_ticket_uid,type:int"`
	Disposition                 string    `bun:"disposition,type:char,nullzero"`
	ReturnToStock               string    `bun:"return_to_stock,type:char,nullzero"`
	ReasonCode                  int32     `bun:"reason_code,type:int,nullzero"`
	LineNo                      int32     `bun:"line_no,type:int"`
	ShipQuantity                float64   `bun:"ship_quantity,type:decimal(19,9)"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30)"`
}
