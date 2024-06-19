package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InvLocMsp struct {
	bun.BaseModel      `bun:"table:inv_loc_msp"`
	InvLocMspUid       int32     `bun:"inv_loc_msp_uid,type:int,pk"`                                  // Unique identifier for the record
	InvMastUid         int32     `bun:"inv_mast_uid,type:int"`                                        // Unique identifier for the item id
	LocationId         float64   `bun:"location_id,type:decimal(19,0)"`                               // Identifier for location id.
	ReceiptProcessFlag string    `bun:"receipt_process_flag,type:char(1),default:('N')"`              // Determines whether the item undergoes secondary processing upon receipt
	ReceiptProcessUid  int32     `bun:"receipt_process_uid,type:int,nullzero"`                        // Identifies the pre-defined rounting for items that undergo secondary processing upon receipt
	RowStatusFlag      int32     `bun:"row_status_flag,type:int,default:(704)"`                       // Status of the row.
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
