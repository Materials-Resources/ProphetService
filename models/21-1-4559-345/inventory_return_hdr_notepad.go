package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InventoryReturnHdrNotepad struct {
	bun.BaseModel                `bun:"table:inventory_return_hdr_notepad"`
	InventoryReturnHdrNotepadUid int32     `bun:"inventory_return_hdr_notepad_uid,type:int,pk"`              // Unique identifier for this record.
	NoteId                       float64   `bun:"note_id,type:decimal(19,0),unique"`                         // What is the unique identifier for this supplier note?
	InventoryReturnHdrUid        int32     `bun:"inventory_return_hdr_uid,type:int"`                         // Unique Identifier for the inventory_return_hdr table that is specific to this note.
	NotepadClassId               string    `bun:"notepad_class_id,type:varchar(8),nullzero"`                 // What is the unique identifier for this notepad cla
	Topic                        string    `bun:"topic,type:varchar(30)"`                                    // The topic of the note for the referenced area.
	Note                         string    `bun:"note,type:text,nullzero"`                                   // What are the contents of the note?
	ActivationDate               time.Time `bun:"activation_date,type:datetime"`                             // When should this note be activated?
	ExpirationDate               time.Time `bun:"expiration_date,type:datetime"`                             // When does this note expire?
	EntryDate                    time.Time `bun:"entry_date,type:datetime"`                                  // Date indicating when this note was entered.
	Mandatory                    string    `bun:"mandatory,type:char(1)"`                                    // Should everyone see this supplier note?
	DeleteFlag                   string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated                  time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	CreatedBy                    string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`      // User who created the record
}
