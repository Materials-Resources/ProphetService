package model

import (
	"github.com/uptrace/bun"
	"time"
)

type OeLinePanel struct {
	bun.BaseModel      `bun:"table:oe_line_panel"`
	OeLinePanelUid     int32     `bun:"oe_line_panel_uid,type:int,pk,identity"`
	DefaultName        string    `bun:"default_name,type:varchar(255),nullzero"`
	DefaultSequenceNo  int32     `bun:"default_sequence_no,type:int"`
	DefaultDisplayInOe string    `bun:"default_display_in_oe,type:char"`
	RealDw             string    `bun:"real_dw,type:char"`
	SystemSettingName  string    `bun:"system_setting_name,type:varchar(128),nullzero"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
