package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CommissionRule struct {
	bun.BaseModel              `bun:"table:commission_rule"`
	CompanyId                  string    `bun:"company_id,type:varchar(8),pk"`
	CommissionRuleId           string    `bun:"commission_rule_id,type:varchar(8),pk"`
	CommissionRuleDesc         string    `bun:"commission_rule_desc,type:varchar(40)"`
	AppliesTo                  string    `bun:"applies_to,type:char"`
	EffectiveDate              time.Time `bun:"effective_date,type:datetime"`
	ExpirationDate             time.Time `bun:"expiration_date,type:datetime"`
	CommissionBasedOn          string    `bun:"commission_based_on,type:char(2)"`
	NegativeCommission         string    `bun:"negative_commission,type:char"`
	StockItems                 string    `bun:"stock_items,type:char"`
	NonStockItems              string    `bun:"non_stock_items,type:char"`
	DirectShipItems            string    `bun:"direct_ship_items,type:char"`
	GroupBy                    string    `bun:"group_by,type:char(2)"`
	ProductGroupId             string    `bun:"product_group_id,type:varchar(8),nullzero"`
	SalesDiscountGroupId       string    `bun:"sales_discount_group_id,type:varchar(8),nullzero"`
	SupplierId                 float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`
	ItemCommissionClass        string    `bun:"item_commission_class,type:varchar(8),nullzero"`
	DeleteFlag                 string    `bun:"delete_flag,type:char"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	CommissionTypeCd           int32     `bun:"commission_type_cd,type:int,default:((1614))"`
	ApplyToMfrRepOrdersCd      int32     `bun:"apply_to_mfr_rep_orders_cd,type:int,default:((1103))"`
	ApplyToServiceOrdersCd     int32     `bun:"apply_to_service_orders_cd,type:int,default:((1103))"`
	CustomerId                 float64   `bun:"customer_id,type:decimal(19,0),nullzero"`
	ManufacturingClassId       string    `bun:"manufacturing_class_id,type:varchar(255),nullzero"`
	InvMastUid                 int32     `bun:"inv_mast_uid,type:int,nullzero"`
	CustomerPartNumber         string    `bun:"customer_part_number,type:varchar(40),nullzero"`
	FreightCodeUid             int32     `bun:"freight_code_uid,type:int,nullzero"`
	SalesLocationId            float64   `bun:"sales_location_id,type:decimal(19,0),nullzero"`
	PrimarySupplierId          float64   `bun:"primary_supplier_id,type:decimal(19,0),nullzero"`
	CommissionRuleUid          int32     `bun:"commission_rule_uid,type:int,identity"`
	AllocatedBackorderFlag     string    `bun:"allocated_backorder_flag,type:char,default:('Y')"`
	SpecialFlag                string    `bun:"special_flag,type:char,default:('Y')"`
	OtherChargeItemFlag        string    `bun:"other_charge_item_flag,type:char,default:('N')"`
	CustomerDays               int32     `bun:"customer_days,type:int,nullzero"`
	ItemDays                   int32     `bun:"item_days,type:int,nullzero"`
	LineThreshold              float64   `bun:"line_threshold,type:decimal(19,2),default:((0.00))"`
	LineThresholdType          string    `bun:"line_threshold_type,type:char,nullzero"`
	AllOtherDispositionsFlag   string    `bun:"all_other_dispositions_flag,type:char,default:('Y')"`
	IncludeManualInvoicesFlag  string    `bun:"include_manual_invoices_flag,type:char,default:('Y')"`
	NewCustomerFlag            string    `bun:"new_customer_flag,type:char,default:('N')"`
	NewItemForCustomerFlag     string    `bun:"new_item_for_customer_flag,type:char,default:('N')"`
	ApplyToProgressBillsCd     int32     `bun:"apply_to_progress_bills_cd,type:int,default:((1103))"`
	CompletedProgBillsOnlyFlag string    `bun:"completed_prog_bills_only_flag,type:char,default:('N')"`
	TerritoryUid               int32     `bun:"territory_uid,type:int,nullzero"`
	ApplyToShiptoFlag          string    `bun:"apply_to_shipto_flag,type:char,default:('N')"`
	NewItemForShipToFlag       string    `bun:"new_item_for_ship_to_flag,type:char,default:('N')"`
	ApplyToSbCreditMemoFlag    string    `bun:"apply_to_sb_credit_memo_flag,type:char,default:('N')"`
	Class4id                   string    `bun:"class_4id,type:varchar(255),nullzero"`
	ApplyToManualArOnlyFlag    string    `bun:"apply_to_manual_ar_only_flag,type:char,nullzero"`
}
