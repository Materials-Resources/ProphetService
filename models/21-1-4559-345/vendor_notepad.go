package model

type VendorNotepad struct {
	bun.BaseModel    `bun:"table:vendor_notepad"`
	NoteId           int32     `bun:"note_id,type:int,pk"`
	CompanyId        string    `bun:"company_id,type:varchar(8)"`
	VendorId         float64   `bun:"vendor_id,type:decimal(19,0)"`
	Topic            string    `bun:"topic,type:varchar(30)"`
	Note             string    `bun:"note,type:text(2147483647),nullzero"`
	ActivationDate   time.Time `bun:"activation_date,type:datetime"`
	ExpirationDate   time.Time `bun:"expiration_date,type:datetime"`
	EntryDate        time.Time `bun:"entry_date,type:datetime"`
	NotepadClass     string    `bun:"notepad_class,type:varchar(8),nullzero"`
	Mandatory        string    `bun:"mandatory,type:char"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
