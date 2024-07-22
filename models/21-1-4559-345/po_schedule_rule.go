package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PoScheduleRule struct {
	bun.BaseModel    `bun:"table:po_schedule_rule"`
	PoHdrUid         int32      `bun:"po_hdr_uid,type:int,pk"`                                    // What PO is this schedule rule for?
	ReleaseType      int32      `bun:"release_type,type:int"`                                     // Rule or specific dates?
	RoundType        int32      `bun:"round_type,type:int"`                                       // Round extra qty to first or last release?
	FirstDate        *time.Time `bun:"first_date,type:datetime"`                                  // Date of first release.
	TotalReleases    *int32     `bun:"total_releases,type:int"`                                   // Total number of releases for the rule.
	FrequencyValue   *int32     `bun:"frequency_value,type:int"`                                  // Numeric value in: Every 1 week, Every 3 months, every 2 years.
	FrequencyType    *string    `bun:"frequency_type,type:varchar(8)"`                            // Time value in: Every 1 week, Every 3 months, every 2 years.
	DefaultToAll     string     `bun:"default_to_all,type:char(1),default:('N')"`                 // Default this schedule to all lines?
	DateCreated      time.Time  `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time  `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
