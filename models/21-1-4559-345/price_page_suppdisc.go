package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PricePageSuppdisc struct {
	bun.BaseModel        `bun:"table:price_page_suppdisc"`
	PricePageSuppdiscUid int32     `bun:"price_page_suppdisc_uid,type:int,autoincrement,identity,pk"` // Unique ID for this price_page_suppdisc record
	PricePageUid         int32     `bun:"price_page_uid,type:int"`                                    // Unique ID for Price Pages
	SupplierId           float64   `bun:"supplier_id,type:decimal(19,0)"`                             // Supplier to associate with this Price Page
	DiscountGroupId      string    `bun:"discount_group_id,type:varchar(8)"`                          // Discount Group to associate with this Price Page
	EffectiveDate        time.Time `bun:"effective_date,type:datetime"`                               // Starting date on which pricing page will be used in pricing calculations.
	ExpirationDate       time.Time `bun:"expiration_date,type:datetime"`                              // Date on which pricing page is no longer used in pricing calculations.
	DateCreated          time.Time `bun:"date_created,type:datetime"`                                 // Indicates the date/time this record was created.
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`                           // Indicates the date/time this record was last modified.
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30)"`                        // ID of the user who last maintained this record
}
