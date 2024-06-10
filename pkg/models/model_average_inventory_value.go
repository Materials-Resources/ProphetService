package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type AverageInventoryValue struct {
	bun.BaseModel            `bun:"table:average_inventory_value"`
	AverageInventoryValueUid int32     `bun:"average_inventory_value_uid,type:int,pk,identity"`
	DemandPeriodUid          int32     `bun:"demand_period_uid,type:int"`
	LocationId               float64   `bun:"location_id,type:decimal(19,0)"`
	InvMastUid               int32     `bun:"inv_mast_uid,type:int"`
	InventoryValue           float64   `bun:"inventory_value,type:decimal(19,9)"`
	NoOfUpdates              int32     `bun:"no_of_updates,type:int"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}

type AverageInventoryValueModel struct {
	bun bun.IDB
}

// GetByInvMastUid returns an slice of AverageInventoryValue by the given InvMastUid
func (m AverageInventoryValueModel) GetByInvMastUid(ctx context.Context, invMastUid int32) (
	[]*AverageInventoryValue, error) {
	var averageInventoryValues []*AverageInventoryValue
	err := m.bun.NewSelect().Model(&averageInventoryValues).Where("inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return averageInventoryValues, nil
}

// Delete deletes the AverageInventoryValue from the database.
func (m AverageInventoryValueModel) Delete(
	ctx context.Context, averageInventoryValue *AverageInventoryValue) error {
	_, err := m.bun.NewDelete().Model(averageInventoryValue).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
