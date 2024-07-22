package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PopupXPopup struct {
	bun.BaseModel       `bun:"table:popup_x_popup"`
	PopupXPopupUid      int32     `bun:"popup_x_popup_uid,type:int,autoincrement,identity,pk"`         // Unique Identifier of the table
	PopupIndexUid       int32     `bun:"popup_index_uid,type:int"`                                     // Unique identifier of the main popup
	LinkedPopupIndexUid int32     `bun:"linked_popup_index_uid,type:int"`                              // Unique identifier of the linked popup
	Order               int32     `bun:"order,type:int"`                                               // Order of the linked popup
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
