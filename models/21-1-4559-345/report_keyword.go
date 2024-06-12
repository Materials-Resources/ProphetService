package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ReportKeyword struct {
	bun.BaseModel    `bun:"table:report_keyword"`
	ReportKeywordUid int32     `bun:"report_keyword_uid,type:int,pk,identity"`
	Keyword          string    `bun:"keyword,type:varchar(250)"`
	Statement        string    `bun:"statement,type:varchar(250)"`
	Alias            string    `bun:"alias,type:varchar(250)"`
	Description      string    `bun:"description,type:varchar(250)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
