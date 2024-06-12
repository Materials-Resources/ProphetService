package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DynachangeMenu struct {
	bun.BaseModel         `bun:"table:dynachange_menu"`
	DynachangeId          float64   `bun:"dynachange_id,type:decimal(6,0),pk"`
	BaseClassItem         string    `bun:"base_class_item,type:varchar(255),pk"`
	PersonalizedClassItem string    `bun:"personalized_class_item,type:varchar(255),pk"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
