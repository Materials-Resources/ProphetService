package model

type MarketingCampaignDetail struct {
	bun.BaseModel              `bun:"table:marketing_campaign_detail"`
	MarketingCampaignDetailUid int32     `bun:"marketing_campaign_detail_uid,type:int,pk,identity"`
	ListId                     float64   `bun:"list_id,type:decimal(19,0)"`
	BoughtFlag                 string    `bun:"bought_flag,type:char"`
	EntityUid                  int32     `bun:"entity_uid,type:int"`
	EntityType                 int32     `bun:"entity_type,type:int"`
	DaysFromPurchase           int32     `bun:"days_from_purchase,type:int"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
