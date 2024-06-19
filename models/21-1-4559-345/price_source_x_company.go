package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PriceSourceXCompany struct {
	bun.BaseModel          `bun:"table:price_source_x_company"`
	PriceSourceXCompanyUid int32     `bun:"price_source_x_company_uid,type:int,pk"` // Internal record identifier
	CompanyId              string    `bun:"company_id,type:varchar(8)"`             // Unique code that identifies a company.
	PricingTypeCd          int16     `bun:"pricing_type_cd,type:tinyint"`           // Indicates this is used for Sales Pricing
	SourceTypeCd           int32     `bun:"source_type_cd,type:int"`                // Indicates the source price (Price 1, Price 2, etc.)
	AvailableInd           string    `bun:"available_ind,type:char(1)"`             // Indicates if this is in use for the company or not
	UsedInZeroPricingInd   string    `bun:"used_in_zero_pricing_ind,type:char(1)"`  // Indicates if this is used in performing zero price resolution or not
	SequenceNumber         int16     `bun:"sequence_number,type:tinyint"`           // Sequence number of this source price within the Company
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`       // Indicates the date/time this record was last modified.
	DateCreated            time.Time `bun:"date_created,type:datetime"`             // Indicates the date/time this record was created.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30)"`    // ID of the user who last maintained this record
}
