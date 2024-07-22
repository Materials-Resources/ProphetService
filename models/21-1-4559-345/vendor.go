package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Vendor struct {
	bun.BaseModel              `bun:"table:vendor"`
	VendorId                   float64    `bun:"vendor_id,type:decimal(19,0),pk"`                           // What is the unique vendor identifier for this row?
	CompanyId                  string     `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	Default1099Type            float64    `bun:"default_1099_type,type:decimal(3,0)"`                       // Vendor 1099 type if applicable.
	ApAccountNo                *string    `bun:"ap_account_no,type:varchar(32)"`                            // AP account number
	DefaultPurchAcctNo         *string    `bun:"default_purch_acct_no,type:varchar(32)"`                    // Purchasing Account
	DefaultTermsId             *string    `bun:"default_terms_id,type:varchar(2)"`                          // Default terms for the vendor.
	Class1Id                   *string    `bun:"class_1id,type:varchar(255)"`                               // What is the contac class?
	Class2Id                   *string    `bun:"class_2id,type:varchar(255)"`                               // What is the class for this order?
	Class3Id                   *string    `bun:"class_3id,type:varchar(255)"`                               // What is the customer class?
	Class4Id                   *string    `bun:"class_4id,type:varchar(255)"`                               // What is the class for this order?
	Class5Id                   *string    `bun:"class_5id,type:varchar(255)"`                               // What is the contac class?
	DeleteFlag                 string     `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated                time.Time  `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified           time.Time  `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy           string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	VendorIdString             string     `bun:"vendor_id_string,type:varchar(22)"`                         // The vendor_id column stored as a string.
	DefaultPayFreightTo        *string    `bun:"default_pay_freight_to,type:char(1)"`                       // Indicates whether pay freight to Carrier or Vendor
	CurrencyId                 float64    `bun:"currency_id,type:decimal(3,0)"`                             // What is the unique currency identifier for this ro
	EdiOrPaper                 string     `bun:"edi_or_paper,type:char(1)"`                                 // This column is unused.
	SecurityInfo               *string    `bun:"security_info,type:varchar(10)"`                            // Password for EDI
	InterchgReceiverId         *string    `bun:"interchg_receiver_id,type:varchar(15)"`                     // EDI Interchange Receiver ID
	IntlSan                    *string    `bun:"intl_san,type:varchar(15)"`                                 // EDI Standard Address Number
	FederalIdNumber            *string    `bun:"federal_id_number,type:varchar(15)"`                        // A number assigned to your company by the government.  This number is used for the IRS 1099 form and for various other payroll purposes.
	DefaultInvoiceDesc         *string    `bun:"default_invoice_desc,type:varchar(30)"`                     // The default description for invoices received from a particular vendor.
	PayableGroupId             *string    `bun:"payable_group_id,type:varchar(8)"`                          // Unique identifier for this payable group.
	TqmTracking                *string    `bun:"tqm_tracking,type:char(1)"`                                 // This column is unused.
	AlwaysTakeTerms            string     `bun:"always_take_terms,type:char(1),default:('N')"`              // When Y, indicates terms should always be taken.
	JobIdRequired              *string    `bun:"job_id_required,type:char(1)"`                              // When Y, indicates that a job id should be specified in voucher entry.
	WorkersCompExpirationDate  *time.Time `bun:"workers_comp_expiration_date,type:datetime"`                // This column is unused.
	GeneralLiabExpirationDate  *time.Time `bun:"general_liab_expiration_date,type:datetime"`                // This column is unused.
	VendorName                 *string    `bun:"vendor_name,type:varchar(255)"`                             // The vendor name corresponding to the vendor ID.
	TradingPartnerName         *string    `bun:"trading_partner_name,type:varchar(255)"`                    // EDI - Name used to identify the trading partner.
	TrackRebates               string     `bun:"track_rebates,type:char(1),default:('N')"`                  // Indicates whether the vendor is enabled for rebate tracking
	RebateAccountNo            *string    `bun:"rebate_account_no,type:varchar(32)"`                        // Vendor rebate account
	RebateAllowanceAccountNo   *string    `bun:"rebate_allowance_account_no,type:varchar(32)"`              // Vendor rebate allowed account
	PurchaseAcctGroupHdrUid    *int32     `bun:"purchase_acct_group_hdr_uid,type:int"`                      // This is used so users can assign a purchase acct group to a vendor.
	CreatedBy                  *string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	AllowRebatePayImpVariance  string     `bun:"allow_rebate_pay_imp_variance,type:char(1),default:('Y')"`  // Determines if imported rebate payments with a variance should be accepted.
	RebatePayImpVarianceType   int32      `bun:"rebate_pay_imp_variance_type,type:int,default:(1612)"`      // Determines how to apply a variance for an imported rebate payment.  Values are defined in the code_p21 table.
	LegacyId                   *string    `bun:"legacy_id,type:varchar(255)"`                               // column to hold vendor ID from legacy app
	WildcardAcctInVoucherEntry string     `bun:"wildcard_acct_in_voucher_entry,type:char(1),default:('N')"` // flag to determine whether or not to wildcard branch in voucher entry
	AttorneyFeeFlag            string     `bun:"attorney_fee_flag,type:char(1),default:('N')"`              // Attorneys Fees option for 1099 processing
	Irs1099State1              *string    `bun:"IRS_1099_state_1,type:varchar(255)"`                        // State 1 (abbreviation) for 1099 form
	Irs1099State1IdNo          *string    `bun:"IRS_1099_state_1_id_no,type:varchar(255)"`                  // State 1 identifictation number
	Irs1099State2              *string    `bun:"IRS_1099_state_2,type:varchar(255)"`                        // State 2 (abbreviation) for 1099 form
	Irs1099State2IdNo          *string    `bun:"IRS_1099_state_2_id_no,type:varchar(255)"`                  // State 2 identifictation number
	TaxpayerIdNoType           int32      `bun:"taxpayer_id_no_type,type:int,default:(300)"`                // Taxpayer ID number type.  Column populated w/values from the code_p21 table.
	WarrantyReceivableAcct     *string    `bun:"warranty_receivable_acct,type:varchar(32)"`                 // Waranty receivable account; used for warranty claims
	WarrantyRevenueAcct        *string    `bun:"warranty_revenue_acct,type:varchar(32)"`                    // Waranty revenue account; used for warranty claims
	WarrantyAllowanceAcct      *string    `bun:"warranty_allowance_acct,type:varchar(32)"`                  // Waranty allowance account; used for warranty claims
	CommissionReceivableAcct   *string    `bun:"commission_receivable_acct,type:varchar(32)"`               // The commission receivable account when manfucturer rep is enabled.
	CommissionRevenueAcct      *string    `bun:"commission_revenue_acct,type:varchar(32)"`                  // The commission revenue account when manfucturer rep is enabled.
	CommissionAllowanceAcct    *string    `bun:"commission_allowance_acct,type:varchar(32)"`                // The commission allowance account when manfucturer rep is enabled.
	TrackCoresFlag             *string    `bun:"track_cores_flag,type:char(1)"`                             // This custom column indicates if the vendor tracks cores or not
	BankCoresFlag              *string    `bun:"bank_cores_flag,type:char(1)"`                              // This custom column indicates if the vendor banks core items, inidcating if the vendor keeps a record of how many cores have been returned and purchased by the distributor..
	IsrTaxWithheldFlag         *string    `bun:"isr_tax_withheld_flag,type:char(1)"`                        // Custom (F30860) - determines whether ISR tax is withheld for this vendor
	IsrTaxPct                  *float64   `bun:"isr_tax_pct,type:decimal(19,9)"`                            // Custom (F30860): ISR tax percentage
	IsrTaxAcct                 *string    `bun:"isr_tax_acct,type:varchar(32)"`                             // Custom (F30860): ISR tax account
	VendorOnHoldFlag           *string    `bun:"vendor_on_hold_flag,type:char(1)"`                          // Indicates whether or not this vendor is on hold
	InvoiceLineVarianceAmt     *float64   `bun:"invoice_line_variance_amt,type:decimal(19,9)"`              // Column to store numeric field to enter an acceptable variance value
	InvoiceLineVarianceType    *string    `bun:"invoice_line_variance_type,type:char(1)"`                   // Column to store drop down menu with choices of Amount or Percentage. A: Amount, P: Percentage
	MaxCustomerSalesDiscount   *float64   `bun:"max_customer_sales_discount,type:decimal(19,9)"`            // Maximum percentage discount that can be used when oe line discount percentage is being set to vendor\items terms.
	VendorTypeCd               int32      `bun:"vendor_type_cd,type:int,default:((2981))"`                  // Indicates whether vendor is of type: expense, inventory or both.
	SeparateChecksFlag         string     `bun:"separate_checks_flag,type:char(1),default:('N')"`           // Custom: Indicates whether this vendor will have a separate check printed for each invoice.
	CreditLimit                *float64   `bun:"credit_limit,type:decimal(19,4)"`                           // Credit Limit given by the Vendor
	AsbMainAccountNo           *string    `bun:"asb_main_account_no,type:varchar(255)"`                     // Custom column for automated short buy main account
}
