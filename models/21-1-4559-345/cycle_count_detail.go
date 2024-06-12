package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CycleCountDetail struct {
	bun.BaseModel            `bun:"table:cycle_count_detail"`
	CycleCountDetailUid      int32     `bun:"cycle_count_detail_uid,type:int,pk"`
	CycleCountHdrUid         int32     `bun:"cycle_count_hdr_uid,type:int"`
	LineNo                   int32     `bun:"line_no,type:int"`
	Bin                      string    `bun:"bin,type:varchar(10),nullzero"`
	InvMastUid               int32     `bun:"inv_mast_uid,type:int"`
	QtyCounted               float64   `bun:"qty_counted,type:decimal(19,9),nullzero"`
	AddedToPhysicalCount     string    `bun:"added_to_physical_count,type:char"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30)"`
	RowStatusFlag            int32     `bun:"row_status_flag,type:int"`
	QtyOnHandAtPhysicalCount float64   `bun:"qty_on_hand_at_physical_count,type:decimal(19,9),nullzero"`
}
