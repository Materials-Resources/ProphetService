package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type AccntGroupMx struct {
	bun.BaseModel    `bun:"table:accnt_group_mx"`
	AccntGroupMxUid  int32     `bun:"accnt_group_mx_uid,type:int,autoincrement,identity,pk"`        // UID for account group. Identity.
	AccntGroupMxId   string    `bun:"accnt_group_mx_id,type:varchar(255)"`                          // account group id
	GroupDescription string    `bun:"group_description,type:varchar(255),nullzero"`                 // Long group description. Not null values defined by SAT at this moment.
	SequenceNo       int32     `bun:"sequence_no,type:int,default:((0))"`                           // Sequence number for ordering records. Will be used in case SAT defines new records.
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:((704))"`                     // Record status flag
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	Level            int32     `bun:"level,type:int,nullzero"`                                      // Level
}
