package model

type InvReclassificationDetail struct {
	bun.BaseModel              `bun:"table:inv_reclassification_detail"`
	InvReclassDetailUid        int32   `bun:"inv_reclass_detail_uid,type:int,pk,identity"`
	LocationId                 float64 `bun:"location_id,type:decimal(19,0)"`
	InvMastUid                 int32   `bun:"inv_mast_uid,type:int"`
	SalesValue                 float64 `bun:"sales_value,type:decimal(19,2),default:(0)"`
	CostValue                  float64 `bun:"cost_value,type:decimal(19,2),default:(0)"`
	UnitsSold                  float64 `bun:"units_sold,type:decimal(19,2),default:(0)"`
	Profit                     float64 `bun:"profit,type:decimal(19,2),default:(0)"`
	InvReclassificationWorkUid int32   `bun:"inv_reclassification_work_uid,type:int"`
	ItemRankValue              float64 `bun:"item_rank_value,type:decimal(19,9),default:(0)"`
	CumulativeRankValue        float64 `bun:"cumulative_rank_value,type:decimal(19,2),default:(0)"`
	CumulativeRankPercent      float64 `bun:"cumulative_rank_percent,type:decimal(19,2),default:(0)"`
	RankBy                     string  `bun:"rank_by,type:char"`
	NewPurchaseClassId         string  `bun:"new_purchase_class_id,type:varchar(8),nullzero"`
	NoOfOrders                 int32   `bun:"no_of_orders,type:int,default:(0)"`
}
