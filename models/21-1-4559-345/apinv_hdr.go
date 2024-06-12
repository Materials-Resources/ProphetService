package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ApinvHdr struct {
	bun.BaseModel              `bun:"table:apinv_hdr"`
	CompanyNo                  string    `bun:"company_no,type:varchar(8)"`
	VendorId                   float64   `bun:"vendor_id,type:decimal(19,0)"`
	InvoiceNo                  string    `bun:"invoice_no,type:varchar(32)"`
	InvoiceDate                time.Time `bun:"invoice_date,type:datetime,nullzero"`
	InvoiceAmount              float64   `bun:"invoice_amount,type:decimal(19,4)"`
	NetDueDate                 time.Time `bun:"net_due_date,type:datetime,nullzero"`
	TermsDueDate               time.Time `bun:"terms_due_date,type:datetime,nullzero"`
	TermsAmount                float64   `bun:"terms_amount,type:decimal(19,4)"`
	CheckNo                    string    `bun:"check_no,type:varchar(255),nullzero"`
	CheckDate                  time.Time `bun:"check_date,type:datetime,nullzero"`
	VoucherNo                  string    `bun:"voucher_no,type:varchar(10),pk"`
	ApAccountNo                string    `bun:"ap_account_no,type:varchar(32)"`
	CashAccountNo              string    `bun:"cash_account_no,type:varchar(32),nullzero"`
	Description                string    `bun:"description,type:varchar(30),nullzero"`
	AmountPaid                 float64   `bun:"amount_paid,type:decimal(19,4)"`
	TermsAmountTaken           float64   `bun:"terms_amount_taken,type:decimal(19,4)"`
	PaidInFull                 string    `bun:"paid_in_full,type:char"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	PoNo                       string    `bun:"po_no,type:varchar(50),nullzero"`
	Period                     float64   `bun:"period,type:decimal(3,0),nullzero"`
	YearForPeriod              float64   `bun:"year_for_period,type:decimal(4,0),nullzero"`
	CurrencyId                 float64   `bun:"currency_id,type:decimal(19,0),nullzero"`
	HomeCurrencyAmt            float64   `bun:"home_currency_amt,type:decimal(19,4),nullzero"`
	ReverseFlag                string    `bun:"reverse_flag,type:char,nullzero"`
	VoucherType                string    `bun:"voucher_type,type:char,nullzero"`
	VoucherDescription         string    `bun:"voucher_description,type:varchar(255),nullzero"`
	MemoAmount                 float64   `bun:"memo_amount,type:decimal(19,4),nullzero"`
	VoucherClass               string    `bun:"voucher_class,type:varchar(8),nullzero"`
	VoucherReferenceNumber     string    `bun:"voucher_reference_number,type:varchar(10),nullzero"`
	PeriodFullyPaid            float64   `bun:"period_fully_paid,type:decimal(3,0),nullzero"`
	YearFullyPaid              float64   `bun:"year_fully_paid,type:decimal(4,0),nullzero"`
	Approved                   string    `bun:"approved,type:char,nullzero"`
	BranchId                   string    `bun:"branch_id,type:varchar(8),nullzero"`
	TermsTakenAccount          string    `bun:"terms_taken_account,type:varchar(32),nullzero"`
	InvoiceAmountDisplay       float64   `bun:"invoice_amount_display,type:decimal(19,4),nullzero"`
	TermsAmountDisplay         float64   `bun:"terms_amount_display,type:decimal(19,4),nullzero"`
	TermsAmountTakenDisplay    float64   `bun:"terms_amount_taken_display,type:decimal(19,4),nullzero"`
	AmountPaidDisplay          float64   `bun:"amount_paid_display,type:decimal(19,4),nullzero"`
	MemoAmountDisplay          float64   `bun:"memo_amount_display,type:decimal(19,4),nullzero"`
	PayableSourceCd            int32     `bun:"payable_source_cd,type:int,nullzero"`
	TermsId                    string    `bun:"terms_id,type:varchar(2),nullzero"`
	SourceTypeCd               int32     `bun:"source_type_cd,type:int,nullzero"`
	AlwaysTakeTerms            string    `bun:"always_take_terms,type:char,default:('N')"`
	VoucherReversalReason      string    `bun:"voucher_reversal_reason,type:varchar(255),nullzero"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DisputedFlag               string    `bun:"disputed_flag,type:char,default:('N')"`
	PreviouslyDisputedFlag     string    `bun:"previously_disputed_flag,type:char,default:('N')"`
	JobId                      string    `bun:"job_id,type:varchar(32),nullzero"`
	JobIdSetByConversionFlag   string    `bun:"job_id_set_by_conversion_flag,type:char,default:('N')"`
	VendorEditedFlag           string    `bun:"vendor_edited_flag,type:char,default:('N')"`
	CustomerId                 float64   `bun:"customer_id,type:decimal(19,0),nullzero"`
	ShipToId                   float64   `bun:"ship_to_id,type:decimal(19,0),nullzero"`
	HandlingFee                float64   `bun:"handling_fee,type:decimal(19,9),nullzero"`
	VoucherRefInvNo            string    `bun:"voucher_ref_inv_no,type:varchar(10),nullzero"`
	ExportedFlag               string    `bun:"exported_flag,type:char,nullzero"`
	FreightVoucherFlag         string    `bun:"freight_voucher_flag,type:char,nullzero"`
	FreightAmount              float64   `bun:"freight_amount,type:decimal(19,9),nullzero"`
	FreightAmountDisplay       float64   `bun:"freight_amount_display,type:decimal(19,9),nullzero"`
	ExchangeRate               float64   `bun:"exchange_rate,type:decimal(19,6),nullzero"`
	FreightDisputeFlag         string    `bun:"freight_dispute_flag,type:char,default:('N')"`
	ExternalClaimId            string    `bun:"external_claim_id,type:varchar(255),nullzero"`
	IvaWithheldAmount          float64   `bun:"iva_withheld_amount,type:decimal(19,4),default:((0))"`
	IvaWithheldAmountDisplay   float64   `bun:"iva_withheld_amount_display,type:decimal(19,4),default:((0))"`
	IvaTermsAmountTaken        float64   `bun:"iva_terms_amount_taken,type:decimal(19,4),default:((0))"`
	IvaTermsAmountTakenDisplay float64   `bun:"iva_terms_amount_taken_display,type:decimal(19,4),default:((0))"`
	VoucherReasonCodeUid       int32     `bun:"voucher_reason_code_uid,type:int,nullzero"`
	PayableSourceUid           int32     `bun:"payable_source_uid,type:int,nullzero"`
	PaymentApprovedFlag        string    `bun:"payment_approved_flag,type:char,nullzero"`
	TermsDiscountPct           float64   `bun:"terms_discount_pct,type:decimal(19,4),nullzero"`
	Notes                      string    `bun:"notes,type:text(2147483647),nullzero"`
	PrintNotes                 string    `bun:"print_notes,type:char,nullzero"`
	WarrantyClaimPaymentsUid   int32     `bun:"warranty_claim_payments_uid,type:int,nullzero"`
	VatOnlyMemoFlag            string    `bun:"vat_only_memo_flag,type:char,default:('N')"`
}
