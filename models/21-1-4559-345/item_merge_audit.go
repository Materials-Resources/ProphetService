package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ItemMergeAudit struct {
	bun.BaseModel     `bun:"table:item_merge_audit"`
	ItemMergeAuditUid int32     `bun:"item_merge_audit_uid,type:int,pk,identity"`
	ItemMergeRun      int32     `bun:"item_merge_run,type:int"`
	SourceItemId      string    `bun:"source_item_id,type:varchar(40)"`
	TargetItemId      string    `bun:"target_item_id,type:varchar(40)"`
	ProcessedFlag     string    `bun:"processed_flag,type:char"`
	Message           string    `bun:"message,type:varchar(255)"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	SourceInvMastUid  int32     `bun:"source_inv_mast_uid,type:int"`
	TargetInvMastUid  int32     `bun:"target_inv_mast_uid,type:int"`
}
