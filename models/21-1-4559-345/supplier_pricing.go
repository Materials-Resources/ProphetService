package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type SupplierPricing struct {
	bun.BaseModel    `bun:"table:supplier_pricing"`
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`                                 // Unique code that identifies a company.
	SupplierId       float64   `bun:"supplier_id,type:decimal(19,0),pk"`                             // What supplier supplies material for this stage?
	PricingMethod    string    `bun:"pricing_method,type:varchar(40)"`                               // Indicates whether to calculate purchase pricing by multiplier or to use the pricing libraries.
	SourcePrice      *string   `bun:"source_price,type:varchar(40)"`                                 // When the pricing_method is Multiplier this indicates the base purchase price.
	Multiplier       *float64  `bun:"multiplier,type:decimal(19,9)"`                                 // When the pricing_method is Multiplier this indicates what to multiply the source_price by to get the purchase price.
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
}
