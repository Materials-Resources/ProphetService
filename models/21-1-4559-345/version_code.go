package model

type VersionCode struct {
	bun.BaseModel    `bun:"table:version_code"`
	VersionCodeUid   int32     `bun:"version_code_uid,type:int,pk"`
	ServentRelease   string    `bun:"servent_release,type:varchar(8)"`
	ReleasePrefix    string    `bun:"release_prefix,type:char(3)"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30)"`
}
