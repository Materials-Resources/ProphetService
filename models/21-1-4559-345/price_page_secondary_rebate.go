package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PricePageSecondaryRebate struct {
	bun.BaseModel               `bun:"table:price_page_secondary_rebate"`
	PricePageSecondaryRebateUid int32     `bun:"price_page_secondary_rebate_uid,type:int,pk"`
	PricePageUid                int32     `bun:"price_page_uid,type:int"`                                      // The unique price page id
	SecondaryRebateTypeCd       *int32    `bun:"secondary_rebate_type_cd,type:int,default:(300)"`              // The type of the secondary rebate (none, source, value)
	SecondaryRebateCalcMethCd   *int32    `bun:"secondary_rebate_calc_meth_cd,type:int"`                       // The method used to calculate the secondary rebate (multiply, divide, add, subtract)
	SecondaryRebateValue        *float64  `bun:"secondary_rebate_value,type:decimal(19,6)"`                    // The decimal value used for calculating the secondary rebate
	SecondaryRebateSourceCd     *int32    `bun:"secondary_rebate_source_cd,type:int"`                          // The source of the secondary rebate
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
