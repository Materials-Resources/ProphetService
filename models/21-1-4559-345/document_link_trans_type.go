package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DocumentLinkTransType struct {
	bun.BaseModel            `bun:"table:document_link_trans_type"`
	DocumentLinkTransTypeUid int32     `bun:"document_link_trans_type_uid,type:int,pk"`                     // unique identifier for each record
	SourceAreaCd             int32     `bun:"source_area_cd,type:int"`                                      // code for where document link was created
	TransType                string    `bun:"trans_type,type:varchar(255)"`                                 // code for inv_trans trans_type
	RowStatusFlag            int32     `bun:"row_status_flag,type:int,default:(704)"`                       // code for status of each record
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
