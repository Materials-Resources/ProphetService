package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceConversion struct {
	bun.BaseModel        `bun:"table:pricing_service_conversion"`
	LayoutId             float64   `bun:"layout_id,type:decimal(19,0),pk"`
	LayoutDetailId       float64   `bun:"layout_detail_id,type:decimal(19,0),pk"`
	ConversionId         float64   `bun:"conversion_id,type:decimal(19,0),pk"`
	SequenceNumber       float64   `bun:"sequence_number,type:decimal(19,0),nullzero"`
	FactorSource         string    `bun:"factor_source,type:varchar(10),nullzero"`
	SourceLayoutId       float64   `bun:"source_layout_id,type:decimal(19,0),nullzero"`
	SourceLayoutDetailId float64   `bun:"source_layout_detail_id,type:decimal(19,0),nullzero"`
	Factor               float64   `bun:"factor,type:decimal(19,4),nullzero"`
	Operand              string    `bun:"operand,type:varchar(10),nullzero"`
	DeleteFlag           string    `bun:"delete_flag,type:char"`
	DateCreated          time.Time `bun:"date_created,type:datetime"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
