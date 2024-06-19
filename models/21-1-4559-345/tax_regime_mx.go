package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type TaxRegimeMx struct {
	bun.BaseModel     `bun:"table:tax_regime_mx"`
	TaxRegimeMxUid    int32     `bun:"tax_regime_mx_uid,type:int,autoincrement,scanonly,pk"`         // Primary Key
	TaxRegimeCd       string    `bun:"tax_regime_cd,type:char(3)"`                                   // Tax regime code
	TaxRegimeDesc     string    `bun:"tax_regime_desc,type:varchar(255)"`                            // Tax decription
	UseForPersonFlag  string    `bun:"use_for_person_flag,type:char(1),nullzero"`                    // Whether this regime is applied for person
	UseForCompanyFlag string    `bun:"use_for_company_flag,type:char(1),nullzero"`                   // Whether this regime is applied for company
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	VersionNo         float64   `bun:"version_no,type:decimal(9,5),default:((1.0))"`                 // Version Number defined by SAT
	RevisionNo        string    `bun:"revision_no,type:varchar(255),default:('0')"`                  // Revision Number defined by SAT
	ValidFromDate     time.Time `bun:"valid_from_date,type:datetime,nullzero"`                       // Valid date from defined by SAT
	ValidUntilDate    time.Time `bun:"valid_until_date,type:datetime,nullzero"`                      // Valid date until defined by SAT
}
