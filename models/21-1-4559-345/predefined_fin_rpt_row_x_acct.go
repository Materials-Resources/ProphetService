package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PredefinedFinRptRowXAcct struct {
	bun.BaseModel           `bun:"table:predefined_fin_rpt_row_x_acct"`
	PredefFinRptRowXAcctUid int32 `bun:"predef_fin_rpt_row_x_acct_uid,type:int,autoincrement,pk"` // Unique identifier of table (Identity Column)
	StatementTypeCd         int32 `bun:"statement_type_cd,type:int"`                              /*
		Financial Statement which the records belongs (Balance Sheet,
		Profit & Loss, etc...)
	*/
	SpreadsheetRowNo int32 `bun:"spreadsheet_row_no,type:int"` /*
		Row of financial statement which the records are associated
		with
	*/
	AccountNumber    string    `bun:"account_number,type:varchar(10)"`                              // Account Number
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
