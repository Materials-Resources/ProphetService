package models

import (
	"context"
	"github.com/uptrace/bun"
)

type InvLocStockStatus struct {
	bun.BaseModel        `bun:"table:inv_loc_stock_status"`
	InvLocStockStatusUid int32   `bun:"inv_loc_stock_status_uid,type:int,pk,identity"`
	InvMastUid           int32   `bun:"inv_mast_uid,type:int"`
	LocationId           float64 `bun:"location_id,type:decimal(19,0)"`
	QtyToTransfer        float64 `bun:"qty_to_transfer,type:decimal(19,9),default:(0)"`
	QtyForProduction     float64 `bun:"qty_for_production,type:decimal(19,9),default:(0)"`
	QtyForProcess        float64 `bun:"qty_for_process,type:decimal(19,9),default:(0)"`
	QtyOnReleaseSchedule float64 `bun:"qty_on_release_schedule,type:decimal(19,9),default:(0)"`
	QtyInProduction      float64 `bun:"qty_in_production,type:decimal(19,9),default:(0)"`
	QtyNonPickable       float64 `bun:"qty_non_pickable,type:decimal(19,9),default:(0)"`
	QtyQuarantined       float64 `bun:"qty_quarantined,type:decimal(19,9),default:(0)"`
	QtyFrozen            float64 `bun:"qty_frozen,type:decimal(19,9),default:(0)"`
	QtyOnSalesOrder      float64 `bun:"qty_on_sales_order,type:decimal(19,9),default:(0)"`
	QtyOnSalesQuote      float64 `bun:"qty_on_sales_quote,type:decimal(19,9),default:(0)"`
	QtyOnSpecialPo       float64 `bun:"qty_on_special_po,type:decimal(19,9),default:(0)"`
	QtyOnDsPo            float64 `bun:"qty_on_ds_po,type:decimal(19,9),default:(0)"`
	QtyOnSpecialPoCost   float64 `bun:"qty_on_special_po_cost,type:decimal(19,9),default:((0))"`
}

type InvLocStockStatusModel struct {
	bun bun.IDB
}

func (m *InvLocStockStatusModel) GetByInvMastUid(
	ctx context.Context, invMastUid int32) ([]*InvLocStockStatus, error) {
	var invLocStockStatus []*InvLocStockStatus
	err := m.bun.NewSelect().Model(&invLocStockStatus).Where("inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return invLocStockStatus, nil
}

// Delete deletes the InvLocStockStatus from the database.
func (m *InvLocStockStatusModel) Delete(ctx context.Context, invLocStockStatus *InvLocStockStatus) error {
	_, err := m.bun.NewDelete().Model(invLocStockStatus).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
