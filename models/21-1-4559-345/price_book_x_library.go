package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PriceBookXLibrary struct {
	bun.BaseModel        `bun:"table:price_book_x_library"`
	PriceBookXLibraryUid int32     `bun:"price_book_x_library_uid,type:int,pk"` // Internal record identifier
	PriceLibraryUid      int32     `bun:"price_library_uid,type:int"`           // Internal record identifier for the Sales Pricing Library
	PriceBookUid         int32     `bun:"price_book_uid,type:int"`              // Internal record identifier for the Sales Price Book
	SequenceNumber       int16     `bun:"sequence_number,type:smallint"`        // Sequence number of this Sales Price Book within the Sales Pricing Library
	RowStatusFlag        int16     `bun:"row_status_flag,type:smallint"`        // Indicates current record status.
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`     // Indicates the date/time this record was last modified.
	DateCreated          time.Time `bun:"date_created,type:datetime"`           // Indicates the date/time this record was created.
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30)"`  // ID of the user who last maintained this record
}
