package model

type PricingTemplateLocation struct {
	bun.BaseModel              `bun:"table:pricing_template_location"`
	PricingTemplateLocationUid int32     `bun:"pricing_template_location_uid,type:int,pk"`
	PricingTemplateUid         int32     `bun:"pricing_template_uid,type:int"`
	LocationId                 float64   `bun:"location_id,type:decimal(19,0)"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
