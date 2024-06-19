package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CreditcardTranstype struct {
	bun.BaseModel          `bun:"table:creditcard_transtype"`
	CreditcardTranstypeUid int32     `bun:"creditcard_transtype_uid,type:int,pk"`     // UID of this table
	TranstypeProtobaseId   int32     `bun:"transtype_protobase_id,type:int"`          // ID number of the transaction from Protobase
	TranstypeName          string    `bun:"transtype_name,type:varchar(50),nullzero"` // Credit card transaction name
	TranstypeSettlable     string    `bun:"transtype_settlable,type:char(1)"`         // A flag to indicate whether the transaction is settlable at Protobase
	DateCreated            time.Time `bun:"date_created,type:datetime"`               // Indicates the date/time this record was created.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30)"`      // ID of the user who last maintained this record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`         // Indicates the date/time this record was last modified.
}
