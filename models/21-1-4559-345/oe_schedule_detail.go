package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type OeScheduleDetail struct {
	bun.BaseModel           `bun:"table:oe_schedule_detail"`
	OrderNo                 string     `bun:"order_no,type:varchar(8),pk"`                                   // Order Number associated with this schedule record.
	ReleaseNo               float64    `bun:"release_no,type:decimal(19,0),pk"`                              // Unique release number for this order.
	ReleaseDate             time.Time  `bun:"release_date,type:datetime"`                                    // Date on which this release should be shipped.
	ExpediteValue           float64    `bun:"expedite_value,type:decimal(19,0)"`                             // Time in terms of expedite_type to expedite this release.
	ExpediteType            string     `bun:"expedite_type,type:varchar(8)"`                                 // Unit to determine expedite_value (Day, Month, etc).
	PickValue               float64    `bun:"pick_value,type:decimal(19,0)"`                                 // Time in terms of pick_type to pick individual release.
	PickType                string     `bun:"pick_type,type:varchar(8)"`                                     // Unit to determine pick_value (Day, Month, etc).
	DateCreated             time.Time  `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified        time.Time  `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy        string     `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	PromiseDateExtendedDesc *string    `bun:"promise_date_extended_desc,type:varchar(8000)"`                 // Notes for the hdr release schedule promise date.
	OriginalPromiseDate     *time.Time `bun:"original_promise_date,type:datetime"`                           // The original promise date.
	PromiseDate             *time.Time `bun:"promise_date,type:datetime"`                                    // The current promise date.
	PromiseDateEditedDate   *time.Time `bun:"promise_date_edited_date,type:datetime"`                        // The date that the promise date was edited.
}
