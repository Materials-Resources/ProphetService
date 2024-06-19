package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ApplicationResourceFile struct {
	bun.BaseModel              `bun:"table:application_resource_file"`
	ApplicationResourceFileUid int32     `bun:"application_resource_file_uid,type:int,autoincrement,scanonly,pk"` // Unique identifier for each record
	FileName                   string    `bun:"file_name,type:varchar(255)"`                                      // Name of the resource file
	FileType                   string    `bun:"file_type,type:varchar(255),default:('IMAGE')"`                    // Type of file (image or other)
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`                   // Date and time the record was originally created
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`             // User who created the record
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`             // Date and time the record was modified
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`     // User who last changed the record
}
