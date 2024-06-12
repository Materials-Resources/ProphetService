package model

import "github.com/uptrace/bun"

type FaultToleranceSalesOrders struct {
	bun.BaseModel `bun:"table:fault_tolerance_sales_orders"`
	FtsoUid       int32  `bun:"ftso_uid,type:int,identity"`
	FtrUid        int32  `bun:"ftr_uid,type:int"`
	FttUid        int32  `bun:"ftt_uid,type:int"`
	FtpCode       string `bun:"ftp_code,type:varchar(255)"`
	TableName     string `bun:"table_name,type:varchar(255)"`
	ColumnName    string `bun:"column_name,type:varchar(255)"`
	BeforeValue   string `bun:"before_value,type:varchar(2000),nullzero"`
	AfterValue    string `bun:"after_value,type:varchar(2000),nullzero"`
	AuditInserted string `bun:"audit_inserted,type:char,default:('N')"`
	LineNo        int32  `bun:"line_no,type:int,default:((0))"`
}
