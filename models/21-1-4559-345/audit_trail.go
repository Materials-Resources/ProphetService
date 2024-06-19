package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type AuditTrail struct {
	bun.BaseModel     `bun:"table:audit_trail"`
	AuditTrailUid     int32     `bun:"audit_trail_uid,type:int,autoincrement,pk"`            // (Identity not for replication) Unique identifier for each record
	SourceAreaCd      int32     `bun:"source_area_cd,type:int"`                              // Code representing transaction or maintenance record where audit_trail record was established
	ColumnChanged     string    `bun:"column_changed,type:varchar(255)"`                     // Column description of value being changed
	Key1Cd            string    `bun:"key1_cd,type:varchar(255)"`                            // Code representing primary column name used to tie audit_trail record to source record
	Key1Value         string    `bun:"key1_value,type:varchar(255)"`                         // Value of primary column name used to tie audit_trail record to source record
	Key2Cd            string    `bun:"key2_cd,type:varchar(255),nullzero"`                   // Code representing secondary column name used to tie audit_trail record to source record
	Key2Value         string    `bun:"key2_value,type:varchar(255),nullzero"`                // Value of secondary column name used to tie audit_trail record to source record
	Key3Cd            string    `bun:"key3_cd,type:varchar(255),nullzero"`                   // Code representing tertiary column name used to tie audit_trail record to source record
	Key3Value         string    `bun:"key3_value,type:varchar(255),nullzero"`                // Value of tertiary column name used to tie audit_trail record to source record
	LineNo            int32     `bun:"line_no,type:int,nullzero"`                            // Line number, where applicable
	InvMastUid        int32     `bun:"inv_mast_uid,type:int,nullzero"`                       // inv_mast_uid, where applicable
	OldValue          string    `bun:"old_value,type:varchar(255),nullzero"`                 // Old Value
	NewValue          string    `bun:"new_value,type:varchar(255),nullzero"`                 // New Value
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`       // Date record created
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"` // Who created the record
	TableChanged      string    `bun:"table_changed,type:varchar(255),nullzero"`
	UidValue          int32     `bun:"uid_value,type:int,nullzero"`
	ColumnDescription string    `bun:"column_description,type:varchar(255),nullzero"`
	AuxiliaryValue    string    `bun:"auxiliary_value,type:varchar(255),nullzero"`
}
