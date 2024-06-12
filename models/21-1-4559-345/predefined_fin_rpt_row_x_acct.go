package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PredefinedFinRptRowXAcct struct {
	bun.BaseModel           `bun:"table:predefined_fin_rpt_row_x_acct"`
	PredefFinRptRowXAcctUid int32     `bun:"predef_fin_rpt_row_x_acct_uid,type:int,pk,identity"`
	StatementTypeCd         int32     `bun:"statement_type_cd,type:int"`
	SpreadsheetRowNo        int32     `bun:"spreadsheet_row_no,type:int"`
	AccountNumber           string    `bun:"account_number,type:varchar(10)"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
