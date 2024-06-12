package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InventoryDefaults struct {
	bun.BaseModel             `bun:"table:inventory_defaults"`
	CompanyId                 string    `bun:"company_id,type:varchar(8),pk"`
	LocationId                float64   `bun:"location_id,type:decimal(19,0),pk"`
	AssetAccount              string    `bun:"asset_account,type:varchar(32)"`
	RevenueAccount            string    `bun:"revenue_account,type:varchar(32),nullzero"`
	CosAccount                string    `bun:"cos_account,type:varchar(32),nullzero"`
	ProtectedStockQty         float64   `bun:"protected_stock_qty,type:decimal(19,9)"`
	InvMin                    float64   `bun:"inv_min,type:decimal(19,9)"`
	InvMax                    float64   `bun:"inv_max,type:decimal(19,9)"`
	SafetyStock               float64   `bun:"safety_stock,type:decimal(4,0),nullzero"`
	Emq                       float64   `bun:"emq,type:decimal(19,9)"`
	Stockable                 string    `bun:"stockable,type:char"`
	ReplenishmentLocation     float64   `bun:"replenishment_location,type:decimal(19,0)"`
	ProductGroupId            string    `bun:"product_group_id,type:varchar(8),nullzero"`
	ReplenishmentMethod       string    `bun:"replenishment_method,type:varchar(8)"`
	Serialized                string    `bun:"serialized,type:char"`
	SalesPricingUnit          string    `bun:"sales_pricing_unit,type:varchar(8)"`
	SalesPricingUnitSize      float64   `bun:"sales_pricing_unit_size,type:decimal(19,4)"`
	SupplierId                float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`
	DivisionId                float64   `bun:"division_id,type:decimal(19,0),nullzero"`
	LeadTimeDays              float64   `bun:"lead_time_days,type:decimal(4,0)"`
	ListPrice                 float64   `bun:"list_price,type:decimal(19,9)"`
	Cost                      float64   `bun:"cost,type:decimal(19,9)"`
	AvgLeadTime               float64   `bun:"avg_lead_time,type:decimal(19,0)"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	PurchasePricingUnit       string    `bun:"purchase_pricing_unit,type:varchar(8)"`
	PurchasePricingUnitSize   float64   `bun:"purchase_pricing_unit_size,type:decimal(19,4)"`
	TrackLots                 string    `bun:"track_lots,type:char"`
	TrackBins                 string    `bun:"track_bins,type:char"`
	DefaultSalesDiscountGroup string    `bun:"default_sales_discount_group,type:varchar(8),nullzero"`
	DefaultPurchaseDiscGroup  string    `bun:"default_purchase_disc_group,type:varchar(8),nullzero"`
	TaxGroupId                string    `bun:"tax_group_id,type:varchar(10),nullzero"`
	DefaultBin                string    `bun:"default_bin,type:varchar(10),nullzero"`
	Buy                       string    `bun:"buy,type:char,default:('Y')"`
	Make                      string    `bun:"make,type:char,default:('N')"`
	ReceiptProcessFlag        string    `bun:"receipt_process_flag,type:char,default:('N')"`
	ReceiptProcessUid         int32     `bun:"receipt_process_uid,type:int,nullzero"`
	DefaultBaseUnit           string    `bun:"default_base_unit,type:varchar(8)"`
	BaseUnitIsItemUom         string    `bun:"base_unit_is_item_uom,type:char,default:('Y')"`
	SafetyStockType           int32     `bun:"safety_stock_type,type:int,nullzero"`
	ServiceLevelMeasure       int32     `bun:"service_level_measure,type:int,nullzero"`
	ServiceLevelPctGoal       float64   `bun:"service_level_pct_goal,type:decimal(19,2),nullzero"`
	ServiceSupplierId         float64   `bun:"service_supplier_id,type:decimal(19,0),nullzero"`
	ServiceDivisionId         float64   `bun:"service_division_id,type:decimal(19,0),nullzero"`
	ServiceProductGroupId     string    `bun:"service_product_group_id,type:varchar(8),nullzero"`
	DefaultPriceFamilyUid     int32     `bun:"default_price_family_uid,type:int,nullzero"`
	AvailForSchDeliveryFlag   string    `bun:"avail_for_sch_delivery_flag,type:char,nullzero"`
	ManufacturingClassId      string    `bun:"manufacturing_class_id,type:varchar(255),nullzero"`
	CommissionClassId         string    `bun:"commission_class_id,type:varchar(255),nullzero"`
	ClassId1                  string    `bun:"class_id1,type:varchar(255),nullzero"`
	ClassId2                  string    `bun:"class_id2,type:varchar(255),nullzero"`
	ClassId3                  string    `bun:"class_id3,type:varchar(255),nullzero"`
	ClassId4                  string    `bun:"class_id4,type:varchar(255),nullzero"`
	ClassId5                  string    `bun:"class_id5,type:varchar(255),nullzero"`
	StandardCostMultiplier    float64   `bun:"standard_cost_multiplier,type:decimal(19,9),nullzero"`
	SpecialCostMultiplier     float64   `bun:"special_cost_multiplier,type:decimal(19,9),nullzero"`
	SupplierCostMultiplier    float64   `bun:"supplier_cost_multiplier,type:decimal(19,9),nullzero"`
	PricingServiceOption      int32     `bun:"pricing_service_option,type:int,nullzero"`
	OverrideProfitLimitsFlag  string    `bun:"override_profit_limits_flag,type:char,nullzero"`
	MinimumOrderProfit        float64   `bun:"minimum_order_profit,type:decimal(19,9),nullzero"`
	MaximumOrderProfit        float64   `bun:"maximum_order_profit,type:decimal(19,9),nullzero"`
	RmaRevenueAccountNo       string    `bun:"rma_revenue_account_no,type:varchar(32),nullzero"`
	VendorRebateAccountNo     string    `bun:"vendor_rebate_account_no,type:varchar(32),nullzero"`
	MasterBinUid              int32     `bun:"master_bin_uid,type:int,nullzero"`
	BinTypeUid                int32     `bun:"bin_type_uid,type:int,nullzero"`
	DrpItemFlag               string    `bun:"drp_item_flag,type:char,nullzero"`
	PurchaseClass             string    `bun:"purchase_class,type:varchar(8),nullzero"`
	DeviationLookbackPds      int32     `bun:"deviation_lookback_pds,type:int,nullzero"`
	DeviationMultOnePd        float64   `bun:"deviation_mult_one_pd,type:decimal(19,9),nullzero"`
	DeviationMultTwoPd        float64   `bun:"deviation_mult_two_pd,type:decimal(19,9),nullzero"`
	DeviationMultThreePd      float64   `bun:"deviation_mult_three_pd,type:decimal(19,9),nullzero"`
	MinSafetyStockDays        float64   `bun:"min_safety_stock_days,type:decimal(19,9),nullzero"`
	MaxSafetyStockDays        float64   `bun:"max_safety_stock_days,type:decimal(19,9),nullzero"`
	CritItemDeviationMult     float64   `bun:"crit_item_deviation_mult,type:decimal(19,9),nullzero"`
	CritMinSafetyStockDays    float64   `bun:"crit_min_safety_stock_days,type:decimal(19,9),nullzero"`
	CritMaxSafetyStockDays    float64   `bun:"crit_max_safety_stock_days,type:decimal(19,9),nullzero"`
}
