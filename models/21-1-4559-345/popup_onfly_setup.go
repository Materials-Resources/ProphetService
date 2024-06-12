package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PopupOnflySetup struct {
	bun.BaseModel      `bun:"table:popup_onfly_setup"`
	PopupOnflySetupUid int32     `bun:"popup_onfly_setup_uid,type:int,pk,identity"`
	SearchKey          string    `bun:"search_key,type:varchar(500)"`
	ParentColumn       string    `bun:"parent_column,type:varchar(100),nullzero"`
	ChildColumn        string    `bun:"child_column,type:varchar(100),nullzero"`
	IsPinned           bool      `bun:"is_pinned,type:bit,default:((0))"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	AditionalColumn    string    `bun:"aditional_column,type:varchar(255),nullzero"`
}
