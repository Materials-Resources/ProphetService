package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchasePricingBook struct {
	bun.BaseModel           `bun:"table:purchase_pricing_book"`
	CompanyId               string    `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	PurchasePricingBookId   string    `bun:"purchase_pricing_book_id,type:varchar(8),pk"`               // Indicates a purchase pricing book.
	PurchasePricingBookDesc string    `bun:"purchase_pricing_book_desc,type:varchar(40)"`               // Description for the purchase pricing book.
	DeleteFlag              string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
