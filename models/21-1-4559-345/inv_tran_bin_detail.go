package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvTranBinDetail struct {
	bun.BaseModel       `bun:"table:inv_tran_bin_detail"`
	CompanyId           string    `bun:"company_id,type:varchar(8)"`                   // Unique code that identifies a company.
	LocationId          float64   `bun:"location_id,type:decimal(19,0)"`               // Where was the used material located?
	Bin                 string    `bun:"bin,type:varchar(10)"`                         // The bin number.
	TransactionNumber   int32     `bun:"transaction_number,type:int,default:((-1))"`   // This column is unused.
	Quantity            float64   `bun:"quantity,type:decimal(19,9)"`                  // The change to quantity on-hand.
	DateCreated         time.Time `bun:"date_created,type:datetime"`                   // Indicates the date/time this record was created.
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`             // Indicates the date/time this record was last modified.
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30)"`          // ID of the user who last maintained this record
	InvTranBinDetailUid int32     `bun:"inv_tran_bin_detail_uid,type:int,pk"`          // Unique identifier for the record.
	DocumentLineBinUid  int32     `bun:"document_line_bin_uid,type:int"`               // What is the unique -  internal identifier for this document line bin?
	InvMastUid          int32     `bun:"inv_mast_uid,type:int"`                        // Unique identifier for the item id.
	QtyAllocated        float64   `bun:"qty_allocated,type:decimal(19,9),default:(0)"` // The change to quantity allocated.
}
