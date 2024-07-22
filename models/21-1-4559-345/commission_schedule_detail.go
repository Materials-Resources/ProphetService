package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CommissionScheduleDetail struct {
	bun.BaseModel               `bun:"table:commission_schedule_detail"`
	CompanyId                   string    `bun:"company_id,type:varchar(8),pk"`                                         // Unique code that identifies a company.
	CommissionScheduleId        string    `bun:"commission_schedule_id,type:varchar(8),pk"`                             // Default Commission Schedule for computations
	CommissionRuleId            string    `bun:"commission_rule_id,type:varchar(8),pk"`                                 // Commission Rule ID
	SortOrder                   float64   `bun:"sort_order,type:decimal(19,0),pk"`                                      // Sorting Order of rules within a schedule
	Inactive                    string    `bun:"inactive,type:char(1)"`                                                 // Indicates if this commission rule is inactive with
	DeleteFlag                  string    `bun:"delete_flag,type:char(1)"`                                              // Indicates whether this record is logically deleted
	DateCreated                 time.Time `bun:"date_created,type:datetime"`                                            // Indicates the date/time this record was created.
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`                                      // Indicates the date/time this record was last modified.
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`             // ID of the user who last maintained this record
	CommissionScheduleDetailUid int32     `bun:"commission_schedule_detail_uid,type:int,autoincrement,unique,identity"` // Identity column for commission_schedule_detail table.
}
