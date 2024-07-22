package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type RibbonTabGroupXRibbonTab struct {
	bun.BaseModel               `bun:"table:ribbon_tab_group_x_ribbon_tab"`
	RibbonTabGroupXRibbonTabUid int32     `bun:"ribbon_tab_group_x_ribbon_tab_uid,type:int,autoincrement,identity,pk"`
	RibbonTabGroupUid           int32     `bun:"ribbon_tab_group_uid,type:int"`
	RibbonTabUid                int32     `bun:"ribbon_tab_uid,type:int"`
	SequenceNo                  int32     `bun:"sequence_no,type:int"`
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
