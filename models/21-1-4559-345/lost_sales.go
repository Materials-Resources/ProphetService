package model

type LostSales struct {
	bun.BaseModel    `bun:"table:lost_sales"`
	LostSalesUid     int32     `bun:"lost_sales_uid,type:int,pk,identity"`
	CompanyId        string    `bun:"company_id,type:varchar(8)"`
	LostSalesId      string    `bun:"lost_sales_id,type:varchar(255)"`
	LostSalesDesc    string    `bun:"lost_sales_desc,type:varchar(255)"`
	AffectUsage      string    `bun:"affect_usage,type:char"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
