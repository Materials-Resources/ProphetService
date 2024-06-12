package model

import (
	"github.com/uptrace/bun"
	"time"
)

type MasterInquiryTabDefault struct {
	bun.BaseModel              `bun:"table:master_inquiry_tab_default"`
	MasterInquiryTabDefaultUid int32     `bun:"master_inquiry_tab_default_uid,type:int,pk,identity"`
	MasterInquiryTypeCd        int32     `bun:"master_inquiry_type_cd,type:int"`
	UserId                     string    `bun:"user_id,type:varchar(30)"`
	DefaultTabIndex            int16     `bun:"default_tab_index,type:smallint"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
