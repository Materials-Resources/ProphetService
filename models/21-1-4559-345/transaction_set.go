package model

import (
	"github.com/uptrace/bun"
	"time"
)

type TransactionSet struct {
	bun.BaseModel      `bun:"table:transaction_set"`
	TransactionSetUid  int32     `bun:"transaction_set_uid,type:int,pk"`
	TransactionSetId   string    `bun:"transaction_set_id,type:varchar(32)"`
	TransactionSetDesc string    `bun:"transaction_set_desc,type:varchar(255)"`
	DateCreated        time.Time `bun:"date_created,type:datetime"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	Enablefor          string    `bun:"enablefor,type:char(6),nullzero"`
	AllowXml           string    `bun:"allow_xml,type:char,default:('N')"`
	P21CodeNo          int32     `bun:"p21_code_no,type:int,nullzero"`
	EdiTransactionId   string    `bun:"edi_transaction_id,type:varchar(255),nullzero"`
}
