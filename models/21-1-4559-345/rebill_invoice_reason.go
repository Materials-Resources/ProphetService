package model

type RebillInvoiceReason struct {
	bun.BaseModel          `bun:"table:rebill_invoice_reason"`
	RebillInvoiceReasonUid int32     `bun:"rebill_invoice_reason_uid,type:int,pk,identity"`
	CompanyId              string    `bun:"company_id,type:varchar(8)"`
	ReasonCode             string    `bun:"reason_code,type:varchar(255)"`
	ReasonDesc             string    `bun:"reason_desc,type:varchar(255)"`
	RowStatus              int16     `bun:"row_status,type:smallint,default:((704))"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
