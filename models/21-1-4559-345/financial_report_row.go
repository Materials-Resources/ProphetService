package model

type FinancialReportRow struct {
	bun.BaseModel         `bun:"table:financial_report_row"`
	FinancialReportRowUid int32     `bun:"financial_report_row_uid,type:int,pk"`
	SpreadsheetRowNo      int32     `bun:"spreadsheet_row_no,type:int"`
	FinReportUid          int32     `bun:"fin_report_uid,type:int"`
	RowDescription        string    `bun:"row_description,type:varchar(255),nullzero"`
	ReverseSignFlag       string    `bun:"reverse_sign_flag,type:char,default:('N')"`
	DeleteFlag            string    `bun:"delete_flag,type:char,default:('N')"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
