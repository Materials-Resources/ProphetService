package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CreditcardTranstype struct {
	bun.BaseModel          `bun:"table:creditcard_transtype"`
	CreditcardTranstypeUid int32     `bun:"creditcard_transtype_uid,type:int,pk"`
	TranstypeProtobaseId   int32     `bun:"transtype_protobase_id,type:int"`
	TranstypeName          string    `bun:"transtype_name,type:varchar(50),nullzero"`
	TranstypeSettlable     string    `bun:"transtype_settlable,type:char"`
	DateCreated            time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30)"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`
}
