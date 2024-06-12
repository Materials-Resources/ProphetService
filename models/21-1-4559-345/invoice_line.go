package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceLine struct {
	bun.BaseModel               `bun:"table:invoice_line"`
	InvoiceNo                   string    `bun:"invoice_no,type:varchar(10),pk"`
	QtyRequested                float64   `bun:"qty_requested,type:decimal(19,9)"`
	QtyShipped                  float64   `bun:"qty_shipped,type:decimal(19,9)"`
	UnitOfMeasure               string    `bun:"unit_of_measure,type:varchar(10),nullzero"`
	ItemId                      string    `bun:"item_id,type:varchar(40),nullzero"`
	ItemDesc                    string    `bun:"item_desc,type:varchar(40)"`
	UnitPrice                   float64   `bun:"unit_price,type:decimal(19,9)"`
	ExtendedPrice               float64   `bun:"extended_price,type:decimal(19,4)"`
	GlRevenueAccountNo          string    `bun:"gl_revenue_account_no,type:varchar(32),nullzero"`
	GlSalseTaxAccountNo         string    `bun:"gl_salse_tax_account_no,type:varchar(32),nullzero"`
	GlCogs                      string    `bun:"gl_cogs,type:varchar(32),nullzero"`
	GlInventory                 string    `bun:"gl_inventory,type:varchar(32),nullzero"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30)"`
	OrderNo                     string    `bun:"order_no,type:varchar(8),nullzero"`
	CogsAmount                  float64   `bun:"cogs_amount,type:decimal(19,4),default:(0)"`
	JobId                       string    `bun:"job_id,type:varchar(32),nullzero"`
	CustomerPartNumber          string    `bun:"customer_part_number,type:varchar(40),nullzero"`
	CompanyId                   string    `bun:"company_id,type:varchar(8),nullzero"`
	TaxItem                     string    `bun:"tax_item,type:char,nullzero"`
	PricingQuantity             float64   `bun:"pricing_quantity,type:decimal(19,9),nullzero"`
	NetQuantity                 float64   `bun:"net_quantity,type:decimal(19,9),nullzero"`
	LineNo                      float64   `bun:"line_no,type:decimal(19,0),pk"`
	SalesCost                   float64   `bun:"sales_cost,type:decimal(19,9),default:(0)"`
	CommissionCost              float64   `bun:"commission_cost,type:decimal(19,9),default:(0)"`
	OtherCost                   float64   `bun:"other_cost,type:decimal(19,9),default:(0)"`
	OeLineNumber                float64   `bun:"oe_line_number,type:decimal(19,0),nullzero"`
	OtherChargeItem             string    `bun:"other_charge_item,type:char,nullzero"`
	ExceptionalSales            string    `bun:"exceptional_sales,type:char"`
	PricingUnit                 string    `bun:"pricing_unit,type:varchar(8),nullzero"`
	InvoiceLineUid              int32     `bun:"invoice_line_uid,type:int"`
	InvoiceLineUidParent        int32     `bun:"invoice_line_uid_parent,type:int"`
	InvMastUid                  int32     `bun:"inv_mast_uid,type:int,nullzero"`
	InvoiceLineType             int32     `bun:"invoice_line_type,type:int,default:(0)"`
	SalesUnitSize               float64   `bun:"sales_unit_size,type:decimal(19,4),nullzero"`
	PricingUnitSize             float64   `bun:"pricing_unit_size,type:decimal(19,4),nullzero"`
	ProductGroupId              string    `bun:"product_group_id,type:varchar(8),nullzero"`
	SupplierId                  float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	JobPriceLineUid             int32     `bun:"job_price_line_uid,type:int,nullzero"`
	CostCenter                  string    `bun:"cost_center,type:char(255),nullzero"`
	BudgetCd                    string    `bun:"budget_cd,type:varchar(255),nullzero"`
	CustomerContractUid         int32     `bun:"customer_contract_uid,type:int,nullzero"`
	GlAppliedLabor              string    `bun:"gl_applied_labor,type:char(32),nullzero"`
	LaborAmount                 float64   `bun:"labor_amount,type:decimal(19,4),nullzero"`
	MacForSpecialItems          float64   `bun:"mac_for_special_items,type:decimal(19,9),nullzero"`
	SpecialPurchaseQtyReceived  float64   `bun:"special_purchase_qty_received,type:decimal(19,9),nullzero"`
	OneTimePriceFlag            string    `bun:"one_time_price_flag,type:char,nullzero"`
	SuggestedRetailPrice        float64   `bun:"suggested_retail_price,type:decimal(19,4),nullzero"`
	HoursWorked                 float64   `bun:"hours_worked,type:decimal(19,9),nullzero"`
	CorePrice                   float64   `bun:"core_price,type:decimal(19,9),nullzero"`
	EnvironmentalFee            float64   `bun:"environmental_fee,type:decimal(19,4),nullzero"`
	AdminFee                    float64   `bun:"admin_fee,type:decimal(19,4),nullzero"`
	CoveredExtendedPrice        float64   `bun:"covered_extended_price,type:decimal(19,9),nullzero"`
	CostPricePageUid            int32     `bun:"cost_price_page_uid,type:int,nullzero"`
	Buyer                       string    `bun:"buyer,type:varchar(255),nullzero"`
	Recipient                   string    `bun:"recipient,type:varchar(255),nullzero"`
	VerifiedFlag                string    `bun:"verified_flag,type:char,nullzero"`
	VerifiedCode                string    `bun:"verified_code,type:varchar(255),nullzero"`
	SkuExceptionalQty           float64   `bun:"sku_exceptional_qty,type:decimal(19,9),nullzero"`
	LastReviewedDate            time.Time `bun:"last_reviewed_date,type:datetime,nullzero"`
	ProcessedFlag               string    `bun:"processed_flag,type:char,default:('N')"`
	CogsMarkupAmount            float64   `bun:"cogs_markup_amount,type:decimal(19,4),nullzero"`
	SalesDiscountGroupId        string    `bun:"sales_discount_group_id,type:varchar(8),nullzero"`
	PriceFamilyUid              int32     `bun:"price_family_uid,type:int,nullzero"`
	UnitPickFee                 float64   `bun:"unit_pick_fee,type:decimal(19,9),nullzero"`
	CostCarrierContractLineUid  int32     `bun:"cost_carrier_contract_line_uid,type:int,nullzero"`
	PriceCarrierContractLineUid int32     `bun:"price_carrier_contract_line_uid,type:int,nullzero"`
	CostCarrierContractZLineUid int32     `bun:"cost_carrier_contract_z_line_uid,type:int,nullzero"`
	TargetPrice                 float64   `bun:"target_price,type:decimal(19,9),nullzero"`
	TaxAmountPaidOnDpApplied    float64   `bun:"tax_amount_paid_on_dp_applied,type:decimal(19,2),nullzero"`
	SentToCarrierDate           time.Time `bun:"sent_to_carrier_date,type:datetime,nullzero"`
	NetBillingFlag              string    `bun:"net_billing_flag,type:char,default:('N')"`
	HazNumOfPackages            string    `bun:"haz_num_of_packages,type:varchar(255),nullzero"`
	PrintPartNo                 string    `bun:"print_part_no,type:varchar(40),nullzero"`
	DiscountItemFlag            string    `bun:"discount_item_flag,type:char,nullzero"`
	DistributorFunding          float64   `bun:"distributor_funding,type:decimal(19,4),nullzero"`
	SupplierFunding             float64   `bun:"supplier_funding,type:decimal(19,4),nullzero"`
	PriorAuthorizationCd        string    `bun:"prior_authorization_cd,type:varchar(255),nullzero"`
	ClaimStartDate              time.Time `bun:"claim_start_date,type:datetime,nullzero"`
	ClaimEndDate                time.Time `bun:"claim_end_date,type:datetime,nullzero"`
	GlDistributorFundingAcctNo  string    `bun:"gl_distributor_funding_acct_no,type:varchar(32),nullzero"`
	GlSupplierFundingAcctNo     string    `bun:"gl_supplier_funding_acct_no,type:varchar(32),nullzero"`
	ExcludeFromEdi844867Flag    string    `bun:"exclude_from_edi_844_867_flag,type:char,nullzero"`
	OtherChargeCreditRebillFlag string    `bun:"other_charge_credit_rebill_flag,type:char,nullzero"`
	SubInvoiceNo                string    `bun:"sub_invoice_no,type:varchar(40),nullzero"`
	CommissionClassId           string    `bun:"commission_class_id,type:varchar(8),nullzero"`
	GlDimenTypeUid              int32     `bun:"gl_dimen_type_uid,type:int,nullzero"`
	GlDimensionId               string    `bun:"gl_dimension_id,type:varchar(255),nullzero"`
	GlDimensionDesc             string    `bun:"gl_dimension_desc,type:varchar(255),nullzero"`
}
