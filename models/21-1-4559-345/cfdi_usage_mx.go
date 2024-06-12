package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CfdiUsageMx struct {
	bun.BaseModel     `bun:"table:cfdi_usage_mx"`
	CfdiUsageMxUid    int32     `bun:"cfdi_usage_mx_uid,type:int,pk,identity"`
	CfdiUsageCd       string    `bun:"cfdi_usage_cd,type:varchar(10)"`
	CfdiUsageDesc     string    `bun:"cfdi_usage_desc,type:varchar(255)"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	UseForPersonFlag  string    `bun:"use_for_person_flag,type:char,default:('Y')"`
	UseForCompanyFlag string    `bun:"use_for_company_flag,type:char,nullzero"`
	VersionNo         float64   `bun:"version_no,type:decimal(9,5),default:((1.0))"`
	RevisionNo        string    `bun:"revision_no,type:varchar(255),default:('0')"`
	ValidFromDate     time.Time `bun:"valid_from_date,type:datetime,nullzero"`
	ValidUntilDate    time.Time `bun:"valid_until_date,type:datetime,nullzero"`
}
