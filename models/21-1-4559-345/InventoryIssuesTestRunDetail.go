package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Inventoryissuestestrundetail struct {
	bun.BaseModel `bun:"table:InventoryIssuesTestRunDetail"`
	Run           int32     `bun:"run,type:int"`
	Testuid       int32     `bun:"TestUID,type:int"`
	Instances     int32     `bun:"instances,type:int"`
	DateStarted   time.Time `bun:"date_started,type:datetime"`
	DateCompleted time.Time `bun:"date_completed,type:datetime"`
	DateDiffMs    int32     `bun:"date_diff_ms,type:int"`
	TotalTestUid  *int32    `bun:"total_test_uid,type:int"`
}
