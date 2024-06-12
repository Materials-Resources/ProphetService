package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchasePricingBook struct {
	bun.BaseModel           `bun:"table:purchase_pricing_book"`
	CompanyId               string    `bun:"company_id,type:varchar(8),pk"`
	PurchasePricingBookId   string    `bun:"purchase_pricing_book_id,type:varchar(8),pk"`
	PurchasePricingBookDesc string    `bun:"purchase_pricing_book_desc,type:varchar(40)"`
	DeleteFlag              string    `bun:"delete_flag,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
