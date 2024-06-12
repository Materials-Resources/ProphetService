package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CellDefinition struct {
	bun.BaseModel    `bun:"table:cell_definition"`
	FinReportId      string    `bun:"fin_report_id,type:varchar(15),pk"`
	Cell             string    `bun:"cell,type:varchar(10),pk"`
	Field            string    `bun:"field,type:char"`
	ColumnNo         float64   `bun:"column_no,type:decimal(19,0)"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	AccountMask      string    `bun:"account_mask,type:varchar(32),nullzero"`
}
