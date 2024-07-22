package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PriceBook struct {
	bun.BaseModel    `bun:"table:price_book"`
	PriceBookUid     int32     `bun:"price_book_uid,type:int,pk"`          // Internal record identifier
	PriceBookId      string    `bun:"price_book_id,type:varchar(20)"`      // User-defined Sales Pricing Book identifier
	Description      string    `bun:"description,type:varchar(255)"`       // Description assigned to the Sales Pricing Book
	RowStatusFlag    int16     `bun:"row_status_flag,type:smallint"`       // Indicates current record status.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`    // Indicates the date/time this record was last modified.
	DateCreated      time.Time `bun:"date_created,type:datetime"`          // Indicates the date/time this record was created.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30)"` // ID of the user who last maintained this record
	CreatedBy        *string   `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	CompanyId        *string   `bun:"company_id,type:varchar(255)"` // Company that this book is for.
}
