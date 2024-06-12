package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchType struct {
	bun.BaseModel    `bun:"table:purch_type"`
	TypeId           string    `bun:"type_id,type:varchar(8),pk"`
	TypeDesc         string    `bun:"type_desc,type:varchar(30)"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
