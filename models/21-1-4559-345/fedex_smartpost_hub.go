package model

type FedexSmartpostHub struct {
	bun.BaseModel         `bun:"table:fedex_smartpost_hub"`
	FedexSmartpostHubUid  int32     `bun:"fedex_smartpost_hub_uid,type:int,pk,identity"`
	FedexSmartpostHubId   int32     `bun:"fedex_smartpost_hub_id,type:int"`
	FedexSmartpostHubDesc string    `bun:"fedex_smartpost_hub_desc,type:varchar(255)"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
