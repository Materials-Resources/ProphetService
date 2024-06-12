package model

import (
	"github.com/uptrace/bun"
	"time"
)

type SortDragdrop struct {
	bun.BaseModel       `bun:"table:sort_dragdrop"`
	SortDragdropUid     int32     `bun:"sort_dragdrop_uid,type:int,pk,identity"`
	Dataobject          string    `bun:"dataobject,type:varchar(255)"`
	UserId              string    `bun:"user_id,type:varchar(30)"`
	SelectedColumnsList string    `bun:"selected_columns_list,type:varchar(8000)"`
	SelectedColumnsSort string    `bun:"selected_columns_sort,type:varchar(8000)"`
	SortBy              string    `bun:"sort_by,type:varchar(255)"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	RowStatusFlag       int32     `bun:"row_status_flag,type:int"`
}
