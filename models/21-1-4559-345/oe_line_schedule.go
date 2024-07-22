package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type OeLineSchedule struct {
	bun.BaseModel           `bun:"table:oe_line_schedule"`
	OrderNo                 string     `bun:"order_no,type:varchar(8),unique"`                               // Order number
	ReleaseNo               float64    `bun:"release_no,type:decimal(19,0),unique"`                          // Release number
	ReleaseDate             time.Time  `bun:"release_date,type:datetime"`                                    // Release date
	ReleaseQty              float64    `bun:"release_qty,type:decimal(19,9)"`                                // Release quantity
	AllocatedQty            float64    `bun:"allocated_qty,type:decimal(19,9)"`                              // Allocated quantity
	ExpediteValue           float64    `bun:"expedite_value,type:decimal(19,0)"`                             // Time before release date the schedule should be expedited.
	ExpediteType            string     `bun:"expedite_type,type:varchar(8)"`                                 // Determines whether the expedite time is in days, weeks, etc.
	PickValue               float64    `bun:"pick_value,type:decimal(19,0)"`                                 // Time before release date the schedule should be picked.
	PickType                string     `bun:"pick_type,type:varchar(8)"`                                     // Determines whether the pick time is in days, weeks, etc.
	Printed                 string     `bun:"printed,type:char(1)"`                                          // Indicates whether this schedule line has been printed
	DateCreated             time.Time  `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified        time.Time  `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy        string     `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	LineNo                  float64    `bun:"line_no,type:decimal(19,0),unique"`                             // What line number of the invoice does this invoice line to sales representative refer to?
	PickDate                *time.Time `bun:"pick_date,type:datetime"`                                       // Date that items on the release schedule should be picked.
	Disposition             *string    `bun:"disposition,type:char(1)"`                                      // Disposition of the open quantity.
	ExpediteDate            *time.Time `bun:"expedite_date,type:datetime"`                                   // Date that the system will allocate the items to an order
	QtyPicked               *float64   `bun:"qty_picked,type:decimal(19,9)"`                                 // Quantity picked
	QtyInvoiced             *float64   `bun:"qty_invoiced,type:decimal(19,9)"`                               // Quantity invoiced
	QtyStaged               *float64   `bun:"qty_staged,type:decimal(19,9)"`                                 // Quantity staged but not shipped.
	QtyCanceled             *float64   `bun:"qty_canceled,type:decimal(19,9)"`                               // Quantity canceled
	OeLineScheduleUid       int32      `bun:"oe_line_schedule_uid,type:int,pk"`                              // Unique identifier for the record.
	InvMastUid              int32      `bun:"inv_mast_uid,type:int"`                                         // Unique identifier for the item id.
	EditFlag                string     `bun:"edit_flag,type:char(1),default:('N')"`                          // Signifies whether the schedule record has been edited.
	AcknowledgementDate     *time.Time `bun:"acknowledgement_date,type:datetime"`                            // According to the spec for Feature 22107, the aknowledgement date is a promised date
	PromiseDateExtendedDesc *string    `bun:"promise_date_extended_desc,type:varchar(8000)"`                 // Notes for the release schedule promise date.
	OriginalPromiseDate     *time.Time `bun:"original_promise_date,type:datetime"`                           // The original promise date.
	PromiseDate             *time.Time `bun:"promise_date,type:datetime"`                                    // The current promise date.
	PromiseDateEditedDate   *time.Time `bun:"promise_date_edited_date,type:datetime"`                        // The date that the promise date was edited.
	ReleaseStatusFlag       string     `bun:"release_status_flag,type:char(1),default:('F')"`                // Custom (F32834): defines the current status of this release.  Valid values are 'F'irm (material will be allocated/shipped as normal) and 'P'lanned (account for possible future requirement).
	PriceRecalculatedFlag   *string    `bun:"price_recalculated_flag,type:char(1)"`                          // Indicates that pricing was recalculated for this release (based on the date of release)..
	Edi830AccumQty          *float64   `bun:"edi830_accum_qty,type:decimal(19,9)"`                           // Custom: cumulative release qty from the EDI 830 planning schedule release import.  Only populated during the import.
}
