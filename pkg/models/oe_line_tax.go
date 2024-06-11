package models

import (
	"database/sql"
	"github.com/uptrace/bun"
	"time"
)

type OeLineTax struct {
	bun.BaseModel    `bun:"table:oe_line_tax"`
	OrderNumber      string        `bun:"order_number,type:varchar(8),pk"`
	LineNumber       float64       `bun:"line_number,type:decimal(19,0),pk"`
	JurisdictionId   string        `bun:"jurisdiction_id,type:varchar(10),pk"`
	Taxable          string        `bun:"taxable,type:char"`
	DeleteFlag       string        `bun:"delete_flag,type:char"`
	DateCreated      time.Time     `bun:"date_created,type:datetime"`
	DateLastModified time.Time     `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string        `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	LineType         sql.NullInt32 `bun:"line_type,type:int,nullzero"`
}
