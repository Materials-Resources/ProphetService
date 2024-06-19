package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ClassNameXFormCode struct {
	bun.BaseModel         `bun:"table:class_name_x_form_code"`
	ClassNameXFormCodeUid int32     `bun:"class_name_x_form_code_uid,type:int,autoincrement,scanonly,pk"` // identity & auto-incremented column
	ClassName             string    `bun:"class_name,type:varchar(255)"`                                  //  Window Class Name
	FormCode              int32     `bun:"form_code,type:int"`                                            // Form code
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`                // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`          // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`          // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`  // User who last changed the record
}
