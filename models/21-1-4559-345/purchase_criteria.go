package model

type PurchaseCriteria struct {
	bun.BaseModel              `bun:"table:purchase_criteria"`
	PoCriteriaId               float64   `bun:"po_criteria_id,type:decimal(19,0),pk"`
	Description                string    `bun:"description,type:varchar(40)"`
	BegSupplierId              float64   `bun:"beg_supplier_id,type:decimal(19,0),nullzero"`
	EndSupplierId              float64   `bun:"end_supplier_id,type:decimal(19,0),nullzero"`
	BegAbcClassId              string    `bun:"beg_abc_class_id,type:varchar(8),nullzero"`
	EndAbcClassId              string    `bun:"end_abc_class_id,type:varchar(8),nullzero"`
	BegPurchaseDiscountGroupId string    `bun:"beg_purchase_discount_group_id,type:varchar(8),nullzero"`
	EndPurchaseDiscountGroupId string    `bun:"end_purchase_discount_group_id,type:varchar(8),nullzero"`
	BegProductGroupId          string    `bun:"beg_product_group_id,type:varchar(8),nullzero"`
	EndProductGroupId          string    `bun:"end_product_group_id,type:varchar(8),nullzero"`
	LocationId                 float64   `bun:"location_id,type:decimal(19,0),nullzero"`
	UsageFactor                float64   `bun:"usage_factor,type:decimal(19,4),nullzero"`
	MinimumDollarsPerItem      float64   `bun:"minimum_dollars_per_item,type:decimal(19,4),nullzero"`
	CarryingCost               float64   `bun:"carrying_cost,type:decimal(19,9),nullzero"`
	OrderPointException        float64   `bun:"order_point_exception,type:decimal(19,9),nullzero"`
	DiscountPercent            float64   `bun:"discount_percent,type:decimal(19,4),nullzero"`
	Roai                       float64   `bun:"roai,type:decimal(19,9),nullzero"`
	NumberOfMonths             float64   `bun:"number_of_months,type:decimal(19,0),nullzero"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	ReplenishmentMethod        string    `bun:"replenishment_method,type:varchar(8),nullzero"`
	PurchaseGroupId            string    `bun:"purchase_group_id,type:varchar(8),nullzero"`
	SafetyStockDays            float64   `bun:"safety_stock_days,type:decimal(19,4),nullzero"`
	LookAheadDays              float64   `bun:"look_ahead_days,type:decimal(19,0),nullzero"`
	PurchaseBy                 string    `bun:"purchase_by,type:char,nullzero"`
	RequestedBy                string    `bun:"requested_by,type:varchar(20),nullzero"`
	IncludeUnapprovedPos       string    `bun:"include_unapproved_pos,type:char"`
	CompanyId                  string    `bun:"company_id,type:varchar(8),nullzero"`
	DeleteFlag                 string    `bun:"delete_flag,type:char"`
	IncludeProductionOrders    string    `bun:"include_production_orders,type:char,nullzero"`
	UseSurplusQty              string    `bun:"use_surplus_qty,type:char,nullzero"`
	DefaultGroupPurchasingType int32     `bun:"default_group_purchasing_type,type:int,nullzero"`
	ExcludeTboQuantity         string    `bun:"exclude_tbo_quantity,type:char,nullzero"`
	ReplenishmentLocationOnly  string    `bun:"replenishment_location_only,type:char,default:('N')"`
	AutogenerateTbo            string    `bun:"autogenerate_tbo,type:char,default:('N')"`
	BuyerSpecificPurchasing    string    `bun:"buyer_specific_purchasing,type:char,default:('N')"`
	RequirementTypeStock       `bun:"requirement_type_stock,type:bit,default:(0)"`
	RequirementTypeNonStock    `bun:"requirement_type_non_stock,type:bit,default:(0)"`
	RequirementTypeSpecial     `bun:"requirement_type_special,type:bit,default:(0)"`
	RequirementTypeRequisition `bun:"requirement_type_requisition,type:bit,default:(0)"`
	RequirementTypeDirectShip  `bun:"requirement_type_direct_ship,type:bit,default:(0)"`
	RequirementTypeBackOrder   `bun:"requirement_type_back_order,type:bit,default:(0)"`
	UseNsAsSourceForTransfer   string  `bun:"use_ns_as_source_for_transfer,type:char,default:('N')"`
	ShowAllItems               string  `bun:"show_all_items,type:char,default:('N')"`
	RdcDetailFlag              string  `bun:"rdc_detail_flag,type:char,default:('S')"`
	InvMastUid                 int32   `bun:"inv_mast_uid,type:int,nullzero"`
	EvaluateRoaiMethod         int32   `bun:"evaluate_roai_method,type:int,default:((1703))"`
	RoaiCutOffMeasure          int32   `bun:"roai_cut_off_measure,type:int,default:((300))"`
	RoaiCutOffPeriods          float64 `bun:"roai_cut_off_periods,type:decimal(19,0),default:((0))"`
	RoaiCutOffDollars          float64 `bun:"roai_cut_off_dollars,type:decimal(19,0),default:((0))"`
	SafetyStockType            int32   `bun:"safety_stock_type,type:int,default:((1261))"`
	ServiceLevelMeasure        int32   `bun:"service_level_measure,type:int,nullzero"`
	ServiceLevelPctGoal        float64 `bun:"service_level_pct_goal,type:decimal(19,2),nullzero"`
	CombineSpWithStockFlag     string  `bun:"combine_sp_with_stock_flag,type:char,default:('N')"`
	BegClassId4                string  `bun:"beg_class_id4,type:varchar(255),nullzero"`
	EndClassId4                string  `bun:"end_class_id4,type:varchar(255),nullzero"`
	UseSupplierDfltPrintFlag   string  `bun:"use_supplier_dflt_print_flag,type:char,nullzero"`
	SupplierIdList             string  `bun:"supplier_id_list,type:text(2147483647),nullzero"`
	SupplierOptionCd           int32   `bun:"supplier_option_cd,type:int,nullzero"`
	BegItemId                  string  `bun:"beg_item_id,type:varchar(40),nullzero"`
	EndItemId                  string  `bun:"end_item_id,type:varchar(40),nullzero"`
	ProductGroupOptionCd       int32   `bun:"product_group_option_cd,type:int,nullzero"`
	ProductGroupIdList         string  `bun:"product_group_id_list,type:text(2147483647),nullzero"`
	SupplierBuyerId            string  `bun:"supplier_buyer_id,type:varchar(16),nullzero"`
	LookAheadDaysCd            int32   `bun:"look_ahead_days_cd,type:int,default:((3045))"`
	BegRouteCodeId             string  `bun:"beg_route_code_id,type:varchar(255),nullzero"`
	EndRouteCodeId             string  `bun:"end_route_code_id,type:varchar(255),nullzero"`
	BegProdOrder               float64 `bun:"beg_prod_order,type:decimal(19,0),default:((0))"`
	EndProdOrder               float64 `bun:"end_prod_order,type:decimal(19,0),default:((99999999))"`
	ProdOrderOptionCd          int32   `bun:"prod_order_option_cd,type:int,default:((1694))"`
	ProdOrderList              string  `bun:"prod_order_list,type:text(2147483647),nullzero"`
	DeviationLookbackPds       int32   `bun:"deviation_lookback_pds,type:int,nullzero"`
	DeviationMultOnePd         float64 `bun:"deviation_mult_one_pd,type:decimal(19,9),nullzero"`
	DeviationMultTwoPd         float64 `bun:"deviation_mult_two_pd,type:decimal(19,9),nullzero"`
	DeviationMultThreePd       float64 `bun:"deviation_mult_three_pd,type:decimal(19,9),nullzero"`
	MinSafetyStockDays         float64 `bun:"min_safety_stock_days,type:decimal(19,9),nullzero"`
	MaxSafetyStockDays         float64 `bun:"max_safety_stock_days,type:decimal(19,9),nullzero"`
	CritItemDeviationMult      float64 `bun:"crit_item_deviation_mult,type:decimal(19,9),nullzero"`
	CritMinSafetyStockDays     float64 `bun:"crit_min_safety_stock_days,type:decimal(19,9),nullzero"`
	CritMaxSafetyStockDays     float64 `bun:"crit_max_safety_stock_days,type:decimal(19,9),nullzero"`
	SupplierGroupId            string  `bun:"supplier_group_id,type:varchar(255),nullzero"`
	ClassId4OptionCd           int32   `bun:"class_id4_option_cd,type:int,nullzero"`
	ClassId4List               string  `bun:"class_id4_list,type:varchar,nullzero"`
}
