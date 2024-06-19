package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PoSchedule struct {
	bun.BaseModel    `bun:"table:po_schedule"`
	PoScheduleUid    int32     `bun:"po_schedule_uid,type:int,pk"`                               // unique identifier for PO schedule records
	PoHdrUid         int32     `bun:"po_hdr_uid,type:int,unique"`                                // Unique identifier for PO records associated with this schedule
	ReleaseNo        int32     `bun:"release_no,type:int,unique"`                                // Release number associated with this PO Schedule
	ReleaseDate      time.Time `bun:"release_date,type:datetime"`                                // Release Date associated with this schedule
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
