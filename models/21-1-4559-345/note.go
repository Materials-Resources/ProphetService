package model

type Note struct {
	bun.BaseModel           `bun:"table:note"`
	NoteUid                 int32     `bun:"note_uid,type:int,pk"`
	DocumentUid             int32     `bun:"document_uid,type:int"`
	NotepadClassId          string    `bun:"notepad_class_id,type:varchar(8),nullzero"`
	Topic                   string    `bun:"topic,type:varchar(255)"`
	Note                    string    `bun:"note,type:text(2147483647),nullzero"`
	ActivationDate          time.Time `bun:"activation_date,type:datetime"`
	ExpirationDate          time.Time `bun:"expiration_date,type:datetime"`
	EntryDate               time.Time `bun:"entry_date,type:datetime"`
	Mandatory               string    `bun:"mandatory,type:char"`
	RowStatusFlag           int32     `bun:"row_status_flag,type:int"`
	NoteTypeCd              int32     `bun:"note_type_cd,type:int"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	CreateOrderLineNoteFlag string    `bun:"create_order_line_note_flag,type:char,default:('N')"`
}
