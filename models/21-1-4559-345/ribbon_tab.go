package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type RibbonTab struct {
	bun.BaseModel    `bun:"table:ribbon_tab"`
	RibbonTabUid     int32     `bun:"ribbon_tab_uid,type:int,autoincrement,identity,pk"`
	TabId            string    `bun:"tab_id,type:varchar(255)"`
	TabText          string    `bun:"tab_text,type:varchar(255)"`
	Description      string    `bun:"description,type:varchar(255),nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
