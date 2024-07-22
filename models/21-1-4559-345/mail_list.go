package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type MailList struct {
	bun.BaseModel    `bun:"table:mail_list"`
	ListId           float64   `bun:"list_id,type:decimal(19,0),pk"`                             // Identifier of mailing list.
	ListDesc         *string   `bun:"list_desc,type:varchar(40)"`                                // What is this mail list for?
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
}
