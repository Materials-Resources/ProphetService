package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CustPartNoNotepad struct {
	bun.BaseModel    `bun:"table:cust_part_no_notepad"`
	NoteId           float64   `bun:"note_id,type:decimal(19,0),pk"`                             // Unique ID for this note.
	CompanyId        string    `bun:"company_id,type:varchar(8)"`                                // Unique code that identifies a company.
	CustomerId       float64   `bun:"customer_id,type:decimal(19,0)"`                            // ID referencing the customer for whom this note applies.
	TheirItemId      string    `bun:"their_item_id,type:varchar(40)"`                            // Your customers part no.
	Topic            string    `bun:"topic,type:varchar(30)"`                                    // The topic of the note for the referenced area.
	Note             *string   `bun:"note,type:text"`                                            // The text of the note.
	ActivationDate   time.Time `bun:"activation_date,type:datetime"`                             // The date when the notepad record becomes activated.
	ExpirationDate   time.Time `bun:"expiration_date,type:datetime"`                             // The date when the notepad record expires.
	EntryDate        time.Time `bun:"entry_date,type:datetime"`                                  // The date when this note was entered.
	NotepadClass     *string   `bun:"notepad_class,type:varchar(8)"`                             // The class for this note.
	Mandatory        string    `bun:"mandatory,type:char(1)"`                                    // Indicates whether this notepad record is mandatory or not.
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`      // User who created the record
}
