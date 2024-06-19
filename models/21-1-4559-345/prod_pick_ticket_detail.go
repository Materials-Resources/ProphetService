package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdPickTicketDetail struct {
	bun.BaseModel           `bun:"table:prod_pick_ticket_detail"`
	ProdPickTicketDetailUid int32     `bun:"prod_pick_ticket_detail_uid,type:int,pk"`                      // unique identify pick ticket detail
	ProdPickTicketHdrUid    int32     `bun:"prod_pick_ticket_hdr_uid,type:int,unique"`                     // Foreign key to prod_pick_ticket_hdr  table
	PtLineNumber            int32     `bun:"pt_line_number,type:int,unique"`                               // prod pick ticket detail line number
	UnitSize                float64   `bun:"unit_size,type:decimal(19,9)"`                                 // unit size
	UnitOfMeasure           string    `bun:"unit_of_measure,type:varchar(8)"`                              // unit of measure
	InvMastUid              int32     `bun:"inv_mast_uid,type:int"`                                        // inv mast uid
	SkuQtyPicked            float64   `bun:"sku_qty_picked,type:decimal(19,9)"`                            // qty picked in terms of sku
	SkuQtyConfirmed         float64   `bun:"sku_qty_confirmed,type:decimal(19,9)"`                         // qty confirmed in terms of sku
	RowStatusFlag           int32     `bun:"row_status_flag,type:int"`                                     // identify the status of this line
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	ProdLineNumber          float64   `bun:"prod_line_number,type:decimal(19,0),default:((0))"`            // store prod order line number
	ProdComponentNumber     float64   `bun:"prod_component_number,type:decimal(19,0),default:((0))"`       // store prod order component number
	SkuQtyCompleted         float64   `bun:"sku_qty_completed,type:decimal(19,9),default:((0))"`           // The qty previously used from this pick ticket toward assemblies that have been completed in production order processing.
	OverUnderPickingFlag    string    `bun:"over_under_picking_flag,type:char(1),default:('N')"`           // Determines if this componen line on the pick ticket is an over or under picked.
}
