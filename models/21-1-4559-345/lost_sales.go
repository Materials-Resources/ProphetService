package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type LostSales struct {
	bun.BaseModel    `bun:"table:lost_sales"`
	LostSalesUid     int32     `bun:"lost_sales_uid,type:int,autoincrement,identity,pk"`            // Unique identifier for lost sales code table
	CompanyId        string    `bun:"company_id,type:varchar(8),unique"`                            // Indicates the company for lost sales code
	LostSalesId      string    `bun:"lost_sales_id,type:varchar(255),unique"`                       // Lost sales reason identifier.
	LostSalesDesc    string    `bun:"lost_sales_desc,type:varchar(255)"`                            // Lost sales reason description.
	AffectUsage      string    `bun:"affect_usage,type:char(1)"`                                    // Indicates if this will affect usage when using this ID
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`                                     // Row status flag for record
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
