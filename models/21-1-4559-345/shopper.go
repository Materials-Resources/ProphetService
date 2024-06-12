package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Shopper struct {
	bun.BaseModel       `bun:"table:shopper"`
	ShopperUid          int32     `bun:"shopper_uid,type:int,pk"`
	WebShopperId        string    `bun:"web_shopper_id,type:nvarchar(128)"`
	WebLoginName        string    `bun:"web_login_name,type:nvarchar(128),default:('')"`
	WebLoginPassword    string    `bun:"web_login_password,type:nvarchar(128),default:('')"`
	WebLoginHint        string    `bun:"web_login_hint,type:nvarchar(128),default:('')"`
	WebEmail            string    `bun:"web_email,type:nvarchar(40),default:('')"`
	DateLastLogon       time.Time `bun:"date_last_logon,type:datetime,nullzero"`
	DateLastLogoff      time.Time `bun:"date_last_logoff,type:datetime,nullzero"`
	AuthorizationTypeId int32     `bun:"authorization_type_id,type:int,nullzero"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:nvarchar(128),default:(user_name())"`
	RowStatusFlag       int16     `bun:"row_status_flag,type:smallint,default:(704)"`
	Random              int32     `bun:"random,type:int,nullzero"`
	DefaultShip2Id      float64   `bun:"default_ship2_id,type:decimal(19,0),nullzero"`
}
