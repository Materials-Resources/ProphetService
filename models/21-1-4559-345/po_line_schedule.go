package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PoLineSchedule struct {
	bun.BaseModel     `bun:"table:po_line_schedule"`
	PoLineScheduleUid int32     `bun:"po_line_schedule_uid,type:int,pk"`
	PoLineUid         int32     `bun:"po_line_uid,type:int"`
	ReleaseNo         int32     `bun:"release_no,type:int"`
	ReleaseDate       time.Time `bun:"release_date,type:datetime"`
	ReleaseQty        float64   `bun:"release_qty,type:decimal(19,9)"`
	QtyReceived       float64   `bun:"qty_received,type:decimal(19,9)"`
	RowStatusFlag     int32     `bun:"row_status_flag,type:int,default:(0)"`
	DateCreated       time.Time `bun:"date_created,type:datetime"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30)"`
	RevisionCode      int32     `bun:"revision_code,type:int,nullzero"`
	ExpediteFlag      string    `bun:"expedite_flag,type:char,nullzero"`
	OeLineScheduleUid int32     `bun:"oe_line_schedule_uid,type:int,nullzero"`
	QtyReady          float64   `bun:"qty_ready,type:decimal(19,9),nullzero"`
}
