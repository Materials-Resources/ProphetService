package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type EscBaseViewAlias struct {
	bun.BaseModel       `bun:"table:esc_base_view_alias"`
	EscBaseViewAliasUid int32     `bun:"esc_base_view_alias_uid,type:int,autoincrement,pk"`            // Table primary key
	EscBaseViewName     string    `bun:"esc_base_view_name,type:varchar(255),unique"`                  // Base views for Baq column data.
	BaseTable           string    `bun:"base_table,type:varchar(255),unique"`                          // P21 base table
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
