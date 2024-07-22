package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvBinAudit struct {
	bun.BaseModel       `bun:"table:inv_bin_audit"`
	InvBinAuditUid      int32      `bun:"inv_bin_audit_uid,type:int,autoincrement,identity"`
	InvBinUid           *int32     `bun:"inv_bin_uid,type:int"`
	InvMastUid          int32      `bun:"inv_mast_uid,type:int"`
	LocationId          float64    `bun:"location_id,type:decimal(19,0)"`
	Bin                 string     `bun:"bin,type:varchar(10)"`
	QtyAllocatedBefore  *float64   `bun:"qty_allocated_before,type:decimal(19,9)"`
	QtyAllocatedAfter   *float64   `bun:"qty_allocated_after,type:decimal(19,9)"`
	QtyAllocatedDelta   *float64   `bun:"qty_allocated_delta,type:decimal(19,9)"`
	QuantityBefore      *float64   `bun:"quantity_before,type:decimal(19,9)"`
	QuantityAfter       *float64   `bun:"quantity_after,type:decimal(19,9)"`
	QuantityDelta       *float64   `bun:"quantity_delta,type:decimal(19,9)"`
	LastMaintainedBy    *string    `bun:"last_maintained_by,type:varchar(255)"`
	DateLastModified    *time.Time `bun:"date_last_modified,type:datetime"`
	SqlCurrentTimestamp *time.Time `bun:"sql_current_timestamp,type:datetime"`
	SqlCurrentUser      *string    `bun:"sql_current_user,type:varchar(255)"`
	ActionType          string     `bun:"action_type,type:varchar(255)"`
}
