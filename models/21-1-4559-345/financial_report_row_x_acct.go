package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type FinancialReportRowXAcct struct {
	bun.BaseModel         `bun:"table:financial_report_row_x_acct"`
	FinReportRowXAcctUid  int32     `bun:"fin_report_row_x_acct_uid,type:int,autoincrement,pk"`          // Unique ID for financial_report_row_x_acct record
	FinancialReportRowUid int32     `bun:"financial_report_row_uid,type:int,unique"`                     // UID from the financial_report_row table
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	CompanyId             string    `bun:"company_id,type:varchar(8),unique"`                            // Company id for account number
	AccountNo             string    `bun:"account_no,type:varchar(32),unique"`                           // Account associated with this record
}
