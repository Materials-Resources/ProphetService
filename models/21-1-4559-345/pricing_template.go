package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingTemplate struct {
	bun.BaseModel       `bun:"table:pricing_template"`
	PricingTemplateUid  int32     `bun:"pricing_template_uid,type:int,pk"`                             // Unique identifier for the record
	PricingTemplateId   string    `bun:"pricing_template_id,type:varchar(255)"`                        // Unique name for this template
	PricingTemplateDesc string    `bun:"pricing_template_desc,type:varchar(255)"`                      // The description for this template
	KeyFieldCd          int32     `bun:"key_field_cd,type:int"`                                        // Determines which Pricing Service column to match on.
	RowStatusFlag       int32     `bun:"row_status_flag,type:int"`                                     // Indicates whether this record is active or not.
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	SecondaryKeyFieldCd int32     `bun:"secondary_key_field_cd,type:int,nullzero"`                     // Secondary pricing service column to match on
}
