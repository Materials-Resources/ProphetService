package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CycleCountDetail struct {
	bun.BaseModel            `bun:"table:cycle_count_detail"`
	CycleCountDetailUid      int32     `bun:"cycle_count_detail_uid,type:int,pk"`                        // Unique identifier for a cycle_count_detail record.
	CycleCountHdrUid         int32     `bun:"cycle_count_hdr_uid,type:int"`                              // Unique id for the cycle count hdr
	LineNo                   int32     `bun:"line_no,type:int"`                                          // What line is this row?
	Bin                      string    `bun:"bin,type:varchar(10),nullzero"`                             // The bin counted
	InvMastUid               int32     `bun:"inv_mast_uid,type:int"`                                     // Unique identifier for the item id.
	QtyCounted               float64   `bun:"qty_counted,type:decimal(19,9),nullzero"`                   // The quantity counted.
	AddedToPhysicalCount     string    `bun:"added_to_physical_count,type:char(1)"`                      // Used when querying the system for existing and previously used inventory cards.
	DateCreated              time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30)"`                       // ID of the user who last maintained this record
	RowStatusFlag            int32     `bun:"row_status_flag,type:int"`                                  // Indicates current record status.
	QtyOnHandAtPhysicalCount float64   `bun:"qty_on_hand_at_physical_count,type:decimal(19,9),nullzero"` // Qty on hand for items not counted at physical count
}
