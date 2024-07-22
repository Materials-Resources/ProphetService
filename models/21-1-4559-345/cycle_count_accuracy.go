package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CycleCountAccuracy struct {
	bun.BaseModel         `bun:"table:cycle_count_accuracy"`
	CycleCountAccuracyUid int32     `bun:"cycle_count_accuracy_uid,type:int,pk"` // Unique identifier for a cycle_count_accuracy record.
	CycleCountHdrUid      int32     `bun:"cycle_count_hdr_uid,type:int"`         // Unique identifier for a cycle count associated with this report.
	AdjustmentNumber      *float64  `bun:"adjustment_number,type:decimal(19,0)"` // The adjustment number associated with the cycle count.
	ItemsOnCount          int32     `bun:"items_on_count,type:int"`              // The number of items counted in the cycle count.
	EditedItemsOnCount    int32     `bun:"edited_items_on_count,type:int"`       // The number of items with a discrepency on the cycle count.
	RowStatusFlag         int32     `bun:"row_status_flag,type:int"`             // Indicates current record status.
	DateCreated           time.Time `bun:"date_created,type:datetime"`           // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`     // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`  // ID of the user who last maintained this record
}
