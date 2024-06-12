package model

type PricingServiceLayoutLoc struct {
	bun.BaseModel              `bun:"table:pricing_service_layout_loc"`
	PricingServiceLayoutLocUid int32     `bun:"pricing_service_layout_loc_uid,type:int,pk,identity"`
	LayoutId                   float64   `bun:"layout_id,type:decimal(19,0)"`
	LocationId                 float64   `bun:"location_id,type:decimal(19,0)"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
