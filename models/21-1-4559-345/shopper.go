package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Shopper struct {
	bun.BaseModel       `bun:"table:shopper"`
	ShopperUid          int32      `bun:"shopper_uid,type:int,pk"`                                                                                                                                                                                                       // Unique Identifier
	WebShopperId        string     `bun:"web_shopper_id,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`                           // Shopper Identifier, ex shopper_administrator, 16642.
	WebLoginName        string     `bun:"web_login_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,default:('')"`              // Shopper name, ex Bubbles Powerpuff.
	WebLoginPassword    string     `bun:"web_login_password,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,default:('')"`          // shopper password.
	WebLoginHint        string     `bun:"web_login_hint,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,default:('')"`              // Hint to help remind the shopper of a forgotten password.
	WebEmail            string     `bun:"web_email,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,default:('')"`                   // Shoppers email address.
	DateLastLogon       *time.Time `bun:"date_last_logon,type:datetime"`                                                                                                                                                                                                 // Date the shopper lasts logged on.
	DateLastLogoff      *time.Time `bun:"date_last_logoff,type:datetime"`                                                                                                                                                                                                // Date the shopper lasts logged off.
	AuthorizationTypeId *int32     `bun:"authorization_type_id,type:int"`                                                                                                                                                                                                // Whether the user has admin rights.
	DateCreated         time.Time  `bun:"date_created,type:datetime,default:(getdate())"`                                                                                                                                                                                // Indicates the date/time this record was created.
	DateLastModified    time.Time  `bun:"date_last_modified,type:datetime,default:(getdate())"`                                                                                                                                                                          // Indicates the date/time this record was last modified.
	LastMaintainedBy    string     `bun:"last_maintained_by,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,default:(user_name())"` // ID of the user who last maintained this record
	RowStatusFlag       int16      `bun:"row_status_flag,type:smallint,default:(704)"`                                                                                                                                                                                   // Indicates current record status.
	Random              *int32     `bun:"random,type:int"`                                                                                                                                                                                                               // unknown.
	DefaultShip2Id      *float64   `bun:"default_ship2_id,type:decimal(19,0)"`                                                                                                                                                                                           // Default ship to id for the shopper/
}
