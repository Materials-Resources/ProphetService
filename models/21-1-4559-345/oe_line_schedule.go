package model

type OeLineSchedule struct {
	bun.BaseModel           `bun:"table:oe_line_schedule"`
	OrderNo                 string    `bun:"order_no,type:varchar(8)"`
	ReleaseNo               float64   `bun:"release_no,type:decimal(19,0)"`
	ReleaseDate             time.Time `bun:"release_date,type:datetime"`
	ReleaseQty              float64   `bun:"release_qty,type:decimal(19,9)"`
	AllocatedQty            float64   `bun:"allocated_qty,type:decimal(19,9)"`
	ExpediteValue           float64   `bun:"expedite_value,type:decimal(19,0)"`
	ExpediteType            string    `bun:"expedite_type,type:varchar(8)"`
	PickValue               float64   `bun:"pick_value,type:decimal(19,0)"`
	PickType                string    `bun:"pick_type,type:varchar(8)"`
	Printed                 string    `bun:"printed,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	LineNo                  float64   `bun:"line_no,type:decimal(19,0)"`
	PickDate                time.Time `bun:"pick_date,type:datetime,nullzero"`
	Disposition             string    `bun:"disposition,type:char,nullzero"`
	ExpediteDate            time.Time `bun:"expedite_date,type:datetime,nullzero"`
	QtyPicked               float64   `bun:"qty_picked,type:decimal(19,9),nullzero"`
	QtyInvoiced             float64   `bun:"qty_invoiced,type:decimal(19,9),nullzero"`
	QtyStaged               float64   `bun:"qty_staged,type:decimal(19,9),nullzero"`
	QtyCanceled             float64   `bun:"qty_canceled,type:decimal(19,9),nullzero"`
	OeLineScheduleUid       int32     `bun:"oe_line_schedule_uid,type:int,pk"`
	InvMastUid              int32     `bun:"inv_mast_uid,type:int"`
	EditFlag                string    `bun:"edit_flag,type:char,default:('N')"`
	AcknowledgementDate     time.Time `bun:"acknowledgement_date,type:datetime,nullzero"`
	PromiseDateExtendedDesc string    `bun:"promise_date_extended_desc,type:varchar(8000),nullzero"`
	OriginalPromiseDate     time.Time `bun:"original_promise_date,type:datetime,nullzero"`
	PromiseDate             time.Time `bun:"promise_date,type:datetime,nullzero"`
	PromiseDateEditedDate   time.Time `bun:"promise_date_edited_date,type:datetime,nullzero"`
	ReleaseStatusFlag       string    `bun:"release_status_flag,type:char,default:('F')"`
	PriceRecalculatedFlag   string    `bun:"price_recalculated_flag,type:char,nullzero"`
	Edi830AccumQty          float64   `bun:"edi830_accum_qty,type:decimal(19,9),nullzero"`
}
