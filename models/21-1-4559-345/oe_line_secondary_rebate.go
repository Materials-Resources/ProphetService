package model

import (
	"github.com/uptrace/bun"
	"time"
)

type OeLineSecondaryRebate struct {
	bun.BaseModel             `bun:"table:oe_line_secondary_rebate"`
	OeLineSecondaryRebateUid  int32     `bun:"oe_line_secondary_rebate_uid,type:int,pk,identity"`
	OeLineUid                 int32     `bun:"oe_line_uid,type:int"`
	SecondaryRebateCalcMethCd int32     `bun:"secondary_rebate_calc_meth_cd,type:int,nullzero"`
	SecondaryRebateSourceCd   int32     `bun:"secondary_rebate_source_cd,type:int,nullzero"`
	SecondaryRebateTypeCd     int32     `bun:"secondary_rebate_type_cd,type:int,nullzero"`
	SecondaryRebateValue      float64   `bun:"secondary_rebate_value,type:decimal(19,9),nullzero"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	OtherCostSourceCd         int32     `bun:"other_cost_source_cd,type:int,nullzero"`
	CostCalculationValue      float64   `bun:"cost_calculation_value,type:decimal(19,9),nullzero"`
	OtherCostSourcePriceOe    float64   `bun:"other_cost_source_price_oe,type:decimal(19,9),nullzero"`
	SecondayRebateSrcPriceOe  float64   `bun:"seconday_rebate_src_price_oe,type:decimal(19,9),nullzero"`
}
