package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ProcessXTransactionDetail struct {
	bun.BaseModel          `bun:"table:process_x_transaction_detail"`
	ProcessXTransDetailUid int32     `bun:"process_x_trans_detail_uid,type:int,pk"`
	ProcessXTransactionUid int32     `bun:"process_x_transaction_uid,type:int"`
	StageUid               int32     `bun:"stage_uid,type:int"`
	SequenceNo             int16     `bun:"sequence_no,type:smallint"`
	StartDate              time.Time `bun:"start_date,type:datetime,nullzero"`
	LagTimeHours           int16     `bun:"lag_time_hours,type:smallint"`
	QtyIn                  float64   `bun:"qty_in,type:decimal(19,9)"`
	QtyOut                 float64   `bun:"qty_out,type:decimal(19,9)"`
	QtyLost                float64   `bun:"qty_lost,type:decimal(19,9)"`
	QtyPartiallyReceived   float64   `bun:"qty_partially_received,type:decimal(19,9)"`
	RowStatusFlag          int16     `bun:"row_status_flag,type:smallint"`
	Cost                   float64   `bun:"cost,type:decimal(19,9)"`
	CostUom                string    `bun:"cost_uom,type:varchar(8),nullzero"`
	StageQtyUom            string    `bun:"stage_qty_uom,type:varchar(8),nullzero"`
	DateCreated            time.Time `bun:"date_created,type:datetime"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	QtySplit               float64   `bun:"qty_split,type:decimal(19,9),default:(0)"`
	CostUnitSize           float64   `bun:"cost_unit_size,type:decimal(19,4)"`
	QtyUnitSize            float64   `bun:"qty_unit_size,type:decimal(19,4)"`
	CostType               int16     `bun:"cost_type,type:smallint,nullzero"`
	PoRequired             string    `bun:"po_required,type:char,default:('N')"`
	AllowPartials          string    `bun:"allow_partials,type:char,default:('Y')"`
	StageWipAccountNumber  string    `bun:"stage_wip_account_number,type:varchar(32),nullzero"`
	StageCd                string    `bun:"stage_cd,type:varchar(255),nullzero"`
	StageDesc              string    `bun:"stage_desc,type:varchar(8000),nullzero"`
	EstimatedHours         float64   `bun:"estimated_hours,type:decimal(6,2),nullzero"`
	SupplierId             float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`
	VendorId               float64   `bun:"vendor_id,type:decimal(19,0),nullzero"`
	MinPoCost              float64   `bun:"min_po_cost,type:decimal(19,9),nullzero"`
	BuyerId                string    `bun:"buyer_id,type:varchar(16),nullzero"`
	DivisionId             float64   `bun:"division_id,type:decimal(19,0),nullzero"`
	StagePoDesc            string    `bun:"stage_po_desc,type:varchar(255),nullzero"`
	PurchasingWeight       float64   `bun:"purchasing_weight,type:decimal(19,9),nullzero"`
	WipInvMastUid          int32     `bun:"wip_inv_mast_uid,type:int,nullzero"`
	BackflushFlag          string    `bun:"backflush_flag,type:char,nullzero"`
	TravelerRequiredFlag   string    `bun:"traveler_required_flag,type:char,default:('N')"`
	PostedGlPoQty          float64   `bun:"posted_gl_po_qty,type:decimal(19,9),default:((0))"`
	ConsolidatableFlag     string    `bun:"consolidatable_flag,type:char,nullzero"`
	ContainerInvMastUid    int32     `bun:"container_inv_mast_uid,type:int,nullzero"`
	ContainerUom           string    `bun:"container_uom,type:varchar(8),nullzero"`
	ContainerFlag          string    `bun:"container_flag,type:char,nullzero"`
	CommentProcess         string    `bun:"comment_process,type:char,default:('N')"`
	Comment                string    `bun:"comment,type:varchar,nullzero"`
}
