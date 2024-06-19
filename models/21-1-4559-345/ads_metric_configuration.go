package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type AdsMetricConfiguration struct {
	bun.BaseModel             `bun:"table:ads_metric_configuration"`
	AdsMetricConfigurationUid int32     `bun:"ads_metric_configuration_uid,type:int,autoincrement,scanonly,pk"` // Unique identifier
	HubMetricKey              string    `bun:"hub_metric_key,type:varchar(255)"`                                // The metric key value from the hub database
	ConfigurationDisplayText  string    `bun:"configuration_display_text,type:varchar(255)"`                    // The display value for the configuration item
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`                  // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`            // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`            // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`    // User who last changed the record
	ConfigurationKey          string    `bun:"configuration_key,type:varchar(255)"`                             // The key value that uniquely identifies the record
}
