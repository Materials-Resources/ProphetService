package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CrmContactInformation struct {
	bun.BaseModel            `bun:"table:crm_contact_information"`
	CrmContactInformationUid int32     `bun:"crm_contact_information_uid,type:int,pk,identity"`
	CompanyId                string    `bun:"company_id,type:varchar(8),nullzero"`
	EntityLinkIdChar         string    `bun:"entity_link_id_char,type:varchar(255),nullzero"`
	EntityLinkIdDec          float64   `bun:"entity_link_id_dec,type:decimal(19,0),nullzero"`
	EntityTypeCd             int32     `bun:"entity_type_cd,type:int"`
	LastHardTouchDate        time.Time `bun:"last_hard_touch_date,type:datetime"`
	ActivityTransNo          string    `bun:"activity_trans_no,type:varchar(10),nullzero"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
