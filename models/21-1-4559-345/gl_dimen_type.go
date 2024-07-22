package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type GlDimenType struct {
	bun.BaseModel    `bun:"table:gl_dimen_type"`
	GlDimenTypeUid   int32     `bun:"gl_dimen_type_uid,type:int,autoincrement,identity,pk"`         // UID - identity
	GlDimenTypeId    string    `bun:"gl_dimen_type_id,type:varchar(255),unique"`                    // Type ID
	GlDimenTypeDesc  string    `bun:"gl_dimen_type_desc,type:varchar(255)"`                         // Description for Type
	RecordTypeCd     int32     `bun:"record_type_cd,type:int,default:((932))"`                      // System Defined or User Defined Code
	DeleteFlag       string    `bun:"delete_flag,type:char(1),default:('N')"`                       // Logical delete flag
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	UseValues        string    `bun:"use_values,type:char(1),default:('N')"`                        // Whether the GL Dimension Type uses values
}
