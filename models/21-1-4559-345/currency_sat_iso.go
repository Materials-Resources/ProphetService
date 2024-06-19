package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CurrencySatIso struct {
	bun.BaseModel              `bun:"table:currency_sat_iso"`
	CurrencySatIsoUid          int32     `bun:"currency_sat_iso_uid,type:int,autoincrement,pk"`               // Id for the table
	CurrencySatIsoCode         string    `bun:"currency_sat_iso_code,type:varchar(255)"`                      // Code for the sat currency
	CurrencySatIsoDesc         string    `bun:"currency_sat_iso_desc,type:varchar(255)"`                      // Description of the currency
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	NumberOfDecimals           int32     `bun:"number_of_decimals,type:int,nullzero"`                         // Number of Decimals for currency
	PercentageVariationAllowed float64   `bun:"percentage_variation_allowed,type:decimal(9,5),nullzero"`      // Percentage Variation Allowed
	VersionNo                  float64   `bun:"version_no,type:decimal(9,5),default:((1.0))"`                 // Version Number defined by SAT
	RevisionNo                 string    `bun:"revision_no,type:varchar(255),default:('0')"`                  // Revision Number defined by SAT
	ValidFromDate              time.Time `bun:"valid_from_date,type:datetime,nullzero"`                       // Valid date from defined by SAT
	ValidUntilDate             time.Time `bun:"valid_until_date,type:datetime,nullzero"`                      // Valid date until defined by SAT
}
