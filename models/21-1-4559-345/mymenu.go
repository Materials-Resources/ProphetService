package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Mymenu struct {
	bun.BaseModel    `bun:"table:mymenu"`
	MymenuUid        int32     `bun:"mymenu_uid,type:int,pk"`
	ConfigurationId  int32     `bun:"configuration_id,type:int"`
	UsersId          string    `bun:"users_id,type:varchar(30)"`
	MenuItemName     string    `bun:"menu_item_name,type:varchar(128)"`
	WindowName       string    `bun:"window_name,type:varchar(128)"`
	WindowTitle      string    `bun:"window_title,type:varchar(128)"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30)"`
	OpenedbyMenuitem *string   `bun:"openedby_menuitem,type:varchar(128)"`
	Menuname         *string   `bun:"menuname,type:varchar(128)"`
	Windowrole       *string   `bun:"WindowRole,type:varchar(128)"`
	Icon             string    `bun:"icon,type:varchar(255),default:('activant.bmp')"` // File name for menu/toolbar icon
	SequenceNo       *int32    `bun:"sequence_no,type:int"`                            // Sequence Order of Menu Items
}
