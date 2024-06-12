package model

type PricingTemplateLocationDflt struct {
	bun.BaseModel                  `bun:"table:pricing_template_location_dflt"`
	PricingTemplateLocationDfltUid int32     `bun:"pricing_template_location_dflt_uid,type:int,pk,identity"`
	PricingTemplateLocationUid     int32     `bun:"pricing_template_location_uid,type:int"`
	PricingTemplateKeyFieldUid     int32     `bun:"pricing_template_key_field_uid,type:int"`
	ColumnId                       float64   `bun:"column_id,type:decimal(19,0)"`
	DefaultValue                   string    `bun:"default_value,type:varchar(50)"`
	DateCreated                    time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                      string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified               time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy               string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
