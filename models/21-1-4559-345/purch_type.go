package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchType struct {
	bun.BaseModel    `bun:"table:purch_type"`
	TypeId           string    `bun:"type_id,type:varchar(8),pk"`                                // Enter the type ID
	TypeDesc         string    `bun:"type_desc,type:varchar(30)"`                                // What is this purchase type for?
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
