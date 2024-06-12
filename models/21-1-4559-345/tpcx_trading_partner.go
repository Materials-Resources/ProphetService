package model

import (
	"github.com/uptrace/bun"
	"time"
)

type TpcxTradingPartner struct {
	bun.BaseModel         `bun:"table:tpcx_trading_partner"`
	TpcxTradingPartnerUid int32     `bun:"tpcx_trading_partner_uid,type:int,pk"`
	TradingPartnerCd      float64   `bun:"trading_partner_cd,type:decimal(19,0),nullzero"`
	TradingPartnerName    string    `bun:"trading_partner_name,type:varchar(255),nullzero"`
	Relationship          int32     `bun:"relationship,type:int"`
	ContactName           string    `bun:"contact_name,type:varchar(255),nullzero"`
	ContactTitle          string    `bun:"contact_title,type:varchar(255),nullzero"`
	ContactEmail          string    `bun:"contact_email,type:varchar(255),nullzero"`
	ContactPhoneNo        string    `bun:"contact_phone_no,type:varchar(255),nullzero"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
