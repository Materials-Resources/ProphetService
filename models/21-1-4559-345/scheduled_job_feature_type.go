package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ScheduledJobFeatureType struct {
	bun.BaseModel               `bun:"table:scheduled_job_feature_type"`
	ScheduledJobFeatureTypeUid  int32     `bun:"scheduled_job_feature_type_uid,type:int,autoincrement,identity,pk"` // Unique identifier for table - primary key.
	ScheduledJobFeatureTypeId   string    `bun:"scheduled_job_feature_type_id,type:varchar(255),unique"`            // Name for the scheduled job feature type.
	ScheduledJobFeatureTypeDesc string    `bun:"scheduled_job_feature_type_desc,type:varchar(255)"`                 // Description of the scheduled job feature type.
	DeleteFlag                  string    `bun:"delete_flag,type:char(1),default:('N')"`                            // Flag to indicate record is marked deleted.
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`                    // Date and time the record was originally created
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`              // User who created the record
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`              // Date and time the record was modified
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`      // User who last changed the record
}
