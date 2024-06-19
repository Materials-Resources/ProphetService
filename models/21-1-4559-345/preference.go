package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Preference struct {
	bun.BaseModel    `bun:"table:preference"`
	PreferenceUid    int32     `bun:"preference_uid,type:int,autoincrement,identity,pk"`            // Unique ID for the given preference
	Name             string    `bun:"name,type:varchar(255)"`                                       // Unique name of the preference
	DefaultValue     string    `bun:"default_value,type:varchar(max)"`                              // Default value of the preference
	DataTypeCd       int32     `bun:"data_type_cd,type:int,default:((1255))"`                       // Data type
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
