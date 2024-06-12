package model

type CustomerCallInvDetail struct {
	bun.BaseModel            `bun:"table:customer_call_inv_detail"`
	CustomerCallInvDetailUid int32     `bun:"customer_call_inv_detail_uid,type:int,pk"`
	CustomerCallUid          int32     `bun:"customer_call_uid,type:int"`
	InvoiceNo                string    `bun:"invoice_no,type:varchar(10)"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30)"`
}
