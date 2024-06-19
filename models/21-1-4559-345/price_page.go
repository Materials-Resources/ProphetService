package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PricePage struct {
	bun.BaseModel              `bun:"table:price_page"`
	PricePageUid               int32     `bun:"price_page_uid,type:int,pk"`                             // Internal record identifier
	PricePageTypeCd            int16     `bun:"price_page_type_cd,type:smallint"`                       // Indicates the type of the Price Page (Item ID, Supplier ID, etc.)
	MfgClassId                 string    `bun:"mfg_class_id,type:varchar(255),nullzero"`                // Manufacturing Class ID (used for Supplier/Mfg Class type pages only)
	SupplierId                 float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`                // Supplier ID (used for Supplier, Supplier/Discount Group, and Supplier/Mfg Class type pages only)
	ProductGroupId             string    `bun:"product_group_id,type:varchar(8),nullzero"`              // Product Group ID (used for Product Group type pages only)
	DiscountGroupId            string    `bun:"discount_group_id,type:varchar(8),nullzero"`             // Discount Group ID (used for Discount Group and Supplier/Discount Group type pages only)
	MajorGroupId               string    `bun:"major_group_id,type:varchar(8),nullzero"`                // Major Group ID (used only when totaling method is Major Group)
	Description                string    `bun:"description,type:varchar(255),nullzero"`                 // User-defined description of the pricing page
	PricingMethodCd            int16     `bun:"pricing_method_cd,type:smallint"`                        // Pricing Method (Source, Price, None)
	SourcePriceCd              int16     `bun:"source_price_cd,type:smallint,nullzero"`                 // Stores the Source Price when Pricing Method is set to Source
	Price                      float64   `bun:"price,type:decimal(19,9),nullzero"`                      // Stores the Price when Pricing Method is set to Price
	EffectiveDate              time.Time `bun:"effective_date,type:datetime"`                           // Date this pricing page takes effect
	ExpirationDate             time.Time `bun:"expiration_date,type:datetime"`                          // Date after which this pricing page is no longer in effect
	ContractNumber             string    `bun:"contract_number,type:varchar(40),nullzero"`              // User-defined contract number
	CalculationMethodCd        int16     `bun:"calculation_method_cd,type:smallint"`                    // Stores calculation method (multiplier, difference, etc.) used for the price
	CalculationValue1          float64   `bun:"calculation_value1,type:decimal(19,6)"`                  // Calculation value for break 1
	CalculationValue2          float64   `bun:"calculation_value2,type:decimal(19,6)"`                  // Calculation value for break 2
	CalculationValue3          float64   `bun:"calculation_value3,type:decimal(19,6)"`                  // Calculation value for break 3
	CalculationValue4          float64   `bun:"calculation_value4,type:decimal(19,6)"`                  // Calculation value for break 4
	CalculationValue5          float64   `bun:"calculation_value5,type:decimal(19,6)"`                  // Calculation value for break 5
	CalculationValue6          float64   `bun:"calculation_value6,type:decimal(19,6)"`                  // Calculation value for break 6
	CalculationValue7          float64   `bun:"calculation_value7,type:decimal(19,6)"`                  // Calculation value for break 7
	CalculationValue8          float64   `bun:"calculation_value8,type:decimal(19,6)"`                  // Calculation value for break 8
	CalculationValue9          float64   `bun:"calculation_value9,type:decimal(19,6)"`                  // Calculation value for break 9
	CalculationValue10         float64   `bun:"calculation_value10,type:decimal(19,6)"`                 // Calculation value for break 10
	CalculationValue11         float64   `bun:"calculation_value11,type:decimal(19,6)"`                 // Calculation value for break 11
	CalculationValue12         float64   `bun:"calculation_value12,type:decimal(19,6)"`                 // Calculation value for break 12
	CalculationValue13         float64   `bun:"calculation_value13,type:decimal(19,6)"`                 // Calculation value for break 13
	CalculationValue14         float64   `bun:"calculation_value14,type:decimal(19,6)"`                 // Calculation value for break 14
	CalculationValue15         float64   `bun:"calculation_value15,type:decimal(19,6)"`                 // Calculation value for break 15
	Break1                     float64   `bun:"break1,type:decimal(19,9)"`                              // Quantity break 1
	Break2                     float64   `bun:"break2,type:decimal(19,9)"`                              // Quantity break 2
	Break3                     float64   `bun:"break3,type:decimal(19,9)"`                              // Quantity break 3
	Break4                     float64   `bun:"break4,type:decimal(19,9)"`                              // Quantity break 4
	Break5                     float64   `bun:"break5,type:decimal(19,9)"`                              // Quantity break 5
	Break6                     float64   `bun:"break6,type:decimal(19,9)"`                              // Quantity break 6
	Break7                     float64   `bun:"break7,type:decimal(19,9)"`                              // Quantity break 7
	Break8                     float64   `bun:"break8,type:decimal(19,9)"`                              // Quantity break 8
	Break9                     float64   `bun:"break9,type:decimal(19,9)"`                              // Quantity break 9
	Break10                    float64   `bun:"break10,type:decimal(19,9)"`                             // Quantity break 10
	Break11                    float64   `bun:"break11,type:decimal(19,9)"`                             // Quantity break 11
	Break12                    float64   `bun:"break12,type:decimal(19,9)"`                             // Quantity break 12
	Break13                    float64   `bun:"break13,type:decimal(19,9)"`                             // Quantity break 13
	Break14                    float64   `bun:"break14,type:decimal(19,9)"`                             // Quantity break 14
	OtherCost1                 float64   `bun:"other_cost1,type:decimal(19,9),nullzero"`                // Other Cost value for break 1
	OtherCost2                 float64   `bun:"other_cost2,type:decimal(19,9),nullzero"`                // Other Cost value for break 2
	OtherCost3                 float64   `bun:"other_cost3,type:decimal(19,9),nullzero"`                // Other Cost value for break 3
	OtherCost4                 float64   `bun:"other_cost4,type:decimal(19,9),nullzero"`                // Other Cost value for break 4
	OtherCost5                 float64   `bun:"other_cost5,type:decimal(19,9),nullzero"`                // Other Cost value for break 5
	OtherCost6                 float64   `bun:"other_cost6,type:decimal(19,9),nullzero"`                // Other Cost value for break 6
	OtherCost7                 float64   `bun:"other_cost7,type:decimal(19,9),nullzero"`                // Other Cost value for break 7
	OtherCost8                 float64   `bun:"other_cost8,type:decimal(19,9),nullzero"`                // Other Cost value for break 8
	OtherCost9                 float64   `bun:"other_cost9,type:decimal(19,9),nullzero"`                // Other Cost value for break 9
	OtherCost10                float64   `bun:"other_cost10,type:decimal(19,9),nullzero"`               // Other Cost value for break 10
	OtherCost11                float64   `bun:"other_cost11,type:decimal(19,9),nullzero"`               // Other Cost value for break 11
	OtherCost12                float64   `bun:"other_cost12,type:decimal(19,9),nullzero"`               // Other Cost value for break 12
	OtherCost13                float64   `bun:"other_cost13,type:decimal(19,9),nullzero"`               // Other Cost value for break 13
	OtherCost14                float64   `bun:"other_cost14,type:decimal(19,9),nullzero"`               // Other Cost value for break 14
	OtherCost15                float64   `bun:"other_cost15,type:decimal(19,9),nullzero"`               // Other Cost value for break 15
	Uom1                       string    `bun:"uom1,type:varchar(8),nullzero"`                          // UOM for break 1
	Uom2                       string    `bun:"uom2,type:varchar(8),nullzero"`                          // UOM for break 2
	Uom3                       string    `bun:"uom3,type:varchar(8),nullzero"`                          // UOM for break 3
	Uom4                       string    `bun:"uom4,type:varchar(8),nullzero"`                          // UOM for break 4
	Uom5                       string    `bun:"uom5,type:varchar(8),nullzero"`                          // UOM for break 5
	Uom6                       string    `bun:"uom6,type:varchar(8),nullzero"`                          // UOM for break 6
	Uom7                       string    `bun:"uom7,type:varchar(8),nullzero"`                          // UOM for break 7
	Uom8                       string    `bun:"uom8,type:varchar(8),nullzero"`                          // UOM for break 8
	Uom9                       string    `bun:"uom9,type:varchar(8),nullzero"`                          // UOM for break 9
	Uom10                      string    `bun:"uom10,type:varchar(8),nullzero"`                         // UOM for break 10
	Uom11                      string    `bun:"uom11,type:varchar(8),nullzero"`                         // UOM for break 11
	Uom12                      string    `bun:"uom12,type:varchar(8),nullzero"`                         // UOM for break 12
	Uom13                      string    `bun:"uom13,type:varchar(8),nullzero"`                         // UOM for break 13
	Uom14                      string    `bun:"uom14,type:varchar(8),nullzero"`                         // UOM for break 14
	TotalingMethodCd           int16     `bun:"totaling_method_cd,type:smallint"`                       // Indicates the Totaling Method (Item, Supplier, Order, etc.)
	TotalingBasisCd            int16     `bun:"totaling_basis_cd,type:smallint"`                        // Indicates the Totaling Basis (Sales Unit, Piece, etc.)
	RowStatusFlag              int16     `bun:"row_status_flag,type:smallint"`                          // Indicates current record status.
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`                       // Indicates the date/time this record was last modified.
	DateCreated                time.Time `bun:"date_created,type:datetime"`                             // Indicates the date/time this record was created.
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30)"`                    // ID of the user who last maintained this record
	OtherCostTypeCd            int16     `bun:"other_cost_type_cd,type:smallint,nullzero"`              // Indicates the type of the Other Cost value (Order, Value, etc.)
	OtherCostValue             float64   `bun:"other_cost_value,type:decimal(19,9),nullzero"`           // Stores a specific Other Cost value when other_cost_type_cd is set to value
	OtherCostSourceCd          int16     `bun:"other_cost_source_cd,type:smallint,nullzero"`            // Stores the Other Cost source when other_cost_type_cd is set to source
	CostCalculationMethodCd    int16     `bun:"cost_calculation_method_cd,type:smallint,nullzero"`      // Stores the calulation method related to Other Cost (multiplier, difference, etc.)
	CostCalculationValue       float64   `bun:"cost_calculation_value,type:decimal(19,4),nullzero"`     // Stores the calculation value related to the Other Cost calculation method
	CommissionCostTypeCd       int16     `bun:"commission_cost_type_cd,type:smallint,nullzero"`         // Indicates the type of  the Commission Cost value (Order, Value, etc.)
	CommissionCostValue        float64   `bun:"commission_cost_value,type:decimal(19,9),nullzero"`      // Stores a specific Commission Cost value when commission_cost_type_cd is set to value
	CommissionCostSourceCd     int16     `bun:"commission_cost_source_cd,type:smallint,nullzero"`       // Stores the Commission Cost Source when commission_cost_type_cd is set to source
	CommissionCostCalcMethodCd int16     `bun:"commission_cost_calc_method_cd,type:smallint,nullzero"`  // Stores the calculation method related to Commission Cost (multiplier, difference, etc.)
	CommissionCostCalcValue    float64   `bun:"commission_cost_calc_value,type:decimal(19,4),nullzero"` // Stores the calculation value related to the Commission Cost calculation method
	PricePageId                string    `bun:"price_page_id,type:varchar(20),nullzero"`                // User-defined Price Page identifier
	CustomerId                 float64   `bun:"customer_id,type:decimal(19,0),nullzero"`                // Customer ID - used for Customer Part Number type pages only
	CustomerPartNo             string    `bun:"customer_part_no,type:varchar(40),nullzero"`             // Customer Part Number - used for Customer Part Number type pages only
	CompanyId                  string    `bun:"company_id,type:varchar(8),nullzero"`                    // Unique code that identifies a company.
	InvMastUid                 int32     `bun:"inv_mast_uid,type:int,nullzero"`                         // Unique identifier for the item id.
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	CommCostCalcValue1         float64   `bun:"comm_cost_calc_value1,type:decimal(19,9),default:(0)"`    // Holds the Commission Cost Calculation Value (number 1 of 15) for the Values tab in Price Page Maintenance
	CommCostCalcValue2         float64   `bun:"comm_cost_calc_value2,type:decimal(19,9),default:(0)"`    // Holds the Commission Cost Calculation Value (number 2 of 15) for the Values tab in Price Page Maintenance
	CommCostCalcValue3         float64   `bun:"comm_cost_calc_value3,type:decimal(19,9),default:(0)"`    // Holds the Commission Cost Calculation Value (number 3 of 15) for the Values tab in Price Page Maintenance
	CommCostCalcValue4         float64   `bun:"comm_cost_calc_value4,type:decimal(19,9),default:(0)"`    // Holds the Commission Cost Calculation Value (number 4 of 15) for the Values tab in Price Page Maintenance
	CommCostCalcValue5         float64   `bun:"comm_cost_calc_value5,type:decimal(19,9),default:(0)"`    // Holds the Commission Cost Calculation Value (number 5 of 15) for the Values tab in Price Page Maintenance
	CommCostCalcValue6         float64   `bun:"comm_cost_calc_value6,type:decimal(19,9),default:(0)"`    // Holds the Commission Cost Calculation Value (number 6 of 15) for the Values tab in Price Page Maintenance
	CommCostCalcValue7         float64   `bun:"comm_cost_calc_value7,type:decimal(19,9),default:(0)"`    // Holds the Commission Cost Calculation Value (number 7 of 15) for the Values tab in Price Page Maintenance
	CommCostCalcValue8         float64   `bun:"comm_cost_calc_value8,type:decimal(19,9),default:(0)"`    // Holds the Commission Cost Calculation Value (number 8 of 15) for the Values tab in Price Page Maintenance
	CommCostCalcValue9         float64   `bun:"comm_cost_calc_value9,type:decimal(19,9),default:(0)"`    // Holds the Commission Cost Calculation Value (number 9 of 15) for the Values tab in Price Page Maintenance
	CommCostCalcValue10        float64   `bun:"comm_cost_calc_value10,type:decimal(19,9),default:(0)"`   // Holds the Commission Cost Calculation Value (number 10 of 15) for the Values tab in Price Page Maintenance
	CommCostCalcValue11        float64   `bun:"comm_cost_calc_value11,type:decimal(19,9),default:(0)"`   // Holds the Commission Cost Calculation Value (number 11 of 15) for the Values tab in Price Page Maintenance
	CommCostCalcValue12        float64   `bun:"comm_cost_calc_value12,type:decimal(19,9),default:(0)"`   // Holds the Commission Cost Calculation Value (number 12 of 15) for the Values tab in Price Page Maintenance
	CommCostCalcValue13        float64   `bun:"comm_cost_calc_value13,type:decimal(19,9),default:(0)"`   // Holds the Commission Cost Calculation Value (number 13 of 15) for the Values tab in Price Page Maintenance
	CommCostCalcValue14        float64   `bun:"comm_cost_calc_value14,type:decimal(19,9),default:(0)"`   // Holds the Commission Cost Calculation Value (number 14 of 15) for the Values tab in Price Page Maintenance
	CommCostCalcValue15        float64   `bun:"comm_cost_calc_value15,type:decimal(19,9),default:(0)"`   // Holds the Commission Cost Calculation Value (number 15 of 15) for the Values tab in Price Page Maintenance
	CalculatorType             string    `bun:"calculator_type,type:char(1),default:('B')"`              // Whether this pricing page is used to calculate price [P], cost [C] or both [B]. Possible values: P,C or B only.
	CurrencyId                 float64   `bun:"currency_id,type:decimal(19,0),nullzero"`                 // Determines the price page currency
	ValuesCurrencyId           float64   `bun:"values_currency_id,type:decimal(19,0),nullzero"`          // Determines the price page values currency
	ApplyPpToMroCd             int16     `bun:"apply_pp_to_mro_cd,type:smallint,default:((1104))"`       // Column indicates if price page will be applied to manufacturer rep orders
	DateLastSentOn832          time.Time `bun:"date_last_sent_on_832,type:datetime,nullzero"`            // Will indicate the last time a pages detail was transmitted on an 832.
	DatePageDeleted            time.Time `bun:"date_page_deleted,type:datetime,nullzero"`                // Will indicate the date a page was deleted.
	PriceFamilyUid             int32     `bun:"price_family_uid,type:int,nullzero"`                      // UID for pricing by price family
	PeerPricePageUid           int32     `bun:"peer_price_page_uid,type:int,nullzero"`                   // Indicates the parent of this price page - if any.
	ContractLineUid            int32     `bun:"contract_line_uid,type:int,nullzero"`                     // Indicates the contract line that created this price page.
	OnContractFlag             string    `bun:"on_contract_flag,type:char(1),default:('N')"`             // The column indicates whether the page is on contract or not.  Default value is N
	StrategicPriceAppliesToCd  int32     `bun:"strategic_price_applies_to_cd,type:int,default:((2636))"` // Whether strategic price applies to Core/Non-Core items
	RoundToNextDollar          int16     `bun:"round_to_next_dollar,type:tinyint,nullzero"`              // Custom - Used for special rounding (next 1 or 9  dollar)
	ApplyFreightFactor         string    `bun:"apply_freight_factor,type:char(1),default:('N')"`         // Indicates that vendor contract freight factor should be applied to other cost
	FreightFactorSourceCd      int32     `bun:"freight_factor_source_cd,type:int,nullzero"`              // Source value to multiply freight factor by
	NoChargeFlag               string    `bun:"no_charge_flag,type:char(1),nullzero"`                    // Custom column to indicate no price charged for the page
	NonStockItemsOnlyFlag      string    `bun:"non_stock_items_only_flag,type:char(1),default:('N')"`    // (Custom) Indicates price page is to be used for non stock items only.
	ApplyPpToSopCd             int16     `bun:"apply_pp_to_sop_cd,type:smallint,default:((1103))"`       // Column indicates if price page will be applied to Service Order Parts
	PriceOverride              string    `bun:"price_override,type:char(1),default:('N')"`               // Flag that states if the users is going to use a rebate margin to allow a rebate
	RebateMargin               float64   `bun:"rebate_margin,type:decimal(19,9),nullzero"`               // The percentage amout after the rebate of the total invoiced amount
	UpperMarginVariance        float64   `bun:"upper_margin_variance,type:decimal(19,9),nullzero"`       // Allowed percentage to exceed the rebate margin
	LowerMarginVariance        float64   `bun:"lower_margin_variance,type:decimal(19,9),nullzero"`       // Allowed percentage to be bellow the rebate margin
	ExcludeOrderLevelDiscFlag  string    `bun:"exclude_order_level_disc_flag,type:char(1),nullzero"`
	RolledItemPricingTypeCd    int32     `bun:"rolled_item_pricing_type_cd,type:int,nullzero"` // Custom: Indicates the pricing type for rolled items.  Full Roll, Partial Roll, or Non-Roll.
}
