package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Activity struct {
	bun.BaseModel    `bun:"table:activity"`
	ActivityId       string    `bun:"activity_id,type:varchar(10),pk"`
	ActivityDesc     string    `bun:"activity_desc,type:varchar(255)"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	HardTouchFlag    string    `bun:"hard_touch_flag,type:char,default:('N')"`
}
