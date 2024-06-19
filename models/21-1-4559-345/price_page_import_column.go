package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PricePageImportColumn struct {
	bun.BaseModel            `bun:"table:price_page_import_column"`
	PricePageImportColumnUid int32     `bun:"price_page_import_column_uid,type:int,autoincrement,pk"`       // Unique identifier for this record
	Name                     string    `bun:"name,type:varchar(50)"`                                        // Datawindow column name for the column
	Description              string    `bun:"description,type:varchar(255)"`                                // Description of the column
	DataType                 string    `bun:"data_type,type:varchar(10)"`                                   // Data type of the column
	Length                   int32     `bun:"length,type:int"`                                              // Data length  of the column
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
