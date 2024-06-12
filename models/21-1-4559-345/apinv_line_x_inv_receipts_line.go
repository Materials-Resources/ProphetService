package model

type ApinvLineXInvReceiptsLine struct {
	bun.BaseModel           `bun:"table:apinv_line_x_inv_receipts_line"`
	ApinvLineXInvRecLineUid int32     `bun:"apinv_line_x_inv_rec_line_uid,type:int,pk,identity"`
	ReceiptNumber           float64   `bun:"receipt_number,type:decimal(19,0)"`
	ReceiptLine             int32     `bun:"receipt_line,type:int"`
	ApinvLineUid            int32     `bun:"apinv_line_uid,type:int,nullzero"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
