package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Note struct {
	bun.BaseModel           `bun:"table:note"`
	NoteUid                 int32     `bun:"note_uid,type:int,pk"`                                         // Note UID
	DocumentUid             int32     `bun:"document_uid,type:int"`                                        // UID of the table that uses the note
	NotepadClassId          string    `bun:"notepad_class_id,type:varchar(8),nullzero"`                    // Class of note
	Topic                   string    `bun:"topic,type:varchar(255)"`                                      // Title of the note
	Note                    string    `bun:"note,type:text,nullzero"`                                      // Body of the note
	ActivationDate          time.Time `bun:"activation_date,type:datetime"`                                // Date that the note begins to be viewable
	ExpirationDate          time.Time `bun:"expiration_date,type:datetime"`                                // Date when the note ends being viewable
	EntryDate               time.Time `bun:"entry_date,type:datetime"`                                     // Date the note is created
	Mandatory               string    `bun:"mandatory,type:char(1)"`                                       // Whether mandatory to view
	RowStatusFlag           int32     `bun:"row_status_flag,type:int"`                                     // Row status
	NoteTypeCd              int32     `bun:"note_type_cd,type:int"`                                        // Type of note - connects with external UID
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	CreateOrderLineNoteFlag string    `bun:"create_order_line_note_flag,type:char(1),default:('N')"`       // Custom (F31997): set to Y -yes when this note is created in response to a new item being created via an assembly decoder item configuration template.  Determines if this note should be copied to an order line note
}
