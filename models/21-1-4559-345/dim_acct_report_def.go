package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DimAcctReportDef struct {
	bun.BaseModel       `bun:"table:dim_acct_report_def"`
	DimAcctReportDefUid int32     `bun:"dim_acct_report_def_uid,type:int,pk,identity"`
	ReportTypeCd        int32     `bun:"report_type_cd,type:int"`
	ReportDescription   string    `bun:"report_description,type:varchar(255)"`
	ViewName            string    `bun:"view_name,type:varchar(255)"`
	RowStatusFlag       int32     `bun:"row_status_flag,type:int"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	FunctionName        string    `bun:"function_name,type:varchar(255),nullzero"`
}
