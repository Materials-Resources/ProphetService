package model

type AdsMetricConfiguration struct {
	bun.BaseModel             `bun:"table:ads_metric_configuration"`
	AdsMetricConfigurationUid int32     `bun:"ads_metric_configuration_uid,type:int,pk,identity"`
	HubMetricKey              string    `bun:"hub_metric_key,type:varchar(255)"`
	ConfigurationDisplayText  string    `bun:"configuration_display_text,type:varchar(255)"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	ConfigurationKey          string    `bun:"configuration_key,type:varchar(255)"`
}
