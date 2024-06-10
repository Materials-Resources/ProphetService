package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type InventorySupplierXLoc struct {
	bun.BaseModel              `bun:"table:inventory_supplier_x_loc"`
	InventorySupplierXLocUid   int32           `bun:"inventory_supplier_x_loc_uid,type:int,pk"`
	InventorySupplierUid       int32           `bun:"inventory_supplier_uid,type:int"`
	LocationId                 float64         `bun:"location_id,type:decimal(19,0)"`
	PrimarySupplier            string          `bun:"primary_supplier,type:char,default:('N')"`
	AverageLeadTime            int16           `bun:"average_lead_time,type:smallint,default:(0)"`
	RowStatusFlag              int32           `bun:"row_status_flag,type:int"`
	DateCreated                time.Time       `bun:"date_created,type:datetime"`
	DateLastModified           time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	LocListPrice               sql.NullFloat64 `bun:"loc_list_price,type:decimal(19,9),default:(0)"`
	LocCost                    sql.NullFloat64 `bun:"loc_cost,type:decimal(19,9),default:(0)"`
	LocContractNumber          sql.NullString  `bun:"loc_contract_number,type:varchar(40),nullzero"`
	FixedLeadTime              sql.NullInt16   `bun:"fixed_lead_time,type:smallint,default:(0)"`
	DefaultDisposition         sql.NullString  `bun:"default_disposition,type:char,nullzero"`
	VmiStatus                  sql.NullInt32   `bun:"vmi_status,type:int,nullzero"`
	CreatedBy                  sql.NullString  `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	OverrideBegDate            sql.NullTime    `bun:"override_beg_date,type:datetime,nullzero"`
	OverrideEndDate            sql.NullTime    `bun:"override_end_date,type:datetime,nullzero"`
	VmiLastSendDate            sql.NullTime    `bun:"vmi_last_send_date,type:datetime,nullzero"`
	OverrideVmiStatus          sql.NullInt32   `bun:"override_vmi_status,type:int,nullzero"`
	OverrideVmiStatusFlag      string          `bun:"override_vmi_status_flag,type:char,default:('N')"`
	KeyVmiIndicatorChangedFlag string          `bun:"key_vmi_indicator_changed_flag,type:char,default:('N')"`
	FutureCost                 sql.NullFloat64 `bun:"future_cost,type:decimal(19,9),nullzero"`
	EffectiveDate              sql.NullTime    `bun:"effective_date,type:datetime,nullzero"`
	ManualLeadTime             sql.NullInt32   `bun:"manual_lead_time,type:int,nullzero"`
	FutureListPrice            sql.NullFloat64 `bun:"future_list_price,type:decimal(19,9),nullzero"`
}

var _ bun.BeforeAppendModelHook = (*InventorySupplierXLoc)(nil)

func (i *InventorySupplierXLoc) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		i.DateCreated = time.Now()
		i.DateLastModified = time.Now()
	case *bun.UpdateQuery:
		i.DateLastModified = time.Now()
	}
	return nil
}

type InventorySupplierXLocModel struct {
	bun bun.IDB
}

func (m InventorySupplierXLocModel) New(
	ctx context.Context, inventorySupplierUid int32, locationId float64) (*InventorySupplierXLoc, error) {
	inventorySupplierXLocUid, err := m.generateInventorySupplierXLocUid(ctx)
	if err != nil {
		return nil, err
	}
	inventorySupplierXLoc := &InventorySupplierXLoc{
		InventorySupplierXLocUid: inventorySupplierXLocUid,
		InventorySupplierUid:     inventorySupplierUid,
		LocationId:               locationId,
	}
	err = m.Insert(inventorySupplierXLoc)
	if err != nil {
		return nil, err
	}
	return inventorySupplierXLoc, nil
}

// Insert inserts the InventorySupplierXLoc into the database.
func (m InventorySupplierXLocModel) Insert(inventorySupplierXLoc *InventorySupplierXLoc) error {
	_, err := m.bun.NewInsert().Model(inventorySupplierXLoc).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// GetByInvMastUidAndLocationId returns a slice of InventorySupplierXLoc by the associated inv_mast_uid and location_id
func (m InventorySupplierXLocModel) GetByInvMastUidAndLocationId(
	ctx context.Context, invMastUid int32,
	locationId float64) ([]*InventorySupplierXLoc, error) {
	var inventorySupplierXLocs []*InventorySupplierXLoc
	err := m.bun.NewSelect().Model(&inventorySupplierXLocs).Join(
		"JOIN inventory_supplier on inventory_supplier."+
			"inventory_supplier_uid = inventory_supplier_x_loc.inventory_supplier_uid").
		Where("location_id = ? AND inventory_supplier.inv_mast_uid = ?", locationId, invMastUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return inventorySupplierXLocs, nil
}

// GetByInventorySupplierUid returns a slice of InventorySupplierXLoc by the given InventorySupplierUid
func (m InventorySupplierXLocModel) GetByInventorySupplierUid(
	ctx context.Context, inventorySupplierUid int32) ([]*InventorySupplierXLoc, error) {
	var inventorySupplierXLocs []*InventorySupplierXLoc
	err := m.bun.NewSelect().Model(&inventorySupplierXLocs).Where(
		"inventory_supplier_uid = ?", inventorySupplierUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return inventorySupplierXLocs, nil
}

// Update updates the InventorySupplierXLoc in the database.
func (m InventorySupplierXLocModel) Update(
	ctx context.Context,
	inventorySupplierXLoc *InventorySupplierXLoc) error {
	_, err := m.bun.NewUpdate().Model(inventorySupplierXLoc).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes the InventorySupplierXLoc from the database.
func (m InventorySupplierXLocModel) Delete(
	ctx context.Context,
	inventorySupplierXLoc *InventorySupplierXLoc) error {
	_, err := m.bun.NewDelete().Model(inventorySupplierXLoc).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m InventorySupplierXLocModel) generateInventorySupplierXLocUid(ctx context.Context) (int32, error) {
	query := `
		DECLARE @id int
		EXEC @oe_line_uid = p21_get_counter 'inventory_supplier_x_loc', 1
		SELECT @id`

	var id int32
	err := m.bun.QueryRowContext(ctx, query).Scan(ctx, &id)
	if err != nil {
		return 0, err
	}
	return id, nil

}
