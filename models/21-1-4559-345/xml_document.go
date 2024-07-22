package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type XmlDocument struct {
	bun.BaseModel         `bun:"table:xml_document"`
	XmlDocumentUid        int32     `bun:"xml_document_uid,type:int,pk"`              // Unique indentifier for records within the table
	DocumentVersion       string    `bun:"document_version,type:varchar(10)"`         // Version of xml document
	DocumentSchema        string    `bun:"document_schema,type:varchar(255)"`         // Type of schema for the xml document
	SchemaSourceCd        int32     `bun:"schema_source_cd,type:int"`                 // Code for the source of the xml document
	TemplateFilename      string    `bun:"template_filename,type:varchar(255)"`       // File name of example/template document used to create the database xml document template
	TransactionSetUid     int32     `bun:"transaction_set_uid,type:int"`              // Foreign key to the transaction_set table, indicating the transaction set associated with the xml document template
	DocumentDesc          string    `bun:"document_desc,type:varchar(255)"`           // Description of the xml document template
	DefaultDocument       string    `bun:"default_document,type:char(1)"`             // Y/N indicator of whether the xml document template is the default for the associated transaction set
	DocumentTemplate      string    `bun:"document_template,type:text"`               // The xml document to be used as a template for export, used as a map for import
	RowStatusFlag         int32     `bun:"row_status_flag,type:int"`                  // Indicates current record status.
	DateCreated           time.Time `bun:"date_created,type:datetime"`                // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`          // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`       // ID of the user who last maintained this record
	RootElement           *string   `bun:"root_element,type:varchar(255)"`            // Root element of xml document
	DocumentSectionPrefix *string   `bun:"document_section_prefix,type:varchar(255)"` // Prefix for each TPCx xml document section
	MajorVersion          int16     `bun:"major_version,type:smallint,default:(0)"`   // Major version of schema
	MinorVersion          int16     `bun:"minor_version,type:smallint,default:(0)"`   // Minor version of schema
	BuildNo               int16     `bun:"build_no,type:smallint,default:(0)"`        // Build number of schema
}
