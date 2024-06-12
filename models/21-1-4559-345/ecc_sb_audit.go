package model

import "github.com/uptrace/bun"

type EccSbAudit struct {
	bun.BaseModel          `bun:"table:ecc_sb_audit"`
	EccSbAuditUid          int32  `bun:"ecc_sb_audit_uid,type:int,pk,identity"`
	EccTransactionTable    string `bun:"ecc_transaction_table,type:varchar(255)"`
	EccInitiatorAuditTable string `bun:"ecc_initiator_audit_table,type:varchar(255)"`
	EccTargetAuditTable    string `bun:"ecc_target_audit_table,type:varchar(255)"`
	SbInitiatorMessage     string `bun:"sb_initiator_message,type:varchar(255)"`
	SbTargetMessage        string `bun:"sb_target_message,type:varchar(255)"`
	SbContract             string `bun:"sb_contract,type:varchar(255)"`
	SbInitiatorQueue       string `bun:"sb_initiator_queue,type:varchar(255)"`
	SbTargetQueue          string `bun:"sb_target_queue,type:varchar(255)"`
	SbInitiatorService     string `bun:"sb_initiator_service,type:varchar(255)"`
	SbTargetService        string `bun:"sb_target_service,type:varchar(255)"`
	AuditTriggerData       string `bun:"audit_trigger_data,type:char"`
	AuditTargetData        string `bun:"audit_target_data,type:char"`
	TransferType           string `bun:"transfer_type,type:varchar(255),nullzero"`
	TableName              string `bun:"table_name,type:varchar(255),nullzero"`
	TriggerName            string `bun:"trigger_name,type:varchar(255),nullzero"`
}
