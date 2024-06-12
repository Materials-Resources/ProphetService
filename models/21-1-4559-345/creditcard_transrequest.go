package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CreditcardTransrequest struct {
	bun.BaseModel             `bun:"table:creditcard_transrequest"`
	CreditcardTransrequestUid int32     `bun:"creditcard_transrequest_uid,type:int,pk"`
	PaymentNumber             float64   `bun:"payment_number,type:decimal(19,0),nullzero"`
	ReqfileLocation           string    `bun:"reqfile_location,type:varchar(255),nullzero"`
	RequesttypeId             int32     `bun:"requesttype_id,type:int"`
	RequestStatus             string    `bun:"request_status,type:varchar(10),nullzero"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
}
