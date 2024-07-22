package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type AlertTypeXToken struct {
	bun.BaseModel      `bun:"table:alert_type_x_token"`
	AlertTypeXTokenUid int32     `bun:"alert_type_x_token_uid,type:int,pk"`                        // Unique Identifier for alert_type_x_token table.
	AlertTypeUid       int32     `bun:"alert_type_uid,type:int,unique"`                            // Unique Identfier from alert_type table.
	TokenUid           int32     `bun:"token_uid,type:int,unique"`                                 // Unique Identifier from token table.
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`      // Indicates the date/time this record was last modified.
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
