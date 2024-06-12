package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceMap struct {
	bun.BaseModel    `bun:"table:pricing_service_map"`
	LayoutId         float64   `bun:"layout_id,type:decimal(19,0),pk"`
	LayoutDetailId   float64   `bun:"layout_detail_id,type:decimal(19,0)"`
	TargetTable      string    `bun:"target_table,type:varchar(32),pk"`
	TargetColumn     string    `bun:"target_column,type:varchar(32),pk"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
