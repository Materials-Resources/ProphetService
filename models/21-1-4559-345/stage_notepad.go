package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type StageNotepad struct {
	bun.BaseModel    `bun:"table:stage_notepad"`
	StageNotepadUid  int32     `bun:"stage_notepad_uid,type:int,pk"`                                // What is the unique identifier for this table
	StageUid         int32     `bun:"stage_uid,type:int"`                                           // What stage does this note belong to?
	NoteId           float64   `bun:"note_id,type:decimal(19,0)"`                                   // What is the unique identifier for this process note?
	NotepadClassId   string    `bun:"notepad_class_id,type:varchar(8),nullzero"`                    // What is the class for this note?
	Topic            string    `bun:"topic,type:varchar(30)"`                                       // Topic - like a Header for the note
	Note             string    `bun:"note,type:text,nullzero"`                                      // What are the contents of the note?
	ActivationDate   time.Time `bun:"activation_date,type:datetime"`                                // When should this note be activated?
	ExpirationDate   time.Time `bun:"expiration_date,type:datetime"`                                // When does this note expire?
	EntryDate        time.Time `bun:"entry_date,type:datetime"`                                     // When was this note entered?
	Mandatory        string    `bun:"mandatory,type:char(1)"`                                       // Is this a Mandatory note? Y or N
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                     // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
