package prophet

import "github.com/uptrace/bun"

type InvReclassificationDetail struct {
	bun.BaseModel              `bun:"table:inv_reclassification_detail"`
	InvReclassDetailUid        int32   `bun:"inv_reclass_detail_uid,type:int,autoincrement,identity,pk"` // Uniquely Identifies each record in the table.
	LocationId                 float64 `bun:"location_id,type:decimal(19,0)"`                            // Location being ranked or reclassified
	InvMastUid                 int32   `bun:"inv_mast_uid,type:int"`                                     // Unique identifier for the item being ranked / reclassified
	SalesValue                 float64 `bun:"sales_value,type:decimal(19,2),default:(0)"`                // Value based upon the specified sales basis
	CostValue                  float64 `bun:"cost_value,type:decimal(19,2),default:(0)"`                 // Value based upon the specified cost basis
	UnitsSold                  float64 `bun:"units_sold,type:decimal(19,2),default:(0)"`                 // Units sold of the item * forecast usage * periods
	Profit                     float64 `bun:"profit,type:decimal(19,2),default:(0)"`                     // Profit * forecast usage * periods
	InvReclassificationWorkUid int32   `bun:"inv_reclassification_work_uid,type:int"`                    // Unique identifier of the record with ranking statistics
	ItemRankValue              float64 `bun:"item_rank_value,type:decimal(19,9),default:(0)"`            // Rank of the item compared to the total
	CumulativeRankValue        float64 `bun:"cumulative_rank_value,type:decimal(19,2),default:(0)"`      // Items rank expressed as an accumulated value
	CumulativeRankPercent      float64 `bun:"cumulative_rank_percent,type:decimal(19,2),default:(0)"`    // Items rank expressed as an accumulated percentage
	RankBy                     string  `bun:"rank_by,type:char(1)"`                                      // Determines which basis is used to rank/reclassify
	NewPurchaseClassId         *string `bun:"new_purchase_class_id,type:varchar(8)"`                     // Suggested new purchase class for reclassification
	NoOfOrders                 int32   `bun:"no_of_orders,type:int,default:(0)"`                         // Used for putaway ranking
}
