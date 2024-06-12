package model

type FedexServiceType struct {
	bun.BaseModel        `bun:"table:fedex_service_type"`
	FedexServiceTypeUid  int32     `bun:"fedex_service_type_uid,type:int,pk"`
	FedexServiceTypeId   string    `bun:"fedex_service_type_id,type:varchar(255)"`
	FedexServiceTypeDesc string    `bun:"fedex_service_type_desc,type:varchar(255)"`
	RowStatusFlag        int32     `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	FedexExpressFlag     string    `bun:"fedex_express_flag,type:char,default:('N')"`
}
