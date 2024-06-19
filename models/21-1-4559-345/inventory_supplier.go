package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InventorySupplier struct {
	bun.BaseModel               `bun:"table:inventory_supplier"`
	InvMastUid                  int32     `bun:"inv_mast_uid,type:int,unique"`                              // Unique identifier for the item id.
	SupplierId                  float64   `bun:"supplier_id,type:decimal(19,0),unique"`                     // What supplier supplies material for this stage?
	DivisionId                  float64   `bun:"division_id,type:decimal(19,0),unique"`                     // What is the unique identifier for this division?
	LeadTimeDays                float64   `bun:"lead_time_days,type:decimal(4,0),default:(0)"`              // What is the lead time -  in days -  for this inventory supplier?
	UpcCode                     string    `bun:"upc_code,type:varchar(14),nullzero"`                        // Unique identifier, used on barcodes.
	CheckDigit                  float64   `bun:"check_digit,type:decimal(1,0),nullzero"`                    // Used to ensure the accuracy of the UPC barcode.
	CatalogName                 string    `bun:"catalog_name,type:varchar(40),nullzero"`                    // Catalog name. Display only.
	CatalogPage                 string    `bun:"catalog_page,type:varchar(6),nullzero"`                     // Catalog page. Display only.
	Msds                        string    `bun:"msds,type:varchar(12),nullzero"`                            // MSDS info. Display only. Preferred method is to use inventory links.
	DeleteFlag                  string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated                 time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	SupplierPartNo              string    `bun:"supplier_part_no,type:varchar(40),nullzero"`                // What the supplier calls this item.
	SupplierSortCode            string    `bun:"supplier_sort_code,type:varchar(10),nullzero"`              // This column is unused.
	ListPrice                   float64   `bun:"list_price,type:decimal(19,9)"`                             // The supplier list price.
	Cost                        float64   `bun:"cost,type:decimal(19,9)"`                                   // The supplier cost.
	ManufacturingClassId        string    `bun:"manufacturing_class_id,type:varchar(255),nullzero"`         // What is the unique identifier for this manufacturing class?
	BackhaulAmount              float64   `bun:"backhaul_amount,type:decimal(19,9),nullzero"`               // Indicates the amount that the supplier credits or
	BackhaulType                string    `bun:"backhaul_type,type:char(1)"`                                // Indicates whether the supplier credits or charges
	InventorySupplierUid        int32     `bun:"inventory_supplier_uid,type:int,pk"`                        // Unique Identifier for this combination of item and supplier.
	ContractNumber              string    `bun:"contract_number,type:varchar(40),nullzero"`                 // contract number assigned at item/supplier level
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	SupplierPurchaseUnit        string    `bun:"supplier_purchase_unit,type:varchar(8),nullzero"`             // Preferred UOM for EDI transactions
	ParkerPartLength            int32     `bun:"parker_part_length,type:int,nullzero"`                        // Field needed for Parker Rebate reporting requirements.
	DateCostLastModified        time.Time `bun:"date_cost_last_modified,type:datetime,nullzero"`              // This column will hold a date when the cost column on this same table was last modified
	UpcCodeSourceTypeCd         int32     `bun:"upc_code_source_type_cd,type:int,default:((2051))"`           // Code that tells how the UPC Code was created (via Main or Wireless application)
	FreightFactor               float64   `bun:"freight_factor,type:decimal(19,9),default:((1))"`             // Custom (F30169): used to calculate the standard cost in item maint, fast-edit, import and pricing service.
	FutureCost                  float64   `bun:"future_cost,type:decimal(19,9),nullzero"`                     // Future cost of an item
	EffectiveDate               time.Time `bun:"effective_date,type:datetime,nullzero"`                       // Date future cost will be effective
	FutureListPrice             float64   `bun:"future_list_price,type:decimal(19,9),nullzero"`               // Custom Column for Future List of an item (47571)
	PriceExpirationDate         time.Time `bun:"price_expiration_date,type:datetime,nullzero"`                // Price expiration date used in order entry to determine if pricing needs to be updated for particular item
	PriceCostLastUpdateDate     time.Time `bun:"price_cost_last_update_date,type:datetime,nullzero"`          // The date the list_price and/or cost column value was last updated
	CostMultiplier              float64   `bun:"cost_multiplier,type:decimal(19,9),nullzero"`                 // column that stored the cost multiplier
	MinimumPurchaseQty          float64   `bun:"minimum_purchase_qty,type:decimal(19,9),default:((0))"`       // Minimum quantity (in purchasing unit) that must be purchased for the item from this supplier
	IncrementalPurchaseQty      float64   `bun:"incremental_purchase_qty,type:decimal(19,9),default:((0))"`   // Incremental qty (in purchasing unit) that must be purchased for the item from this supplier above the minimum
	MinimumUnitsForPurchase     float64   `bun:"minimum_units_for_purchase,type:decimal(19,4),nullzero"`      // The vendor minimum number of units for purchase.
	UomForMinimumUnits          string    `bun:"uom_for_minimum_units,type:varchar(8),nullzero"`              // The unit of measure associated with this minimum number of units
	DfltPurchaseUnit            string    `bun:"dflt_purchase_unit,type:varchar(8),nullzero"`                 // Default purchase UOM for a supplier, to be used instead of item defaults
	DfltPurchaseUnitSize        float64   `bun:"dflt_purchase_unit_size,type:decimal(19,4),nullzero"`         // Unit size related related to default purchase unit
	DfltPurchasePricingUnit     string    `bun:"dflt_purchase_pricing_unit,type:varchar(8),nullzero"`         // Default purchase pricing UOM for a supplier, to be used instead of item defaults
	DfltPurchasePricingUnitSize float64   `bun:"dflt_purchase_pricing_unit_size,type:decimal(19,4),nullzero"` // Unit size related related to default purchase unit
	PrimarySupplierFlag         string    `bun:"primary_supplier_flag,type:char(1),nullzero"`                 // Default Supplier for Locations
	RecordUsageActualLocFlag    string    `bun:"record_usage_actual_loc_flag,type:char(1),nullzero"`          // Record Usage at source location for this supplier.
	EanCode                     string    `bun:"ean_code,type:varchar(255),nullzero"`
	BuybackSupplierPartNo       string    `bun:"buyback_supplier_part_no,type:varchar(255),nullzero"` // Custom: Indicates the supplier part number to be used for Pegmost Buyback Export for this item/supplier combo.
	BuybackUom                  string    `bun:"buyback_uom,type:varchar(255),nullzero"`              // Custom: Indicates the UOM to be used for Pegmost Buyback Export for this item/supplier combo.
}
