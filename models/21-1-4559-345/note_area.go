package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type NoteArea struct {
	bun.BaseModel    `bun:"table:note_area"`
	NoteId           float64   `bun:"note_id,type:decimal(19,0),pk"`                                 // What is the unique identifier for this note?
	Area             string    `bun:"area,type:varchar(40),pk"`                                      // The area that the area is used.
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
}
