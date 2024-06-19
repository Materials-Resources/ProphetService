package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingTemplateKeyField struct {
	bun.BaseModel              `bun:"table:pricing_template_key_field"`
	PricingTemplateKeyFieldUid int32     `bun:"pricing_template_key_field_uid,type:int,pk"`                   // Unique identifier for the record
	PricingTemplateUid         int32     `bun:"pricing_template_uid,type:int"`                                // Unique identifier for the associated pricing_template record
	CompanyId                  string    `bun:"company_id,type:varchar(8),nullzero"`                          // The company_id when using the Product Group key field
	KeyFieldId                 string    `bun:"key_field_id,type:varchar(255)"`                               // Specified key field for which defaults should be applied
	KeyFieldDesc               string    `bun:"key_field_desc,type:varchar(255)"`                             // Key field description.
	AppendCd                   int32     `bun:"append_cd,type:int,default:((300))"`                           // Indicates whether to append a value to an Item ID.
	NumberOfCharacters         int32     `bun:"number_of_characters,type:int,nullzero"`                       // Indiciates the number of characters to be appended to the Item ID.
	ItemPrefixSuffix           string    `bun:"item_prefix_suffix,type:varchar(255),nullzero"`                // The value to be appended to the Item ID.
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	SequenceNo                 int32     `bun:"sequence_no,type:int,default:((1))"`                           // Sequence in which key values should be searched
	SecondaryKeyFieldId        string    `bun:"secondary_key_field_id,type:varchar(255),nullzero"`            // Value of the secondary key type
	SecondaryKeyFieldDesc      string    `bun:"secondary_key_field_desc,type:varchar(255),nullzero"`          // Description of the secondary key value
}
