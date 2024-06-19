package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type XmlStylesheet struct {
	bun.BaseModel    `bun:"table:xml_stylesheet"`
	XmlStylesheetUid int32     `bun:"xml_stylesheet_uid,type:int,autoincrement,scanonly,pk"`        // Table UID
	DocumentVersion  string    `bun:"document_version,type:varchar(255)"`                           // Version of the stylesheet
	XmlStylesheetCd  int32     `bun:"xml_stylesheet_cd,type:int"`                                   // Code value of the stylesheet
	XmlStylesheet    string    `bun:"xml_stylesheet,type:text"`                                     // The stylesheet
	DocumentDesc     string    `bun:"document_desc,type:varchar(255)"`                              // Description of the stylesheet
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:((704))"`                     // Status of record
}
