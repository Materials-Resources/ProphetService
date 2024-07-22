package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ApplicationSecurity struct {
	bun.BaseModel          `bun:"table:application_security"`
	ApplicationSecurityUid int32     `bun:"application_security_uid,type:int,autoincrement,identity,pk"`  // Unique ID for table, primary key
	InternalName           string    `bun:"internal_name,type:varchar(255)"`                              // Internal name of security setting
	DisplayName            string    `bun:"display_name,type:varchar(255)"`                               // External (display) name for security setting
	ScopeTypeCd            int32     `bun:"scope_type_cd,type:int"`                                       // Code value to determine the scope of this setting
	ValueTypeCd            int32     `bun:"value_type_cd,type:int"`                                       // Code value to determine the value type for this setting
	DefaultCodeValue       *int32    `bun:"default_code_value,type:int"`                                  // Default code value for setting
	DefaultDecimalValue    *float64  `bun:"default_decimal_value,type:decimal(19,9)"`                     // Default decimal value for setting
	DefaultStringValue     *string   `bun:"default_string_value,type:varchar(255)"`                       // Default varchar value for setting
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	ConfigurationId        int32     `bun:"configuration_id,type:int,default:((0))"`                      // Customer configuration number.  Used to create security option for a specified configuration.
}
