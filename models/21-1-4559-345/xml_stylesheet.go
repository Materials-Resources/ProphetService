package model

import (
	"github.com/uptrace/bun"
	"time"
)

type XmlStylesheet struct {
	bun.BaseModel    `bun:"table:xml_stylesheet"`
	XmlStylesheetUid int32     `bun:"xml_stylesheet_uid,type:int,pk,identity"`
	DocumentVersion  string    `bun:"document_version,type:varchar(255)"`
	XmlStylesheetCd  int32     `bun:"xml_stylesheet_cd,type:int"`
	XmlStylesheet    string    `bun:"xml_stylesheet,type:text(2147483647)"`
	DocumentDesc     string    `bun:"document_desc,type:varchar(255)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:((704))"`
}
