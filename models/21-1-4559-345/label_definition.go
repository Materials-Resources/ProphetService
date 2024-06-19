package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type LabelDefinition struct {
	bun.BaseModel       `bun:"table:label_definition"`
	LabelDefinitionUid  int32     `bun:"label_definition_uid,type:int,pk"`                             // Unique Identifier for this record
	CompanyId           string    `bun:"company_id,type:varchar(8),unique"`                            // Company ID to which this label definition is defined
	LabelDefinitionId   string    `bun:"label_definition_id,type:varchar(255),unique"`                 // Label Definition ID
	LabelDefinitionDesc string    `bun:"label_definition_desc,type:varchar(255),nullzero"`             // Description of this Label Definition
	LabelTypeCd         int32     `bun:"label_type_cd,type:int,unique"`                                // Label Type - from standard label types code_p21 record
	RowStatusFlag       int32     `bun:"row_status_flag,type:int"`                                     // Status of this record
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
