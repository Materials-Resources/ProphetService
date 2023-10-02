package model

import (
	"database/sql"

	"github.com/uptrace/bun"
)

type InvLoc struct {
	bun.BaseModel `bun:"table:inv_loc"`

	LocationId float64 `bun:"location_id,type:decimal(19),pk"`
	InvMastUid int     `bun:"inv_mast_uid,type:int,pk"`
	CompanyId  string  `bun:"company_id,type:varchar(8),pk"`

	PrimaryBin sql.NullString `bun:"primary_bin,type:varchar(10)"`
}
