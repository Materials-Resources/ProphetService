package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DocumentLineBin struct {
	bun.BaseModel         `bun:"table:document_line_bin"`
	DocumentNo            float64   `bun:"document_no,type:decimal(19,0)"`
	LineNo                int32     `bun:"line_no,type:int"`
	DocumentType          string    `bun:"document_type,type:char(2)"`
	BinCd                 string    `bun:"bin_cd,type:varchar(40)"`
	QtyApplied            float64   `bun:"qty_applied,type:decimal(19,9)"`
	UnitQuantity          float64   `bun:"unit_quantity,type:decimal(19,9)"`
	UnitOfMeasure         string    `bun:"unit_of_measure,type:varchar(8)"`
	DocumentCd            string    `bun:"document_cd,type:varchar(10),nullzero"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`
	DocumentLineBinUid    int32     `bun:"document_line_bin_uid,type:int,pk"`
	UnitSize              float64   `bun:"unit_size,type:decimal(19,4)"`
	QtyToChange           float64   `bun:"qty_to_change,type:decimal(19,9)"`
	SubLineNo             int32     `bun:"sub_line_no,type:int"`
	RfQtyPicked           float64   `bun:"rf_qty_picked,type:decimal(19,9),default:(0)"`
	QtyFromTags           float64   `bun:"qty_from_tags,type:decimal(19,9),default:((-1))"`
	SourceDlbUid          int32     `bun:"source_dlb_uid,type:int,nullzero"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	PrintedFlag           string    `bun:"printed_flag,type:char,nullzero"`
	WorkOrderUid          int32     `bun:"work_order_uid,type:int,nullzero"`
	PickStatus            string    `bun:"pick_status,type:varchar(2),nullzero"`
	AssignedWorkstationId string    `bun:"assigned_workstation_id,type:varchar(255),nullzero"`
}
