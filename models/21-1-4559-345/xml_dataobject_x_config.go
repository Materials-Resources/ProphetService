package model

import (
	"github.com/uptrace/bun"
	"time"
)

type XmlDataobjectXConfig struct {
	bun.BaseModel           `bun:"table:xml_dataobject_x_config"`
	XmlDataobjectXConfigUid int32     `bun:"xml_dataobject_x_config_uid,type:int,pk,identity"`
	XmlDataobjectUid        int32     `bun:"xml_dataobject_uid,type:int"`
	ConfigurationId         int32     `bun:"configuration_id,type:int"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
