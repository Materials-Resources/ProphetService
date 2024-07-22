package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type UserPreference struct {
	bun.BaseModel     `bun:"table:user_preference"`
	UserPreferenceUid int32     `bun:"user_preference_uid,type:int,autoincrement,identity,pk"`       // Unique ID for the given user preference
	UserId            string    `bun:"user_id,type:varchar(30),unique"`                              // User ID
	PreferenceUid     int32     `bun:"preference_uid,type:int,unique"`                               // Preference ID - relates back to the preference table.
	Value             string    `bun:"value,type:varchar(max)"`                                      // Value of the preference
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
