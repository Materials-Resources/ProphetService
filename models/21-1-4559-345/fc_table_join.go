package model

import (
	"github.com/uptrace/bun"
	"time"
)

type FcTableJoin struct {
	bun.BaseModel    `bun:"table:fc_table_join"`
	FcTableJoinUid   int32     `bun:"fc_table_join_uid,type:int,identity"`
	BaseTable        string    `bun:"base_table,type:varchar(255)"`
	BaseDataobject   string    `bun:"base_dataobject,type:varchar(255),nullzero"`
	JoinTable        string    `bun:"join_table,type:varchar(255)"`
	JoinSyntax       string    `bun:"join_syntax,type:varchar(8000)"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:((704))"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
