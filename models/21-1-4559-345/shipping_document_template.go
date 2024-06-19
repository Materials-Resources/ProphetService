package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ShippingDocumentTemplate struct {
	bun.BaseModel               `bun:"table:shipping_document_template"`
	ShippingDocumentTemplateUid int32     `bun:"shipping_document_template_uid,type:int,autoincrement,pk"`     // Unique identifier for table.
	DocumentType                string    `bun:"document_type,type:varchar(255)"`                              // Document Identifier
	DocumentTemplate            string    `bun:"document_template,type:text"`                                  // XML template string of the document
	CarrierTypeCd               int32     `bun:"carrier_type_cd,type:int"`                                     // Carrier tpye, UPS, FEDEX, etc.
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
