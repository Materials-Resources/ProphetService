package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvBin struct {
	bun.BaseModel    `bun:"table:inv_bin"`
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	LocationId       float64   `bun:"location_id,type:decimal(19,0),pk"`                         // What is the unique location identifier for this ro
	Bin              string    `bun:"bin,type:varchar(10),pk"`                                   // Which bin holds the material?
	Quantity         float64   `bun:"quantity,type:decimal(19,9),default:(0)"`                   // Bin quantity on-hand
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	InvMastUid       int32     `bun:"inv_mast_uid,type:int,pk"`                                  // Unique identifier for the item id.
	QtyAllocated     float64   `bun:"qty_allocated,type:decimal(19,9),default:(0)"`              // Bin quantity allocated.
	InvBinUid        int32     `bun:"inv_bin_uid,type:int,unique"`                               // Uniquer row ID for the table.
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:(1037)"`                   // row status; whether inv_bin is available or not pickable
}
