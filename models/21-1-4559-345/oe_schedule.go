package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type OeSchedule struct {
	bun.BaseModel     `bun:"table:oe_schedule"`
	OrderNumber       string    `bun:"order_number,type:varchar(8),pk"`                               // Order Number associated with this schedule record.
	ReleaseType       string    `bun:"release_type,type:varchar(30)"`                                 // Method of creating release schedules (Specific Dates/Rule).
	RoundType         string    `bun:"round_type,type:varchar(30)"`                                   // Release to round quantities on (First/Last).
	FirstDate         time.Time `bun:"first_date,type:datetime,nullzero"`                             // First release date for Rule type schedules.
	TotalReleases     float64   `bun:"total_releases,type:decimal(19,0),nullzero"`                    // Total number of releases for Rule type schedules.
	FrequencyValue    float64   `bun:"frequency_value,type:decimal(19,0),nullzero"`                   // Time in terms of frequency_type between releases for Rule type schedules
	FrequencyType     string    `bun:"frequency_type,type:varchar(8),nullzero"`                       // Unit to determine frequency_value for Rule type schedules (Day, Month, etc).
	ExpediteValue     float64   `bun:"expedite_value,type:decimal(19,0),nullzero"`                    // Time in terms of expedite_type to expedite individual schedules.
	ExpediteType      string    `bun:"expedite_type,type:varchar(8),nullzero"`                        // Unit to determine expedite_value (Day, Month, etc).
	PickValue         float64   `bun:"pick_value,type:decimal(19,0),nullzero"`                        // Time in terms of pick_type to pick individual schedules.
	PickType          string    `bun:"pick_type,type:varchar(8),nullzero"`                            // Unit to determine pick_value (Day, Month, etc).
	DateCreated       time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	DefaultToAll      string    `bun:"default_to_all,type:char(1)"`                                   // Value to default current record values for all lines on this Order.
	ApplyToLineMethod int32     `bun:"apply_to_line_method,type:int,nullzero"`                        // Custom - Which method should be used to apply this schedule to line items?
}
