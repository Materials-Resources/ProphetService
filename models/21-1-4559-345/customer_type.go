package model

type CustomerType struct {
	bun.BaseModel    `bun:"table:customer_type"`
	CustomerTypeUid  int32     `bun:"customer_type_uid,type:int,pk"`
	CustomerType     string    `bun:"customer_type,type:varchar(40)"`
	RowStatusFlag    int16     `bun:"row_status_flag,type:smallint"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(suser_sname())"`
}
