package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PricePageTypeXCompany struct {
	bun.BaseModel            `bun:"table:price_page_type_x_company"`
	PricePageTypeXCompanyUid int32     `bun:"price_page_type_x_company_uid,type:int,pk"` // Internal record identifier
	CompanyId                string    `bun:"company_id,type:varchar(8)"`                // Unique code that identifies a company.
	PricePageTypeCd          int16     `bun:"price_page_type_cd,type:smallint"`          // Indicates the price page type (Item ID, Supplier ID, etc.)
	SequenceNumber           int16     `bun:"sequence_number,type:tinyint"`              // Indicates the sequence in which to process the loc
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`          // Indicates the date/time this record was last modified.
	DateCreated              time.Time `bun:"date_created,type:datetime"`                // Indicates the date/time this record was created.
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30)"`       // ID of the user who last maintained this record
}
