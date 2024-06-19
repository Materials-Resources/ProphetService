package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ItemCatalogDef struct {
	bun.BaseModel     `bun:"table:item_catalog_def"`
	ItemCatalogDefUid int32     `bun:"item_catalog_def_uid,type:int,autoincrement,scanonly,pk"`      // Unique identifier of each record
	ColumnId          string    `bun:"column_id,type:varchar(255)"`                                  // Two letter identifier for item_catalog column
	ColumnName        string    `bun:"column_name,type:varchar(255),nullzero"`                       // User-defined column name
	ItemSearchFlag    string    `bun:"item_search_flag,type:char(1),default:('N')"`                  // Indicates whether to include column in item search popup
	DisplayColumnId   string    `bun:"display_column_id,type:varchar(255)"`                          // Column ID displayed to the user
	RowStatusFlag     int32     `bun:"row_status_flag,type:int,default:((704))"`                     // Status of each record
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
