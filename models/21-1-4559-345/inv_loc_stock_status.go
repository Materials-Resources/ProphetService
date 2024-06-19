package gen

import "github.com/uptrace/bun"

type InvLocStockStatus struct {
	bun.BaseModel        `bun:"table:inv_loc_stock_status"`
	InvLocStockStatusUid int32   `bun:"inv_loc_stock_status_uid,type:int,autoincrement,identity,pk"` // Uniquely Identifies each record in the table.
	InvMastUid           int32   `bun:"inv_mast_uid,type:int,unique"`                                // Uniquely indentifies the item
	LocationId           float64 `bun:"location_id,type:decimal(19,0),unique"`                       // Unique indentifier for the location
	QtyToTransfer        float64 `bun:"qty_to_transfer,type:decimal(19,9),default:(0)"`              // Qty on transfer at the source location
	QtyForProduction     float64 `bun:"qty_for_production,type:decimal(19,9),default:(0)"`           // Qty on production order components
	QtyForProcess        float64 `bun:"qty_for_process,type:decimal(19,9),default:(0)"`              // Qty on secondary processes as the raw item
	QtyOnReleaseSchedule float64 `bun:"qty_on_release_schedule,type:decimal(19,9),default:(0)"`      // Qty on release schedules
	QtyInProduction      float64 `bun:"qty_in_production,type:decimal(19,9),default:(0)"`            // Qty currently being made on production orders
	QtyNonPickable       float64 `bun:"qty_non_pickable,type:decimal(19,9),default:(0)"`             // Qty currently in bins marked not pickable
	QtyQuarantined       float64 `bun:"qty_quarantined,type:decimal(19,9),default:(0)"`              // Qty currently in bins marked quarantined
	QtyFrozen            float64 `bun:"qty_frozen,type:decimal(19,9),default:(0)"`                   // Bin frozen quantity (quantity on cycle count)
	QtyOnSalesOrder      float64 `bun:"qty_on_sales_order,type:decimal(19,9),default:(0)"`           // Current Qty on Sales Orders
	QtyOnSalesQuote      float64 `bun:"qty_on_sales_quote,type:decimal(19,9),default:(0)"`           // The quantity of inventory that has been quoted
	QtyOnSpecialPo       float64 `bun:"qty_on_special_po,type:decimal(19,9),default:(0)"`            // Qty currently on special order POs
	QtyOnDsPo            float64 `bun:"qty_on_ds_po,type:decimal(19,9),default:(0)"`                 // Qty currently on direct ship POs
	QtyOnSpecialPoCost   float64 `bun:"qty_on_special_po_cost,type:decimal(19,9),default:((0))"`     // column holds qty in special inventory for item using specific cost
}
