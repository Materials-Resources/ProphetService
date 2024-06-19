package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type OeLineSecondaryRebate struct {
	bun.BaseModel             `bun:"table:oe_line_secondary_rebate"`
	OeLineSecondaryRebateUid  int32     `bun:"oe_line_secondary_rebate_uid,type:int,autoincrement,scanonly,pk"` // Unique id for the row.
	OeLineUid                 int32     `bun:"oe_line_uid,type:int"`                                            // Indicates the associated oe_line record.
	SecondaryRebateCalcMethCd int32     `bun:"secondary_rebate_calc_meth_cd,type:int,nullzero"`                 // The method used to calculate the secondary rebate (Multiply, Divide, Add, Subtract).
	SecondaryRebateSourceCd   int32     `bun:"secondary_rebate_source_cd,type:int,nullzero"`                    // The source of the secondary rebate.
	SecondaryRebateTypeCd     int32     `bun:"secondary_rebate_type_cd,type:int,nullzero"`                      // The type of the secondary rebate (None, Source, Value).
	SecondaryRebateValue      float64   `bun:"secondary_rebate_value,type:decimal(19,9),nullzero"`              // The decimal value used for calculating the secondary rebate.
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`                  // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`            // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`            // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`    // User who last changed the record
	OtherCostSourceCd         int32     `bun:"other_cost_source_cd,type:int,nullzero"`                          // Indicates the source for the other cost.
	CostCalculationValue      float64   `bun:"cost_calculation_value,type:decimal(19,9),nullzero"`              // Indicates the multiplier that would be used in conjunction with calculated other cost.
	OtherCostSourcePriceOe    float64   `bun:"other_cost_source_price_oe,type:decimal(19,9),nullzero"`          // Indicates the source price at OE for the other cost.
	SecondayRebateSrcPriceOe  float64   `bun:"seconday_rebate_src_price_oe,type:decimal(19,9),nullzero"`        // Indicates the source price at OE for the secondary rebate.
}
