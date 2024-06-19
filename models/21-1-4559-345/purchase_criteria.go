package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchaseCriteria struct {
	bun.BaseModel              `bun:"table:purchase_criteria"`
	PoCriteriaId               float64   `bun:"po_criteria_id,type:decimal(19,0),pk"`                          // Criteria ID seen by user when a criteria is saved.
	Description                string    `bun:"description,type:varchar(40)"`                                  // Description of the GPOR criteria set.
	BegSupplierId              float64   `bun:"beg_supplier_id,type:decimal(19,0),nullzero"`                   // Used when giving a range of supplier IDs, GPOR will only get requirements for the range.
	EndSupplierId              float64   `bun:"end_supplier_id,type:decimal(19,0),nullzero"`                   // Used when giving a range of supplier IDs, GPOR will only get requirements for the range.
	BegAbcClassId              string    `bun:"beg_abc_class_id,type:varchar(8),nullzero"`                     // Used when giving a range of ABC Class IDs, GPOR will only get requirements for the range.
	EndAbcClassId              string    `bun:"end_abc_class_id,type:varchar(8),nullzero"`                     // Used when giving a range of ABC Class IDs, GPOR will only get requirements for the range.
	BegPurchaseDiscountGroupId string    `bun:"beg_purchase_discount_group_id,type:varchar(8),nullzero"`       // Not used.
	EndPurchaseDiscountGroupId string    `bun:"end_purchase_discount_group_id,type:varchar(8),nullzero"`       // Not Used.
	BegProductGroupId          string    `bun:"beg_product_group_id,type:varchar(8),nullzero"`                 // Used when giving a range of Product Group IDs, GPOR will only get requirements for the range.
	EndProductGroupId          string    `bun:"end_product_group_id,type:varchar(8),nullzero"`                 // Used when giving a range of Product Group IDs, GPOR will only get requirements for the range.
	LocationId                 float64   `bun:"location_id,type:decimal(19,0),nullzero"`                       // Main location ID where you want the requirements.
	UsageFactor                float64   `bun:"usage_factor,type:decimal(19,4),nullzero"`                      // Used in some replenishment methods, calculating if a requirement exists for an item.
	MinimumDollarsPerItem      float64   `bun:"minimum_dollars_per_item,type:decimal(19,4),nullzero"`          // Used in some replenishment methods, calculating if a requirement exists for an item.
	CarryingCost               float64   `bun:"carrying_cost,type:decimal(19,9),nullzero"`                     // Used in some replenishment methods, calculating if a requirement exists for an item.
	OrderPointException        float64   `bun:"order_point_exception,type:decimal(19,9),nullzero"`             // Used in some replenishment methods, calculating if a requirement exists for an item.
	DiscountPercent            float64   `bun:"discount_percent,type:decimal(19,4),nullzero"`                  // Used in some replenishment methods, calculating how much needs to be purchased for an item.
	Roai                       float64   `bun:"roai,type:decimal(19,9),nullzero"`                              // Used in some replenishment methods, calculating how much needs to be purchased for an item.
	NumberOfMonths             float64   `bun:"number_of_months,type:decimal(19,0),nullzero"`                  // Used in some replenishment methods, calculating how much needs to be purchased for an item.
	DateCreated                time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	ReplenishmentMethod        string    `bun:"replenishment_method,type:varchar(8),nullzero"`                 // Populate if you only want items that match a replenishment method.
	PurchaseGroupId            string    `bun:"purchase_group_id,type:varchar(8),nullzero"`                    // The Purchase group,used to group locations together for group purchasing.
	SafetyStockDays            float64   `bun:"safety_stock_days,type:decimal(19,4),nullzero"`                 // Used in some replenishment methods, calculating if a requirement exists for an item.
	LookAheadDays              float64   `bun:"look_ahead_days,type:decimal(19,0),nullzero"`                   // How long into the future are we purchasing for?
	PurchaseBy                 string    `bun:"purchase_by,type:char(1),nullzero"`                             // Purchasing for one location, or a group?
	RequestedBy                string    `bun:"requested_by,type:varchar(20),nullzero"`                        // Buyer ID for the PO that will be generated.  Sometimes used to determine requirements.
	IncludeUnapprovedPos       string    `bun:"include_unapproved_pos,type:char(1)"`                           // determines whether previously generated unapproved POs are included in an automatic purchasing session.
	CompanyId                  string    `bun:"company_id,type:varchar(8),nullzero"`                           // Unique code that identifies a company.
	DeleteFlag                 string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	IncludeProductionOrders    string    `bun:"include_production_orders,type:char(1),nullzero"`               // Include requirements generated from production orders?
	UseSurplusQty              string    `bun:"use_surplus_qty,type:char(1),nullzero"`                         // Allows users the option to include surplus qty from other locations within the group.
	DefaultGroupPurchasingType int32     `bun:"default_group_purchasing_type,type:int,nullzero"`               // If group purchasing, by location or RDC?
	ExcludeTboQuantity         string    `bun:"exclude_tbo_quantity,type:char(1),nullzero"`                    // override for a purchasing setting in System Settings that determines how Transfer Backorder quantities factor into generating group purchasing quantities.
	ReplenishmentLocationOnly  string    `bun:"replenishment_location_only,type:char(1),default:('N')"`        // Indicates whether items will be purchased only when the replenishment location is either the purchasing location, or in the purchasing location group.
	AutogenerateTbo            string    `bun:"autogenerate_tbo,type:char(1),default:('N')"`
	BuyerSpecificPurchasing    string    `bun:"buyer_specific_purchasing,type:char(1),default:('N')"` // Indicates whether the Buyer Specific Purchasing feature should be used
	RequirementTypeStock       bool      `bun:"requirement_type_stock,type:bit,default:(0)"`          // Specifies if the critieria generation should include stock items
	RequirementTypeNonStock    bool      `bun:"requirement_type_non_stock,type:bit,default:(0)"`      // Specifies if the criteria generation should include non stock items
	RequirementTypeSpecial     bool      `bun:"requirement_type_special,type:bit,default:(0)"`        // Specifies if the criteria generation should include items on special orders
	RequirementTypeRequisition bool      `bun:"requirement_type_requisition,type:bit,default:(0)"`    // Specifies if the criteria generation should include requisition items
	RequirementTypeDirectShip  bool      `bun:"requirement_type_direct_ship,type:bit,default:(0)"`    // Specifies if the criteria generation should include items on direct ship order
	RequirementTypeBackOrder   bool      `bun:"requirement_type_back_order,type:bit,default:(0)"`     // Specifies if the criteria generation should include items on back orders
	UseNsAsSourceForTransfer   string    `bun:"use_ns_as_source_for_transfer,type:char(1),default:('N')"`
	ShowAllItems               string    `bun:"show_all_items,type:char(1),default:('N')"`
	RdcDetailFlag              string    `bun:"rdc_detail_flag,type:char(1),default:('S')"`
	InvMastUid                 int32     `bun:"inv_mast_uid,type:int,nullzero"`                         // Foreign Key to inv_mast table
	EvaluateRoaiMethod         int32     `bun:"evaluate_roai_method,type:int,default:((1703))"`         // ROAI option to use in GPOR.  Options include None (300), Buy Ahead (1703), and Next Break (1702).
	RoaiCutOffMeasure          int32     `bun:"roai_cut_off_measure,type:int,default:((300))"`          // Provides upper limit of additional stock to suggest for ROAI Next Break.
	RoaiCutOffPeriods          float64   `bun:"roai_cut_off_periods,type:decimal(19,0),default:((0))"`  // Upper limit in periods of additional stock to suggest for ROAI Next Break.
	RoaiCutOffDollars          float64   `bun:"roai_cut_off_dollars,type:decimal(19,0),default:((0))"`  // Upper limit in dollars of additional stock to suggest for ROAI Next Break.
	SafetyStockType            int32     `bun:"safety_stock_type,type:int,default:((1261))"`            // Method used to determine safety stock
	ServiceLevelMeasure        int32     `bun:"service_level_measure,type:int,nullzero"`                // Customer service level (stock out back order)
	ServiceLevelPctGoal        float64   `bun:"service_level_pct_goal,type:decimal(19,2),nullzero"`     // Percent of service level goal for safety stock
	CombineSpWithStockFlag     string    `bun:"combine_sp_with_stock_flag,type:char(1),default:('N')"`  // Option to combine specials with regular stock/non-stocks on POs.
	BegClassId4                string    `bun:"beg_class_id4,type:varchar(255),nullzero"`               // This column will hold the beginning Class ID 4 criteria for PO Requirements
	EndClassId4                string    `bun:"end_class_id4,type:varchar(255),nullzero"`               // This column will hold the ending Class ID 4 criteria for PO Requirements
	UseSupplierDfltPrintFlag   string    `bun:"use_supplier_dflt_print_flag,type:char(1),nullzero"`     // Indicates whether POs generated using the criteria should be sent using the supplier's default delivery method (Print/Fax/Email/EDI.)
	SupplierIdList             string    `bun:"supplier_id_list,type:text,nullzero"`                    // List of distinct supplier IDs that will be used for requirements
	SupplierOptionCd           int32     `bun:"supplier_option_cd,type:int,nullzero"`                   // Determines if the criteria will use a Supplier Range or a list of Supplier IDs to generate requirements.
	BegItemId                  string    `bun:"beg_item_id,type:varchar(40),nullzero"`                  // Starting item ID when running GPOR with a range of items.
	EndItemId                  string    `bun:"end_item_id,type:varchar(40),nullzero"`                  // Used when giving a range of supplier IDs, GPOR will only get requirements for the range.
	ProductGroupOptionCd       int32     `bun:"product_group_option_cd,type:int,nullzero"`              // Determines if the criteria will use a range or list of Product Group IDs to generate requirements.
	ProductGroupIdList         string    `bun:"product_group_id_list,type:text,nullzero"`               // The list of product group IDs to be used for to generate requirements
	SupplierBuyerId            string    `bun:"supplier_buyer_id,type:varchar(16),nullzero"`            // When Buyer Specific Purchasing is used, this is the Buyer ID whose Suppliers will be used.
	LookAheadDaysCd            int32     `bun:"look_ahead_days_cd,type:int,default:((3045))"`           // Code indicating static/dynamic look ahead
	BegRouteCodeId             string    `bun:"beg_route_code_id,type:varchar(255),nullzero"`           // Starting shipping route as PORG criteria with specials or DS selected
	EndRouteCodeId             string    `bun:"end_route_code_id,type:varchar(255),nullzero"`           // Ending shipping route as PORG criteria with specials or DS selected
	BegProdOrder               float64   `bun:"beg_prod_order,type:decimal(19,0),default:((0))"`        // Beginning Prod Order number for limiting GPOR items
	EndProdOrder               float64   `bun:"end_prod_order,type:decimal(19,0),default:((99999999))"` // End Prod Order number for limiting GPOR items
	ProdOrderOptionCd          int32     `bun:"prod_order_option_cd,type:int,default:((1694))"`         // Determine how you want to limit by Prod Orders, Range or List
	ProdOrderList              string    `bun:"prod_order_list,type:text,nullzero"`                     // List of Prod Order numbers for limiting GPOR items
	DeviationLookbackPds       int32     `bun:"deviation_lookback_pds,type:int,nullzero"`               // When calculating safety stock by Deviation Multiplier, the number of periods to look backwards to determine the deviation.
	DeviationMultOnePd         float64   `bun:"deviation_mult_one_pd,type:decimal(19,9),nullzero"`      // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when one of the lookback periods has forecast usage less than actual usage.
	DeviationMultTwoPd         float64   `bun:"deviation_mult_two_pd,type:decimal(19,9),nullzero"`      // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when two of the lookback periods has forecast usage less than actual usage.
	DeviationMultThreePd       float64   `bun:"deviation_mult_three_pd,type:decimal(19,9),nullzero"`    // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when three or more of the lookback periods has forecast usage less than actual usage.
	MinSafetyStockDays         float64   `bun:"min_safety_stock_days,type:decimal(19,9),nullzero"`      // Minimum fence of safety stock days to use which overrides safety stock days calculation.
	MaxSafetyStockDays         float64   `bun:"max_safety_stock_days,type:decimal(19,9),nullzero"`      // Maximum fence of safety stock days to use which overrides safety stock days calculation.
	CritItemDeviationMult      float64   `bun:"crit_item_deviation_mult,type:decimal(19,9),nullzero"`   // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when the item/location is flagged as a critical item/location.
	CritMinSafetyStockDays     float64   `bun:"crit_min_safety_stock_days,type:decimal(19,9),nullzero"` // Minimum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.
	CritMaxSafetyStockDays     float64   `bun:"crit_max_safety_stock_days,type:decimal(19,9),nullzero"` // Maximum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.
	SupplierGroupId            string    `bun:"supplier_group_id,type:varchar(255),nullzero"`           // Used when giving a group of supplier IDs, GPOR will only get requirements for the suppliers within the group.
	ClassId4OptionCd           int32     `bun:"class_id4_option_cd,type:int,nullzero"`                  // Determines if the criteria will use item class id4 to generate requirements.
	ClassId4List               string    `bun:"class_id4_list,type:varchar(max),nullzero"`              // List of distinct item class id4s that will be used for requirements
}
