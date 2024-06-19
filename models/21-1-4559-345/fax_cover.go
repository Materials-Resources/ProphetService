package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type FaxCover struct {
	bun.BaseModel    `bun:"table:fax_cover"`
	FaxCoverUid      int32     `bun:"fax_cover_uid,type:int,autoincrement,identity,pk"`             // UID
	FaxCoverId       string    `bun:"fax_cover_id,type:varchar(255)"`                               // ID of the cover sheet in VSI fax
	FaxCoverDesc     string    `bun:"fax_cover_desc,type:varchar(255),nullzero"`                    // Description of the coversheet
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`                                     // Active \ Delete
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
