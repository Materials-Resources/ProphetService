package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ActivantLayoutDef struct {
	bun.BaseModel        `bun:"table:activant_layout_def"`
	ActivantLayoutDefUid int32     `bun:"activant_layout_def_uid,type:int,pk,identity"`
	ColumnId             string    `bun:"column_id,type:varchar(255)"`
	DisplayColumnId      string    `bun:"display_column_id,type:varchar(255)"`
	ColumnName           string    `bun:"column_name,type:varchar(255),nullzero"`
	ItemColumnName       string    `bun:"item_column_name,type:varchar(255),nullzero"`
	RowStatusFlag        int32     `bun:"row_status_flag,type:int,default:((704))"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DisplaySequenceNo    int32     `bun:"display_sequence_no,type:int"`
}
