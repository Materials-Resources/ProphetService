package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type SortDragdrop struct {
	bun.BaseModel       `bun:"table:sort_dragdrop"`
	SortDragdropUid     int32     `bun:"sort_dragdrop_uid,type:int,autoincrement,identity,pk"`         // Unique Identifier
	Dataobject          string    `bun:"dataobject,type:varchar(255),unique"`                          // Datawindow Object Name
	UserId              string    `bun:"user_id,type:varchar(30),unique"`                              // User ID
	SelectedColumnsList string    `bun:"selected_columns_list,type:varchar(8000)"`                     // Selected Column List (Comma Delimited)
	SelectedColumnsSort string    `bun:"selected_columns_sort,type:varchar(8000)"`                     // Selected Column Sort List (Comma Delimited)
	SortBy              string    `bun:"sort_by,type:varchar(255)"`                                    // Type of Sort
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	RowStatusFlag       int32     `bun:"row_status_flag,type:int"`                                     // Row Status
}
