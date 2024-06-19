package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CommissionRule struct {
	bun.BaseModel              `bun:"table:commission_rule"`
	CompanyId                  string    `bun:"company_id,type:varchar(8),pk"`                              // Unique code that identifies a company.
	CommissionRuleId           string    `bun:"commission_rule_id,type:varchar(8),pk"`                      // Commission Rule ID
	CommissionRuleDesc         string    `bun:"commission_rule_desc,type:varchar(40)"`                      // Commission Rule Description~r~n
	AppliesTo                  string    `bun:"applies_to,type:char(1)"`                                    // Applies To
	EffectiveDate              time.Time `bun:"effective_date,type:datetime"`                               // When does this purchase pricing page go into effec
	ExpirationDate             time.Time `bun:"expiration_date,type:datetime"`                              // When does this note expire?
	CommissionBasedOn          string    `bun:"commission_based_on,type:char(2)"`                           // Indicates default basis for computing commissions
	NegativeCommission         string    `bun:"negative_commission,type:char(1)"`                           // Indicates if negative commission is computed for R
	StockItems                 string    `bun:"stock_items,type:char(1)"`                                   // Indicates if this rule applies to stock items
	NonStockItems              string    `bun:"non_stock_items,type:char(1)"`                               // Indicates if this rule applies to non-stock items
	DirectShipItems            string    `bun:"direct_ship_items,type:char(1)"`                             // Indicates if this rule applies to direct ship item
	GroupBy                    string    `bun:"group_by,type:char(2)"`                                      // Indicates the grouping when using this rule for co
	ProductGroupId             string    `bun:"product_group_id,type:varchar(8),nullzero"`                  // Product group associated with this rule
	SalesDiscountGroupId       string    `bun:"sales_discount_group_id,type:varchar(8),nullzero"`           // Discount group this rule groups by (filled in when
	SupplierId                 float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`                    // What supplier supplies material for this stage?
	ItemCommissionClass        string    `bun:"item_commission_class,type:varchar(8),nullzero"`             // Item Commission Class this rule groups by (filled
	DeleteFlag                 string    `bun:"delete_flag,type:char(1)"`                                   // Indicates whether this record is logically deleted
	DateCreated                time.Time `bun:"date_created,type:datetime"`                                 // Indicates the date/time this record was created.
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`                           // Indicates the date/time this record was last modified.
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`  // ID of the user who last maintained this record
	CommissionTypeCd           int32     `bun:"commission_type_cd,type:int,default:((1614))"`               // What kind of orders commission is for
	ApplyToMfrRepOrdersCd      int32     `bun:"apply_to_mfr_rep_orders_cd,type:int,default:((1103))"`       // Indicates whether to apply this rule to Manufacturer Rep Orders
	ApplyToServiceOrdersCd     int32     `bun:"apply_to_service_orders_cd,type:int,default:((1103))"`       // Indicates whether to apply this rule to service orders
	CustomerId                 float64   `bun:"customer_id,type:decimal(19,0),nullzero"`                    // Customer ID associated with this rule.
	ManufacturingClassId       string    `bun:"manufacturing_class_id,type:varchar(255),nullzero"`          // Manufacturing Class ID associated with this rule.
	InvMastUid                 int32     `bun:"inv_mast_uid,type:int,nullzero"`                             // Unique Item ID associated with this rule.
	CustomerPartNumber         string    `bun:"customer_part_number,type:varchar(40),nullzero"`             // Customer Part No associated with this rule.
	FreightCodeUid             int32     `bun:"freight_code_uid,type:int,nullzero"`                         // Freight code associated with this rule.
	SalesLocationId            float64   `bun:"sales_location_id,type:decimal(19,0),nullzero"`              // Sales Location ID associated with this rule.
	PrimarySupplierId          float64   `bun:"primary_supplier_id,type:decimal(19,0),nullzero"`            // Primary Supplier ID associated with this rule.
	CommissionRuleUid          int32     `bun:"commission_rule_uid,type:int,autoincrement,unique,identity"` // Unique identifier for a record
	AllocatedBackorderFlag     string    `bun:"allocated_backorder_flag,type:char(1),default:('Y')"`        // Indicates whether rule applies to only lines where stock item is allocated or backordered.
	SpecialFlag                string    `bun:"special_flag,type:char(1),default:('Y')"`                    // Indicates whether rule applies to only lines where disposition is special.
	OtherChargeItemFlag        string    `bun:"other_charge_item_flag,type:char(1),default:('N')"`          // Indicates whether rule applies to only lines where item is other charge.
	CustomerDays               int32     `bun:"customer_days,type:int,nullzero"`                            // The number of days that this rule is valid from when the first order with this customer was entered.
	ItemDays                   int32     `bun:"item_days,type:int,nullzero"`                                // The number of days this rule is valid from when the first order with this item was entered.
	LineThreshold              float64   `bun:"line_threshold,type:decimal(19,2),default:((0.00))"`         // The threshold which the line must meet in order for the rule to apply.
	LineThresholdType          string    `bun:"line_threshold_type,type:char(1),nullzero"`                  // The type of threshold (percent or dollar).
	AllOtherDispositionsFlag   string    `bun:"all_other_dispositions_flag,type:char(1),default:('Y')"`     // This column dtermines whether to include any disposition that was not specifically requested.
	IncludeManualInvoicesFlag  string    `bun:"include_manual_invoices_flag,type:char(1),default:('Y')"`    // This column determines whether to include manual invoices, disclude them, or ONLY include them
	NewCustomerFlag            string    `bun:"new_customer_flag,type:char(1),default:('N')"`               // Indicates whether rule applies only to new customers (of a specified number of days).
	NewItemForCustomerFlag     string    `bun:"new_item_for_customer_flag,type:char(1),default:('N')"`      // Indicates whether rule applies only to new items for this customer (of a specified number of days).
	ApplyToProgressBillsCd     int32     `bun:"apply_to_progress_bills_cd,type:int,default:((1103))"`       // Indicates whether commission rule applies to Progress Billing invoices.
	CompletedProgBillsOnlyFlag string    `bun:"completed_prog_bills_only_flag,type:char(1),default:('N')"`  // Indicates whether the commission rule applies to only progress bills that have been completed versus those that are still in progress.
	TerritoryUid               int32     `bun:"territory_uid,type:int,nullzero"`                            // The territory UID for which this commission rule applies.
	ApplyToShiptoFlag          string    `bun:"apply_to_shipto_flag,type:char(1),default:('N')"`            // Indicates whether rule applies only to new Ship to (of a specified number of days).
	NewItemForShipToFlag       string    `bun:"new_item_for_ship_to_flag,type:char(1),default:('N')"`       // Indicates whether rule applies only to new items for this Ship to (of a specified number of days).
	ApplyToSbCreditMemoFlag    string    `bun:"apply_to_sb_credit_memo_flag,type:char(1),default:('N')"`    // Indicates if this rule apply to ServiceBench Credit Memos.
	Class4Id                   string    `bun:"class_4id,type:varchar(255),nullzero"`                       // Field used with the Order Class 4 Group By option in Commission Rules to determine if the salesrep on a particular sales order is eligible for commission.
	ApplyToManualArOnlyFlag    string    `bun:"apply_to_manual_ar_only_flag,type:char(1),nullzero"`         // apply_to_manual_ar_only_flag
}
