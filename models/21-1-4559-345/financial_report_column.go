package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type FinancialReportColumn struct {
	bun.BaseModel            `bun:"table:financial_report_column"`
	FinancialReportColumnUid int32     `bun:"financial_report_column_uid,type:int,autoincrement,identity,pk"` // Unique ID for financial report column
	SpreadsheetColumn        string    `bun:"spreadsheet_column,type:varchar(2),unique"`                      // The alpha-numeric value of the column for this record
	SpreadsheetColumnNo      int32     `bun:"spreadsheet_column_no,type:int"`                                 // The numeric value of the column for this record
	FinReportUid             int32     `bun:"fin_report_uid,type:int,unique"`                                 // The uid from the fin_report table
	ColumnsUid               int32     `bun:"columns_uid,type:int"`                                           // UID from the columns table
	AmountType               string    `bun:"amount_type,type:char(1)"`                                       // The value of the amount type for this record
	DeleteFlag               string    `bun:"delete_flag,type:char(1),default:('N')"`                         // Delete flag for this record
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`                 // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`           // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`           // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`   // User who last changed the record
}
