package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Ribbon struct {
	bun.BaseModel    `bun:"table:ribbon"`
	RibbonUid        int32     `bun:"ribbon_uid,type:int,autoincrement,scanonly,pk"`
	Name             string    `bun:"name,type:varchar(255)"`
	Description      string    `bun:"description,type:varchar(255),nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
