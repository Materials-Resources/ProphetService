package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Activity struct {
	bun.BaseModel    `bun:"table:activity"`
	ActivityId       string    `bun:"activity_id,type:varchar(10),pk"`                           // This is the system generated unique identifier for an activity trans.
	ActivityDesc     string    `bun:"activity_desc,type:varchar(255)"`                           // Description of the activity.
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	HardTouchFlag    string    `bun:"hard_touch_flag,type:char(1),default:('N')"`                // Indicates that the activity involves direct contact with the customer, supplier, etc
}
