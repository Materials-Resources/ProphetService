package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type SystemAlerts struct {
	bun.BaseModel     `bun:"table:system_alerts"`
	SystemAlertsUid   int32     `bun:"system_alerts_uid,type:int,pk"`       // Unique identifier for the alert.
	UserId            string    `bun:"user_id,type:varchar(30)"`            // This column is unused.
	SystemAlertTypeCd int32     `bun:"system_alert_type_cd,type:int"`       // Will identify the type of system alert being issued.
	Subject           string    `bun:"subject,type:varchar(255)"`           // Will store the subject of a system alert.
	Notes             string    `bun:"notes,type:varchar(255)"`             // This will be the body or main text of a system alert
	RowStatusFlag     int32     `bun:"row_status_flag,type:int"`            // Indicates current record status.
	DateCreated       time.Time `bun:"date_created,type:datetime"`          // Indicates the date/time this record was created.
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`    // Indicates the date/time this record was last modified.
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30)"` // ID of the user who last maintained this record
}
