package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DimAcctReportDef struct {
	bun.BaseModel       `bun:"table:dim_acct_report_def"`
	DimAcctReportDefUid int32     `bun:"dim_acct_report_def_uid,type:int,autoincrement,pk"`            // Table UID
	ReportTypeCd        int32     `bun:"report_type_cd,type:int,unique"`                               // Report Type Code
	ReportDescription   string    `bun:"report_description,type:varchar(255)"`                         // Report Description
	ViewName            string    `bun:"view_name,type:varchar(255),unique"`                           // Report View Name
	RowStatusFlag       int32     `bun:"row_status_flag,type:int"`                                     // Records Starus Code
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	FunctionName        string    `bun:"function_name,type:varchar(255),nullzero"`                     // Report Function Name
}
