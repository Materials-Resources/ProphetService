package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceHdrNotepad struct {
	bun.BaseModel        `bun:"table:invoice_hdr_notepad"`
	InvoiceNo            string    `bun:"invoice_no,type:varchar(10),pk"`                       // Invoice number to which these invoice notes apply.
	NoteId               float64   `bun:"note_id,type:decimal(19,0),pk"`                        // Identifies the note.
	Topic                string    `bun:"topic,type:varchar(30)"`                               // The topic of the note for the referenced area.
	Note                 *string   `bun:"note,type:text"`                                       // Text of the note.
	ActivationDate       time.Time `bun:"activation_date,type:datetime"`                        // Starting date of the note.
	ExpirationDate       time.Time `bun:"expiration_date,type:datetime"`                        // Expiration date of the note.
	EntryDate            time.Time `bun:"entry_date,type:datetime"`                             // Entry date of the note.
	NotepadClassId       *string   `bun:"notepad_class_id,type:varchar(8)"`                     // A user-defined ID code which can be associated with an invoice note
	Mandatory            string    `bun:"mandatory,type:char(1)"`                               // Is the note mandatory?
	DeleteFlag           string    `bun:"delete_flag,type:char(1)"`                             // Indicates whether this record is logically deleted
	DateCreated          time.Time `bun:"date_created,type:datetime"`                           // Indicates the date/time this record was created.
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`                     // Indicates the date/time this record was last modified.
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30)"`                  // ID of the user who last maintained this record
	InvoiceHdrNotepadUid int32     `bun:"invoice_hdr_notepad_uid,type:int,unique"`              // Unique ID of each invoice header notepad record.
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"` // User who created the record
}
