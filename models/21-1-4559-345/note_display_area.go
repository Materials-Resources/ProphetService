package model

type NoteDisplayArea struct {
	bun.BaseModel       `bun:"table:note_display_area"`
	NoteDisplayAreasUid int32     `bun:"note_display_areas_uid,type:int,pk,identity"`
	TopicId             int32     `bun:"topic_id,type:int"`
	DisplayArea         string    `bun:"display_area,type:varchar(40)"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
