package prophet_21_1_4559

import (
	"time"

	"github.com/uptrace/bun"
)

type ItemIdChangeHistory struct {
	bun.BaseModel          `bun:"table:item_id_change_history"`
	InvMastUid             int32     `bun:"inv_mast_uid,type:int"`
	OldItemId              string    `bun:"old_item_id,type:varchar(40)"`
	NewItemId              string    `bun:"new_item_id,type:varchar(40)"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	ItemIdChangeHistoryUid int32     `bun:"item_id_change_history_uid,type:int,pk,identity"`
}
