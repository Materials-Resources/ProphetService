package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type FieldChooserArea struct {
	bun.BaseModel       `bun:"table:field_chooser_area"`
	FieldChooserAreaUid int32     `bun:"field_chooser_area_uid,type:int,autoincrement,scanonly,pk"`    // Unique identifier for each record
	Area                int32     `bun:"area,type:int"`                                                // Field Chooser area / context code
	Dataobject          string    `bun:"dataobject,type:varchar(255)"`                                 // Dataobject associated with a field chooser  area
	WindowName          string    `bun:"window_name,type:varchar(255),nullzero"`                       // window name
	ConfigurationId     int32     `bun:"configuration_id,type:int,nullzero"`                           // Configuration in which record applies
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
