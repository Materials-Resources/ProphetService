package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InvTranLotDetail struct {
	bun.BaseModel       `bun:"table:inv_tran_lot_detail"`
	CompanyId           string    `bun:"company_id,type:varchar(8)"`                                    // Unique code that identifies a company.
	LocationId          float64   `bun:"location_id,type:decimal(19,0)"`                                // What is the unique location identifier for this record
	Lot                 string    `bun:"lot,type:varchar(40)"`                                          // Lot number
	TransactionNumber   int32     `bun:"transaction_number,type:int,default:((-1))"`                    // Transaction number from inv_tran. Not used after CC v7.0.
	Quantity            float64   `bun:"quantity,type:decimal(19,9)"`                                   // Change to the quantity on-hand.
	DateCreated         time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	QtyAllocated        float64   `bun:"qty_allocated,type:decimal(19,9),nullzero"`                     // Change to the quantity allocated.
	InvTranLotDetailUid int32     `bun:"inv_tran_lot_detail_uid,type:int,pk"`                           // Internal unique identifier for a inv_tran_lot_detail row.
	DocumentLineLotUid  int32     `bun:"document_line_lot_uid,type:int"`                                // What is the unique identifier for this document line lot?
	InvMastUid          int32     `bun:"inv_mast_uid,type:int"`                                         // Unique identifier for the item id.
	SkuCost             float64   `bun:"sku_cost,type:decimal(19,9),nullzero"`                          // Lot cost
}
