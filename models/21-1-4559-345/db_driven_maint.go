package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DbDrivenMaint struct {
	bun.BaseModel    `bun:"table:db_driven_maint"`
	DbDrivenMaintUid int32     `bun:"db_driven_maint_uid,type:int,pk,identity"`
	TableName        string    `bun:"table_name,type:varchar(255)"`
	DisplayName      string    `bun:"display_name,type:varchar(255)"`
	BoClass          string    `bun:"bo_class,type:varchar(255),nullzero"`
	BrClass          string    `bun:"br_class,type:varchar(255),nullzero"`
	VdwClass         string    `bun:"vdw_class,type:varchar(255),nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
