package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type BusinessRuleEventExtdInfo struct {
	bun.BaseModel                `bun:"table:business_rule_event_extd_info"`
	BusinessRuleEventExtdInfoUid int32     `bun:"business_rule_event_extd_info_uid,type:int,autoincrement,identity,pk"`                                                                                                                                       // UID for this table
	BusinessRuleEventUid         int32     `bun:"business_rule_event_uid,type:int"`                                                                                                                                                                           // FK column to buisness_rule_event table
	TriggeredFrom                string    `bun:"triggered_from,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`        // Describes from where and how the event is triggered and
	Keys                         string    `bun:"keys,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`                  // Key Columns that are being used, mainly applicable to update hooks.
	Context                      string    `bun:"context,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`               // Enumerates the context values, mainly applicable to update hooks.
	SpecialReturnValues          string    `bun:"special_return_values,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"` // If hook acts on or affects very specific fields, they should be listed here.
	SeeAlso                      string    `bun:"see_also,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`              // List of other related hooks.
	DateCreated                  time.Time `bun:"date_created,type:datetime,default:(getdate())"`                                                                                                                                                             // Date and time the record was originally created
	CreatedBy                    string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`                                                                                                                                                       // User who created the record
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`                                                                                                                                                       // Date and time the record was modified
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`                                                                                                                                               // User who last changed the record
}
