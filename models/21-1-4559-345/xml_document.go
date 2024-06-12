package model

type XmlDocument struct {
	bun.BaseModel         `bun:"table:xml_document"`
	XmlDocumentUid        int32     `bun:"xml_document_uid,type:int,pk"`
	DocumentVersion       string    `bun:"document_version,type:varchar(10)"`
	DocumentSchema        string    `bun:"document_schema,type:varchar(255)"`
	SchemaSourceCd        int32     `bun:"schema_source_cd,type:int"`
	TemplateFilename      string    `bun:"template_filename,type:varchar(255)"`
	TransactionSetUid     int32     `bun:"transaction_set_uid,type:int"`
	DocumentDesc          string    `bun:"document_desc,type:varchar(255)"`
	DefaultDocument       string    `bun:"default_document,type:char"`
	DocumentTemplate      string    `bun:"document_template,type:text(2147483647)"`
	RowStatusFlag         int32     `bun:"row_status_flag,type:int"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`
	RootElement           string    `bun:"root_element,type:varchar(255),nullzero"`
	DocumentSectionPrefix string    `bun:"document_section_prefix,type:varchar(255),nullzero"`
	MajorVersion          int16     `bun:"major_version,type:smallint,default:(0)"`
	MinorVersion          int16     `bun:"minor_version,type:smallint,default:(0)"`
	BuildNo               int16     `bun:"build_no,type:smallint,default:(0)"`
}
