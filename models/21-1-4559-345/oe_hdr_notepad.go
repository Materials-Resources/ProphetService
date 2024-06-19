package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type OeHdrNotepad struct {
	bun.BaseModel    `bun:"table:oe_hdr_notepad"`
	NoteId           float64   `bun:"note_id,type:decimal(19,0),pk"`                                 // What is the unique identifier for this supplier note?
	OrderNo          string    `bun:"order_no,type:varchar(8),nullzero"`                             // What order does this note belong to?
	NotepadClassId   string    `bun:"notepad_class_id,type:varchar(8),nullzero"`                     // What is the class for this note?
	Topic            string    `bun:"topic,type:varchar(30)"`                                        // The topic of the note for the referenced area.
	Note             string    `bun:"note,type:text,nullzero"`                                       // What are the contents of the note?
	ActivationDate   time.Time `bun:"activation_date,type:datetime"`                                 // Date when the note should be activated.
	ExpirationDate   time.Time `bun:"expiration_date,type:datetime"`                                 // When does this note expire?
	EntryDate        time.Time `bun:"entry_date,type:datetime"`                                      // date the activity was entered
	Mandatory        string    `bun:"mandatory,type:char(1)"`                                        // Indicates whether the user should be forced to see
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	DefaultInOe      string    `bun:"default_in_oe,type:char(1)"`                                    // This column is unused.
	GeneralNote      string    `bun:"general_note,type:char(1),nullzero"`                            // Determines if a note is a general note rather than
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`          // User who created the record
}
