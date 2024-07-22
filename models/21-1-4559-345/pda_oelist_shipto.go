package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PdaOelistShipto struct {
	bun.BaseModel      `bun:"table:pda_oelist_shipto"`
	PdaOelistShiptoUid int32     `bun:"pda_oelist_shipto_uid,type:int,pk"`   // Unique ID of the record.
	PdaOelistHdrUid    int32     `bun:"pda_oelist_hdr_uid,type:int"`         // Pointer to the hdr record.
	ShipToId           float64   `bun:"ship_to_id,type:decimal(19,0)"`       // Shipto ID this record is for.
	DateCreated        time.Time `bun:"date_created,type:datetime"`          // Indicates the date/time this record was created.
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`    // Indicates the date/time this record was last modified.
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30)"` // ID of the user who last maintained this record
}
