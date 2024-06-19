package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ReportParm struct {
	bun.BaseModel    `bun:"table:report_parm"`
	ReportParmUid    int32     `bun:"report_parm_uid,type:int,autoincrement,identity,pk"`           // identity
	ReportHdrUid     int32     `bun:"report_hdr_uid,type:int"`                                      // key to primary table
	ReportLevel      int32     `bun:"report_level,type:int"`                                        // level of parms in crystal report (0-primary, 1-1st subreport, ect)
	ParmSeqNo        int32     `bun:"parm_seq_no,type:int"`                                         // order of parms per level
	ParmName         string    `bun:"parm_name,type:varchar(255)"`                                  // actual name of crystal parm
	ParmValue        string    `bun:"parm_value,type:varchar(max)"`                                 // value to be supplied to crystal report
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	ParmReportName   string    `bun:"parm_report_name,type:varchar(255),default:('')"`              // reporting parameter report name
}
