package model

type OeLineLastMarginPrice struct {
	bun.BaseModel            `bun:"table:oe_line_last_margin_price"`
	OeLineLastMarginPriceUid int32     `bun:"oe_line_last_margin_price_uid,type:int,pk,identity"`
	OeLineUid                int32     `bun:"oe_line_uid,type:int"`
	PrevSkuPrice             float64   `bun:"prev_sku_price,type:decimal(19,9),nullzero"`
	PrevQtyOrdered           float64   `bun:"prev_qty_ordered,type:decimal(19,9),nullzero"`
	PrevProfitPercent        float64   `bun:"prev_profit_percent,type:decimal(19,9),nullzero"`
	OneTimePrice             string    `bun:"one_time_price,type:char"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	CostBasisCdUsed          int32     `bun:"cost_basis_cd_used,type:int,nullzero"`
	ItemPricedByLastMargin   string    `bun:"item_priced_by_last_margin,type:char,default:('Y')"`
}
