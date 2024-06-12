package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PricePageImportColumn struct {
	bun.BaseModel            `bun:"table:price_page_import_column"`
	PricePageImportColumnUid int32     `bun:"price_page_import_column_uid,type:int,pk,identity"`
	Name                     string    `bun:"name,type:varchar(50)"`
	Description              string    `bun:"description,type:varchar(255)"`
	DataType                 string    `bun:"data_type,type:varchar(10)"`
	Length                   int32     `bun:"length,type:int"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
