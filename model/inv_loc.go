package model

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type InvLoc struct {
	bun.BaseModel `bun:"table:inv_loc"`

	LocationId        float64         `bun:"location_id,type:decimal(19),pk"`
	QtyOnHand         sql.NullFloat64 `bun:"qty_on_hand,type:decimal(19, 9),default:0"`
	QtyInProcess      float64         `bun:"qty_in_process,type:decimal(19, 9),default:0"`
	DateCreated       time.Time       `bun:"date_created,type:datetime"`
	DateLastModified  time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy  string          `bun:"last_maintained_by,type:varchar(30)"`
	CompanyId         string          `bun:"company_id,type:varchar(8),pk"`
	LastRecPo         sql.NullFloat64 `bun:"last_rec_po,type:decimal(19, 9)"`
	LastRecPoWithDisc sql.NullFloat64 `bun:"last_rec_po_with_disc,type:decimal(19, 9)"`
	GlAccountNo       sql.NullString  `bun:"gl_account_no,type:varchar(32)"`

	InvMastUid int `bun:"inv_mast_uid,type:int,pk"`

	PrimaryBin sql.NullString `bun:"primary_bin,type:varchar(10)"`
}
