package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DwSyntaxCacheWindow struct {
	bun.BaseModel          `bun:"table:dw_syntax_cache_window"`
	DwSyntaxCacheWindowUid int32     `bun:"dw_syntax_cache_window_uid,type:int,autoincrement,pk"`         // Unique identifier for this record
	WindowName             string    `bun:"window_name,type:varchar(255)"`                                // Window name that dynachange uses.
	EnableCachingFlag      string    `bun:"enable_caching_flag,type:char(1)"`                             // Determines if datawindow syntax caching should happen for this window.
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
