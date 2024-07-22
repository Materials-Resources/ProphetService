package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ContactSalesrep struct {
	bun.BaseModel      `bun:"table:contact_salesrep"`
	ContactId          string    `bun:"contact_id,type:varchar(16),unique"`                        // What contact deals with this sales representative?
	SalesrepId         string    `bun:"salesrep_id,type:varchar(16),unique"`                       // Who is the actual sales representative for this intersection?
	DeleteFlag         string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated        time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	ContactSalesrepUid int32     `bun:"contact_salesrep_uid,type:int,pk"`                          // Unique identifier for record.
}
