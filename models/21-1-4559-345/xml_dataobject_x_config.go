package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type XmlDataobjectXConfig struct {
	bun.BaseModel           `bun:"table:xml_dataobject_x_config"`
	XmlDataobjectXConfigUid int32     `bun:"xml_dataobject_x_config_uid,type:int,autoincrement,scanonly,pk"` // Unique identifier for xml_dataobject_x_config
	XmlDataobjectUid        int32     `bun:"xml_dataobject_uid,type:int"`                                    // Foreign Key: xml_dataobject_uid from xml_dataobject table
	ConfigurationId         int32     `bun:"configuration_id,type:int"`                                      // Customer configuration id that corresponds to the dataobject
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`                 // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`           // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`           // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`   // User who last changed the record
}
