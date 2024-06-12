package model

import (
	"github.com/uptrace/bun"
	"time"
)

type TransSetXXmlDataobject struct {
	bun.BaseModel     `bun:"table:trans_set_x_xml_dataobject"`
	TransactionSetUid int32     `bun:"transaction_set_uid,type:int,pk"`
	XmlDataobjectUid  int32     `bun:"xml_dataobject_uid,type:int,pk"`
	ProcessSeq        int16     `bun:"process_seq,type:smallint"`
	DateCreated       time.Time `bun:"date_created,type:datetime"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30)"`
}
