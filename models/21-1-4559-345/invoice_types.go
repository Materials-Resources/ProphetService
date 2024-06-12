package model

type InvoiceTypes struct {
	bun.BaseModel          `bun:"table:invoice_types"`
	InvoiceTypeId          string    `bun:"invoice_type_id,type:varchar(2),pk"`
	InvoiceTypeDescription string    `bun:"invoice_type_description,type:varchar(30)"`
	DeleteFlag             string    `bun:"delete_flag,type:char"`
	DateCreated            time.Time `bun:"date_created,type:datetime"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	Object                 string    `bun:"object,type:varchar(50),nullzero"`
}
