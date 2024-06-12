package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DirectionRecentSearch struct {
	bun.BaseModel            `bun:"table:direction_recent_search"`
	DirectionRecentSearchUid int32     `bun:"direction_recent_search_uid,type:int,pk,identity"`
	UserId                   string    `bun:"user_id,type:varchar(30)"`
	Name                     string    `bun:"name,type:varchar(255),nullzero"`
	Address1                 string    `bun:"address1,type:varchar(255),nullzero"`
	City                     string    `bun:"city,type:varchar(255),nullzero"`
	State                    string    `bun:"state,type:varchar(255),nullzero"`
	Zip                      string    `bun:"zip,type:varchar(255),nullzero"`
	Country                  string    `bun:"country,type:varchar(255),nullzero"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
