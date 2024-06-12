package model

import (
	"github.com/uptrace/bun"
	"time"
)

type FcDataobjectTable struct {
	bun.BaseModel        `bun:"table:fc_dataobject_table"`
	FcDataobjectTableUid int32     `bun:"fc_dataobject_table_uid,type:int,pk"`
	FcDataobjectUid      int32     `bun:"fc_dataobject_uid,type:int"`
	TableName            string    `bun:"table_name,type:varchar(255)"`
	JoinSyntax           string    `bun:"join_syntax,type:varchar(8000),nullzero"`
	RowStatusFlag        int32     `bun:"row_status_flag,type:int"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
