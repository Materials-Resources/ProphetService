package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdPickTicketHdr struct {
	bun.BaseModel            `bun:"table:prod_pick_ticket_hdr"`
	ProdPickTicketHdrUid     int32     `bun:"prod_pick_ticket_hdr_uid,type:int,pk"`
	ProdPickTicketNumber     int32     `bun:"prod_pick_ticket_number,type:int"`
	ProdOrderNumber          float64   `bun:"prod_order_number,type:decimal(19,0)"`
	LocationId               float64   `bun:"location_id,type:decimal(19,0)"`
	PrintDate                time.Time `bun:"print_date,type:datetime"`
	RowStatusFlag            int32     `bun:"row_status_flag,type:int"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	RetrievedByWms           string    `bun:"retrieved_by_wms,type:char,default:('N')"`
	ConfirmableRowStatusFlag int32     `bun:"confirmable_row_status_flag,type:int,default:((1980))"`
}
