package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type XmlDocumentElement struct {
	bun.BaseModel          `bun:"table:xml_document_element"`
	XmlDocumentElementUid  int32     `bun:"xml_document_element_uid,type:int,pk"`      // Unique identifier for records within the table
	XmlDocumentUid         int32     `bun:"xml_document_uid,type:int"`                 // Foreign key to associated parent record in the xml_document table
	ElementName            string    `bun:"element_name,type:varchar(255)"`            // Name of xml document element described by each record
	ElementTypeCoded       *string   `bun:"element_type_coded,type:varchar(255)"`      // Value of data within an element, of a named pair of elements.  The code identifies the meaning of the associated code value
	ValueElementName       *string   `bun:"value_element_name,type:varchar(255)"`      // The name of the element which contains the value for a named value pair.
	ConvertedElementValue  *string   `bun:"converted_element_value,type:varchar(255)"` // Conversion value for coded element.  1 is the converted_element_value when a CurrencyCoded element has the value USD
	IdentifyingNodeName    string    `bun:"identifying_node_name,type:varchar(255)"`   // The name of the node element that, along with the other node identifiers, uniquely identifies an element within the document tree structure
	ListofNodeName         *string   `bun:"listof_node_name,type:varchar(255)"`        // Name of the node element that indicates the subordinate element group is a repeating group
	ElementTypeCd          int32     `bun:"element_type_cd,type:int"`                  // Code to identify whether record describes an element or attribute of an xml document
	XmlDataobjectColumnUid int32     `bun:"xml_dataobject_column_uid,type:int"`        // Foreign key to xml_dataobject_column table, and identifies the column that is the source or target for the element value
	ElementSeq             int32     `bun:"element_seq,type:int"`                      // Sequence number of the element within the document
	DocumentSectionCd      int32     `bun:"document_section_cd,type:int"`              // Code to identify which of the three major parts of the document in which the element exists - header, detail, summary
	RequiredCd             int32     `bun:"required_cd,type:int"`                      // Code to identify whether the element is required
	ValueId                *float64  `bun:"value_id,type:decimal(19,0)"`               // Foreign key to pricing_service_value table, to allow translation of code to CC database value
	DateCreated            time.Time `bun:"date_created,type:datetime"`                // Indicates the date/time this record was created.
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`          // Indicates the date/time this record was last modified.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30)"`       // ID of the user who last maintained this record
	ListofParentNodeName   *string   `bun:"listof_parent_node_name,type:varchar(255)"` // Used with nested lists, the name of the node element that indicates the subordinate element group is a repeating group
}
