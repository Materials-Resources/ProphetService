package model

import (
	"github.com/uptrace/bun"
	"time"
)

type TaxRegimeMx struct {
	bun.BaseModel     `bun:"table:tax_regime_mx"`
	TaxRegimeMxUid    int32     `bun:"tax_regime_mx_uid,type:int,pk,identity"`
	TaxRegimeCd       string    `bun:"tax_regime_cd,type:char(3)"`
	TaxRegimeDesc     string    `bun:"tax_regime_desc,type:varchar(255)"`
	UseForPersonFlag  string    `bun:"use_for_person_flag,type:char,nullzero"`
	UseForCompanyFlag string    `bun:"use_for_company_flag,type:char,nullzero"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	VersionNo         float64   `bun:"version_no,type:decimal(9,5),default:((1.0))"`
	RevisionNo        string    `bun:"revision_no,type:varchar(255),default:('0')"`
	ValidFromDate     time.Time `bun:"valid_from_date,type:datetime,nullzero"`
	ValidUntilDate    time.Time `bun:"valid_until_date,type:datetime,nullzero"`
}
