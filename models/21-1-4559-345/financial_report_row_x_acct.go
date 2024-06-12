package model

type FinancialReportRowXAcct struct {
	bun.BaseModel         `bun:"table:financial_report_row_x_acct"`
	FinReportRowXAcctUid  int32     `bun:"fin_report_row_x_acct_uid,type:int,pk,identity"`
	FinancialReportRowUid int32     `bun:"financial_report_row_uid,type:int"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	CompanyId             string    `bun:"company_id,type:varchar(8)"`
	AccountNo             string    `bun:"account_no,type:varchar(32)"`
}
