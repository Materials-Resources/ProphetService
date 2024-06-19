package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PopupFieldValue struct {
	bun.BaseModel      `bun:"table:popup_field_value"`
	PopupFieldValueUid int32     `bun:"popup_field_value_uid,type:int,autoincrement,scanonly,pk"`     // Unique identifier of the popup field value
	PopupFieldUid      int32     `bun:"popup_field_uid,type:int"`                                     // Unique identifier of the field
	Value              string    `bun:"value,type:varchar(255),nullzero"`                             // The value to display in the control
	Condition          string    `bun:"condition,type:varchar(5000),nullzero"`                        // The condition that will be applied when the value is selected
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
