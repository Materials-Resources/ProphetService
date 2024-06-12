package model

import (
	"github.com/uptrace/bun"
	"time"
)

type OeScheduleDetail struct {
	bun.BaseModel           `bun:"table:oe_schedule_detail"`
	OrderNo                 string    `bun:"order_no,type:varchar(8),pk"`
	ReleaseNo               float64   `bun:"release_no,type:decimal(19,0),pk"`
	ReleaseDate             time.Time `bun:"release_date,type:datetime"`
	ExpediteValue           float64   `bun:"expedite_value,type:decimal(19,0)"`
	ExpediteType            string    `bun:"expedite_type,type:varchar(8)"`
	PickValue               float64   `bun:"pick_value,type:decimal(19,0)"`
	PickType                string    `bun:"pick_type,type:varchar(8)"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	PromiseDateExtendedDesc string    `bun:"promise_date_extended_desc,type:varchar(8000),nullzero"`
	OriginalPromiseDate     time.Time `bun:"original_promise_date,type:datetime,nullzero"`
	PromiseDate             time.Time `bun:"promise_date,type:datetime,nullzero"`
	PromiseDateEditedDate   time.Time `bun:"promise_date_edited_date,type:datetime,nullzero"`
}
