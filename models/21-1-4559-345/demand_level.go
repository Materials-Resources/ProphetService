package model

type DemandLevel struct {
	bun.BaseModel       `bun:"table:demand_level"`
	DemandLevelUid      int32     `bun:"demand_level_uid,type:int,pk,identity"`
	ServiceLevelPctGoal float64   `bun:"service_level_pct_goal,type:decimal(19,2),nullzero"`
	StockOutDeviation   float64   `bun:"stock_out_deviation,type:decimal(19,2),nullzero"`
	BackorderDeviation  float64   `bun:"backorder_deviation,type:decimal(19,2),nullzero"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
