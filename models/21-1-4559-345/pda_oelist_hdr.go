package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PdaOelistHdr struct {
	bun.BaseModel         `bun:"table:pda_oelist_hdr"`
	PdaOelistHdrUid       int32     `bun:"pda_oelist_hdr_uid,type:int,pk"`
	OelistListId          string    `bun:"oelist_list_id,type:varchar(255)"`
	OelistListDescription string    `bun:"oelist_list_description,type:varchar(255),nullzero"`
	ListDate              time.Time `bun:"list_date,type:datetime"`
	CompanyId             string    `bun:"company_id,type:varchar(8)"`
	LocationId            float64   `bun:"location_id,type:decimal(19,0)"`
	SalesrepId            string    `bun:"salesrep_id,type:varchar(16)"`
	RowStatusFlag         int32     `bun:"row_status_flag,type:int"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`
}
