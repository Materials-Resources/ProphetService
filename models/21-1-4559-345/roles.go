package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Roles struct {
	bun.BaseModel           `bun:"table:roles"`
	RoleUid                 int32     `bun:"role_uid,type:int,pk"`
	Role                    string    `bun:"role,type:varchar(128)"`
	DeleteFlag              string    `bun:"delete_flag,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	MinimumMarginPercentage float64   `bun:"minimum_margin_percentage,type:decimal(19,9),nullzero"`
}
