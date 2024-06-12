package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PdaOelistShipto struct {
	bun.BaseModel      `bun:"table:pda_oelist_shipto"`
	PdaOelistShiptoUid int32     `bun:"pda_oelist_shipto_uid,type:int,pk"`
	PdaOelistHdrUid    int32     `bun:"pda_oelist_hdr_uid,type:int"`
	ShipToId           float64   `bun:"ship_to_id,type:decimal(19,0)"`
	DateCreated        time.Time `bun:"date_created,type:datetime"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30)"`
}
