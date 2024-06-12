package model

type ArReceiptsDetail struct {
	bun.BaseModel            `bun:"table:ar_receipts_detail"`
	ReceiptNumber            float64   `bun:"receipt_number,type:decimal(19,0),pk"`
	CustomerId               float64   `bun:"customer_id,type:decimal(19,0)"`
	InvoiceNo                string    `bun:"invoice_no,type:varchar(10),pk"`
	CompanyId                string    `bun:"company_id,type:varchar(8)"`
	PaymentAmount            float64   `bun:"payment_amount,type:decimal(19,4)"`
	TermsAmount              float64   `bun:"terms_amount,type:decimal(19,4)"`
	AllowedAmount            float64   `bun:"allowed_amount,type:decimal(19,4)"`
	DeleteFlag               string    `bun:"delete_flag,type:char"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	CurrencyVarianceAmtHome  float64   `bun:"currency_variance_amt_home,type:decimal(19,2),default:(0)"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	TaxTermsTakenAmt         float64   `bun:"tax_terms_taken_amt,type:decimal(19,9),nullzero"`
	OtherChargeTermsTakenAmt float64   `bun:"other_charge_terms_taken_amt,type:decimal(19,9),nullzero"`
	TaxAmountPaid            float64   `bun:"tax_amount_paid,type:decimal(19,2),default:((0))"`
}
