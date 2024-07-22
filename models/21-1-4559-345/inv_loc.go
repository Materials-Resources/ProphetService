package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvLoc struct {
	bun.BaseModel                   `bun:"table:inv_loc"`
	LocationId                      float64    `bun:"location_id,type:decimal(19,0),pk"`                         // Unique identifier for this location.
	QtyOnHand                       *float64   `bun:"qty_on_hand,type:decimal(19,9),default:(0)"`                // Qty of this item currently at this location.
	QtyInProcess                    float64    `bun:"qty_in_process,type:decimal(19,9),default:(0)"`             // Qty of this item that is on PO.
	DateCreated                     time.Time  `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified                time.Time  `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy                string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	CompanyId                       string     `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	LastRecPo                       *float64   `bun:"last_rec_po,type:decimal(19,9)"`                            // The last received PO cost of the item.
	LastRecPoWithDisc               *float64   `bun:"last_rec_po_with_disc,type:decimal(19,9)"`                  // Cost of an item the last time it was received, after any discounts were taken.
	GlAccountNo                     *string    `bun:"gl_account_no,type:varchar(32)"`                            // General Ledger Account number
	PurchOrTransfer                 *string    `bun:"purch_or_transfer,type:char(1)"`                            // Indicates if this is a purchase or transfer.
	NextDueInPoCost                 *float64   `bun:"next_due_in_po_cost,type:decimal(19,9)"`                    // Cost of the next expected PO in.
	NextDueInPoDate                 *time.Time `bun:"next_due_in_po_date,type:datetime"`                         // Date that we expect a PO with this item to arrive.
	RevenueAccountNo                *string    `bun:"revenue_account_no,type:varchar(32)"`                       // A general ledger account that tracks the value of revenue.
	CosAccountNo                    *string    `bun:"cos_account_no,type:varchar(32)"`                           // A general ledger account that tracks the value of Cost of Goods Sold.
	Sellable                        *string    `bun:"sellable,type:char(1)"`                                     // Indicates if this item sellable at this location
	MovingAverageCost               *float64   `bun:"moving_average_cost,type:decimal(19,9)"`                    // Average cost of this item - updated when material is received.
	StandardCost                    *float64   `bun:"standard_cost,type:decimal(19,9)"`                          // The standard cost of the item.
	ProtectedStockQty               *float64   `bun:"protected_stock_qty,type:decimal(19,9)"`                    // The minimum stock quantity that should be maintained for an item.
	InvMin                          *float64   `bun:"inv_min,type:decimal(19,9)"`                                // Used in calculating qty to order, different use depending on one of 4 replenishment methods.
	InvMax                          *float64   `bun:"inv_max,type:decimal(19,9)"`                                // Used in calculating qty to order, different use depending on one of 4 replenishment methods.
	SafetyStock                     *float64   `bun:"safety_stock,type:decimal(19,9)"`                           // How many days of stock do we want to carry for safety?  Used in dynamic replenishment methods.
	Stockable                       *string    `bun:"stockable,type:char(1)"`                                    // Flag that determines if we carry stock of this item at this location.
	ReplenishmentLocation           *float64   `bun:"replenishment_location,type:decimal(19,0)"`                 // The item assigned purchasing method, which dictates when more needs to be ordered to replenish stock and exactly how much to order - four replenishment methods:  Order Point/Order Quantity, Min/Max, Up To, and Economic Order Quantity.
	MonthsInSeason                  *float64   `bun:"months_in_season,type:decimal(2,0)"`                        // Number of months this item is in season at this location.
	AverageMonthlyUsage             *float64   `bun:"average_monthly_usage,type:decimal(19,9)"`                  // Indicates the average monthly usage for this item at this location.
	ProductGroupId                  *string    `bun:"product_group_id,type:varchar(8)"`                          // A user-defined code that represents a group of products, usually with something in common.
	PurchaseClass                   *string    `bun:"purchase_class,type:varchar(8)"`                            // What is the purchase class?
	CycleCountingCategory           *float64   `bun:"cycle_counting_category,type:decimal(19,0)"`                // This column is no longer being used and is scheduled for removal in later version.
	PurchaseDiscountGroup           *string    `bun:"purchase_discount_group,type:varchar(8)"`                   // The purchase Discount Group ID that is automatically assigned when an item is added to a new location.
	SalesDiscountGroup              *string    `bun:"sales_discount_group,type:varchar(8)"`                      // The sales Discount Group ID that is automatically assigned when an item is added to a new location.
	NoCharge                        *string    `bun:"no_charge,type:char(1)"`                                    // An option on the Item Maintenance Location tab that indicates whether an item is given to customers at no charge.
	Price1                          *float64   `bun:"price1,type:decimal(19,9)"`                                 // A range of prices for an item, arranged in a series of ten columns.
	Price2                          *float64   `bun:"price2,type:decimal(19,9)"`                                 // A range of prices for an item, arranged in a series of ten columns.
	Price3                          *float64   `bun:"price3,type:decimal(19,9)"`                                 // A range of prices for an item, arranged in a series of ten columns.
	Price4                          *float64   `bun:"price4,type:decimal(19,9)"`                                 // A range of prices for an item, arranged in a series of ten columns.
	Price5                          *float64   `bun:"price5,type:decimal(19,9)"`                                 // A range of prices for an item, arranged in a series of ten columns.
	Price6                          *float64   `bun:"price6,type:decimal(19,9)"`                                 // A range of prices for an item, arranged in a series of ten columns.
	Price7                          *float64   `bun:"price7,type:decimal(19,9)"`                                 // A range of prices for an item, arranged in a series of ten columns.
	Price8                          *float64   `bun:"price8,type:decimal(19,9)"`                                 // A range of prices for an item, arranged in a series of ten columns.
	Price9                          *float64   `bun:"price9,type:decimal(19,9)"`                                 // A range of prices for an item, arranged in a series of ten columns.
	Price10                         *float64   `bun:"price10,type:decimal(19,9)"`                                // A range of prices for an item, arranged in a series of ten columns.
	ReplenishmentMethod             *string    `bun:"replenishment_method,type:varchar(8)"`                      // Either Minmax, opoq, upto or EOQ, way that GPOR determines how much stock to order.
	OrderQuantity                   *float64   `bun:"order_quantity,type:decimal(19,9),default:(0)"`             // The quantity on purchase order.
	QtyAllocated                    *float64   `bun:"qty_allocated,type:decimal(19,9),default:(0)"`              // The quantity of material that has been allocated to this inventory transfer.
	QtyBackordered                  *float64   `bun:"qty_backordered,type:decimal(19,9),default:(0)"`            // The quantity of the item that was backordered.
	QtyInTransit                    *float64   `bun:"qty_in_transit,type:decimal(19,9),default:(0)"`             // Qty of this item that is on its way to this location through transfer.
	TrackBins                       *string    `bun:"track_bins,type:char(1)"`                                   // Indicates if we track bins at this loc for this item.
	DefaultInOe                     *string    `bun:"default_in_oe,type:char(1)"`                                // Items with this checkbox checked are automatically entered as a line item on every order entered in the system
	PeriodFirstStocked              *float64   `bun:"period_first_stocked,type:decimal(3,0)"`                    // Identifies the period when this item was first stocked at this location.
	YearFirstStocked                *float64   `bun:"year_first_stocked,type:decimal(4,0)"`                      // Identifies the year when this item was first stocked at this location.
	UsageLock                       *string    `bun:"usage_lock,type:char(1)"`                                   // Indicates whether or not the forecast usage data maintained by the system is used to calculate Forecast Usage for an item.
	UsageLockPeriod                 *float64   `bun:"usage_lock_period,type:decimal(3,0)"`                       // The period a usage lock will end.
	UsageLockYear                   *float64   `bun:"usage_lock_year,type:decimal(4,0)"`                         // The year a usage lock will end.
	PrimaryBin                      *string    `bun:"primary_bin,type:varchar(10)"`                              // Indicates the bin where an item is usually found.
	TaxGroupId                      *string    `bun:"tax_group_id,type:varchar(10)"`                             // Indicates the tax group identification.
	QtyReservedDueIn                *float64   `bun:"qty_reserved_due_in,type:decimal(19,9),default:(0)"`        // The portion of qty_in_transit that is reserved.
	DateLastCounted                 *string    `bun:"date_last_counted,type:smalldatetime"`                      // Date the inventory at this location last counted.
	InvMastUid                      int32      `bun:"inv_mast_uid,type:int,pk"`                                  // Unique identifier for the item id.
	Requisition                     string     `bun:"requisition,type:char(1),default:('N')"`                    // Is the note area an available display area for a requistion note?
	DeadstockFlag                   string     `bun:"deadstock_flag,type:char(1),default:('N')"`                 // Indicates deadstock and can be available for sale in CommerceCenter using Trading Partner Connect.
	LotBinIntegration               *string    `bun:"lot_bin_integration,type:char(1)"`                          // Indicates that the item uses lot/bin tracking.
	AddToCycleCount                 string     `bun:"add_to_cycle_count,type:char(1),default:('N')"`             // Indicates that the item should appear in the next cycle count.
	DefaultShipment                 int32      `bun:"default_shipment,type:int,default:(1138)"`                  // Indicates when an other charge default item should appear with every shipment or with the first shipment only.
	PrimarySupplierId               *float64   `bun:"primary_supplier_id,type:decimal(19,0)"`                    // Denormalized from inventory_supplier_x_loc. Indicates the primary supplier for the location.
	OnObtFlag                       *string    `bun:"on_obt_flag,type:char(1),default:('N')"`                    // Indicates whether the item is currently on an order based transfer.
	OnReleaseScheduleFlag           *string    `bun:"on_release_schedule_flag,type:char(1),default:('N')"`       // Indicates whether the item is currently on a release schedule.
	OnBackorderFlag                 *string    `bun:"on_backorder_flag,type:char(1),default:('N')"`              // Indicates whether the item is currently on backorder.
	LastSaleDate                    time.Time  `bun:"last_sale_date,type:datetime,default:('01/01/1990')"`       // Last date order line is created for the item,location combination.
	LastPurchaseDate                *time.Time `bun:"last_purchase_date,type:datetime"`                          // Stores the date of the last Purchase Order for an item/location.  Po_line.date_created is used.
	Buy                             string     `bun:"buy,type:char(1),default:('Y')"`                            // indicate that item can be procured at the location by a purchase order or transfer.
	Make                            string     `bun:"make,type:char(1),default:('N')"`                           // indicate that item can be procured at the location by a production or MSP order.
	CreatedBy                       *string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	PeriodsToSupplyMin              *int16     `bun:"periods_to_supply_min,type:tinyint,default:(0)"`         // This is the minimum number of periods to supply when using variable control replenishment method.
	PeriodsToSupplyMax              *int16     `bun:"periods_to_supply_max,type:tinyint,default:(0)"`         // This is the maximum number of periods to supply. Used in calculations of variable control replenishment methods
	PutawayRank                     *string    `bun:"putaway_rank,type:varchar(8)"`                           // Putaway Rank corresponds to ABC class
	InvLastChangedDate              *time.Time `bun:"inv_last_changed_date,type:datetime"`                    // Indicates the date the item has had stock activity.
	IvaTaxableFlag                  *string    `bun:"iva_taxable_flag,type:char(1)"`                          // IVA taxable item indicator
	ItemPutawayAttributeUid         *int32     `bun:"item_putaway_attribute_uid,type:int"`                    // Putaway attribute ID for this item/location
	SplittableFlag                  *string    `bun:"splittable_flag,type:char(1)"`                           // Determines whether or not this item allows OE line splitting at this location.
	MinReplenishmentQty             *float64   `bun:"min_replenishment_qty,type:decimal(19,9)"`               // The minimum quantity that should be purchased when replenishing this item at this location.
	Discontinued                    string     `bun:"discontinued,type:char(1),default:('N')"`                // Flag to indicate if item is discontinued at a location.
	AllowDsDiscontinuedItems        string     `bun:"allow_ds_discontinued_items,type:char(1),default:('N')"` // Flag to indicate if item can have a DS disposition if item is discountinued.
	AllowSpDiscontinuedItems        string     `bun:"allow_sp_discontinued_items,type:char(1),default:('N')"` // Flag to indicate if item can have a SP disposition if item is discountinued.
	SafetyStockType                 *int32     `bun:"safety_stock_type,type:int"`                             // Method used to determine safety stock
	ServiceLevelMeasure             *int32     `bun:"service_level_measure,type:int"`                         // Customer service level (stock out back order)
	ServiceLevelPctGoal             *float64   `bun:"service_level_pct_goal,type:decimal(19,2)"`              // Percent of service level goal for safety stock
	DemandPatternCd                 *int32     `bun:"demand_pattern_cd,type:int"`                             // The demand pattern exhibited by historical usage for this item/location.  Used for improved forecasting in Advanced Demand Forecasting functionality.
	DemandPatternEvaluationDate     *time.Time `bun:"demand_pattern_evaluation_date,type:datetime"`           // The last time the demand pattern was evaluated by the system for this item/location.
	DemandForecastFormulaUid        *int32     `bun:"demand_forecast_formula_uid,type:int"`                   // Effective formula used for forecasting with advanced demand.
	PatternManuallyEditedFlag       *string    `bun:"pattern_manually_edited_flag,type:char(1)"`              // Used for advanced demand.  Flag identifying whether the user has manually chosen the demand pattern instead of using the system recommended demand pattern.
	DemandPatternBehaviorCd         *int32     `bun:"demand_pattern_behavior_cd,type:int"`                    // Used for advanced demand.  Code signifying whether item has a particular demand pattern or instead behaves like another item/location.
	PatternLikeInvMastUid           *int32     `bun:"pattern_like_inv_mast_uid,type:int"`                     // Used for advanced demand.  If the item/location behaves like another item/location, this is the target inv_mast_uid.
	PatternLikeLocationId           *float64   `bun:"pattern_like_location_id,type:decimal(19,0)"`            // Used for advanced demand.  If the item/location behaves like another item/location, this is the target location_id.
	BehavesLikeLockFlag             *string    `bun:"behaves_like_lock_flag,type:char(1)"`                    // Flags whether to lock item on behaves like so it is not considered in demand pattern identification.
	BehavesLikeLockPeriod           *int32     `bun:"behaves_like_lock_period,type:int"`                      // Expiration period for behaves like lock.
	BehavesLikeLockYear             *int32     `bun:"behaves_like_lock_year,type:int"`                        // Expiration year for behaves like lock.
	LandedCostAccountNo             *string    `bun:"landed_cost_account_no,type:varchar(255)"`               // Landed Cost Income Account.
	PriceFamilyUid                  *int32     `bun:"price_family_uid,type:int"`                              // The default price_family for this item at this location - unique identifier for price_family table.
	DeleteFlag                      *string    `bun:"delete_flag,type:char(1)"`                               // Indicates if record is marked as deleted.
	CardlockProductId               *string    `bun:"cardlock_product_id,type:varchar(255)"`                  // Product id within Cardlock system.
	DefaultSellingUnit              *string    `bun:"default_selling_unit,type:varchar(8)"`                   // Indicates default sales / order unit of measure for this location
	LocItemTypeCd                   *int32     `bun:"loc_item_type_cd,type:int"`                              // The Item Type at this location
	LocCustParentInvMastUid         *int32     `bun:"loc_cust_parent_inv_mast_uid,type:int"`                  // The Parent Item of a fillable item at this location
	MainBulkLocationFlag            *string    `bun:"main_bulk_location_flag,type:char(1)"`                   // Flag to determine if the location is a main bulk location
	DfltSourceLocFlag               *string    `bun:"dflt_source_loc_flag,type:char(1)"`                      // Custom column to indicate the location id is the item's default source location
	CycleCountFlag                  *string    `bun:"cycle_count_flag,type:char(1)"`                          // Determines if the cycle count should pick up this row when the criteria window specifies it.
	PromotionalFlag                 *string    `bun:"promotional_flag,type:char(1)"`                          // Custom column to indicate the item is promotional per location
	RmaRevenueAccountNo             *string    `bun:"rma_revenue_account_no,type:varchar(32)"`                // Default value for rma account
	FutureStandardCost              *float64   `bun:"future_standard_cost,type:decimal(19,9)"`                // This column indicates the future standard cost for this item at this location effective from the effective date
	EffectiveDate                   *time.Time `bun:"effective_date,type:datetime"`                           // This column indicates the effective date from which the future standard cost would be effective
	AssemblyPromptOption            *int32     `bun:"assembly_prompt_option,type:int"`                        // Used to indicate whether to build, buy, or prompt the user when processing an item at the location.
	RestrictedFlag                  *string    `bun:"restricted_flag,type:char(1)"`                           // Indicates whether visibility of an item is restricted within the application for a location.
	DrpItemFlag                     string     `bun:"drp_item_flag,type:char(1),default:('N')"`
	SavedDemandForecastFormulaUid   *int32     `bun:"saved_demand_forecast_formula_uid,type:int"` // The demand formula will be stored here temporarily in cases when forecasting must temporarily overwrite the demand formula for an item during calculation
	DemandFormulaComputedYearPeriod *int32     `bun:"demand_formula_computed_year_period,type:int"`
	DemandPatternComputedYearPeriod *int32     `bun:"demand_pattern_computed_year_period,type:int"`
	YearPeriodLastForecast          *int32     `bun:"year_period_last_forecast,type:int"`
	VendorRebateAccountNo           *string    `bun:"vendor_rebate_account_no,type:varchar(32)"`              // A general ledger account that tracks the value of Vendor Rebates
	TransferUsagePercent            *float64   `bun:"transfer_usage_percent,type:decimal(19,9)"`              // (Custom functionality) Define the precentage of quantity that should be recorded as usage for transfers of the item from this location
	ManufacturerReqStockFlag        string     `bun:"manufacturer_req_stock_flag,type:char(1),default:('N')"` // Checkbox on the Replenishment tab in Item Maintenance
	FcBin                           *string    `bun:"fc_bin,type:varchar(10)"`                                // Front counter bin location
	PumpoffItem                     *string    `bun:"pumpoff_item,type:varchar(40)"`                          // Reference to a pumpoff item code.
	PumpoffUom                      *string    `bun:"pumpoff_uom,type:varchar(8)"`                            // The UOM used by a pumpoff item.
	OnFutureOrderFlag               *string    `bun:"on_future_order_flag,type:char(1),default:('N')"`        // Tracks whether this item location combination has an open future order line (F disposition) as used in the Alternate OE Allocation functionality,
	OnFutureProdOrderFlag           *string    `bun:"on_future_prod_order_flag,type:char(1),default:('N')"`   // Tracks whether this item location combination has an open future production order line (F disposition).
	MaxLiability                    *float64   `bun:"max_liability,type:decimal(19,9)"`                       // Maximum liability for the item at the location.
	MaxTransferQty                  *float64   `bun:"max_transfer_qty,type:decimal(19,9)"`                    // Maximum transfer quantity for the item at the location.
	AltTaxGroupId                   *string    `bun:"alt_tax_group_id,type:varchar(10)"`
	ReturnableFlag                  *string    `bun:"returnable_flag,type:char(1),default:('Y')"`    // Custom column to indicate items that are not returnable at location level
	IqsItemFlag                     *string    `bun:"iqs_item_flag,type:char(1)"`                    // Used to identify that the item requires IQS processing on a specific location.
	SlabTrackBinsFlag               *string    `bun:"slab_track_bins_flag,type:char(1)"`             // Slab Track Bins Flag to specify that a slab item track bins.
	Edi832DiscontinuedSentFlag      *string    `bun:"edi_832_discontinued_sent_flag,type:char(1)"`   // Custom (F63577): determines if a discontinued item has been included in an EDI 832 price list export.
	BuyerId                         *string    `bun:"buyer_id,type:varchar(16)"`                     // Custom: Buyer ID associated with purchasing this item at this location.
	TaRentalRevenueAccountNo        *string    `bun:"ta_rental_revenue_account_no,type:varchar(32)"` // TrackAbout Rental Revenue Account
	CycleCountNo                    *float64   `bun:"cycle_count_no,type:decimal(19,0)"`             // Cycle count number last edited
	CycleCountQty                   *float64   `bun:"cycle_count_qty,type:decimal(19,9)"`            // Quantity adjusted in the last cycle count
	TaRentalTaxGroupId              *string    `bun:"ta_rental_tax_group_id,type:varchar(10)"`       // TrackAbout Rental Invoice Tax Group ID
	TransferConversionUom           *string    `bun:"transfer_conversion_uom,type:varchar(8)"`       // (Custom F69980) Indicates the UOM that transfer quantity should be rounded (converted) to when this location is the destination location.
	EccEnabledFlag                  string     `bun:"ecc_enabled_flag,type:char(1),default:('N')"`   // Enable Item for ECC
	InvMastUidDupUsage              *int32     `bun:"inv_mast_uid_dup_usage,type:int"`               // Link to inv_mast record for duplicate item usage per location
	TransferOrderPoint              *float64   `bun:"transfer_order_point,type:decimal(19,9)"`       // Custom: Indicates order point used for transfer generation.
	TransferOrderQty                *float64   `bun:"transfer_order_qty,type:decimal(19,9)"`         // Custom: Indicates order quantity for transfer generation.
	OsmiItemFlag                    *string    `bun:"osmi_item_flag,type:char(1),default:('N')"`     // Osmi Item
	LocalItemFlag                   *string    `bun:"local_item_flag,type:char(1),default:('N')"`    // Local Item
	CriticalItemFlag                *string    `bun:"critical_item_flag,type:char(1)"`               // Flag whether, for purposes of safety stock calculations, this item/location should be treated as a critical item/location.  Critical item/locations can use different safety stock multipliers.
	DeviationLookbackPds            *int32     `bun:"deviation_lookback_pds,type:int"`               // When calculating safety stock by Deviation Multiplier, the number of periods to look backwards to determine the deviation.
	DeviationMultOnePd              *float64   `bun:"deviation_mult_one_pd,type:decimal(19,9)"`      // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when one of the lookback periods has forecast usage less than actual usage.
	DeviationMultTwoPd              *float64   `bun:"deviation_mult_two_pd,type:decimal(19,9)"`      // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when two of the lookback periods has forecast usage less than actual usage.
	DeviationMultThreePd            *float64   `bun:"deviation_mult_three_pd,type:decimal(19,9)"`    // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when three or more of the lookback periods has forecast usage less than actual usage.
	MinSafetyStockDays              *float64   `bun:"min_safety_stock_days,type:decimal(19,9)"`      // Minimum fence of safety stock days to use which overrides safety stock days calculation.
	MaxSafetyStockDays              *float64   `bun:"max_safety_stock_days,type:decimal(19,9)"`      // Maximum fence of safety stock days to use which overrides safety stock days calculation.
	CritItemDeviationMult           *float64   `bun:"crit_item_deviation_mult,type:decimal(19,9)"`   // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when the item/location is flagged as a critical item/location.
	CritMinSafetyStockDays          *float64   `bun:"crit_min_safety_stock_days,type:decimal(19,9)"` // Minimum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.
	CritMaxSafetyStockDays          *float64   `bun:"crit_max_safety_stock_days,type:decimal(19,9)"` // Maximum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.
}
