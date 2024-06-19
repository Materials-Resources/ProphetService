package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type AlertType struct {
	bun.BaseModel    `bun:"table:alert_type"`
	AlertTypeUid     int32     `bun:"alert_type_uid,type:int,pk"`                                // Unique Identifier of record
	TypeCd           int32     `bun:"type_cd,type:int"`                                          // Code to identify type of alert
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`                                  // Indicates current record status.
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`      // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	ModuleCd         int32     `bun:"module_cd,type:int"`                                        // Identifies to what module an alert belongs.
	ConfigurationId  int32     `bun:"configuration_id,type:int"`                                 // A company id that the alert_type is available.
	ViewName         string    `bun:"view_name,type:varchar(255)"`                               // View that is executed by the alert.
	JobId            string    `bun:"job_id,type:varchar(255),nullzero"`                         // This column stores the job_id of the SQLServer job associate with a specific alert type.  Such as order entry.
}
