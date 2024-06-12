package model

type PoLineNotepad struct {
	bun.BaseModel    `bun:"table:po_line_notepad"`
	NoteId           float64   `bun:"note_id,type:decimal(19,0),pk"`
	PoNo             float64   `bun:"po_no,type:decimal(19,0)"`
	LineNo           float64   `bun:"line_no,type:decimal(19,0)"`
	NotepadClass     string    `bun:"notepad_class,type:varchar(8),nullzero"`
	Topic            string    `bun:"topic,type:varchar(30)"`
	Note             string    `bun:"note,type:text(2147483647),nullzero"`
	ActivationDate   time.Time `bun:"activation_date,type:datetime"`
	ExpirationDate   time.Time `bun:"expiration_date,type:datetime"`
	EntryDate        time.Time `bun:"entry_date,type:datetime"`
	Mandatory        string    `bun:"mandatory,type:char"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
}
