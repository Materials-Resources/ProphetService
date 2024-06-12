package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InventorySupplier struct {
	bun.BaseModel               `bun:"table:inventory_supplier"`
	InvMastUid                  int32     `bun:"inv_mast_uid,type:int"`
	SupplierId                  float64   `bun:"supplier_id,type:decimal(19,0)"`
	DivisionId                  float64   `bun:"division_id,type:decimal(19,0)"`
	LeadTimeDays                float64   `bun:"lead_time_days,type:decimal(4,0),default:(0)"`
	UpcCode                     string    `bun:"upc_code,type:varchar(14),nullzero"`
	CheckDigit                  float64   `bun:"check_digit,type:decimal(1,0),nullzero"`
	CatalogName                 string    `bun:"catalog_name,type:varchar(40),nullzero"`
	CatalogPage                 string    `bun:"catalog_page,type:varchar(6),nullzero"`
	Msds                        string    `bun:"msds,type:varchar(12),nullzero"`
	DeleteFlag                  string    `bun:"delete_flag,type:char"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	SupplierPartNo              string    `bun:"supplier_part_no,type:varchar(40),nullzero"`
	SupplierSortCode            string    `bun:"supplier_sort_code,type:varchar(10),nullzero"`
	ListPrice                   float64   `bun:"list_price,type:decimal(19,9)"`
	Cost                        float64   `bun:"cost,type:decimal(19,9)"`
	ManufacturingClassId        string    `bun:"manufacturing_class_id,type:varchar(255),nullzero"`
	BackhaulAmount              float64   `bun:"backhaul_amount,type:decimal(19,9),nullzero"`
	BackhaulType                string    `bun:"backhaul_type,type:char"`
	InventorySupplierUid        int32     `bun:"inventory_supplier_uid,type:int,pk"`
	ContractNumber              string    `bun:"contract_number,type:varchar(40),nullzero"`
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	SupplierPurchaseUnit        string    `bun:"supplier_purchase_unit,type:varchar(8),nullzero"`
	ParkerPartLength            int32     `bun:"parker_part_length,type:int,nullzero"`
	DateCostLastModified        time.Time `bun:"date_cost_last_modified,type:datetime,nullzero"`
	UpcCodeSourceTypeCd         int32     `bun:"upc_code_source_type_cd,type:int,default:((2051))"`
	FreightFactor               float64   `bun:"freight_factor,type:decimal(19,9),default:((1))"`
	FutureCost                  float64   `bun:"future_cost,type:decimal(19,9),nullzero"`
	EffectiveDate               time.Time `bun:"effective_date,type:datetime,nullzero"`
	FutureListPrice             float64   `bun:"future_list_price,type:decimal(19,9),nullzero"`
	PriceExpirationDate         time.Time `bun:"price_expiration_date,type:datetime,nullzero"`
	PriceCostLastUpdateDate     time.Time `bun:"price_cost_last_update_date,type:datetime,nullzero"`
	CostMultiplier              float64   `bun:"cost_multiplier,type:decimal(19,9),nullzero"`
	MinimumPurchaseQty          float64   `bun:"minimum_purchase_qty,type:decimal(19,9),default:((0))"`
	IncrementalPurchaseQty      float64   `bun:"incremental_purchase_qty,type:decimal(19,9),default:((0))"`
	MinimumUnitsForPurchase     float64   `bun:"minimum_units_for_purchase,type:decimal(19,4),nullzero"`
	UomForMinimumUnits          string    `bun:"uom_for_minimum_units,type:varchar(8),nullzero"`
	DfltPurchaseUnit            string    `bun:"dflt_purchase_unit,type:varchar(8),nullzero"`
	DfltPurchaseUnitSize        float64   `bun:"dflt_purchase_unit_size,type:decimal(19,4),nullzero"`
	DfltPurchasePricingUnit     string    `bun:"dflt_purchase_pricing_unit,type:varchar(8),nullzero"`
	DfltPurchasePricingUnitSize float64   `bun:"dflt_purchase_pricing_unit_size,type:decimal(19,4),nullzero"`
	PrimarySupplierFlag         string    `bun:"primary_supplier_flag,type:char,nullzero"`
	RecordUsageActualLocFlag    string    `bun:"record_usage_actual_loc_flag,type:char,nullzero"`
	EanCode                     string    `bun:"ean_code,type:varchar(255),nullzero"`
	BuybackSupplierPartNo       string    `bun:"buyback_supplier_part_no,type:varchar(255),nullzero"`
	BuybackUom                  string    `bun:"buyback_uom,type:varchar(255),nullzero"`
}
