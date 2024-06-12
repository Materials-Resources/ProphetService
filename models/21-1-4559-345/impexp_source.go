package model

type ImpexpSource struct {
	bun.BaseModel    `bun:"table:impexp_source"`
	ImpexpSourceUid  int32     `bun:"impexp_source_uid,type:int,pk"`
	ImpexpSourceId   string    `bun:"impexp_source_id,type:varchar(50)"`
	ImpexpSourceDesc string    `bun:"impexp_source_desc,type:varchar(50)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,nullzero"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,nullzero"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	ImpexpType       string    `bun:"impexp_type,type:char,nullzero"`
}
