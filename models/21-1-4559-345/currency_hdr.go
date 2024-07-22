package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CurrencyHdr struct {
	bun.BaseModel              `bun:"table:currency_hdr"`
	CurrencyId                 float64   `bun:"currency_id,type:decimal(19,0),pk"`                         // What is the unique currency identifier for this ro
	CurrencyDesc               string    `bun:"currency_desc,type:varchar(20)"`                            // What currency is this?
	DeleteFlag                 string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated                time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	CurrencyMask               *string   `bun:"currency_mask,type:varchar(50)"`                            // What display mask should be used for this currency?
	AvailableForOrdersInvoices string    `bun:"available_for_orders_invoices,type:char(1),default:('Y')"`  // Specifies a currency is avilable for orders and invoices.
	IsoCurrencyCd              *int32    `bun:"iso_currency_cd,type:int"`                                  // ISO Currency Code
}
