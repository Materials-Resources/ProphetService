package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type RibbonTabGroup struct {
	bun.BaseModel     `bun:"table:ribbon_tab_group"`
	RibbonTabGroupUid int32     `bun:"ribbon_tab_group_uid,type:int,autoincrement,pk"`
	TabGroupId        string    `bun:"tab_group_id,type:varchar(255)"`
	TabGroupText      string    `bun:"tab_group_text,type:varchar(255)"`
	Description       string    `bun:"description,type:varchar(255)"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
