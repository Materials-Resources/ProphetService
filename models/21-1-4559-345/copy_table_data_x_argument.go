package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CopyTableDataXArgument struct {
	bun.BaseModel             `bun:"table:copy_table_data_x_argument"`
	CopyTableDataXArgumentUid int32     `bun:"copy_table_data_x_argument_uid,type:int,pk,identity"`
	CopyTableDataXProcessUid  int32     `bun:"copy_table_data_x_process_uid,type:int"`
	ArgumentName              string    `bun:"argument_name,type:varchar(255)"`
	DefaultValue              string    `bun:"default_value,type:varchar(255)"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
