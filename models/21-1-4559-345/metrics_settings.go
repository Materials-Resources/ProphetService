package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type MetricsSettings struct {
	bun.BaseModel    `bun:"table:metrics_settings"`
	SettingUid       int32     `bun:"setting_uid,type:int,autoincrement,scanonly,pk"`                                                                                                                                                                                  // setting_uid
	MajorKey         string    `bun:"major_key,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`                                  // major_key
	MinorKey         string    `bun:"minor_key,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`                                  // minor_key
	Value            string    `bun:"value,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`                             // value
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`                                                                                                                                                                                  // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`                                                                                                                                                                            // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,default:(suser_sname())"` // User who last changed the record
}
