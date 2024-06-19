package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InventoryReturnLineNotepad struct {
	bun.BaseModel           `bun:"table:inventory_return_line_notepad"`
	InvReturnLineNotepadUid int32     `bun:"inv_return_line_notepad_uid,type:int,pk"`                   // Unique Identifier for record
	NoteId                  float64   `bun:"note_id,type:decimal(19,0),unique"`                         // What is the unique identifier for this supplier note?
	InventoryReturnLineUid  int32     `bun:"inventory_return_line_uid,type:int"`                        // Unique Identifier from inventory_return_line table.
	NotepadClassId          string    `bun:"notepad_class_id,type:varchar(8),nullzero"`                 // What is the class for this note?
	Topic                   string    `bun:"topic,type:varchar(30)"`                                    // The topic of the note for the referenced area.
	Note                    string    `bun:"note,type:text,nullzero"`                                   // What are the contents of the note?
	ActivationDate          time.Time `bun:"activation_date,type:datetime"`                             // When should this note be activated?
	ExpirationDate          time.Time `bun:"expiration_date,type:datetime"`                             // When does this note expire?
	EntryDate               time.Time `bun:"entry_date,type:datetime"`                                  // date the activity was entered
	Mandatory               string    `bun:"mandatory,type:char(1)"`                                    // Should this note be seen by everyone?
	DeleteFlag              string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`      // User who created the record
}
