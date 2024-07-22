package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type TagHoldClass struct {
	bun.BaseModel    `bun:"table:tag_hold_class"`
	TagHoldClassUid  int32     `bun:"tag_hold_class_uid,type:int,pk"`                               // Unique ID for tag and hold class
	TagHoldClassId   string    `bun:"tag_hold_class_id,type:varchar(255),unique"`                   // Descriptive unique ID for tag and hold class
	TagHoldClassDesc string    `bun:"tag_hold_class_desc,type:varchar(255)"`                        // Description of tag and hold class
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`                                     // Row status flag
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
