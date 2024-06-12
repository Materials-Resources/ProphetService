package model

type CustomObjectsBackup struct {
	bun.BaseModel    `bun:"table:custom_objects_backup"`
	CustomObjectsUid int32     `bun:"custom_objects_uid,type:int"`
	UsersId          string    `bun:"users_id,type:varchar(30)"`
	Object           string    `bun:"object,type:varchar(128)"`
	ConfigurationId  int32     `bun:"configuration_id,type:int"`
	ModString        string    `bun:"mod_string,type:text(2147483647)"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30)"`
	VersionId        string    `bun:"version_id,type:varchar(128)"`
	VersionDesc      string    `bun:"version_desc,type:varchar(255)"`
	RoleId           int32     `bun:"role_id,type:int"`
	Type             string    `bun:"type,type:char"`
	ObjectType       string    `bun:"object_type,type:char"`
	ApplyToAll       int32     `bun:"apply_to_all,type:int"`
	DefaultValues    string    `bun:"default_values,type:text(2147483647),nullzero"`
}
