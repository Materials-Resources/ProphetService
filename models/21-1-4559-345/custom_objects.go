package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomObjects struct {
	bun.BaseModel    `bun:"table:custom_objects"`
	CustomObjectsUid int32     `bun:"custom_objects_uid,type:int,pk"`
	UsersId          string    `bun:"users_id,type:varchar(30)"`
	Object           string    `bun:"object,type:varchar(128)"`
	ConfigurationId  int32     `bun:"configuration_id,type:int"`
	ModString        string    `bun:"mod_string,type:text(2147483647)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	VersionId        string    `bun:"version_id,type:varchar(128)"`
	VersionDesc      string    `bun:"version_desc,type:varchar(255)"`
	RoleId           int32     `bun:"role_id,type:int"`
	Type             string    `bun:"type,type:char"`
	ObjectType       string    `bun:"object_type,type:char"`
	ApplyToAll       int32     `bun:"apply_to_all,type:int,default:(2)"`
	DefaultValues    string    `bun:"default_values,type:text(2147483647),nullzero"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:((704))"`
	DesignUid        int32     `bun:"design_uid,type:int,nullzero"`
	MigratedToWeb    string    `bun:"migrated_to_web,type:char,default:('N')"`
}
