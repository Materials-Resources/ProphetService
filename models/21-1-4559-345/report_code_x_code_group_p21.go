package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ReportCodeXCodeGroupP21 struct {
	bun.BaseModel              `bun:"table:report_code_x_code_group_p21"`
	ReportCodeXCodeGroupP21Uid int32     `bun:"report_code_x_code_group_p21_uid,type:int,pk,identity"`
	ReportCodeGroupP21Uid      int32     `bun:"report_code_group_p21_uid,type:int"`
	ReportCodeP21Uid           int32     `bun:"report_code_p21_uid,type:int"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
