package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DbDrivenMaintKey struct {
	bun.BaseModel       `bun:"table:db_driven_maint_key"`
	DbDrivenMaintKeyUid int32     `bun:"db_driven_maint_key_uid,type:int,pk,identity"`
	DbDrivenMaintUid    int32     `bun:"db_driven_maint_uid,type:int"`
	KeyColumnName       string    `bun:"key_column_name,type:varchar(255)"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
