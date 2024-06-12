package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CycleCountAccuracy struct {
	bun.BaseModel         `bun:"table:cycle_count_accuracy"`
	CycleCountAccuracyUid int32     `bun:"cycle_count_accuracy_uid,type:int,pk"`
	CycleCountHdrUid      int32     `bun:"cycle_count_hdr_uid,type:int"`
	AdjustmentNumber      float64   `bun:"adjustment_number,type:decimal(19,0),nullzero"`
	ItemsOnCount          int32     `bun:"items_on_count,type:int"`
	EditedItemsOnCount    int32     `bun:"edited_items_on_count,type:int"`
	RowStatusFlag         int32     `bun:"row_status_flag,type:int"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`
}
