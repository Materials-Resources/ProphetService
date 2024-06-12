package model

type Remittances struct {
	bun.BaseModel    `bun:"table:remittances"`
	RemittanceNumber float64   `bun:"remittance_number,type:decimal(19,0),pk"`
	OrderNo          string    `bun:"order_no,type:varchar(8)"`
	CashDrawerId     string    `bun:"cash_drawer_id,type:varchar(8)"`
	CompanyId        string    `bun:"company_id,type:varchar(8)"`
	PaymentNumber    float64   `bun:"payment_number,type:decimal(19,0)"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	InCashDrawer     string    `bun:"in_cash_drawer,type:char"`
	CashDrawerSeqNo  int32     `bun:"cash_drawer_seq_no,type:int,default:(0)"`
}
