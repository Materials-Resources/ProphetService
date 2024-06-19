package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ItemMergeAudit struct {
	bun.BaseModel     `bun:"table:item_merge_audit"`
	ItemMergeAuditUid int32     `bun:"item_merge_audit_uid,type:int,autoincrement,pk"`       // Identifier
	ItemMergeRun      int32     `bun:"item_merge_run,type:int"`                              // Item Merge Run Number
	SourceItemId      string    `bun:"source_item_id,type:varchar(40)"`                      // Source Item ID
	TargetItemId      string    `bun:"target_item_id,type:varchar(40)"`                      // Target Item ID
	ProcessedFlag     string    `bun:"processed_flag,type:char(1)"`                          // Processed Flag
	Message           string    `bun:"message,type:varchar(255)"`                            // Message returned by the Item Merge Process
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`       // Date and time the record was originally created
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"` // User who created the record
	SourceInvMastUid  int32     `bun:"source_inv_mast_uid,type:int"`                         // Indicates unique identifier of the source item.
	TargetInvMastUid  int32     `bun:"target_inv_mast_uid,type:int"`                         // Indicates unique identifier of the target item.
}
