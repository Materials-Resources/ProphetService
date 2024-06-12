package model

type PoSchedule struct {
	bun.BaseModel    `bun:"table:po_schedule"`
	PoScheduleUid    int32     `bun:"po_schedule_uid,type:int,pk"`
	PoHdrUid         int32     `bun:"po_hdr_uid,type:int"`
	ReleaseNo        int32     `bun:"release_no,type:int"`
	ReleaseDate      time.Time `bun:"release_date,type:datetime"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
