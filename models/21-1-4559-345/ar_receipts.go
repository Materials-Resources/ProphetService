package model

type ArReceipts struct {
	bun.BaseModel      `bun:"table:ar_receipts"`
	ReceiptNumber      float64   `bun:"receipt_number,type:decimal(19,0),pk"`
	PaymentNumber      float64   `bun:"payment_number,type:decimal(19,0)"`
	CompanyId          string    `bun:"company_id,type:varchar(8)"`
	RemitterId         float64   `bun:"remitter_id,type:decimal(19,0)"`
	DateReceived       time.Time `bun:"date_received,type:datetime"`
	BankNo             float64   `bun:"bank_no,type:decimal(19,0),nullzero"`
	GlBankAccountNo    string    `bun:"gl_bank_account_no,type:varchar(32),nullzero"`
	ArAccountNo        string    `bun:"ar_account_no,type:varchar(32),nullzero"`
	TermsAccountNo     string    `bun:"terms_account_no,type:varchar(32),nullzero"`
	AllowedAccountNo   string    `bun:"allowed_account_no,type:varchar(32),nullzero"`
	Approved           string    `bun:"approved,type:char"`
	DeleteFlag         string    `bun:"delete_flag,type:char"`
	DateCreated        time.Time `bun:"date_created,type:datetime"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	ReversePaymentFlag string    `bun:"reverse_payment_flag,type:char,default:('N')"`
	ClearedBank        string    `bun:"cleared_bank,type:char,default:('N')"`
	ClearedPeriod      int16     `bun:"cleared_period,type:smallint,nullzero"`
	ClearedYear        int16     `bun:"cleared_year,type:smallint,nullzero"`
}
