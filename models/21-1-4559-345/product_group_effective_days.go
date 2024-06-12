package model

type ProductGroupEffectiveDays struct {
	bun.BaseModel             `bun:"table:product_group_effective_days"`
	ProdGroupEffectiveDaysUid int32     `bun:"prod_group_effective_days_uid,type:int,pk"`
	ProductGroupId            string    `bun:"product_group_id,type:varchar(8)"`
	CompanyId                 string    `bun:"company_id,type:varchar(8)"`
	EffectiveDays             int32     `bun:"effective_days,type:int,default:((0))"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	JobEntryFlag              string    `bun:"job_entry_flag,type:char,default:('Y')"`
	OrderEntryFlag            string    `bun:"order_entry_flag,type:char,default:('Y')"`
	JobMinGrossProfit         float64   `bun:"job_min_gross_profit,type:decimal(19,9),default:((0))"`
	OrderMinGrossProfit       float64   `bun:"order_min_gross_profit,type:decimal(19,9),default:((0))"`
}
