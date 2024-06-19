package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type EccCustomColumnTransferTypeXTable struct {
	bun.BaseModel                        `bun:"table:ecc_custom_column_transfer_type_x_table"`
	EccCustomColumnTransferTypeXTableUid int32     `bun:"ecc_custom_column_transfer_type_x_table_uid,type:int,autoincrement,pk"`
	EccTransferType                      string    `bun:"ecc_transfer_type,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,unique"`
	UdTableName                          string    `bun:"ud_table_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,unique"`
	JoinSql                              string    `bun:"join_sql,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`
	DateCreated                          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified                     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy                     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	SbUpdateJoin                         string    `bun:"sb_update_join,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`
	SbInsertSelect                       string    `bun:"sb_insert_select,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`
	ApplyTableJoins                      string    `bun:"apply_table_joins,type:char(1),nullzero"`
	TriggerKeyColumn                     string    `bun:"trigger_key_column,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`
	TriggerDeletedJoin                   string    `bun:"trigger_deleted_join,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`
	ProcedureKeyColumn                   string    `bun:"procedure_key_column,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`
}
