package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ApinvHdr struct {
	bun.BaseModel              `bun:"table:apinv_hdr"`
	CompanyNo                  string    `bun:"company_no,type:varchar(8)"`                                // Unique code that identifies a company.
	VendorId                   float64   `bun:"vendor_id,type:decimal(19,0)"`                              // What is the unique vendor identifier for this row?
	InvoiceNo                  string    `bun:"invoice_no,type:varchar(32)"`                               // What invoice number is this payment detail for?
	InvoiceDate                time.Time `bun:"invoice_date,type:datetime,nullzero"`                       // Date on the voucher
	InvoiceAmount              float64   `bun:"invoice_amount,type:decimal(19,4)"`                         // Amount voucher is being created for
	NetDueDate                 time.Time `bun:"net_due_date,type:datetime,nullzero"`                       // This is the net due date
	TermsDueDate               time.Time `bun:"terms_due_date,type:datetime,nullzero"`                     // This is the terms due date
	TermsAmount                float64   `bun:"terms_amount,type:decimal(19,4)"`                           // This is the terms amount
	CheckNo                    string    `bun:"check_no,type:varchar(255),nullzero"`                       // This is the check number
	CheckDate                  time.Time `bun:"check_date,type:datetime,nullzero"`                         // What is the date of the payment -  as marked on the
	VoucherNo                  string    `bun:"voucher_no,type:varchar(10),pk"`                            // Unique, system_generated number
	ApAccountNo                string    `bun:"ap_account_no,type:varchar(32)"`                            // What is the accounts payable account number for this voucher
	CashAccountNo              string    `bun:"cash_account_no,type:varchar(32),nullzero"`                 // Enter the cash account number
	Description                string    `bun:"description,type:varchar(30),nullzero"`                     // How would you describe this repeating journal entry?
	AmountPaid                 float64   `bun:"amount_paid,type:decimal(19,4)"`                            // What is the actual amount paid?
	TermsAmountTaken           float64   `bun:"terms_amount_taken,type:decimal(19,4)"`                     // Enter the terms amount taken
	PaidInFull                 string    `bun:"paid_in_full,type:char(1)"`                                 // Indicate whether or not to pay the invoice in full
	DateCreated                time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	PoNo                       string    `bun:"po_no,type:varchar(50),nullzero"`                           // Purchase Order associated with this voucher, if any
	Period                     float64   `bun:"period,type:decimal(3,0),nullzero"`                         // Which period does this quota apply to?
	YearForPeriod              float64   `bun:"year_for_period,type:decimal(4,0),nullzero"`                // Enter a year
	CurrencyId                 float64   `bun:"currency_id,type:decimal(19,0),nullzero"`                   // What is the unique currency identifier for this ro
	HomeCurrencyAmt            float64   `bun:"home_currency_amt,type:decimal(19,4),nullzero"`             // Please Enter The Home Currency Amount.
	ReverseFlag                string    `bun:"reverse_flag,type:char(1),nullzero"`                        // Was this voucher reversed?
	VoucherType                string    `bun:"voucher_type,type:char(1),nullzero"`                        // Enter the voucher type
	VoucherDescription         string    `bun:"voucher_description,type:varchar(255),nullzero"`            // This appears on the AP Credit Debit Memo Entry win
	MemoAmount                 float64   `bun:"memo_amount,type:decimal(19,4),nullzero"`                   // What is the memo amount for this payment detail?
	VoucherClass               string    `bun:"voucher_class,type:varchar(8),nullzero"`                    // Used to seperate vouchers into various categories
	VoucherReferenceNumber     string    `bun:"voucher_reference_number,type:varchar(10),nullzero"`        // This is the voucher reference number
	PeriodFullyPaid            float64   `bun:"period_fully_paid,type:decimal(3,0),nullzero"`              // In what period was this voucher is fully paid?
	YearFullyPaid              float64   `bun:"year_fully_paid,type:decimal(4,0),nullzero"`                // In which year was this invoice considered fully paid?
	Approved                   string    `bun:"approved,type:char(1),nullzero"`                            // Has this product order been approved?
	BranchId                   string    `bun:"branch_id,type:varchar(8),nullzero"`                        // What branch issued this Purchase Order?
	TermsTakenAccount          string    `bun:"terms_taken_account,type:varchar(32),nullzero"`             // GL account used for posting terms amounts
	InvoiceAmountDisplay       float64   `bun:"invoice_amount_display,type:decimal(19,4),nullzero"`        // Holds the invoice amount displayed on the window.
	TermsAmountDisplay         float64   `bun:"terms_amount_display,type:decimal(19,4),nullzero"`          // Terms amount in Foreign Currency (will be home currency if not using Foreign Currency)
	TermsAmountTakenDisplay    float64   `bun:"terms_amount_taken_display,type:decimal(19,4),nullzero"`    // Terms amount that was taken in Foreign Currency (will be home currency if not using Foreign Currency)
	AmountPaidDisplay          float64   `bun:"amount_paid_display,type:decimal(19,4),nullzero"`           // Holds the values that are displayed on the window.
	MemoAmountDisplay          float64   `bun:"memo_amount_display,type:decimal(19,4),nullzero"`           // Holds the values that are displayed on the window.
	PayableSourceCd            int32     `bun:"payable_source_cd,type:int,nullzero"`                       // Indicates where a payable record was created (Return or other)
	TermsId                    string    `bun:"terms_id,type:varchar(2),nullzero"`                         // Terms associated with this voucher
	SourceTypeCd               int32     `bun:"source_type_cd,type:int,nullzero"`                          // Code (from code_p21 table) which indicates where the voucher was created
	AlwaysTakeTerms            string    `bun:"always_take_terms,type:char(1),default:('N')"`              // This column indicates whether terms on the voucher should automatically calculate even when its paid past the terms due date.
	VoucherReversalReason      string    `bun:"voucher_reversal_reason,type:varchar(255),nullzero"`        // Reason for reversing the voucher.
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DisputedFlag               string    `bun:"disputed_flag,type:char(1),default:('N')"`                        // Indicates whether the voucher is in dispute status.
	PreviouslyDisputedFlag     string    `bun:"previously_disputed_flag,type:char(1),default:('N')"`             // Indicates whether the voucher was once under dispute
	JobId                      string    `bun:"job_id,type:varchar(32),nullzero"`                                // Job ID that is associated with this voucher header.
	JobIdSetByConversionFlag   string    `bun:"job_id_set_by_conversion_flag,type:char(1),default:('N')"`        // Tells whether or not the job_id field was set by a conversion to voucher process.
	VendorEditedFlag           string    `bun:"vendor_edited_flag,type:char(1),default:('N')"`                   // This column indicates whether or not the vendor has been edited before saving the voucher
	CustomerId                 float64   `bun:"customer_id,type:decimal(19,0),nullzero"`                         // Custom: indicates the customer to which the AR invoice is created.
	ShipToId                   float64   `bun:"ship_to_id,type:decimal(19,0),nullzero"`                          // Custom: indicates the shio to to which the AR invoice is created.
	HandlingFee                float64   `bun:"handling_fee,type:decimal(19,9),nullzero"`                        // Custom: voucher amount multiplied by Mark-Up percentage.
	VoucherRefInvNo            string    `bun:"voucher_ref_inv_no,type:varchar(10),nullzero"`                    // Custom: reference to AR invoice
	ExportedFlag               string    `bun:"exported_flag,type:char(1),nullzero"`                             // Determines whether or not this record has been exported.
	FreightVoucherFlag         string    `bun:"freight_voucher_flag,type:char(1),nullzero"`                      // Indicates that a voucher is for a freight bill.
	FreightAmount              float64   `bun:"freight_amount,type:decimal(19,9),nullzero"`                      // Freight amount for the voucher (Home)
	FreightAmountDisplay       float64   `bun:"freight_amount_display,type:decimal(19,9),nullzero"`              // Freight amount for the voucher (Foreign)
	ExchangeRate               float64   `bun:"exchange_rate,type:decimal(19,6),nullzero"`                       // Stores the exchange rate used.
	FreightDisputeFlag         string    `bun:"freight_dispute_flag,type:char(1),default:('N')"`                 // Identifies when a Voucher has been set in Dispute due to Freight Variance.
	ExternalClaimId            string    `bun:"external_claim_id,type:varchar(255),nullzero"`                    // This field should also be available in the Credit/Dedit memo
	IvaWithheldAmount          float64   `bun:"iva_withheld_amount,type:decimal(19,4),default:((0))"`            // IVA Withheld Amount
	IvaWithheldAmountDisplay   float64   `bun:"iva_withheld_amount_display,type:decimal(19,4),default:((0))"`    // IVA Withheld Amount (Foreign Currency)
	IvaTermsAmountTaken        float64   `bun:"iva_terms_amount_taken,type:decimal(19,4),default:((0))"`         // IVA Terms Amount
	IvaTermsAmountTakenDisplay float64   `bun:"iva_terms_amount_taken_display,type:decimal(19,4),default:((0))"` // IVA Terms Amount (Foreign Currency)
	VoucherReasonCodeUid       int32     `bun:"voucher_reason_code_uid,type:int,nullzero"`                       // Reason code for a voucher
	PayableSourceUid           int32     `bun:"payable_source_uid,type:int,nullzero"`                            // Transaction Associated with the Voucher (Landed Cost Driver, Inventory Receipt, etc...)
	PaymentApprovedFlag        string    `bun:"payment_approved_flag,type:char(1),nullzero"`                     // Payment approved by president flag.
	TermsDiscountPct           float64   `bun:"terms_discount_pct,type:decimal(19,4),nullzero"`                  // Percentage of terms discount
	Notes                      string    `bun:"notes,type:text,nullzero"`                                        // Memo notes for vouchers on associated checks.
	PrintNotes                 string    `bun:"print_notes,type:char(1),nullzero"`                               // Flag that indicates if notes would be printed on the checks.
	WarrantyClaimPaymentsUid   int32     `bun:"warranty_claim_payments_uid,type:int,nullzero"`                   // Custom column used to link the voucher to warranty claim payment
	VatOnlyMemoFlag            string    `bun:"vat_only_memo_flag,type:char(1),default:('N')"`                   // Indicates that this debit/credit memo contains only VAT charges.
}
