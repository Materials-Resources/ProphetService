package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DirectionRecentSearch struct {
	bun.BaseModel            `bun:"table:direction_recent_search"`
	DirectionRecentSearchUid int32     `bun:"direction_recent_search_uid,type:int,autoincrement,identity,pk"` // Unique identifier for the table
	UserId                   string    `bun:"user_id,type:varchar(30)"`                                       // Addresses are displayed for individual users
	Name                     *string   `bun:"name,type:varchar(255)"`                                         // Address name
	Address1                 *string   `bun:"address1,type:varchar(255)"`                                     // Address 1
	City                     *string   `bun:"city,type:varchar(255)"`                                         // Address city
	State                    *string   `bun:"state,type:varchar(255)"`                                        // Address state
	Zip                      *string   `bun:"zip,type:varchar(255)"`                                          // Address zip
	Country                  *string   `bun:"country,type:varchar(255)"`                                      // Address country
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`                 // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`           // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`           // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`   // User who last changed the record
}
