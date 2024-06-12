package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingTemplateKeyField struct {
	bun.BaseModel              `bun:"table:pricing_template_key_field"`
	PricingTemplateKeyFieldUid int32     `bun:"pricing_template_key_field_uid,type:int,pk"`
	PricingTemplateUid         int32     `bun:"pricing_template_uid,type:int"`
	CompanyId                  string    `bun:"company_id,type:varchar(8),nullzero"`
	KeyFieldId                 string    `bun:"key_field_id,type:varchar(255)"`
	KeyFieldDesc               string    `bun:"key_field_desc,type:varchar(255)"`
	AppendCd                   int32     `bun:"append_cd,type:int,default:((300))"`
	NumberOfCharacters         int32     `bun:"number_of_characters,type:int,nullzero"`
	ItemPrefixSuffix           string    `bun:"item_prefix_suffix,type:varchar(255),nullzero"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	SequenceNo                 int32     `bun:"sequence_no,type:int,default:((1))"`
	SecondaryKeyFieldId        string    `bun:"secondary_key_field_id,type:varchar(255),nullzero"`
	SecondaryKeyFieldDesc      string    `bun:"secondary_key_field_desc,type:varchar(255),nullzero"`
}
