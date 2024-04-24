package data

import (
	"github.com/uptrace/bun"
	"time"
)

type QuoteLine struct {
	bun.BaseModel    `bun:"table:quote_line"`
	QuoteLineUid     int32     `bun:"quote_line_uid,type:int,pk"`
	OeLineUid        int32     `bun:"oe_line_uid,type:int"`
	LineCompleteFlag string    `bun:"line_complete_flag,type:char,default:('N')"`
	QtyConverted     float64   `bun:"qty_converted,type:decimal(19,9)"`
	UomConverted     string    `bun:"uom_converted,type:varchar(8)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
}
