package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DemandLevel struct {
	bun.BaseModel       `bun:"table:demand_level"`
	DemandLevelUid      int32     `bun:"demand_level_uid,type:int,autoincrement,identity,pk"`          // Unique Identifier for table.
	ServiceLevelPctGoal float64   `bun:"service_level_pct_goal,type:decimal(19,2),nullzero"`           // Percent of service level for safety stock.
	StockOutDeviation   float64   `bun:"stock_out_deviation,type:decimal(19,2),nullzero"`              // Standard deviation for stock outs.
	BackorderDeviation  float64   `bun:"backorder_deviation,type:decimal(19,2),nullzero"`              // Standard deviation for backorders
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
