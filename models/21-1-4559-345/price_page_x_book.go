package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PricePageXBook struct {
	bun.BaseModel     `bun:"table:price_page_x_book"`
	PricePageXBookUid int32     `bun:"price_page_x_book_uid,type:int,pk"`   // Internal record identifier
	PriceBookUid      int32     `bun:"price_book_uid,type:int"`             // Internal record identifier for the Sales Price Book
	PricePageUid      int32     `bun:"price_page_uid,type:int"`             // Internal record identifier for the Sales Pricing Page
	RowStatusFlag     int16     `bun:"row_status_flag,type:smallint"`       // Indicates current record status.
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`    // Indicates the date/time this record was last modified.
	DateCreated       time.Time `bun:"date_created,type:datetime"`          // Indicates the date/time this record was created.
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30)"` // ID of the user who last maintained this record
}
