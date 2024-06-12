package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PriceSourceXCompany struct {
	bun.BaseModel          `bun:"table:price_source_x_company"`
	PriceSourceXCompanyUid int32     `bun:"price_source_x_company_uid,type:int,pk"`
	CompanyId              string    `bun:"company_id,type:varchar(8)"`
	PricingTypeCd          int16     `bun:"pricing_type_cd,type:tinyint"`
	SourceTypeCd           int32     `bun:"source_type_cd,type:int"`
	AvailableInd           string    `bun:"available_ind,type:char"`
	UsedInZeroPricingInd   string    `bun:"used_in_zero_pricing_ind,type:char"`
	SequenceNumber         int16     `bun:"sequence_number,type:tinyint"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`
	DateCreated            time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30)"`
}
