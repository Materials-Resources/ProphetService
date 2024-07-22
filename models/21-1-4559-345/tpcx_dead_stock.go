package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type TpcxDeadStock struct {
	bun.BaseModel         `bun:"table:tpcx_dead_stock"`
	TpcxDeadStockUid      int32     `bun:"tpcx_dead_stock_uid,type:int,pk"`                              // Primary Key
	TpcxTradingPartnerUid int32     `bun:"tpcx_trading_partner_uid,type:int,unique"`                     // Pointer to the trading partner record that this dead stock item belongs to.
	InvMastUid            int32     `bun:"inv_mast_uid,type:int,unique"`                                 // Pointer to the inventory master record for this item.
	TpcxItemDesc          *string   `bun:"tpcx_item_desc,type:varchar(255)"`                             // Selling Trading Partners description for the item.
	TpcxUpcCd             *string   `bun:"tpcx_upc_cd,type:varchar(255)"`                                // UPC or EAC code for the item
	TpcxQtyAvailable      float64   `bun:"tpcx_qty_available,type:decimal(19,9)"`                        // SKU qty available that the trading partner has available to sell.
	TpcxPrice             float64   `bun:"tpcx_price,type:decimal(19,9)"`                                // Sellers price for the item
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
