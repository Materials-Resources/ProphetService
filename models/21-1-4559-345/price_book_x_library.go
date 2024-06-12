package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PriceBookXLibrary struct {
	bun.BaseModel        `bun:"table:price_book_x_library"`
	PriceBookXLibraryUid int32     `bun:"price_book_x_library_uid,type:int,pk"`
	PriceLibraryUid      int32     `bun:"price_library_uid,type:int"`
	PriceBookUid         int32     `bun:"price_book_uid,type:int"`
	SequenceNumber       int16     `bun:"sequence_number,type:smallint"`
	RowStatusFlag        int16     `bun:"row_status_flag,type:smallint"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`
	DateCreated          time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30)"`
}
