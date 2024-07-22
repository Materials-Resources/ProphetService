package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ReportEmailSubjectXToken struct {
	bun.BaseModel               `bun:"table:report_email_subject_x_token"`
	ReportEmailSubjectXTokenUid int32     `bun:"report_email_subject_x_token_uid,type:int,autoincrement,identity,pk"` // Identity
	TokenUid                    int32     `bun:"token_uid,type:int"`                                                  // Token
	AvailableAreaCode           int32     `bun:"available_area_code,type:int"`                                        // Area Code where token is available
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`                      // Date and time the record was originally created
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`                // User who created the record
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`                // Date and time the record was modified
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`        // User who last changed the record
}
