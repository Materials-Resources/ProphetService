package model

import (
	"github.com/uptrace/bun"
	"time"
)

type UnitOfMeasure struct {
	bun.BaseModel         `bun:"table:unit_of_measure"`
	UnitId                string    `bun:"unit_id,type:varchar(8),pk"`
	UnitDescription       string    `bun:"unit_description,type:varchar(30)"`
	DeleteFlag            string    `bun:"delete_flag,type:char"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	DisplayDescription    string    `bun:"display_description,type:varchar(30),nullzero"`
	RepackagingCostFactor float64   `bun:"repackaging_cost_factor,type:decimal(19,9),nullzero"`
	MaritimeUom           string    `bun:"maritime_uom,type:varchar(255),nullzero"`
	PackagingUnitFlag     string    `bun:"packaging_unit_flag,type:char,default:('N')"`
	DimensionScale        string    `bun:"dimension_scale,type:varchar(2),nullzero"`
}
