package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DbDrivenMaint struct {
	bun.BaseModel    `bun:"table:db_driven_maint"`
	DbDrivenMaintUid int32     `bun:"db_driven_maint_uid,type:int,autoincrement,identity,pk"`       // UID for this table.
	TableName        string    `bun:"table_name,type:varchar(255)"`                                 // DB table that this maint window is going to update.
	DisplayName      string    `bun:"display_name,type:varchar(255)"`                               // Name as it will be displayed in P21 (used in windw title bar).
	BoClass          *string   `bun:"bo_class,type:varchar(255)"`                                   // Optional - BO class to be used for maint window.
	BrClass          *string   `bun:"br_class,type:varchar(255)"`                                   // Optional - BR class to be used for maint window.
	VdwClass         *string   `bun:"vdw_class,type:varchar(255)"`                                  // Optional - VDW class to be used for maint window.
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
