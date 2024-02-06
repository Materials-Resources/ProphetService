package model

import (
	"time"

	"github.com/uptrace/bun"
)

type PricePageXBook struct {
	bun.BaseModel     `bun:"table:price_page_x_book"`
	PricePageXBookUid int32     `bun:"price_page_x_book_uid,type:int,pk"`
	PriceBookUid      int32     `bun:"price_book_uid,type:int"`
	PricePageUid      int32     `bun:"price_page_uid,type:int"`
	RowStatusFlag     int16     `bun:"row_status_flag,type:smallint"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`
	DateCreated       time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30)"`

	PricePage *PricePage `bun:"rel:belongs-to,join:price_page_uid=price_page_uid"`
}
