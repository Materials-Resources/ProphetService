package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InventoryDefaults struct {
	bun.BaseModel             `bun:"table:inventory_defaults"`
	CompanyId                 string    `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	LocationId                float64   `bun:"location_id,type:decimal(19,0),pk"`                         // These defaults are for what location?
	AssetAccount              string    `bun:"asset_account,type:varchar(32)"`                            // Asset Account default for an item at this location
	RevenueAccount            string    `bun:"revenue_account,type:varchar(32),nullzero"`                 // Revenue Account default for an item at this location
	CosAccount                string    `bun:"cos_account,type:varchar(32),nullzero"`                     // COS Account default for an item at this location
	ProtectedStockQty         float64   `bun:"protected_stock_qty,type:decimal(19,9)"`                    // Default protected stock quantity. Not used.
	InvMin                    float64   `bun:"inv_min,type:decimal(19,9)"`                                // Min value for replenishment method calculation.
	InvMax                    float64   `bun:"inv_max,type:decimal(19,9)"`                                // Max value for replenishment method calculation.
	SafetyStock               float64   `bun:"safety_stock,type:decimal(4,0),nullzero"`                   // Default safety stock.
	Emq                       float64   `bun:"emq,type:decimal(19,9)"`                                    // Default economic make quantity. Not used.
	Stockable                 string    `bun:"stockable,type:char(1)"`                                    // Are default items stockable?
	ReplenishmentLocation     float64   `bun:"replenishment_location,type:decimal(19,0)"`                 // What location replenishes items at this location?
	ProductGroupId            string    `bun:"product_group_id,type:varchar(8),nullzero"`                 // Default product group for an item at this location.
	ReplenishmentMethod       string    `bun:"replenishment_method,type:varchar(8)"`                      // Default method for replenishement minmax, opoq, eoq or upto.
	Serialized                string    `bun:"serialized,type:char(1)"`                                   // Does the item track serial numbers?
	SalesPricingUnit          string    `bun:"sales_pricing_unit,type:varchar(8)"`                        // default uom for sales
	SalesPricingUnitSize      float64   `bun:"sales_pricing_unit_size,type:decimal(19,4)"`                // size of default uom for this item
	SupplierId                float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`                   // What supplier is participating in this relationship?
	DivisionId                float64   `bun:"division_id,type:decimal(19,0),nullzero"`                   // What is the unique identifier for this division?
	LeadTimeDays              float64   `bun:"lead_time_days,type:decimal(4,0)"`                          // What is the lead time -  in days -  for this inventory supplier?
	ListPrice                 float64   `bun:"list_price,type:decimal(19,9)"`                             // Used to store the list price in comparison to the Deleted on or before 10/10/98 dstrait
	Cost                      float64   `bun:"cost,type:decimal(19,9)"`                                   // Default supplier cost.
	AvgLeadTime               float64   `bun:"avg_lead_time,type:decimal(19,0)"`                          // What is the average lead time for this supplier?
	DateCreated               time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	PurchasePricingUnit       string    `bun:"purchase_pricing_unit,type:varchar(8)"`                     // default uom for pricing
	PurchasePricingUnitSize   float64   `bun:"purchase_pricing_unit_size,type:decimal(19,4)"`             // size of default pricing uom for this item.
	TrackLots                 string    `bun:"track_lots,type:char(1)"`                                   // Does this item track lots
	TrackBins                 string    `bun:"track_bins,type:char(1)"`                                   // Does this item track bins
	DefaultSalesDiscountGroup string    `bun:"default_sales_discount_group,type:varchar(8),nullzero"`     // Default sales discount group.
	DefaultPurchaseDiscGroup  string    `bun:"default_purchase_disc_group,type:varchar(8),nullzero"`      // Default purchase discount group.
	TaxGroupId                string    `bun:"tax_group_id,type:varchar(10),nullzero"`                    // Indicates the tax group identification.
	DefaultBin                string    `bun:"default_bin,type:varchar(10),nullzero"`                     // Used as the default bin for an item created on the
	Buy                       string    `bun:"buy,type:char(1),default:('Y')"`                            // indicate that item can be procered at the location by a purchase order or transfer.
	Make                      string    `bun:"make,type:char(1),default:('N')"`                           // indicate that item can be procured at the location by a production or MSP order.
	ReceiptProcessFlag        string    `bun:"receipt_process_flag,type:char(1),default:('N')"`           // Determines whether the item undergoes secondary processing upon receipt
	ReceiptProcessUid         int32     `bun:"receipt_process_uid,type:int,nullzero"`                     // Identifies the pre-defined rounting for items that undergo secondary processing upon receipt
	DefaultBaseUnit           string    `bun:"default_base_unit,type:varchar(8)"`                         // Default base unit
	BaseUnitIsItemUom         string    `bun:"base_unit_is_item_uom,type:char(1),default:('Y')"`          // When enabled this will add the base unit to item_uom
	SafetyStockType           int32     `bun:"safety_stock_type,type:int,nullzero"`                       // Method used to determine safety stock
	ServiceLevelMeasure       int32     `bun:"service_level_measure,type:int,nullzero"`                   // Customer service level (stock out back order)
	ServiceLevelPctGoal       float64   `bun:"service_level_pct_goal,type:decimal(19,2),nullzero"`        // Percent of service level goal for safety stock
	ServiceSupplierId         float64   `bun:"service_supplier_id,type:decimal(19,0),nullzero"`           // Default Supplier ID used when creating an on-the-fly service item at this location.
	ServiceDivisionId         float64   `bun:"service_division_id,type:decimal(19,0),nullzero"`           // Default Division ID used when creating an on-the-fly service item at this location.
	ServiceProductGroupId     string    `bun:"service_product_group_id,type:varchar(8),nullzero"`         // Default Product Group ID used when creating an on-the-fly service item at this location.
	DefaultPriceFamilyUid     int32     `bun:"default_price_family_uid,type:int,nullzero"`                // The default price family identifier that pertains to this location and company - unique identifier for price_family table.
	AvailForSchDeliveryFlag   string    `bun:"avail_for_sch_delivery_flag,type:char(1),nullzero"`         // Determines whether or not an item will be available for scheduling on the Degree Days and Calendar Based Tabs of Ship To Maintenance.
	ManufacturingClassId      string    `bun:"manufacturing_class_id,type:varchar(255),nullzero"`         // Default manufacturing class ID
	CommissionClassId         string    `bun:"commission_class_id,type:varchar(255),nullzero"`            // Default commission class ID
	ClassId1                  string    `bun:"class_id1,type:varchar(255),nullzero"`                      // Default item class 1
	ClassId2                  string    `bun:"class_id2,type:varchar(255),nullzero"`                      // Default item class 2
	ClassId3                  string    `bun:"class_id3,type:varchar(255),nullzero"`                      // Default item class 3
	ClassId4                  string    `bun:"class_id4,type:varchar(255),nullzero"`                      // Default item class 4
	ClassId5                  string    `bun:"class_id5,type:varchar(255),nullzero"`                      // Default item class 5
	StandardCostMultiplier    float64   `bun:"standard_cost_multiplier,type:decimal(19,9),nullzero"`      // Default standard cost multiplier
	SpecialCostMultiplier     float64   `bun:"special_cost_multiplier,type:decimal(19,9),nullzero"`       // Default special cost multiplier
	SupplierCostMultiplier    float64   `bun:"supplier_cost_multiplier,type:decimal(19,9),nullzero"`      // Default supplier cost multiplier
	PricingServiceOption      int32     `bun:"pricing_service_option,type:int,nullzero"`                  // Default pricing service option
	OverrideProfitLimitsFlag  string    `bun:"override_profit_limits_flag,type:char(1),nullzero"`         // Flag that will be the default for wether or not profit limits will come from this item
	MinimumOrderProfit        float64   `bun:"minimum_order_profit,type:decimal(19,9),nullzero"`          // This will be the minimum order profit percentage for this item
	MaximumOrderProfit        float64   `bun:"maximum_order_profit,type:decimal(19,9),nullzero"`          // This will be the maximum order profit percentage for this item
	RmaRevenueAccountNo       string    `bun:"rma_revenue_account_no,type:varchar(32),nullzero"`          // Default value for rma revenu account.
	VendorRebateAccountNo     string    `bun:"vendor_rebate_account_no,type:varchar(32),nullzero"`        // A general ledger account that tracks the value of Vendor Rebates
	MasterBinUid              int32     `bun:"master_bin_uid,type:int,nullzero"`                          // Indicates master bin for this location default.
	BinTypeUid                int32     `bun:"bin_type_uid,type:int,nullzero"`                            // Indicates bin type for this location default.
	DrpItemFlag               string    `bun:"drp_item_flag,type:char(1),nullzero"`                       // Save the default value to set for same column name while creating new items.
	PurchaseClass             string    `bun:"purchase_class,type:varchar(8),nullzero"`                   // purchase_class
	DeviationLookbackPds      int32     `bun:"deviation_lookback_pds,type:int,nullzero"`                  // When calculating safety stock by Deviation Multiplier, the number of periods to look backwards to determine the deviation.
	DeviationMultOnePd        float64   `bun:"deviation_mult_one_pd,type:decimal(19,9),nullzero"`         // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when one of the lookback periods has forecast usage less than actual usage.
	DeviationMultTwoPd        float64   `bun:"deviation_mult_two_pd,type:decimal(19,9),nullzero"`         // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when two of the lookback periods has forecast usage less than actual usage.
	DeviationMultThreePd      float64   `bun:"deviation_mult_three_pd,type:decimal(19,9),nullzero"`       // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when three or more of the lookback periods has forecast usage less than actual usage.
	MinSafetyStockDays        float64   `bun:"min_safety_stock_days,type:decimal(19,9),nullzero"`         // Minimum fence of safety stock days to use which overrides safety stock days calculation.
	MaxSafetyStockDays        float64   `bun:"max_safety_stock_days,type:decimal(19,9),nullzero"`         // Maximum fence of safety stock days to use which overrides safety stock days calculation.
	CritItemDeviationMult     float64   `bun:"crit_item_deviation_mult,type:decimal(19,9),nullzero"`      // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when the item/location is flagged as a critical item/location.
	CritMinSafetyStockDays    float64   `bun:"crit_min_safety_stock_days,type:decimal(19,9),nullzero"`    // Minimum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.
	CritMaxSafetyStockDays    float64   `bun:"crit_max_safety_stock_days,type:decimal(19,9),nullzero"`    // Maximum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.
}
