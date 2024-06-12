package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ShippingDocumentTemplate struct {
	bun.BaseModel               `bun:"table:shipping_document_template"`
	ShippingDocumentTemplateUid int32     `bun:"shipping_document_template_uid,type:int,pk,identity"`
	DocumentType                string    `bun:"document_type,type:varchar(255)"`
	DocumentTemplate            string    `bun:"document_template,type:text(2147483647)"`
	CarrierTypeCd               int32     `bun:"carrier_type_cd,type:int"`
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
