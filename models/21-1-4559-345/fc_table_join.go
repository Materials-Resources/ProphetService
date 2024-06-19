package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type FcTableJoin struct {
	bun.BaseModel    `bun:"table:fc_table_join"`
	FcTableJoinUid   int32     `bun:"fc_table_join_uid,type:int,autoincrement,identity"`            // Unique identifier for each record
	BaseTable        string    `bun:"base_table,type:varchar(255)"`                                 // table to which a secondary table is joined
	BaseDataobject   string    `bun:"base_dataobject,type:varchar(255),nullzero"`                   // dataobject to which join_syntax is added
	JoinTable        string    `bun:"join_table,type:varchar(255)"`                                 // secondary table
	JoinSyntax       string    `bun:"join_syntax,type:varchar(8000)"`                               // syntax to join join_table to base_table
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:((704))"`                     // status of each record
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
