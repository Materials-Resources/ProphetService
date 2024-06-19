package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchasePriceLibraryDetail struct {
	bun.BaseModel          `bun:"table:purchase_price_library_detail"`
	CompanyId              string    `bun:"company_id,type:varchar(8),pk"` // Unique code that identifies a company.
	PurchasePriceLibraryId string    `bun:"purchase_price_library_id,type:varchar(8),pk"`
	PurchasePricingBookId  string    `bun:"purchase_pricing_book_id,type:varchar(8),pk"`               // Indicates a single purchase pricing book used by the library.
	SequenceNumber         float64   `bun:"sequence_number,type:decimal(19,0)"`                        // Indicates the sequence in which to process the loc
	DeleteFlag             string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated            time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
