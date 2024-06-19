package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ReportHdr struct {
	bun.BaseModel        `bun:"table:report_hdr"`
	ReportHdrUid         int32     `bun:"report_hdr_uid,type:int,autoincrement,pk"`                     // identity
	ReportName           string    `bun:"report_name,type:varchar(255)"`                                // rpt file name of crystal report
	ReportSourcePath     string    `bun:"report_source_path,type:varchar(255)"`                         // fully qualifier path to report file
	PostPurgeFlag        string    `bun:"post_purge_flag,type:char(1)"`                                 // flag to indicate purge of parm data after report is executed
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	NumberOfReportCopies int32     `bun:"number_of_report_copies,type:int,default:((1))"`               // Number of copies to print for this report run
	ReportCollate        int32     `bun:"report_collate,type:int,default:((1))"`                        // Whether or not to collate from printer select
}
