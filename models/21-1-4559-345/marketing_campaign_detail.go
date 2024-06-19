package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type MarketingCampaignDetail struct {
	bun.BaseModel              `bun:"table:marketing_campaign_detail"`
	MarketingCampaignDetailUid int32     `bun:"marketing_campaign_detail_uid,type:int,autoincrement,identity,pk"` // Unique identifier for the table
	ListId                     float64   `bun:"list_id,type:decimal(19,0)"`                                       // The key from the mail_list table
	BoughtFlag                 string    `bun:"bought_flag,type:char(1)"`                                         // Indicates whether the item must have been bought or not bought
	EntityUid                  int32     `bun:"entity_uid,type:int"`                                              // Either the product_group_uid or inv_mast_uid
	EntityType                 int32     `bun:"entity_type,type:int"`                                             // Code indicating either Item or Product Group
	DaysFromPurchase           int32     `bun:"days_from_purchase,type:int"`                                      // Number of days since Item or PG was last purchased
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`                   // Date and time the record was originally created
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`             // User who created the record
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`             // Date and time the record was modified
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`     // User who last changed the record
}
