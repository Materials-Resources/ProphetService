package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ReportCodeGroupP21 struct {
	bun.BaseModel         `bun:"table:report_code_group_p21"`
	ReportCodeGroupP21Uid int32     `bun:"report_code_group_p21_uid,type:int,pk,identity"`
	CodeGroupNo           int32     `bun:"code_group_no,type:int"`
	GroupTypeNo           int32     `bun:"group_type_no,type:int,nullzero"`
	CodeGroupDescription  string    `bun:"code_group_description,type:varchar(255)"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
