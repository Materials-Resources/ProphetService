package model

type ScheduledImportDetails struct {
	bun.BaseModel             `bun:"table:scheduled_import_details"`
	ScheduledImportDetailsUid int32     `bun:"scheduled_import_details_uid,type:int,pk"`
	ScheduledImportDefUid     int32     `bun:"scheduled_import_def_uid,type:int"`
	Fileprefix                string    `bun:"fileprefix,type:varchar(20),nullzero"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
