package model

import (
	"github.com/uptrace/bun"
	"time"
)

type TaxGroupHdr struct {
	bun.BaseModel       `bun:"table:tax_group_hdr"`
	TaxGroupId          string    `bun:"tax_group_id,type:varchar(10),pk"`
	CompanyId           string    `bun:"company_id,type:varchar(8),pk"`
	TaxGroupDescription string    `bun:"tax_group_description,type:varchar(40)"`
	DeleteFlag          string    `bun:"delete_flag,type:char"`
	DateCreated         time.Time `bun:"date_created,type:datetime"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	OverrideFcTaxFlag   string    `bun:"override_fc_tax_flag,type:char,default:('N')"`
}
