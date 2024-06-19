package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type OeLineLastMarginPrice struct {
	bun.BaseModel            `bun:"table:oe_line_last_margin_price"`
	OeLineLastMarginPriceUid int32     `bun:"oe_line_last_margin_price_uid,type:int,autoincrement,scanonly,pk"` // UID for table
	OeLineUid                int32     `bun:"oe_line_uid,type:int"`                                             // Link to OE Line
	PrevSkuPrice             float64   `bun:"prev_sku_price,type:decimal(19,9),nullzero"`                       // Previous SKU Price
	PrevQtyOrdered           float64   `bun:"prev_qty_ordered,type:decimal(19,9),nullzero"`                     // Previous Qty Ordered
	PrevProfitPercent        float64   `bun:"prev_profit_percent,type:decimal(19,9),nullzero"`                  // Previous Profit Percent
	OneTimePrice             string    `bun:"one_time_price,type:char(1)"`                                      // Whether to use this for margin in the future
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`                   // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`             // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`             // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`     // User who last changed the record
	CostBasisCdUsed          int32     `bun:"cost_basis_cd_used,type:int,nullzero"`                             // The cost basis that was used to determine the previous profit percent
	ItemPricedByLastMargin   string    `bun:"item_priced_by_last_margin,type:char(1),default:('Y')"`            // Indicates whether the oe_line used last margin pricing to calculate the price.
}
