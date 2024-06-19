package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ScriptCommand struct {
	bun.BaseModel      `bun:"table:script_command"`
	ScriptCommandUid   int32     `bun:"script_command_uid,type:int,autoincrement,scanonly,pk"`                                                                                                                                                             // Script command UID
	CommandName        string    `bun:"command_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`                 // Command name
	CommandDescription string    `bun:"command_description,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"` // Command description
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`                                                                                                                                                                    // Date and time the record was originally created
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`                                                                                                                                                              // User who created the record
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`                                                                                                                                                              // Date and time the record was modified
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`                                                                                                                                                      // User who last changed the record
}
