package model

type RemittancesDetail struct {
	bun.BaseModel           `bun:"table:remittances_detail"`
	RemittanceNumber        float64   `bun:"remittance_number,type:decimal(19,0),pk"`
	InvoiceNo               string    `bun:"invoice_no,type:varchar(10),pk"`
	PaymentAmount           float64   `bun:"payment_amount,type:decimal(19,4)"`
	TermsAmount             float64   `bun:"terms_amount,type:decimal(19,4)"`
	AllowedAmount           float64   `bun:"allowed_amount,type:decimal(19,4)"`
	DeleteFlag              string    `bun:"delete_flag,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	CurrencyVarianceAmtHome float64   `bun:"currency_variance_amt_home,type:decimal(19,2),default:(0)"`
	TaxAmountPaid           float64   `bun:"tax_amount_paid,type:decimal(19,2),default:((0))"`
}
