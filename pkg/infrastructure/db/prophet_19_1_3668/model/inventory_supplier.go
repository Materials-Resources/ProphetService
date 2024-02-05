package model

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type InventorySupplier struct {
	bun.BaseModel `bun:"table:inventory_supplier"`

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

	InvMast InvMast `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}
