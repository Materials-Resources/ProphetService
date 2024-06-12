package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CurrencyHdr struct {
	bun.BaseModel              `bun:"table:currency_hdr"`
	CurrencyId                 float64   `bun:"currency_id,type:decimal(19,0),pk"`
	CurrencyDesc               string    `bun:"currency_desc,type:varchar(20)"`
	DeleteFlag                 string    `bun:"delete_flag,type:char"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	CurrencyMask               string    `bun:"currency_mask,type:varchar(50),nullzero"`
	AvailableForOrdersInvoices string    `bun:"available_for_orders_invoices,type:char,default:('Y')"`
	IsoCurrencyCd              int32     `bun:"iso_currency_cd,type:int,nullzero"`
}
