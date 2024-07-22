package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PopupOnflySetup struct {
	bun.BaseModel      `bun:"table:popup_onfly_setup"`
	PopupOnflySetupUid int32     `bun:"popup_onfly_setup_uid,type:int,autoincrement,identity,pk"`     // unic id of the table
	SearchKey          string    `bun:"search_key,type:varchar(500),unique"`                          // this is the name of the popupsOn Fly
	ParentColumn       *string   `bun:"parent_column,type:varchar(100),unique"`                       // name of the column of the parent used to the relationship
	ChildColumn        *string   `bun:"child_column,type:varchar(100),unique"`                        // name of the column of the child used to the relationship
	IsPinned           bool      `bun:"is_pinned,type:bit,default:((0)),unique"`                      // if this column is pined on the parent
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	AditionalColumn    *string   `bun:"aditional_column,type:varchar(255)"`                           // Sql Definition  of the new column added
}
