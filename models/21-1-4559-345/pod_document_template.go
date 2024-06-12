package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PodDocumentTemplate struct {
	bun.BaseModel          `bun:"table:pod_document_template"`
	PodDocumentTemplateUid int32     `bun:"pod_document_template_uid,type:int,pk,identity"`
	DocumentCd             int32     `bun:"document_cd,type:int"`
	DocumentTemplate       string    `bun:"document_template,type:text(2147483647),nullzero"`
	RowStatusFlag          int32     `bun:"row_status_flag,type:int,default:((704))"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DocumentVersion        string    `bun:"document_version,type:varchar(255),nullzero"`
}
