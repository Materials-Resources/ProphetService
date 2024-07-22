package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type FaultToleranceAuditTrail struct {
	bun.BaseModel `bun:"table:fault_tolerance_audit_trail"`
	FtatUid       int32     `bun:"ftat_uid,type:int,autoincrement,identity"`
	FtrUid        int32     `bun:"ftr_uid,type:int,pk"`
	FtaUid        int32     `bun:"fta_uid,type:int,pk"`
	FtpUid        int32     `bun:"ftp_uid,type:int,pk"`
	FttUid        int32     `bun:"ftt_uid,type:int,pk"`
	BeforeValue   *string   `bun:"before_value,type:varchar(2000)"`
	AfterValue    *string   `bun:"after_value,type:varchar(2000)"`
	DateCreated   time.Time `bun:"date_created,type:datetime"`
	AuditUserName *string   `bun:"audit_user_name,type:varchar(255)"`
	AuditLogin    *string   `bun:"audit_login,type:varchar(255)"`
	ColumnName    *string   `bun:"column_name,type:varchar(255)"`
	TableName     *string   `bun:"table_name,type:varchar(255)"`
	LineNo        string    `bun:"line_no,type:varchar(255),default:((0)),pk"`
}
