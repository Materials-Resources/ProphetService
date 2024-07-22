package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type VendorNotepad struct {
	bun.BaseModel    `bun:"table:vendor_notepad"`
	NoteId           int32     `bun:"note_id,type:int,pk"`                                          // unique note id
	CompanyId        string    `bun:"company_id,type:varchar(8)"`                                   // company id
	VendorId         float64   `bun:"vendor_id,type:decimal(19,0)"`                                 // vendor id
	Topic            string    `bun:"topic,type:varchar(30)"`                                       // topic for note
	Note             *string   `bun:"note,type:text"`                                               // note
	ActivationDate   time.Time `bun:"activation_date,type:datetime"`                                // date of activation
	ExpirationDate   time.Time `bun:"expiration_date,type:datetime"`                                // date of expiration
	EntryDate        time.Time `bun:"entry_date,type:datetime"`                                     // date entered
	NotepadClass     *string   `bun:"notepad_class,type:varchar(8)"`                                // class the notepad belongs to
	Mandatory        string    `bun:"mandatory,type:char(1)"`                                       // is the note mandatory or not
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                     // flag to indicate note is deleted or not
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
