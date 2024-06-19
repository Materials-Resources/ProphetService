package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type BusinessRuleDataElement struct {
	bun.BaseModel              `bun:"table:business_rule_data_element"`
	BusinessRuleDataElementUid int32     `bun:"business_rule_data_element_uid,type:int,autoincrement,identity,pk"` // Unique identifier
	BusinessRuleUid            int32     `bun:"business_rule_uid,type:int"`                                        // The business rule this record this element applies to
	FieldName                  string    `bun:"field_name,type:varchar(255)"`                                      // The field for the data element
	ClassName                  string    `bun:"class_name,type:varchar(255),nullzero"`                             // The class for the data element
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`                    // Date and time the record was originally created
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`              // User who created the record
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`              // Date and time the record was modified
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`      // User who last changed the record
	FieldAlias                 string    `bun:"field_alias,type:varchar(255),nullzero"`                            // Alias for the field name
}
