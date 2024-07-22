package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ScanPack struct {
	bun.BaseModel            `bun:"table:scan_pack"`
	ScanPackUid              int32     `bun:"scan_pack_uid,type:int,pk"`                                    // Unique identifier for the table.
	RowStatusFlag            int32     `bun:"row_status_flag,type:int"`                                     // Indicates row status (active/delete/complete)
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	SourceCd                 int32     `bun:"source_cd,type:int,default:((3612))"`                          // Indicates the source of the Scan and Pack Shipment. Valid codes: 1122 (Shipping) 3612 (Scan and Pack Entry)
	PrntUcc128LblDisplayFlag string    `bun:"prnt_ucc128_lbl_display_flag,type:char(1),default:('N')"`      // Print flag for GS1-128 labels
	LockFlag                 string    `bun:"lock_flag,type:char(1),default:('N')"`                         // Custom: Indicates that PTs cannot be added to shipment.
}
