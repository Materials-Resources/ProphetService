package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ReportEmailDefaultsXToken struct {
	bun.BaseModel                `bun:"table:report_email_defaults_x_token"`
	ReportEmailDefaultsXTokenUid int32     `bun:"report_email_defaults_x_token_uid,type:int,pk,identity"`
	TokenUid                     int32     `bun:"token_uid,type:int"`
	AvailableAreaCode            int32     `bun:"available_area_code,type:int"`
	DateCreated                  time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                    string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
