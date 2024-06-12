package model

type JobPriceCustShiptoOrdlim struct {
	bun.BaseModel               `bun:"table:job_price_cust_shipto_ordlim"`
	JobPriceCustshiptoordlimUid int32     `bun:"job_price_custshiptoordlim_uid,type:int,pk,identity"`
	JobPriceCustShiptoUid       int32     `bun:"job_price_cust_shipto_uid,type:int"`
	JobPriceHdrBudgetPrdUid     int32     `bun:"job_price_hdr_budget_prd_uid,type:int"`
	OrderLimit                  int32     `bun:"order_limit,type:int,default:((0))"`
	OrderLimitUsed              int32     `bun:"order_limit_used,type:int"`
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
