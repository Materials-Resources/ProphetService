package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PoLineSchedule struct {
	bun.BaseModel     `bun:"table:po_line_schedule"`
	PoLineScheduleUid int32     `bun:"po_line_schedule_uid,type:int,pk"`       // Unique Identifier
	PoLineUid         int32     `bun:"po_line_uid,type:int,unique"`            // ID for the PO_line
	ReleaseNo         int32     `bun:"release_no,type:int,unique"`             // Ordered PO release ID
	ReleaseDate       time.Time `bun:"release_date,type:datetime"`             // Date this release should be received.
	ReleaseQty        float64   `bun:"release_qty,type:decimal(19,9)"`         // Amount due in this release
	QtyReceived       float64   `bun:"qty_received,type:decimal(19,9)"`        // Amount actually received from supplier.
	RowStatusFlag     int32     `bun:"row_status_flag,type:int,default:(0)"`   // Indicates current record status.
	DateCreated       time.Time `bun:"date_created,type:datetime"`             // Indicates the date/time this record was created.
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`       // Indicates the date/time this record was last modified.
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30)"`    // ID of the user who last maintained this record
	RevisionCode      int32     `bun:"revision_code,type:int,nullzero"`        // What has happened to this release?  Date changed, Qty Changed?
	ExpediteFlag      string    `bun:"expedite_flag,type:char(1),nullzero"`    // Indicates whether this schedule line should be included on the expedite request.
	OeLineScheduleUid int32     `bun:"oe_line_schedule_uid,type:int,nullzero"` // Unique Identifier of oe_line_schedule table
	QtyReady          float64   `bun:"qty_ready,type:decimal(19,9),nullzero"`  // Quantity ready for container building
}
