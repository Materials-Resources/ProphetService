package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type RibbonTool struct {
	bun.BaseModel      `bun:"table:ribbon_tool"`
	RibbonToolUid      int32     `bun:"ribbon_tool_uid,type:int,autoincrement,identity,pk"`
	ToolId             string    `bun:"tool_id,type:varchar(255)"`
	ToolText           string    `bun:"tool_text,type:varchar(255)"`
	Description        string    `bun:"description,type:varchar(255)"`
	ErpMenu            string    `bun:"erp_menu,type:varchar(255),nullzero"`
	ErpEventMessage    string    `bun:"erp_event_message,type:varchar(255),nullzero"`
	ErpMenuAttribute   string    `bun:"erp_menu_attribute,type:varchar(255),nullzero"`
	ToolTypeCd         int32     `bun:"tool_type_cd,type:int"`
	DefaultToolSize    int32     `bun:"default_tool_size,type:int"`
	ImageFile          string    `bun:"image_file,type:varchar(255),nullzero"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(user_name())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(user_name())"`
	UsedInUiserverFlag string    `bun:"used_in_uiserver_flag,type:char(1),default:('N')"` // flags whether tool is used in the UIServer
}
