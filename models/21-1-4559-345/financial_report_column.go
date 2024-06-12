package model

import (
	"github.com/uptrace/bun"
	"time"
)

type FinancialReportColumn struct {
	bun.BaseModel            `bun:"table:financial_report_column"`
	FinancialReportColumnUid int32     `bun:"financial_report_column_uid,type:int,pk,identity"`
	SpreadsheetColumn        string    `bun:"spreadsheet_column,type:varchar(2)"`
	SpreadsheetColumnNo      int32     `bun:"spreadsheet_column_no,type:int"`
	FinReportUid             int32     `bun:"fin_report_uid,type:int"`
	ColumnsUid               int32     `bun:"columns_uid,type:int"`
	AmountType               string    `bun:"amount_type,type:char"`
	DeleteFlag               string    `bun:"delete_flag,type:char,default:('N')"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
