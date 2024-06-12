package model

type NotepadClass struct {
	bun.BaseModel    `bun:"table:notepad_class"`
	NotepadClassId   string    `bun:"notepad_class_id,type:varchar(8),pk"`
	NotepadClassDesc string    `bun:"notepad_class_desc,type:varchar(40)"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
}
