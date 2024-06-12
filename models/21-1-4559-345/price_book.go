package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PriceBook struct {
	bun.BaseModel    `bun:"table:price_book"`
	PriceBookUid     int32     `bun:"price_book_uid,type:int,pk"`
	PriceBookId      string    `bun:"price_book_id,type:varchar(20)"`
	Description      string    `bun:"description,type:varchar(255)"`
	RowStatusFlag    int16     `bun:"row_status_flag,type:smallint"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30)"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	CompanyId        string    `bun:"company_id,type:varchar(255),nullzero"`
}
