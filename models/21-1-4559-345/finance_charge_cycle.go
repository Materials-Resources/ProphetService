package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type FinanceChargeCycle struct {
	bun.BaseModel    `bun:"table:finance_charge_cycle"`
	FcCycleId        string    `bun:"fc_cycle_id,type:char(1),pk"`                               // What is the unique finance charge cycle identifier
	FcCycleDesc      string    `bun:"fc_cycle_desc,type:varchar(40)"`                            // What is this finance charge cycle for?
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
