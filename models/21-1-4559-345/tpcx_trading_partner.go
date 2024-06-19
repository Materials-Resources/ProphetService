package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type TpcxTradingPartner struct {
	bun.BaseModel         `bun:"table:tpcx_trading_partner"`
	TpcxTradingPartnerUid int32     `bun:"tpcx_trading_partner_uid,type:int,pk"`                         // Primary Key
	TradingPartnerCd      float64   `bun:"trading_partner_cd,type:decimal(19,0),nullzero"`               // Trading Partner ID, Supplier ID in CC systems
	TradingPartnerName    string    `bun:"trading_partner_name,type:varchar(255),nullzero"`              // Trading Partner Name
	Relationship          int32     `bun:"relationship,type:int"`                                        // Numberic code value representing the type of relationship between us and the trading partner
	ContactName           string    `bun:"contact_name,type:varchar(255),nullzero"`                      // Name of main contact at the trading partner
	ContactTitle          string    `bun:"contact_title,type:varchar(255),nullzero"`                     // Title of the contact at the trading partner
	ContactEmail          string    `bun:"contact_email,type:varchar(255),nullzero"`                     // Email address of the Trading Partners contact.
	ContactPhoneNo        string    `bun:"contact_phone_no,type:varchar(255),nullzero"`                  // Phone number of the Trading Partners contact
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
