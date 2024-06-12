package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ApplicationSecurity struct {
	bun.BaseModel          `bun:"table:application_security"`
	ApplicationSecurityUid int32     `bun:"application_security_uid,type:int,pk,identity"`
	InternalName           string    `bun:"internal_name,type:varchar(255)"`
	DisplayName            string    `bun:"display_name,type:varchar(255)"`
	ScopeTypeCd            int32     `bun:"scope_type_cd,type:int"`
	ValueTypeCd            int32     `bun:"value_type_cd,type:int"`
	DefaultCodeValue       int32     `bun:"default_code_value,type:int,nullzero"`
	DefaultDecimalValue    float64   `bun:"default_decimal_value,type:decimal(19,9),nullzero"`
	DefaultStringValue     string    `bun:"default_string_value,type:varchar(255),nullzero"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	ConfigurationId        int32     `bun:"configuration_id,type:int,default:((0))"`
}
