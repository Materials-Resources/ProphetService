package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CurrencySatIso struct {
	bun.BaseModel              `bun:"table:currency_sat_iso"`
	CurrencySatIsoUid          int32     `bun:"currency_sat_iso_uid,type:int,pk,identity"`
	CurrencySatIsoCode         string    `bun:"currency_sat_iso_code,type:varchar(255)"`
	CurrencySatIsoDesc         string    `bun:"currency_sat_iso_desc,type:varchar(255)"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	NumberOfDecimals           int32     `bun:"number_of_decimals,type:int,nullzero"`
	PercentageVariationAllowed float64   `bun:"percentage_variation_allowed,type:decimal(9,5),nullzero"`
	VersionNo                  float64   `bun:"version_no,type:decimal(9,5),default:((1.0))"`
	RevisionNo                 string    `bun:"revision_no,type:varchar(255),default:('0')"`
	ValidFromDate              time.Time `bun:"valid_from_date,type:datetime,nullzero"`
	ValidUntilDate             time.Time `bun:"valid_until_date,type:datetime,nullzero"`
}
