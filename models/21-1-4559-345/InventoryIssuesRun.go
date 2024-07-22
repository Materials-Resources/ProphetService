package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Inventoryissuesrun struct {
	bun.BaseModel    `bun:"table:InventoryIssuesRun"`
	Run              int32      `bun:"Run,type:int,pk"`
	InvMastUid       *int32     `bun:"inv_mast_uid,type:int"`
	LocationId       *int32     `bun:"location_id,type:int"`
	DateCreated      time.Time  `bun:"date_created,type:datetime"`
	LastMaintainedBy string     `bun:"last_maintained_by,type:varchar(255)"`
	RunByUser        *string    `bun:"run_by_user,type:varchar(255)"`
	DateRunEnded     *time.Time `bun:"date_run_ended,type:datetime"`
	DateRunStarted   *time.Time `bun:"date_run_started,type:datetime"`
	Server           *string    `bun:"server,type:varchar(255)"`
	DbName           *string    `bun:"db_name,type:varchar(255)"`
	DbVersion        *string    `bun:"db_version,type:varchar(255)"`
	Testuid          *int32     `bun:"TestUID,type:int"`
	Instances        *int32     `bun:"instances,type:int"`
	DateDiffMs       *int32     `bun:"date_diff_ms,type:int"`
	TotalTestUid     *int32     `bun:"total_test_uid,type:int"`
}
