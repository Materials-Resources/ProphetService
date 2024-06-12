package model

import (
	"github.com/uptrace/bun"
	"time"
)

type FcDataobjectColumn struct {
	bun.BaseModel         `bun:"table:fc_dataobject_column"`
	FcDataobjectColumnUid int32     `bun:"fc_dataobject_column_uid,type:int,pk"`
	FcDataobjectTableUid  int32     `bun:"fc_dataobject_table_uid,type:int"`
	UserColumnName        string    `bun:"user_column_name,type:varchar(255)"`
	ColumnName            string    `bun:"column_name,type:varchar(255)"`
	UpdateableFlag        string    `bun:"updateable_flag,type:char,default:('N')"`
	UpdatewhereclauseFlag string    `bun:"updatewhereclause_flag,type:char,default:('N')"`
	DataType              string    `bun:"data_type,type:varchar(255),nullzero"`
	TypeCd                int32     `bun:"type_cd,type:int"`
	Expression            string    `bun:"expression,type:varchar(8000),nullzero"`
	RowStatusFlag         int32     `bun:"row_status_flag,type:int"`
	Description           string    `bun:"description,type:varchar(8000),nullzero"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
