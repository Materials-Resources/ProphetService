package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdPickTicketDetail struct {
	bun.BaseModel           `bun:"table:prod_pick_ticket_detail"`
	ProdPickTicketDetailUid int32     `bun:"prod_pick_ticket_detail_uid,type:int,pk"`
	ProdPickTicketHdrUid    int32     `bun:"prod_pick_ticket_hdr_uid,type:int"`
	PtLineNumber            int32     `bun:"pt_line_number,type:int"`
	UnitSize                float64   `bun:"unit_size,type:decimal(19,9)"`
	UnitOfMeasure           string    `bun:"unit_of_measure,type:varchar(8)"`
	InvMastUid              int32     `bun:"inv_mast_uid,type:int"`
	SkuQtyPicked            float64   `bun:"sku_qty_picked,type:decimal(19,9)"`
	SkuQtyConfirmed         float64   `bun:"sku_qty_confirmed,type:decimal(19,9)"`
	RowStatusFlag           int32     `bun:"row_status_flag,type:int"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	ProdLineNumber          float64   `bun:"prod_line_number,type:decimal(19,0),default:((0))"`
	ProdComponentNumber     float64   `bun:"prod_component_number,type:decimal(19,0),default:((0))"`
	SkuQtyCompleted         float64   `bun:"sku_qty_completed,type:decimal(19,9),default:((0))"`
	OverUnderPickingFlag    string    `bun:"over_under_picking_flag,type:char,default:('N')"`
}
