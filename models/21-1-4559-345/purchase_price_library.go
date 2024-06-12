package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchasePriceLibrary struct {
	bun.BaseModel            `bun:"table:purchase_price_library"`
	CompanyId                string    `bun:"company_id,type:varchar(8),pk"`
	PurchasePriceLibraryId   string    `bun:"purchase_price_library_id,type:varchar(8),pk"`
	PurchasePriceLibraryDesc string    `bun:"purchase_price_library_desc,type:varchar(40)"`
	PurchasePriceLibraryType string    `bun:"purchase_price_library_type,type:varchar(40)"`
	SourcePrice              string    `bun:"source_price,type:varchar(40),nullzero"`
	Multiplier               float64   `bun:"multiplier,type:decimal(19,9),nullzero"`
	Inactive                 string    `bun:"inactive,type:char"`
	DeleteFlag               string    `bun:"delete_flag,type:char"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
