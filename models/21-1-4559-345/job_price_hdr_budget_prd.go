package model

type JobPriceHdrBudgetPrd struct {
	bun.BaseModel           `bun:"table:job_price_hdr_budget_prd"`
	JobPriceHdrBudgetPrdUid int32     `bun:"job_price_hdr_budget_prd_uid,type:int,pk"`
	JobPriceHdrUid          int32     `bun:"job_price_hdr_uid,type:int"`
	Period                  float64   `bun:"period,type:decimal(3,0)"`
	YearForPeriod           float64   `bun:"year_for_period,type:decimal(4,0)"`
	BeginningDate           time.Time `bun:"beginning_date,type:datetime"`
	EndingDate              time.Time `bun:"ending_date,type:datetime"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
