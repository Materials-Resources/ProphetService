package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type InventorySupplier struct {
	bun.BaseModel               `bun:"table:inventory_supplier"`
	InvMastUid                  int32           `bun:"inv_mast_uid,type:int"`
	SupplierId                  float64         `bun:"supplier_id,type:decimal(19,0)"`
	DivisionId                  float64         `bun:"division_id,type:decimal(19,0)"`
	LeadTimeDays                float64         `bun:"lead_time_days,type:decimal(4,0),default:(0)"`
	UpcCode                     sql.NullString  `bun:"upc_code,type:varchar(14),nullzero"`
	CheckDigit                  sql.NullFloat64 `bun:"check_digit,type:decimal(1,0),nullzero"`
	CatalogName                 sql.NullString  `bun:"catalog_name,type:varchar(40),nullzero"`
	CatalogPage                 sql.NullString  `bun:"catalog_page,type:varchar(6),nullzero"`
	Msds                        sql.NullString  `bun:"msds,type:varchar(12),nullzero"`
	DeleteFlag                  string          `bun:"delete_flag,type:char"`
	DateCreated                 time.Time       `bun:"date_created,type:datetime"`
	DateLastModified            time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	SupplierPartNo              sql.NullString  `bun:"supplier_part_no,type:varchar(40),nullzero"`
	SupplierSortCode            sql.NullString  `bun:"supplier_sort_code,type:varchar(10),nullzero"`
	ListPrice                   float64         `bun:"list_price,type:decimal(19,9)"`
	Cost                        float64         `bun:"cost,type:decimal(19,9)"`
	ManufacturingClassId        sql.NullString  `bun:"manufacturing_class_id,type:varchar(255),nullzero"`
	BackhaulAmount              sql.NullFloat64 `bun:"backhaul_amount,type:decimal(19,9),nullzero"`
	BackhaulType                string          `bun:"backhaul_type,type:char"`
	InventorySupplierUid        int32           `bun:"inventory_supplier_uid,type:int,pk"`
	ContractNumber              sql.NullString  `bun:"contract_number,type:varchar(40),nullzero"`
	CreatedBy                   sql.NullString  `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	SupplierPurchaseUnit        sql.NullString  `bun:"supplier_purchase_unit,type:varchar(8),nullzero"`
	ParkerPartLength            sql.NullInt32   `bun:"parker_part_length,type:int,nullzero"`
	DateCostLastModified        sql.NullTime    `bun:"date_cost_last_modified,type:datetime,nullzero"`
	UpcCodeSourceTypeCd         int32           `bun:"upc_code_source_type_cd,type:int,default:((2051))"`
	FreightFactor               float64         `bun:"freight_factor,type:decimal(19,9),default:((1))"`
	FutureCost                  sql.NullFloat64 `bun:"future_cost,type:decimal(19,9),nullzero"`
	EffectiveDate               sql.NullTime    `bun:"effective_date,type:datetime,nullzero"`
	FutureListPrice             sql.NullFloat64 `bun:"future_list_price,type:decimal(19,9),nullzero"`
	PriceExpirationDate         sql.NullTime    `bun:"price_expiration_date,type:datetime,nullzero"`
	PriceCostLastUpdateDate     sql.NullTime    `bun:"price_cost_last_update_date,type:datetime,nullzero"`
	CostMultiplier              sql.NullFloat64 `bun:"cost_multiplier,type:decimal(19,9),nullzero"`
	MinimumPurchaseQty          float64         `bun:"minimum_purchase_qty,type:decimal(19,9),default:((0))"`
	IncrementalPurchaseQty      float64         `bun:"incremental_purchase_qty,type:decimal(19,9),default:((0))"`
	MinimumUnitsForPurchase     sql.NullFloat64 `bun:"minimum_units_for_purchase,type:decimal(19,4),nullzero"`
	UomForMinimumUnits          sql.NullString  `bun:"uom_for_minimum_units,type:varchar(8),nullzero"`
	DfltPurchaseUnit            sql.NullString  `bun:"dflt_purchase_unit,type:varchar(8),nullzero"`
	DfltPurchaseUnitSize        sql.NullFloat64 `bun:"dflt_purchase_unit_size,type:decimal(19,4),nullzero"`
	DfltPurchasePricingUnit     sql.NullString  `bun:"dflt_purchase_pricing_unit,type:varchar(8),nullzero"`
	DfltPurchasePricingUnitSize sql.NullFloat64 `bun:"dflt_purchase_pricing_unit_size,type:decimal(19,4),nullzero"`
	PrimarySupplierFlag         sql.NullString  `bun:"primary_supplier_flag,type:char,nullzero"`
	RecordUsageActualLocFlag    sql.NullString  `bun:"record_usage_actual_loc_flag,type:char,nullzero"`
	EanCode                     sql.NullString  `bun:"ean_code,type:varchar(255),nullzero"`
	BuybackSupplierPartNo       sql.NullString  `bun:"buyback_supplier_part_no,type:varchar(255),nullzero"`
	BuybackUom                  sql.NullString  `bun:"buyback_uom,type:varchar(255),nullzero"`

	InventorySupplierXLoc *InventorySupplierXLoc `bun:"rel:has-one,join:inventory_supplier_uid=inventory_supplier_uid"`
	InvMast               InvMast                `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}

var _ bun.BeforeAppendModelHook = (*InventorySupplier)(nil)

func (i *InventorySupplier) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		i.DateCreated = time.Now()
		i.DateLastModified = time.Now()
		i.CreatedBy = sql.NullString{String: i.LastMaintainedBy, Valid: true}
	case *bun.UpdateQuery:
		i.DateLastModified = time.Now()
	}
	return nil
}

type InventorySupplierModel struct {
	bun bun.IDB
}

func (m *InventorySupplierModel) New(
	ctx context.Context, invMastUid int32, supplierId,
	divisionId, leadTimeDays float64, lastMaintainedBy, supplierPartNo string, listPrice,
	cost, minimumPurchaseQty, incrementalPurchaseQty float64) (
	*InventorySupplier,
	error) {

	inventorySupplierUid, err := m.generateInventorySupplierUid(ctx)
	if err != nil {
		return nil, err
	}

	inventorySupplier := &InventorySupplier{
		InventorySupplierUid:   inventorySupplierUid,
		InvMastUid:             invMastUid,
		SupplierId:             supplierId,
		DivisionId:             divisionId,
		LeadTimeDays:           leadTimeDays,
		DeleteFlag:             "N",
		SupplierPartNo:         sql.NullString{String: supplierPartNo, Valid: true},
		LastMaintainedBy:       lastMaintainedBy,
		ListPrice:              listPrice,
		Cost:                   cost,
		MinimumPurchaseQty:     minimumPurchaseQty,
		IncrementalPurchaseQty: incrementalPurchaseQty,
	}

	err = m.Insert(ctx, inventorySupplier)
	if err != nil {
		return nil, err
	}
	return inventorySupplier, nil

}
func (m *InventorySupplierModel) Insert(ctx context.Context, inventorySupplier *InventorySupplier) error {
	_, err := m.bun.NewInsert().Model(inventorySupplier).Exec(ctx)
	if err != nil {
		return err
	}
	return nil

}
func (m *InventorySupplierModel) Get(ctx context.Context) error {
	return nil
}

func (m *InventorySupplierModel) GetBySupplierIdDivisionIdInvMastUid(
	ctx context.Context, supplierId, divisionId float64, invMastUid int32) (*InventorySupplier, error) {
	inventorySupplier := new(InventorySupplier)
	err := m.bun.NewSelect().Model(inventorySupplier).Where(
		"supplier_id = ? AND division_id = ? AND inv_mast_uid = ?", supplierId, divisionId, invMastUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return inventorySupplier, nil
}

func (m *InventorySupplierModel) GetByInvMastUid(
	ctx context.Context, invMastUid int32) ([]*InventorySupplier, error) {
	var inventorySuppliers []*InventorySupplier
	err := m.bun.NewSelect().Model(&inventorySuppliers).Where("inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return nil, err

	}
	return inventorySuppliers, nil
}

func (m *InventorySupplierModel) Update(ctx context.Context, inventorySupplier *InventorySupplier) error {
	_, err := m.bun.NewUpdate().Model(inventorySupplier).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *InventorySupplierModel) Delete(ctx context.Context, inventorySupplier *InventorySupplier) error {
	_, err := m.bun.NewDelete().Model(inventorySupplier).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *InventorySupplierModel) generateInventorySupplierUid(ctx context.Context) (int32, error) {
	query := `DECLARE @id int
			EXEC @id = p21_get_counter 'inventory_supplier', 1
			SELECT @id`
	var inventorySupplierUid int32
	err := m.bun.QueryRowContext(ctx, query).Scan(ctx, &inventorySupplierUid)
	if err != nil {
		return 0, err
	}
	return inventorySupplierUid, nil
}
