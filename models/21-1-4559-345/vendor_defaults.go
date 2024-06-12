package model

type VendorDefaults struct {
	bun.BaseModel    `bun:"table:vendor_defaults"`
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`
	ApAccountNumber  string    `bun:"ap_account_number,type:varchar(32)"`
	TermsId          string    `bun:"terms_id,type:varchar(2)"`
	EdiOrPaper       string    `bun:"edi_or_paper,type:char"`
	CurrencyId       float64   `bun:"currency_id,type:decimal(19,0)"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
