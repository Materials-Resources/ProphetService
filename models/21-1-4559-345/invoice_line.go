package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceLine struct {
	bun.BaseModel               `bun:"table:invoice_line"`
	InvoiceNo                   string     `bun:"invoice_no,type:varchar(10),pk"`                 // Invoice number that is associated with invoice line
	QtyRequested                float64    `bun:"qty_requested,type:decimal(19,9)"`               // Quantity from the original sales order.
	QtyShipped                  float64    `bun:"qty_shipped,type:decimal(19,9)"`                 // How many/much of the item were/was actually shipped?
	UnitOfMeasure               *string    `bun:"unit_of_measure,type:varchar(10)"`               // What is the unit of measure for this row?
	ItemId                      *string    `bun:"item_id,type:varchar(40)"`                       // The item id of the invoice line item.
	ItemDesc                    string     `bun:"item_desc,type:varchar(40)"`                     // How would you describe this item or material?
	UnitPrice                   float64    `bun:"unit_price,type:decimal(19,9)"`                  // What is the unit price for this line item?
	ExtendedPrice               float64    `bun:"extended_price,type:decimal(19,4)"`              // The total dollar amount of the invoice line item.
	GlRevenueAccountNo          *string    `bun:"gl_revenue_account_no,type:varchar(32)"`         // The revenue account of the invoice line item.
	GlSalseTaxAccountNo         *string    `bun:"gl_salse_tax_account_no,type:varchar(32)"`       // The sales tax account of the invoice line item.
	GlCogs                      *string    `bun:"gl_cogs,type:varchar(32)"`                       // The cost of goods sold account of the invoice line item.
	GlInventory                 *string    `bun:"gl_inventory,type:varchar(32)"`                  // The inventory account of the invoice line item.
	DateCreated                 time.Time  `bun:"date_created,type:datetime"`                     // Indicates the date/time this record was created.
	DateLastModified            time.Time  `bun:"date_last_modified,type:datetime"`               // Indicates the date/time this record was last modified.
	LastMaintainedBy            string     `bun:"last_maintained_by,type:varchar(30)"`            // ID of the user who last maintained this record
	OrderNo                     *string    `bun:"order_no,type:varchar(8)"`                       // What order does this note belong to?
	CogsAmount                  float64    `bun:"cogs_amount,type:decimal(19,4),default:(0)"`     // What is the Cost of Goods amount for this invoice
	JobId                       *string    `bun:"job_id,type:varchar(32)"`                        // This column is unused.
	CustomerPartNumber          *string    `bun:"customer_part_number,type:varchar(40)"`          // The customers part number which has been assigned to an item ID
	CompanyId                   *string    `bun:"company_id,type:varchar(8)"`                     // Unique code that identifies a company.
	TaxItem                     *string    `bun:"tax_item,type:char(1)"`                          // Indicates whether this line is a tax jurisdiction.
	PricingQuantity             *float64   `bun:"pricing_quantity,type:decimal(19,9)"`            // Quantity in terms of sales pricing units.
	NetQuantity                 *float64   `bun:"net_quantity,type:decimal(19,9)"`                // The total quantity
	LineNo                      float64    `bun:"line_no,type:decimal(19,0),pk"`                  // Indicates the unique line number for this invoice.
	SalesCost                   float64    `bun:"sales_cost,type:decimal(19,9),default:(0)"`      // The total cost of the items on an order.
	CommissionCost              float64    `bun:"commission_cost,type:decimal(19,9),default:(0)"` // The commission cost of the invoice line item.
	OtherCost                   float64    `bun:"other_cost,type:decimal(19,9),default:(0)"`      // The other cost of the invoice line item. Usually used for rebates.
	OeLineNumber                *float64   `bun:"oe_line_number,type:decimal(19,0)"`              // What order line does this invoice line correspond to?
	OtherChargeItem             *string    `bun:"other_charge_item,type:char(1)"`                 // Indicates that the invoice line item is a charge rather than material.
	ExceptionalSales            string     `bun:"exceptional_sales,type:char(1)"`                 // Used to determine if a sale is exceptional.
	PricingUnit                 *string    `bun:"pricing_unit,type:varchar(8)"`                   // Maintains the pricing unit for the invoice line.
	InvoiceLineUid              int32      `bun:"invoice_line_uid,type:int"`                      // Unique identifier for the invoice line item.
	InvoiceLineUidParent        int32      `bun:"invoice_line_uid_parent,type:int"`               // What is the parent of this invoice line -  if any?
	InvMastUid                  *int32     `bun:"inv_mast_uid,type:int"`                          // Unique identifier for the item id.
	InvoiceLineType             int32      `bun:"invoice_line_type,type:int,default:(0)"`         // Id indicating the type of invoice line (tax, regular, incoming freight, outgoing freight...)
	SalesUnitSize               *float64   `bun:"sales_unit_size,type:decimal(19,4)"`             // The unit size of the sales unit.
	PricingUnitSize             *float64   `bun:"pricing_unit_size,type:decimal(19,4)"`           // Maintains the pricing unit size.
	ProductGroupId              *string    `bun:"product_group_id,type:varchar(8)"`               // The product group of the invoice line item.
	SupplierId                  *float64   `bun:"supplier_id,type:decimal(19,0)"`                 // What supplier supplies material for this stage?
	CreatedBy                   *string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	JobPriceLineUid             *int32     `bun:"job_price_line_uid,type:int"`                      // Unique ID for associated Job/Contract Line Items.
	CostCenter                  *string    `bun:"cost_center,type:char(255)"`                       // Holds cost center data.  Copied from oe_line table on invoicing.
	BudgetCd                    *string    `bun:"budget_cd,type:varchar(255)"`                      // Custom (F23038): User defined budget code - informational only
	CustomerContractUid         *int32     `bun:"customer_contract_uid,type:int"`                   // Contains the customer_contract_uid associated with this invoice line
	GlAppliedLabor              *string    `bun:"gl_applied_labor,type:char(32)"`                   // Service Order applied labor account
	LaborAmount                 *float64   `bun:"labor_amount,type:decimal(19,4)"`                  // Amount column relating to gl_applied_labor
	MacForSpecialItems          *float64   `bun:"mac_for_special_items,type:decimal(19,9)"`         // Stores the moving average cost at the time of shipping for the item on the invoice line. Currently only used by custom.
	SpecialPurchaseQtyReceived  *float64   `bun:"special_purchase_qty_received,type:decimal(19,9)"` // Stores the received quantity of the special purchase item on the invoice line. Currently only used by custom.
	OneTimePriceFlag            *string    `bun:"one_time_price_flag,type:char(1)"`                 // Whether this invoice line is used with last margin pricing
	SuggestedRetailPrice        *float64   `bun:"suggested_retail_price,type:decimal(19,4)"`        // Price to be displayed on an invoice next to actual prices
	HoursWorked                 *float64   `bun:"hours_worked,type:decimal(19,9)"`                  // Hours worked for service labor item
	CorePrice                   *float64   `bun:"core_price,type:decimal(19,9)"`                    // Custom column for the extended price for core item
	EnvironmentalFee            *float64   `bun:"environmental_fee,type:decimal(19,4)"`             // Environmental Fee related for the invoice line
	AdminFee                    *float64   `bun:"admin_fee,type:decimal(19,4)"`                     // Admin Fee related to invoice line
	CoveredExtendedPrice        *float64   `bun:"covered_extended_price,type:decimal(19,9)"`        // Amount of extended price covered by warranty
	CostPricePageUid            *int32     `bun:"cost_price_page_uid,type:int"`                     // Price page used to calculate other and commission cost
	Buyer                       *string    `bun:"buyer,type:varchar(255)"`                          // Holds buyer data. Copied from oe_line table on invoicing.
	Recipient                   *string    `bun:"recipient,type:varchar(255)"`                      // Holds recipient data. Copied from oe_line table on invoicing.
	VerifiedFlag                *string    `bun:"verified_flag,type:char(1)"`                       // Indicates line has been verified - custom
	VerifiedCode                *string    `bun:"verified_code,type:varchar(255)"`                  // The code entered for a verified line item. - custom
	SkuExceptionalQty           *float64   `bun:"sku_exceptional_qty,type:decimal(19,9)"`           // Custom column to indicate the exceptional amount of sales
	LastReviewedDate            *time.Time `bun:"last_reviewed_date,type:datetime"`                 // Indicates last reviewed date of the usage period - for demand review window only
	ProcessedFlag               string     `bun:"processed_flag,type:char(1),default:('N')"`        // Custom column for 47261 - indicates if invoice line has been processed by outside service
	CogsMarkupAmount            *float64   `bun:"cogs_markup_amount,type:decimal(19,4)"`            // Mark Up Amount added by Inter-Company Transactions
	SalesDiscountGroupId        *string    `bun:"sales_discount_group_id,type:varchar(8)"`          // Sales Discount Group
	PriceFamilyUid              *int32     `bun:"price_family_uid,type:int"`                        // Price Family UID
	UnitPickFee                 *float64   `bun:"unit_pick_fee,type:decimal(19,9)"`                 // The pick fee when making an invoice
	CostCarrierContractLineUid  *int32     `bun:"cost_carrier_contract_line_uid,type:int"`          // Specifies carrier contract line used to cost this line.  FK to carrier_contract_line.
	PriceCarrierContractLineUid *int32     `bun:"price_carrier_contract_line_uid,type:int"`         // Specifies carrier contract line used to price this line.  FK to carrier_contract_line.
	CostCarrierContractZLineUid *int32     `bun:"cost_carrier_contract_z_line_uid,type:int"`        // Specifies carrier contract z quote line used to cost this line.  FK to carrier_contract_z_line
	TargetPrice                 *float64   `bun:"target_price,type:decimal(19,9)"`                  // column that stored the Target Price
	TaxAmountPaidOnDpApplied    *float64   `bun:"tax_amount_paid_on_dp_applied,type:decimal(19,2)"` // Tax Amt Paid on a shipping invoice with IVA tax that has a DP Applied to it.
	SentToCarrierDate           *time.Time `bun:"sent_to_carrier_date,type:datetime"`               // Date that this line was exported with Carrier Claim export.
	NetBillingFlag              string     `bun:"net_billing_flag,type:char(1),default:('N')"`      // Net Billing Flag
	HazNumOfPackages            *string    `bun:"haz_num_of_packages,type:varchar(255)"`
	PrintPartNo                 *string    `bun:"print_part_no,type:varchar(40)"`
	DiscountItemFlag            *string    `bun:"discount_item_flag,type:char(1)"`
	DistributorFunding          *float64   `bun:"distributor_funding,type:decimal(19,4)"`          // Funding qty for Distributor funding account
	SupplierFunding             *float64   `bun:"supplier_funding,type:decimal(19,4)"`             // Supplier qty for Supplier funding account
	PriorAuthorizationCd        *string    `bun:"prior_authorization_cd,type:varchar(255)"`        // Person at the supplier who authorized the claim
	ClaimStartDate              *time.Time `bun:"claim_start_date,type:datetime"`                  // Claim Start Date
	ClaimEndDate                *time.Time `bun:"claim_end_date,type:datetime"`                    // Claim End Date
	GlDistributorFundingAcctNo  *string    `bun:"gl_distributor_funding_acct_no,type:varchar(32)"` // Specify the accounts to which the claim amounts of Distributor will be posted
	GlSupplierFundingAcctNo     *string    `bun:"gl_supplier_funding_acct_no,type:varchar(32)"`    // Specify the accounts to which the claim amounts of Supplier will be posted
	ExcludeFromEdi844867Flag    *string    `bun:"exclude_from_edi_844_867_flag,type:char(1)"`      // Flag to exclude lines from a rebilled invoice on EDI 844 and 867
	OtherChargeCreditRebillFlag *string    `bun:"other_charge_credit_rebill_flag,type:char(1)"`    // Determines if the other charge item was added via RMB on Credit and Rebill window
	SubInvoiceNo                *string    `bun:"sub_invoice_no,type:varchar(40)"`                 // Sub Invoice number for Pro Forma Invoice
	CommissionClassId           *string    `bun:"commission_class_id,type:varchar(8)"`             // The commission class of the item at the time the invoice was created
	GlDimenTypeUid              *int32     `bun:"gl_dimen_type_uid,type:int"`                      // UID for GL dimension type
	GlDimensionId               *string    `bun:"gl_dimension_id,type:varchar(255)"`               // Value for dimension type
	GlDimensionDesc             *string    `bun:"gl_dimension_desc,type:varchar(255)"`             // Desc of dimension type
}
