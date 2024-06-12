package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerCall struct {
	bun.BaseModel       `bun:"table:customer_call"`
	CustomerCallUid     int32     `bun:"customer_call_uid,type:int,pk"`
	CompanyId           string    `bun:"company_id,type:varchar(8)"`
	CustomerId          float64   `bun:"customer_id,type:decimal(19,0)"`
	TopicUid            int32     `bun:"topic_uid,type:int"`
	ContactId           string    `bun:"contact_id,type:varchar(16)"`
	NumberOfTimesCalled int16     `bun:"number_of_times_called,type:smallint"`
	PromisedDate        time.Time `bun:"promised_date,type:datetime,nullzero"`
	PromisedAmount      float64   `bun:"promised_amount,type:decimal(19,4),nullzero"`
	CallBackDate        time.Time `bun:"call_back_date,type:datetime"`
	Notes               string    `bun:"notes,type:text(2147483647)"`
	RowStatusFlag       int16     `bun:"row_status_flag,type:smallint"`
	DateCreated         time.Time `bun:"date_created,type:datetime"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30)"`
}
