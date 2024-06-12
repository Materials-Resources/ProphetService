package model

import (
	"github.com/uptrace/bun"
	"time"
)

type QuoteHdr struct {
	bun.BaseModel    `bun:"table:quote_hdr"`
	QuoteHdrUid      int32     `bun:"quote_hdr_uid,type:int,pk"`
	OeHdrUid         int32     `bun:"oe_hdr_uid,type:int"`
	CompleteFlag     string    `bun:"complete_flag,type:char,default:('N')"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	Date843Sent      time.Time `bun:"date_843_sent,type:datetime,nullzero"`
	ExpirationDate   time.Time `bun:"expiration_date,type:datetime,nullzero"`
}
