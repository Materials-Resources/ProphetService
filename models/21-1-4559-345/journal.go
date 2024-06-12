package model

type Journal struct {
	bun.BaseModel      `bun:"table:journal"`
	JournalId          string    `bun:"journal_id,type:varchar(2),pk"`
	JournalDescription string    `bun:"journal_description,type:varchar(20)"`
	DeleteFlag         string    `bun:"delete_flag,type:char"`
	DateCreated        time.Time `bun:"date_created,type:datetime"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	JournalType        int32     `bun:"journal_type,type:int"`
}
