package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CustDefaults struct {
	bun.BaseModel              `bun:"table:cust_defaults"`
	ArAccountNo                string    `bun:"ar_account_no,type:varchar(32)"`                             // Account number representing an Accounts Receivable account.
	RevenueAccountNo           string    `bun:"revenue_account_no,type:varchar(32)"`                        // Account that shows the gross increase in your companys income.
	CosAccountNo               string    `bun:"cos_account_no,type:varchar(32)"`                            // Cost of Goods Sold account that tracks the cost of products sold.
	AllowedAccountNo           string    `bun:"allowed_account_no,type:varchar(32)"`                        // Account consisting of excess amounts that your company writes off from an invoice.
	TermsAccountNo             string    `bun:"terms_account_no,type:varchar(32)"`                          // Account consisting of discounts taken against invoices based upon the terms agreement.
	FreightAccountNo           string    `bun:"freight_account_no,type:varchar(32)"`                        // Account that posts the freight amount when an invoice is created for this customer.
	BrokerageAccountNo         string    `bun:"brokerage_account_no,type:varchar(32)"`                      // Account that keeps track of your brokerage costs.
	PriceFileId                *float64  `bun:"price_file_id,type:decimal(19,0)"`                           // This column is unused.
	CompanyId                  string    `bun:"company_id,type:varchar(8),pk"`                              // Unique code that identifies a company.
	LocationId                 *float64  `bun:"location_id,type:decimal(19,0)"`                             // An ID that corresponds to a location.
	DeleteFlag                 string    `bun:"delete_flag,type:char(1)"`                                   // Indicates whether this record is logically deleted
	DateCreated                time.Time `bun:"date_created,type:datetime"`                                 // Indicates the date/time this record was created.
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`                           // Indicates the date/time this record was last modified.
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`  // ID of the user who last maintained this record
	Fob                        *string   `bun:"fob,type:varchar(20)"`                                       // (Free On Board) The point in the delivery process when freight costs and liability become the responsibility of the customer.
	DefaultDisposition         *string   `bun:"default_disposition,type:char(1)"`                           // Disposition setting automatically assigned to items ordered by a specific customer when the order cannot be filled.
	DeferredRevenueAccountNo   *string   `bun:"deferred_revenue_account_no,type:varchar(32)"`               // The default account number to post deferred revenue for a new customer
	PickTicketType             *string   `bun:"pick_ticket_type,type:varchar(2)"`                           // Indicates the type of pick ticket that prints for a specific ship to.
	FinanceChgRevenueAccountNo *string   `bun:"finance_chg_revenue_account_no,type:varchar(32)"`            // An account used to track a customers finance charges.
	InvoiceType                *string   `bun:"invoice_type,type:varchar(2)"`                               // Indicates what type of invoice is used.
	AcceptableWaitTime         float64   `bun:"acceptable_wait_time,type:decimal(5,0)"`                     // Maximum number of days a customer agrees to wait for ordered material.
	CreditLimit                float64   `bun:"credit_limit,type:decimal(22,4),default:(0)"`                // Maximum allowed amount of a customers outstanding debt at any given time.
	PendingPaymentAccountNo    *string   `bun:"pending_payment_account_no,type:varchar(32)"`                // An account used to track disputed zero dollar invoices.
	TermsId                    *string   `bun:"terms_id,type:varchar(2)"`                                   // Default terms used when creating new customer records.
	PricingMethodCd            *int16    `bun:"pricing_method_cd,type:tinyint"`                             // Default pricing method code for new customers
	SourcePriceCd              *int32    `bun:"source_price_cd,type:int"`                                   // Default source for Multiplier pricing method code for new customers
	Multiplier                 *float64  `bun:"multiplier,type:decimal(19,9)"`                              // Default multiplier for Multiplier pricing method code for new customers
	PriceLibraryId             *string   `bun:"price_library_id,type:varchar(20)"`                          // Default price library id for Price Library pricing method code for new customers
	SalesrepId                 *string   `bun:"salesrep_id,type:varchar(16)"`                               // Default salesre rep id for new customers
	DefaultBranch              *string   `bun:"default_branch,type:varchar(8)"`                             // Default branch id for new customers
	PreferredLocationId        *float64  `bun:"preferred_location_id,type:decimal(19,0)"`                   // Preferred location
	CustomerTypeUid            *int32    `bun:"customer_type_uid,type:int"`                                 // Custom column - ties back to customer_type table indicating what type of customer this is (ie customer, prospect, resale).
	AllowAdvanceBilling        *string   `bun:"allow_advance_billing,type:char(1),default:('N')"`           // Provides default value to whether or not advance billing is allowed
	AdvanceBillAccountNo       *string   `bun:"advance_bill_account_no,type:varchar(32)"`                   // Default account no to post advance bills
	DataIdentifierGroupUid     *int32    `bun:"data_identifier_group_uid,type:int"`                         // Customer defaults Data Identifier Group UID (typically used in the generation of sales-order related documents requiring data identifiers)
	CostCenterTrackingOption   *int32    `bun:"cost_center_tracking_option,type:int"`                       // Determines the cost center tracking options. Values from code table.
	SignatureRequired          string    `bun:"signature_required,type:char(1),default:('N')"`              // Indicates that this ship-to requires a signature on delivery
	PackingBasis               *string   `bun:"packing_basis,type:varchar(16)"`                             // The packing basis for the default customer
	InvoiceSurchargePct        float64   `bun:"invoice_surcharge_pct,type:decimal(19,9),default:(0)"`       // The default Invoice surcharge for all new customers
	LaborRate                  *float64  `bun:"labor_rate,type:decimal(19,9)"`                              // Labor rate for newly added customers
	TaxableFlag                string    `bun:"taxable_flag,type:char(1),default:('N')"`                    // Indicates whether a new customer will be taxable
	AllowLineItemFreightFlag   string    `bun:"allow_line_item_freight_flag,type:char(1),default:('N')"`    // Indicates if this customer allows freight on a per item level
	CarrierId                  *float64  `bun:"carrier_id,type:decimal(19,0)"`                              // Default carrier id for a ship-to.
	ShippingRouteUid           *int32    `bun:"shipping_route_uid,type:int"`                                // Default shipping route for a ship-to.
	ServiceTermsId             *string   `bun:"service_terms_id,type:varchar(2)"`                           // Terms for Service Orders
	DownpaymentPercentage      float64   `bun:"downpayment_percentage,type:decimal(19,2),default:((0.00))"` // Percentage of amount ordered that will be requested as downpayment.
	ReqPymtUponReleaseOfItems  string    `bun:"req_pymt_upon_release_of_items,type:char(1),default:('N')"`  // Indicates whether full pymt is required when ordered items are invoiced.
	IncludeDpSummaryOnInvoices string    `bun:"include_dp_summary_on_invoices,type:char(1),default:('N')"`  // Indicates whether printed invoices should include a downpayment summary.
	JobNumberRequiredFlag      string    `bun:"job_number_required_flag,type:char(1),default:('N')"`        // Indicates whether newly added customers should default to requiring a job number.
	UseLastMarginPricingFlag   string    `bun:"use_last_margin_pricing_flag,type:char(1),default:('N')"`    // Default customer to using last margin pricing
	InvoiceCompCostCdTier1     int32     `bun:"invoice_comp_cost_cd_tier1,type:int,default:((2110))"`       // Default cost used to compare invoice profit vs order profit
	DaysUntilQuoteExpires      *int32    `bun:"days_until_quote_expires,type:int"`                          // Number of days until a newly entered quote expires
	InvoiceCompCostCdTier2     int32     `bun:"invoice_comp_cost_cd_tier2,type:int,default:((2111))"`       // Default 2nd tier cost used to compare invoice profit vs order profit
	InvoiceCompCostCdTier3     int32     `bun:"invoice_comp_cost_cd_tier3,type:int,default:((1195))"`       // Default 3rd tier cost used to compare invoice profit vs order profit
	CustomerTypeCd             *int32    `bun:"customer_type_cd,type:int"`                                  // Indicates whether a new customer should default as a customer or prospect.
	ConsInvSummaryFilename     *string   `bun:"cons_inv_summary_filename,type:varchar(255)"`                // Filename specified for Consolidated Invoice Summary form.
	ConsInvDetailFilename      *string   `bun:"cons_inv_detail_filename,type:varchar(255)"`                 // Filename specified for Consolidated Invoice Detail form
	InvoiceFilename            *string   `bun:"invoice_filename,type:varchar(255)"`                         // Filename specified for Invoice Form
	StatementFilename          *string   `bun:"statement_filename,type:varchar(255)"`                       // Filename specified for Statement Form
	UseVendorItemTermsFlag     *string   `bun:"use_vendor_item_terms_flag,type:char(1)"`                    // If 'Y' and system setting line_item_terms is 'customer terms or vendor\items', line terms will use vendor\item terms. If 'N', use customer terms always for this customer.
	RmaRevenueAccountNo        *string   `bun:"rma_revenue_account_no,type:varchar(32)"`                    // Default value for rma  account
	DealerWrrtyClaimsAccountNo *string   `bun:"dealer_wrrty_claims_account_no,type:varchar(32)"`            // Default value for dealer warranty claims account
	ArBatchType                string    `bun:"ar_batch_type,type:char(1),default:('C')"`                   // Custom column which defines whether to use the customer invoice batch or a system default invoice batch in Order Entry.
	ConsBackordersFlag         *string   `bun:"cons_backorders_flag,type:char(1)"`                          // Indicates whether backorders can be consolidated for this customer
	TaxGroupId                 *string   `bun:"tax_group_id,type:varchar(10)"`                              // VAT Tax Group ID
	CreditStatus               *string   `bun:"credit_status,type:varchar(8)"`                              // Credit Status default for new customers
	CurrencyId                 float64   `bun:"currency_id,type:decimal(19,0)"`                             // Default Currency ID
	ApplyConvenienceFeeFlag    *string   `bun:"apply_convenience_fee_flag,type:char(1)"`                    // Custom: Flag to indicate if a customer should be charged a convenience fee when making a payment with a credit card.
	ApplyFuelSurchargesCd      *int32    `bun:"apply_fuel_surcharges_cd,type:int"`                          // A custom column contains code of Yes (1103), No (1104), or Rep (2338). If its Yes, surcharges will be calculated based on info in class maint. If its No, no surcharge will be applied. If Rep, a surcharge will be calculated based on Sales Reps info.
	AppliedFuelchargesToDsFlag *string   `bun:"applied_fuelcharges_to_ds_flag,type:char(1)"`                // A custom column which indicates whether a surcharge will be calculated and applied to invoices via Direct Ship Confirmation.
	SuppressZeroDollarFlag     *string   `bun:"suppress_zero_dollar_flag,type:char(1)"`                     // A custom column which indicates whether the system will suppress zero dollar fuel surcharge lines from appearing on any forms.
	Class5Id                   *string   `bun:"class_5id,type:varchar(255)"`                                // A user-defined code that identifies a group of customers.
}
