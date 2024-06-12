package model

import (
	"github.com/uptrace/bun"
	"time"
)

type JurisdictionTaxIsTaxable struct {
	bun.BaseModel        `bun:"table:jurisdiction_tax_is_taxable"`
	JurisdictionId       string    `bun:"jurisdiction_id,type:varchar(10),pk"`
	TaxingJurisdictionId string    `bun:"taxing_jurisdiction_id,type:varchar(10),pk"`
	DeleteFlag           string    `bun:"delete_flag,type:char"`
	DateCreated          time.Time `bun:"date_created,type:datetime"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
