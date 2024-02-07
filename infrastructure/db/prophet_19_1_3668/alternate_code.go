package prophet_19_1_3668

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
	LastMaintainedBy   string         `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	InvMastUid         int32          `bun:"inv_mast_uid,type:int,pk"`
	AlternateCodeDesc  sql.NullString `bun:"alternate_code_desc,type:varchar(40),nullzero"`
	AlternateCodeUid   int32          `bun:"alternate_code_uid,type:int"`
	CreatedBy          string         `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	SourceTypeCd       int32          `bun:"source_type_cd,type:int,default:((2051))"`
	DefaultUom         sql.NullString `bun:"default_uom,type:varchar(8),nullzero"`
	ExcludeFromB2bFlag sql.NullString `bun:"exclude_from_b2b_flag,type:char,nullzero"`
}
