package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CommissionSchedule struct {
	bun.BaseModel          `bun:"table:commission_schedule"`
	CompanyId              string    `bun:"company_id,type:varchar(8),pk"`                                  // Unique code that identifies a company.
	CommissionScheduleId   string    `bun:"commission_schedule_id,type:varchar(8),pk"`                      // User-defined code that identifies a commission schedule.
	CommissionScheduleDesc string    `bun:"commission_schedule_desc,type:varchar(40)"`                      // Description of Commission Schedule ID.
	CommissionScheduleType string    `bun:"commission_schedule_type,type:char(1)"`                          // Determines how the system will search for a rule to apply within a commission schedule.
	Inactive               string    `bun:"inactive,type:char(1)"`                                          // Indicates if this commission rule is inactive.
	DeleteFlag             string    `bun:"delete_flag,type:char(1)"`                                       // Indicates whether this record is logically deleted
	DateCreated            time.Time `bun:"date_created,type:datetime"`                                     // Indicates the date/time this record was created.
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`                               // Indicates the date/time this record was last modified.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`      // ID of the user who last maintained this record
	AppliesTo              string    `bun:"applies_to,type:char(1)"`                                        // Determines whether to apply the commission rule to the entire invoice or the line.
	CommissionScheduleUid  int32     `bun:"commission_schedule_uid,type:int,autoincrement,unique,scanonly"` // Unique identifier for a record
}
