package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type FcDataobjectTable struct {
	bun.BaseModel        `bun:"table:fc_dataobject_table"`
	FcDataobjectTableUid int32     `bun:"fc_dataobject_table_uid,type:int,pk"`                          // Unique Identifier for fc_dataobject_table record.
	FcDataobjectUid      int32     `bun:"fc_dataobject_uid,type:int"`                                   // Identified of the associated fc_dataobject record.
	TableName            string    `bun:"table_name,type:varchar(255)"`                                 // The database table associated with a field chooser modification to a datawindow.
	JoinSyntax           *string   `bun:"join_syntax,type:varchar(8000)"`                               // The JOIN syntax for a modification to a field chooser datawindow.
	RowStatusFlag        int32     `bun:"row_status_flag,type:int"`                                     // Status flag indicating whether the record is active or not.
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
