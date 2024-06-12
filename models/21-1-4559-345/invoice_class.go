package model

type InvoiceClass struct {
	bun.BaseModel    `bun:"table:invoice_class"`
	InvoiceClass     string    `bun:"invoice_class,type:varchar(8),pk"`
	InvoiceClassDesc string    `bun:"invoice_class_desc,type:varchar(30)"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
