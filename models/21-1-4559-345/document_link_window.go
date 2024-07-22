package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DocumentLinkWindow struct {
	bun.BaseModel         `bun:"table:document_link_window"`
	DocumentLinkWindowUid int32     `bun:"document_link_window_uid,type:int,pk"`
	AreaCd                int32     `bun:"area_cd,type:int"`
	WindowClass           string    `bun:"window_class,type:varchar(255)"`
	RowStatusFlag         int32     `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
