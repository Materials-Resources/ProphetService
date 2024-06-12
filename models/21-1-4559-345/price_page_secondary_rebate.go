package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PricePageSecondaryRebate struct {
	bun.BaseModel               `bun:"table:price_page_secondary_rebate"`
	PricePageSecondaryRebateUid int32     `bun:"price_page_secondary_rebate_uid,type:int,pk"`
	PricePageUid                int32     `bun:"price_page_uid,type:int"`
	SecondaryRebateTypeCd       int32     `bun:"secondary_rebate_type_cd,type:int,default:(300)"`
	SecondaryRebateCalcMethCd   int32     `bun:"secondary_rebate_calc_meth_cd,type:int,nullzero"`
	SecondaryRebateValue        float64   `bun:"secondary_rebate_value,type:decimal(19,6),nullzero"`
	SecondaryRebateSourceCd     int32     `bun:"secondary_rebate_source_cd,type:int,nullzero"`
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
