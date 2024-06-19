package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ActivantLayoutDef struct {
	bun.BaseModel        `bun:"table:activant_layout_def"`
	ActivantLayoutDefUid int32     `bun:"activant_layout_def_uid,type:int,autoincrement,scanonly,pk"`   // Unique identifier for each record
	ColumnId             string    `bun:"column_id,type:varchar(255)"`                                  // Name of column in item_catalog table
	DisplayColumnId      string    `bun:"display_column_id,type:varchar(255)"`                          // Two letter identifier for item_catalog column
	ColumnName           string    `bun:"column_name,type:varchar(255),nullzero"`                       // User-defined column name
	ItemColumnName       string    `bun:"item_column_name,type:varchar(255),nullzero"`                  // Inventory table column name to be updated
	RowStatusFlag        int32     `bun:"row_status_flag,type:int,default:((704))"`                     // Status of current record
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	DisplaySequenceNo    int32     `bun:"display_sequence_no,type:int"`                                 // Sequence number in which to display records
}
