package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ExtensibilityWindow struct {
	bun.BaseModel          `bun:"table:extensibility_window"`
	ExtensibilityWindowUid int32     `bun:"extensibility_window_uid,type:int,pk,identity"`
	WindowName             string    `bun:"window_name,type:varchar(255)"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	EnableFlag             string    `bun:"enable_flag,type:char,default:('Y')"`
}
