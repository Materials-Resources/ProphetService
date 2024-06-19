package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdPickTicketHdr struct {
	bun.BaseModel            `bun:"table:prod_pick_ticket_hdr"`
	ProdPickTicketHdrUid     int32     `bun:"prod_pick_ticket_hdr_uid,type:int,pk"`                         // identify production order pick ticket
	ProdPickTicketNumber     int32     `bun:"prod_pick_ticket_number,type:int,unique"`                      // production order pick ticket number
	ProdOrderNumber          float64   `bun:"prod_order_number,type:decimal(19,0)"`                         // store production order number where this pick ticket created from
	LocationId               float64   `bun:"location_id,type:decimal(19,0)"`                               // location where pick component from
	PrintDate                time.Time `bun:"print_date,type:datetime"`                                     // date when print the pick ticket
	RowStatusFlag            int32     `bun:"row_status_flag,type:int"`                                     // identify the status of the pick ticket
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	RetrievedByWms           string    `bun:"retrieved_by_wms,type:char(1),default:('N')"`                  // Determines whether or not this record has been pulled into an external WMS system
	ConfirmableRowStatusFlag int32     `bun:"confirmable_row_status_flag,type:int,default:((1980))"`        // Indicates current picking status.
}
