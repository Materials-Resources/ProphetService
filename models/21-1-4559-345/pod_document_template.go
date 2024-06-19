package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PodDocumentTemplate struct {
	bun.BaseModel          `bun:"table:pod_document_template"`
	PodDocumentTemplateUid int32     `bun:"pod_document_template_uid,type:int,autoincrement,pk"`          // Unique row identifier
	DocumentCd             int32     `bun:"document_cd,type:int,unique"`                                  // Identifies document application
	DocumentTemplate       string    `bun:"document_template,type:text,nullzero"`                         // XML document template
	RowStatusFlag          int32     `bun:"row_status_flag,type:int,default:((704))"`                     // Row status
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	DocumentVersion        string    `bun:"document_version,type:varchar(255),nullzero"`                  // Tiebreaker if multiple version of the same document id are present
}
