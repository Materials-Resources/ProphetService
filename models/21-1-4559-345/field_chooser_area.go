package model

import (
	"github.com/uptrace/bun"
	"time"
)

type FieldChooserArea struct {
	bun.BaseModel       `bun:"table:field_chooser_area"`
	FieldChooserAreaUid int32     `bun:"field_chooser_area_uid,type:int,pk,identity"`
	Area                int32     `bun:"area,type:int"`
	Dataobject          string    `bun:"dataobject,type:varchar(255)"`
	WindowName          string    `bun:"window_name,type:varchar(255),nullzero"`
	ConfigurationId     int32     `bun:"configuration_id,type:int,nullzero"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
