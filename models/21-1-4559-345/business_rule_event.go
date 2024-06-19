package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type BusinessRuleEvent struct {
	bun.BaseModel        `bun:"table:business_rule_event"`
	BusinessRuleEventUid int32     `bun:"business_rule_event_uid,type:int,autoincrement,identity,pk"`   // UID for this table.
	PublishedEventName   string    `bun:"published_event_name,type:varchar(255)"`                       // Event name as appears in the publish statement in P21 code.
	DisplayName          string    `bun:"display_name,type:varchar(255)"`                               // Event name as it will be shown to the user.
	Description          string    `bun:"description,type:varchar(255)"`                                // Explanation of what event is and where it's called from.
	InternalOnlyFlag     string    `bun:"internal_only_flag,type:char(1),default:('N')"`                // Determines if the event is available for users to subscribe to.
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	AllowNewRowsFlag     string    `bun:"allow_new_rows_flag,type:char(1),default:('N')"`               // Indicates whether event support adding new rows in P21
}
