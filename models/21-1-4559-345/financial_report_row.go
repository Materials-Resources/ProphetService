package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type FinancialReportRow struct {
	bun.BaseModel         `bun:"table:financial_report_row"`
	FinancialReportRowUid int32     `bun:"financial_report_row_uid,type:int,pk"`                         // Unique ID for financial report row
	SpreadsheetRowNo      int32     `bun:"spreadsheet_row_no,type:int,unique"`                           // Row number on the spreadsheet of this record
	FinReportUid          int32     `bun:"fin_report_uid,type:int,unique"`                               // UID for the fin_report_table
	RowDescription        string    `bun:"row_description,type:varchar(255),nullzero"`                   // Description for the row on this record
	ReverseSignFlag       string    `bun:"reverse_sign_flag,type:char(1),default:('N')"`                 // Determine whether to reverse the sign of the amount
	DeleteFlag            string    `bun:"delete_flag,type:char(1),default:('N')"`                       // Delete flag for this record
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
