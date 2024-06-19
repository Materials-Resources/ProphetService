package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Salutation struct {
	bun.BaseModel    `bun:"table:salutation"`
	Salutation       string    `bun:"salutation,type:varchar(4),pk"`                             // What is the unique ID for the salutation?
	FullName         string    `bun:"full_name,type:varchar(30)"`                                // What is the text for the salutation?
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	RowStatusFlag    int16     `bun:"row_status_flag,type:smallint,default:(704)"`               // Indicates current record status.
}
