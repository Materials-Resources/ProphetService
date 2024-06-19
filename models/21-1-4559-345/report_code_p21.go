package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ReportCodeP21 struct {
	bun.BaseModel    `bun:"table:report_code_p21"`
	ReportCodeP21Uid int32     `bun:"report_code_p21_uid,type:int,autoincrement,scanonly,pk"`       // Primary key
	CodeNo           int32     `bun:"code_no,type:int"`                                             // Code of the option
	LanguageId       string    `bun:"language_id,type:varchar(8)"`                                  // id of the lauges of the option
	CodeDescription  string    `bun:"code_description,type:varchar(255)"`                           // Description displayed for the user
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
