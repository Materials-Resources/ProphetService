package model

type VendorPurchaseAcct struct {
	bun.BaseModel             `bun:"table:vendor_purchase_acct"`
	CompanyId                 string    `bun:"company_id,type:varchar(8),pk"`
	VendorId                  float64   `bun:"vendor_id,type:decimal(18,0),pk"`
	PurchaseAccountNo         string    `bun:"purchase_account_no,type:varchar(80),pk"`
	DeleteFlag                string    `bun:"delete_flag,type:char"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	CoreAcctFlag              string    `bun:"core_acct_flag,type:char,nullzero"`
	FreightEstimatedAcctFlag  string    `bun:"freight_estimated_acct_flag,type:char,nullzero"`
	FreightDifferenceAcctFlag string    `bun:"freight_difference_acct_flag,type:char,nullzero"`
	BranchId                  string    `bun:"branch_id,type:varchar(8),nullzero"`
	AllocationPercent         float64   `bun:"allocation_percent,type:decimal(19,9),nullzero"`
	PurchaseDesc              string    `bun:"purchase_desc,type:varchar(255),nullzero"`
	PurchaseCompanyId         string    `bun:"purchase_company_id,type:varchar(8),nullzero"`
}
