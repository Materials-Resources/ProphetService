package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type SystemSetting struct {
	bun.BaseModel    `bun:"table:system_setting"`
	SystemSettingUid int32     `bun:"system_setting_uid,type:int,pk"`                            // Unique Identifier
	ConfigurationId  int32     `bun:"configuration_id,type:int,unique"`                          // Configuration number for the system setting, 0 for baseline.
	ModuleCd         int32     `bun:"module_cd,type:int"`                                        // CommerceCenter module this setting is associated with.
	Name             string    `bun:"name,type:varchar(128),unique"`                             // Name of the setting.
	Value            string    `bun:"value,type:varchar(255)"`                                   // Curerent value for the setting
	DataTypeCd       int32     `bun:"data_type_cd,type:int"`                                     // Data type of the setting, ex, string, integer, etc.
	DataTypeLength   int32     `bun:"data_type_length,type:int"`                                 // If applicable, length of data type, length of string for example.
	DataTypeScale    int32     `bun:"data_type_scale,type:int"`                                  // Precision of decimal values if applicable.
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`      // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
