package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerEmailDefaults struct {
	bun.BaseModel            `bun:"table:customer_email_defaults"`
	CustomerEmailDefaultsUid int32     `bun:"customer_email_defaults_uid,type:int,pk,identity"`
	CompanyId                string    `bun:"company_id,type:varchar(8)"`
	CustomerId               float64   `bun:"customer_id,type:decimal(19,0)"`
	FormTypeCd               int32     `bun:"form_type_cd,type:int"`
	Subject                  string    `bun:"subject,type:varchar(255),nullzero"`
	Copy                     string    `bun:"copy,type:varchar(255),nullzero"`
	BlindCopy                string    `bun:"blind_copy,type:varchar(255),nullzero"`
	Memo                     string    `bun:"memo,type:varchar(255),nullzero"`
	CcUser                   string    `bun:"cc_user,type:char,nullzero"`
	CcSalesrep               string    `bun:"cc_salesrep,type:char,nullzero"`
	CcTaker                  string    `bun:"cc_taker,type:char,nullzero"`
	BcUser                   string    `bun:"bc_user,type:char,nullzero"`
	BcSalesrep               string    `bun:"bc_salesrep,type:char,nullzero"`
	BcTaker                  string    `bun:"bc_taker,type:char,nullzero"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
