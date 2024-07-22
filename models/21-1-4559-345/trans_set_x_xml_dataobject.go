package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type TransSetXXmlDataobject struct {
	bun.BaseModel     `bun:"table:trans_set_x_xml_dataobject"`
	TransactionSetUid int32     `bun:"transaction_set_uid,type:int,pk"`     // Foreign key to transaction_set table, this column associates an xml_dataobject record with a transaction set record
	XmlDataobjectUid  int32     `bun:"xml_dataobject_uid,type:int,pk"`      // Foreign key to xml_dataobject table, associates an xml_dataobject record with  transaction set record
	ProcessSeq        int16     `bun:"process_seq,type:smallint"`           // Sequence number indicating the sequence in which a table is processed during xml document conversion
	DateCreated       time.Time `bun:"date_created,type:datetime"`          // Indicates the date/time this record was created.
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`    // Indicates the date/time this record was last modified.
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30)"` // ID of the user who last maintained this record
}
