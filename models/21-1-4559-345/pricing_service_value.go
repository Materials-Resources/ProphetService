package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceValue struct {
	bun.BaseModel    `bun:"table:pricing_service_value"`
	ValueId          float64   `bun:"value_id,type:decimal(19,0),pk"`
	ValueDesc        string    `bun:"value_desc,type:varchar(50),nullzero"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	UseExactMatch    string    `bun:"use_exact_match,type:char,default:('Y')"`
	ProcessType      string    `bun:"process_type,type:varchar(10),default:('PSDESIGN')"`
}
