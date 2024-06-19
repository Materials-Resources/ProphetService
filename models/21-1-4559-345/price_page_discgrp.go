package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PricePageDiscgrp struct {
	bun.BaseModel       `bun:"table:price_page_discgrp"`
	PricePageDiscgrpUid int32     `bun:"price_page_discgrp_uid,type:int,autoincrement,pk"` // Unique ID for this price_page_discgrp record.
	PricePageUid        int32     `bun:"price_page_uid,type:int"`                          // Unique ID for Price Pages.
	DiscountGroupId     string    `bun:"discount_group_id,type:varchar(8)"`                // Discount Group to associate with this Price Page.
	EffectiveDate       time.Time `bun:"effective_date,type:datetime"`                     // Starting date on which pricing page will be used in pricing calculations.
	ExpirationDate      time.Time `bun:"expiration_date,type:datetime"`                    // Date on which pricing page is no longer used in pricing calculations.
	DateCreated         time.Time `bun:"date_created,type:datetime"`                       // Indicates the date/time this record was created.
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`                 // Indicates the date/time this record was last modified.
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30)"`              // ID of the user who last maintained this record
}
