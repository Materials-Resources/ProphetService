package model

type StageNotepad struct {
	bun.BaseModel    `bun:"table:stage_notepad"`
	StageNotepadUid  int32     `bun:"stage_notepad_uid,type:int,pk"`
	StageUid         int32     `bun:"stage_uid,type:int"`
	NoteId           float64   `bun:"note_id,type:decimal(19,0)"`
	NotepadClassId   string    `bun:"notepad_class_id,type:varchar(8),nullzero"`
	Topic            string    `bun:"topic,type:varchar(30)"`
	Note             string    `bun:"note,type:text(2147483647),nullzero"`
	ActivationDate   time.Time `bun:"activation_date,type:datetime"`
	ExpirationDate   time.Time `bun:"expiration_date,type:datetime"`
	EntryDate        time.Time `bun:"entry_date,type:datetime"`
	Mandatory        string    `bun:"mandatory,type:char"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
