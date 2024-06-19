package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type GporRun struct {
	bun.BaseModel                 `bun:"table:gpor_run"`
	GporRunUid                    int32     `bun:"gpor_run_uid,type:int,autoincrement,pk"`
	GporRunHdrUid                 int32     `bun:"gpor_run_hdr_uid,type:int"`
	SequenceNo                    int16     `bun:"sequence_no,type:smallint"`
	ItemId                        string    `bun:"item_id,type:varchar(40)"`
	LocationId                    float64   `bun:"location_id,type:decimal(19,0)"`
	SupplierId                    float64   `bun:"supplier_id,type:decimal(19,0)"`
	PurchaseClass                 string    `bun:"purchase_class,type:char(8),nullzero"`
	PurchaseDiscountGroup         string    `bun:"purchase_discount_group,type:char(8),nullzero"`
	ProductGroupId                string    `bun:"product_group_id,type:char(8),nullzero"`
	ItemDesc                      string    `bun:"item_desc,type:varchar(40),nullzero"`
	ReviewCycle                   int16     `bun:"review_cycle,type:smallint,nullzero"`
	LeadTimeSource                string    `bun:"lead_time_source,type:varchar(8),nullzero"`
	SafetyStockDays               float64   `bun:"safety_stock_days,type:decimal(18,4),nullzero"`
	AverageLeadTime               int16     `bun:"average_lead_time,type:smallint,nullzero"`
	QtyOnHand                     float64   `bun:"qty_on_hand,type:decimal(19,9),nullzero"`
	QtyAllocated                  float64   `bun:"qty_allocated,type:decimal(19,9),nullzero"`
	QtyBackordered                float64   `bun:"qty_backordered,type:decimal(19,9),nullzero"`
	RecommendedQtyToOrder         float64   `bun:"recommended_qty_to_order,type:decimal(19,9),nullzero"`
	EoqUptoOrderPoint             float64   `bun:"eoq_upto_order_point,type:decimal(19,9),nullzero"`
	InvMin                        float64   `bun:"inv_min,type:decimal(19,9),nullzero"`
	InvMax                        float64   `bun:"inv_max,type:decimal(19,9),nullzero"`
	DivisionId                    float64   `bun:"division_id,type:decimal(19,0),nullzero"`
	DaysInPeriod                  int16     `bun:"days_in_period,type:smallint,nullzero"`
	CompanyId                     string    `bun:"company_id,type:char(8),nullzero"`
	PurchasePricingUnit           string    `bun:"purchase_pricing_unit,type:char(8),nullzero"`
	PurchasePricingUnitSize       float64   `bun:"purchase_pricing_unit_size,type:decimal(19,4),nullzero"`
	DefaultCarrier                float64   `bun:"default_carrier,type:decimal(19,0),nullzero"`
	RecommendedQtyToOrderCalc     float64   `bun:"recommended_qty_to_order_calc,type:decimal(19,9),nullzero"`
	DailyUsage                    float64   `bun:"daily_usage,type:decimal(19,9),nullzero"`
	ReplenishmentMethod           string    `bun:"replenishment_method,type:char(8),nullzero"`
	CurrencyId                    int16     `bun:"currency_id,type:smallint,nullzero"`
	PeriodUsage                   float64   `bun:"period_usage,type:decimal(19,9),nullzero"`
	UnitSize                      float64   `bun:"unit_size,type:decimal(19,4),nullzero"`
	UpcCode                       string    `bun:"upc_code,type:char(14),nullzero"`
	ControlValue                  string    `bun:"control_value,type:char(16),nullzero"`
	EoqUptoOrderQuantity          float64   `bun:"eoq_upto_order_quantity,type:decimal(19,9),nullzero"`
	EachCost                      float64   `bun:"each_cost,type:decimal(19,9),nullzero"`
	PeriodsToSupply               float64   `bun:"periods_to_supply,type:decimal(10,4),nullzero"`
	LeadTimeSafetyFactor          float64   `bun:"lead_time_safety_factor,type:decimal(19,4),nullzero"`
	BuyAheadQty                   float64   `bun:"buy_ahead_qty,type:decimal(19,9),nullzero"`
	QtyReservedDueIn              float64   `bun:"qty_reserved_due_in,type:decimal(19,9),nullzero"`
	ItemWeight                    float64   `bun:"item_weight,type:decimal(19,6),nullzero"`
	ItemNetWeight                 float64   `bun:"item_net_weight,type:decimal(19,6),nullzero"`
	ItemCube                      float64   `bun:"item_cube,type:decimal(19,6),nullzero"`
	ItemPurchasingWeight          float64   `bun:"item_purchasing_weight,type:decimal(19,6),nullzero"`
	TargetValue                   float64   `bun:"target_value,type:decimal(19,9),nullzero"`
	OrderQuantity                 float64   `bun:"order_quantity,type:decimal(19,9),nullzero"`
	UnitOfMeasure                 string    `bun:"unit_of_measure,type:char(8),nullzero"`
	UnapprovedPoQty               float64   `bun:"unapproved_po_qty,type:decimal(19,9),nullzero"`
	Stockable                     string    `bun:"stockable,type:char(1),nullzero"`
	BlanketPoQty                  float64   `bun:"blanket_PO_qty,type:decimal(19,9),nullzero"`
	QtyInTransit                  float64   `bun:"qty_in_transit,type:decimal(19,9),nullzero"`
	PricingUnit                   string    `bun:"pricing_unit,type:char(8),nullzero"`
	PricingUnitSize               float64   `bun:"pricing_unit_size,type:decimal(19,4),nullzero"`
	NetStock                      float64   `bun:"net_stock,type:decimal(19,9),nullzero"`
	TboQty                        float64   `bun:"tbo_qty,type:decimal(19,9),nullzero"`
	NumberOfDemandPeriods         int16     `bun:"number_of_demand_periods,type:smallint,nullzero"`
	RequiredDate                  time.Time `bun:"required_date,type:datetime,nullzero"`
	DueDate                       time.Time `bun:"due_date,type:datetime,nullzero"`
	ComputedSkuWeight             float64   `bun:"computed_sku_weight,type:decimal(19,6),nullzero"`
	QtyOnReleaseSchedule          float64   `bun:"qty_on_release_schedule,type:decimal(19,9),nullzero"`
	InvMastUid                    int32     `bun:"inv_mast_uid,type:int,nullzero"`
	ExtendedDesc                  string    `bun:"extended_desc,type:varchar(255),nullzero"`
	SafetyStockFactor             float64   `bun:"safety_stock_factor,type:decimal(19,4),nullzero"`
	UsageInUnits                  float64   `bun:"usage_in_units,type:decimal(19,9),nullzero"`
	CarryingCost                  float64   `bun:"carrying_cost,type:decimal(19,9),nullzero"`
	CostToOrder                   float64   `bun:"cost_to_order,type:decimal(19,9),nullzero"`
	PrimaryLocationStockable      string    `bun:"primary_location_stockable,type:char(1),nullzero"`
	QtyInProcess                  float64   `bun:"qty_in_process,type:decimal(19,9),nullzero"`
	UserLineNo                    string    `bun:"user_line_no,type:varchar(8),nullzero"`
	Pendingbackorderqty           float64   `bun:"pendingBackorderqty,type:decimal(19,9),nullzero"`
	LastSaleDate                  time.Time `bun:"last_sale_date,type:datetime,nullzero"`
	FutureOrderQty                float64   `bun:"future_order_qty,type:decimal(19,9),nullzero"`
	DefaultTransferUnit           string    `bun:"default_transfer_unit,type:char(8),nullzero"`
	TransferUnitSize              float64   `bun:"transfer_unit_size,type:decimal(19,4),nullzero"`
	DateCreated                   time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                     string    `bun:"created_by,type:varchar(40),default:(suser_sname())"`
	DateLastModified              time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy              string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	CoverBackorders               int32     `bun:"cover_backorders,type:int,default:(0)"` // 1 if covering backorders, 0 if not
	CustomerId                    float64   `bun:"customer_id,type:decimal(18,0),nullzero"`
	CustomerName                  string    `bun:"customer_name,type:varchar(255),nullzero"`
	OrderNo                       string    `bun:"order_no,type:varchar(8),nullzero"`
	OrderLineNo                   float64   `bun:"order_line_no,type:decimal(18,0),nullzero"`
	BaseQty                       float64   `bun:"base_qty,type:decimal(18,9),nullzero"`
	DeliveryInstructions          string    `bun:"delivery_instructions,type:varchar(255),nullzero"`
	LotBillInd                    string    `bun:"lot_bill_ind,type:char(1),default:('N')"`
	OeLineUid                     int32     `bun:"oe_line_uid,type:int,nullzero"`
	ParentOeLineUid               int32     `bun:"parent_oe_line_uid,type:int,nullzero"`
	DetailType                    int32     `bun:"detail_type,type:int,nullzero"`
	LotGroupItemId                string    `bun:"lot_group_item_id,type:varchar(40),nullzero"`
	LotGroupOeLineUid             int32     `bun:"lot_group_oe_line_uid,type:int,nullzero"`
	NumberOfPos                   int32     `bun:"number_of_pos,type:int,nullzero"`
	ReqGroupCd                    string    `bun:"req_group_cd,type:char(1),default:('O')"`
	ComponentNumber               float64   `bun:"component_number,type:decimal(18,0),default:(0)"`
	TransferCandidateInd          string    `bun:"transfer_candidate_ind,type:char(1),default:('N')"`        // If Y, then record generated as a possible candidate to tranfer. If N, then not a candidate for transfer
	AvailableToTransferQty        float64   `bun:"available_to_transfer_qty,type:decimal(19,9),default:(0)"` // Used to store amount that can potentially be transferred to another location. Not always populated. depends on how GPOR was run
	InventorySupplierXLocUid      int32     `bun:"inventory_supplier_x_loc_uid,type:int,nullzero"`
	XferNoOfPeriodsToSupply       int16     `bun:"xfer_no_of_periods_to_supply,type:tinyint,nullzero"` // this is the number of periods to review when periods supply is chosen as the available to transfer option
	AvailableToTransferCalcCd     int32     `bun:"available_to_transfer_calc_cd,type:int,nullzero"`    // Integer value that represents the algorithm used to calculate the available to transfer quantity from GPOR
	MakeBuyInd                    string    `bun:"make_buy_ind,type:char(1),default:('B')"`
	FreightControlValue           string    `bun:"freight_control_value,type:varchar(16),nullzero"`
	FreightTargetValue            float64   `bun:"freight_target_value,type:decimal(19,9),nullzero"`
	TpcxStatus                    int32     `bun:"tpcx_status,type:int,nullzero"`
	SupplierCost                  float64   `bun:"supplier_cost,type:decimal(19,9),default:(0)"`
	LocationCost                  float64   `bun:"location_cost,type:decimal(19,9),default:(0)"`
	LocationListPrice             float64   `bun:"location_list_price,type:decimal(19,9),default:(0)"`
	QtyOnProdOrder                float64   `bun:"qty_on_prod_order,type:decimal(19,9),default:(0)"`
	PeriodsToSupplyMin            int16     `bun:"periods_to_supply_min,type:tinyint,nullzero"`
	PeriodsToSupplyMax            int16     `bun:"periods_to_supply_max,type:tinyint,nullzero"`
	RequirementLocationStockable  string    `bun:"requirement_location_stockable,type:char(1),nullzero"`
	MinReplenishmentQty           float64   `bun:"min_replenishment_qty,type:decimal(19,9),default:(0.00)"`       // Minimum qty that should be purchased for this item at this location.
	PurchaseGroupTotalForecast    float64   `bun:"purchase_group_total_forecast,type:decimal(9,0),default:((0))"` // The sum total of the period forecast for all locations in the purchase group.  Used in location based group purchasing.
	CombineStockNsSpecialFlag     string    `bun:"combine_stock_ns_special_flag,type:char(1),default:('N')"`      // Supplier level setting of whether to combine special and stock items on same PO.
	StdDeviation                  float64   `bun:"std_deviation,type:decimal(19,2),nullzero"`                     // Standard deviation for service level safety stock
	MeanAbsolutePercentError      float64   `bun:"mean_absolute_percent_error,type:decimal(19,4),nullzero"`       // Mean Average Percent Error
	SafetyStockTypeCd             int32     `bun:"safety_stock_type_cd,type:int,default:((1261))"`                // Safety stock type code
	OeLineServicePartDisp         string    `bun:"oe_line_service_part_disp,type:char(1),nullzero"`               // If this requirement line is associated with a service order part then this column contains the disposition of that part line.
	UptoCalculation               float64   `bun:"upto_calculation,type:decimal(19,9),nullzero"`
	QtyQuarantined                float64   `bun:"qty_quarantined,type:decimal(19,9),default:((0))"` // Column will hold quantity that is in quarantined bin for an item.
	PassThroughFlag               string    `bun:"pass_through_flag,type:char(1),nullzero"`          // This column will be similar to po_line.pass_through_flag -  it will be used to group lines that should be on different POs
	TwelveMonthAvgPeriodUsage     float64   `bun:"twelve_month_avg_period_usage,type:decimal(19,9),nullzero"`
	RemnantQty                    float64   `bun:"remnant_qty,type:decimal(19,9),nullzero"`                           // The quantity of an item in remnants which is to be removed from the net stock
	UseRevisionsFlag              string    `bun:"use_revisions_flag,type:varchar(1),nullzero"`                       // indicating whether item uses revisions
	RevisionLevel                 string    `bun:"revision_level,type:varchar(255),nullzero"`                         // indicating item's current revision if it uses revisions
	ItemRevisionUid               int32     `bun:"item_revision_uid,type:int,nullzero"`                               // uid for revision_level
	QtyAvailableInOtherRevisions  float64   `bun:"qty_available_in_other_revisions,type:decimal(19,9),default:((0))"` // Column to hold quantity in other revisions that are not current.
	RecommendedQtyToOrderSku      float64   `bun:"recommended_qty_to_order_sku,type:decimal(19,9),nullzero"`
	SalesOrderProductionOrderQty  float64   `bun:"sales_order_production_order_qty,type:decimal(19,9),default:((0))"` // Open Quantity for Sales Orders Assemblies on 'P' Disposition
	HighVelocityLevel             float64   `bun:"high_velocity_level,type:decimal(19,9),nullzero"`                   // High velocity level at location that was used instead of SS
	MidVelocityLevel              float64   `bun:"mid_velocity_level,type:decimal(19,9),nullzero"`                    // Mid velocity level at location that was used instead of SS
	OrderPointDays                float64   `bun:"order_point_days,type:decimal(19,9),nullzero"`                      // Sum of different days (safety stock, review cycle, lead time) which determines how far we need to look out in purchasing for the order point.
	DrpOpForecast                 float64   `bun:"drp_op_forecast,type:decimal(19,9),nullzero"`                       // Prorated portion of future forecasts applied in the calculation of order point for DRP items.
	DrpOqForecast                 float64   `bun:"drp_oq_forecast,type:decimal(19,9),nullzero"`                       // Prorated portion of future forecasts applied in the calculation of order quantity for DRP items.
	DrpItemFlag                   string    `bun:"drp_item_flag,type:char(1),nullzero"`                               // Whether this line is treated as a DRP line in purchasing.
	DrpPeriodsToSupplyMinForecast float64   `bun:"drp_periods_to_supply_min_forecast,type:decimal(19,9),nullzero"`
	DrpPeriodsToSupplyMaxForecast float64   `bun:"drp_periods_to_supply_max_forecast,type:decimal(19,9),nullzero"`
	LookAheadDate                 time.Time `bun:"look_ahead_date,type:datetime,nullzero"` // Look ahead date for the GPOR run
	ReasonWhyNotDrpCd             int32     `bun:"reason_why_not_drp_cd,type:int,nullzero"`
	OrderQuantityDate             time.Time `bun:"order_quantity_date,type:datetime,nullzero"`       // column to hold the eoq order date
	OrderQuantityPeriod           int32     `bun:"order_quantity_period,type:int,nullzero"`          // column to hold the eoq order period
	OrderQuantityYearForPeriod    int32     `bun:"order_quantity_year_for_period,type:int,nullzero"` // column to hold the eoq order year
	QtyOnFutureOrdersFlag         string    `bun:"qty_on_future_orders_flag,type:char(1),nullzero"`
	MaxLiability                  float64   `bun:"max_liability,type:decimal(19,9),nullzero"`                  // Maximum liability for the item at the location.
	MaxTransferQty                float64   `bun:"max_transfer_qty,type:decimal(19,9),nullzero"`               // Maximum transfer quantity for the item at the location.
	AllFutureOrderQty             float64   `bun:"all_future_order_qty,type:decimal(19,9),nullzero"`           // The total quantity of this item in F dispostion across all open orders - as opposed to just the quantity that factors into the purchase requirement
	CriticalItemFlag              string    `bun:"critical_item_flag,type:char(1),nullzero"`                   // Flag whether, for purposes of safety stock calculations, this item/location should be treated as a critical item/location. Critical item/locations can use different safety stock multipliers.
	DeviationLookbackPds          int32     `bun:"deviation_lookback_pds,type:int,nullzero"`                   // When calculating safety stock by Deviation Multiplier, the number of periods to look backwards to determine the deviation.
	DeviationMultOnePd            float64   `bun:"deviation_mult_one_pd,type:decimal(19,9),nullzero"`          // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when one of the lookback periods has forecast usage less than actual usage.
	DeviationMultTwoPd            float64   `bun:"deviation_mult_two_pd,type:decimal(19,9),nullzero"`          // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when two of the lookback periods has forecast usage less than actual usage.
	DeviationMultThreePd          float64   `bun:"deviation_mult_three_pd,type:decimal(19,9),nullzero"`        // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when three or more of the lookback periods has forecast usage less than actual usage.
	MinSafetyStockDays            float64   `bun:"min_safety_stock_days,type:decimal(19,9),nullzero"`          // Minimum fence of safety stock days to use which overrides safety stock days calculation.
	MaxSafetyStockDays            float64   `bun:"max_safety_stock_days,type:decimal(19,9),nullzero"`          // Maximum fence of safety stock days to use which overrides safety stock days calculation.
	CritItemDeviationMult         float64   `bun:"crit_item_deviation_mult,type:decimal(19,9),nullzero"`       // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when the item/location is flagged as a critical item/location.
	CritMinSafetyStockDays        float64   `bun:"crit_min_safety_stock_days,type:decimal(19,9),nullzero"`     // Minimum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.
	CritMaxSafetyStockDays        float64   `bun:"crit_max_safety_stock_days,type:decimal(19,9),nullzero"`     // Maximum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.
	MedianForecastDeviation       float64   `bun:"median_forecast_deviation,type:decimal(19,9),nullzero"`      // When calculating safety stock by Deviation Multiplier, this will be the calculated median forecast deviation across the deviation_lookback_pds
	EffectiveDeviationMultiplier  float64   `bun:"effective_deviation_multiplier,type:decimal(19,9),nullzero"` // When calculating safety stock by Deviation Multiplier, this will be the actual deviation multiplier used to calculate safety stock days (based on whether item is marked as critical and/or how many periods have deviation)
	EffectiveSafetyStockDays      float64   `bun:"effective_safety_stock_days,type:decimal(19,9),nullzero"`    // The calculated effective safety stock days used in the recommended qty to order calculation after all fences and factors are applied.
}
