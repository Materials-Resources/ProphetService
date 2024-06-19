package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type RibbonToolXRibbonTabGroup struct {
	bun.BaseModel                `bun:"table:ribbon_tool_x_ribbon_tab_group"`
	RibbonToolXRibbonTabGroupUid int32     `bun:"ribbon_tool_x_ribbon_tab_group_uid,type:int,autoincrement,scanonly,pk"`
	RibbonToolUid                int32     `bun:"ribbon_tool_uid,type:int"`
	RibbonTabGroupUid            int32     `bun:"ribbon_tab_group_uid,type:int"`
	SequenceNo                   int32     `bun:"sequence_no,type:int"`
	DateCreated                  time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                    string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	EnabledFlag                  bool      `bun:"enabled_flag,type:bit,default:((1))"`
	EnabledForVersionCd          int32     `bun:"enabled_for_version_cd,type:int,default:((1423))"` // Indicates whether the tool should display for desktop, web or both
}
