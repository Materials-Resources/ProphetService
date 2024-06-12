package model

import (
	"github.com/uptrace/bun"
	"time"
)

type XmlDocumentElement struct {
	bun.BaseModel          `bun:"table:xml_document_element"`
	XmlDocumentElementUid  int32     `bun:"xml_document_element_uid,type:int,pk"`
	XmlDocumentUid         int32     `bun:"xml_document_uid,type:int"`
	ElementName            string    `bun:"element_name,type:varchar(255)"`
	ElementTypeCoded       string    `bun:"element_type_coded,type:varchar(255),nullzero"`
	ValueElementName       string    `bun:"value_element_name,type:varchar(255),nullzero"`
	ConvertedElementValue  string    `bun:"converted_element_value,type:varchar(255),nullzero"`
	IdentifyingNodeName    string    `bun:"identifying_node_name,type:varchar(255)"`
	ListofNodeName         string    `bun:"listof_node_name,type:varchar(255),nullzero"`
	ElementTypeCd          int32     `bun:"element_type_cd,type:int"`
	XmlDataobjectColumnUid int32     `bun:"xml_dataobject_column_uid,type:int"`
	ElementSeq             int32     `bun:"element_seq,type:int"`
	DocumentSectionCd      int32     `bun:"document_section_cd,type:int"`
	RequiredCd             int32     `bun:"required_cd,type:int"`
	ValueId                float64   `bun:"value_id,type:decimal(19,0),nullzero"`
	DateCreated            time.Time `bun:"date_created,type:datetime"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30)"`
	ListofParentNodeName   string    `bun:"listof_parent_node_name,type:varchar(255),nullzero"`
}
