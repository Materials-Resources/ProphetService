package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchasePriceLibraryDetail struct {
	bun.BaseModel          `bun:"table:purchase_price_library_detail"`
	CompanyId              string    `bun:"company_id,type:varchar(8),pk"`
	PurchasePriceLibraryId string    `bun:"purchase_price_library_id,type:varchar(8),pk"`
	PurchasePricingBookId  string    `bun:"purchase_pricing_book_id,type:varchar(8),pk"`
	SequenceNumber         float64   `bun:"sequence_number,type:decimal(19,0)"`
	DeleteFlag             string    `bun:"delete_flag,type:char"`
	DateCreated            time.Time `bun:"date_created,type:datetime"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
