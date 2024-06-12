package model

type LostSalesTransaction struct {
	bun.BaseModel           `bun:"table:lost_sales_transaction"`
	LostSalesTransactionUid int32     `bun:"lost_sales_transaction_uid,type:int,pk,identity"`
	LostSalesUid            int32     `bun:"lost_sales_uid,type:int"`
	AffectUsage             string    `bun:"affect_usage,type:char"`
	TransactionCodeNo       int32     `bun:"transaction_code_no,type:int"`
	TransactionNo           int32     `bun:"transaction_no,type:int"`
	LineNo                  int32     `bun:"line_no,type:int,nullzero"`
	SubLineNo               int32     `bun:"sub_line_no,type:int,nullzero"`
	SkuQtyChange            float64   `bun:"sku_qty_change,type:decimal(19,9),nullzero"`
	UsageProcessedFlag      string    `bun:"usage_processed_flag,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
