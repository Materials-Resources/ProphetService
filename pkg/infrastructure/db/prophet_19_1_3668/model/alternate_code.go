package model

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type AlternateCode struct {
	bun.BaseModel `bun:"table:alternate_code"`

	AlternateCode      string         `bun:"alternate_code,type:varchar(40),pk"`
	DeleteFlag         string         `bun:"delete_flag,type:char"`
	DateCreated        time.Time      `bun:"date_created,type:datetime"`
	DateLastModified   time.Time      `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy   string         `bun:"last_maintained_by,type:varchar(30)"`
	InvMastUid         int32          `bun:"inv_mast_uid,type:int,pk"`
	AlternateCodeDesc  sql.NullString `bun:"alternate_code_desc,type:varchar(40)"`
	AlternateCodeUid   int32          `bun:"alternate_code_uid,type:int"`
	CreatedBy          string         `bun:"created_by,type:varchar(255)"`
	SourceTypeCd       int32          `bun:"source_type_cd,type:int"`
	DefaultUom         sql.NullString `bun:"default_uom,type:varchar(8)"`
	ExcludeFromB2bFlag sql.NullString `bun:"exclude_from_b2b_flag,type:char"`
}
