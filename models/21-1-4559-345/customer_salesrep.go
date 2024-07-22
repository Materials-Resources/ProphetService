package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerSalesrep struct {
	bun.BaseModel        `bun:"table:customer_salesrep"`
	CustomerSalesrepUid  int32     `bun:"customer_salesrep_uid,type:int,pk"`                            // Unique id for the record
	CompanyId            string    `bun:"company_id,type:varchar(8),unique"`                            // Unique id assigned to a company
	CustomerId           float64   `bun:"customer_id,type:decimal(19,0),unique"`                        // Unique id assigned to a customer
	SalesrepId           string    `bun:"salesrep_id,type:varchar(16),unique"`                          // Unique id assigned to an outside salesrep
	CommissionPercentage float64   `bun:"commission_percentage,type:decimal(19,4)"`                     // Percentage of commission split among the inside salesreps
	PrimarySalesrepFlag  string    `bun:"primary_salesrep_flag,type:char(1),default:('N')"`             // Indicate whether the salesrep is a primary salesrep for the customer
	RowStatusFlag        int32     `bun:"row_status_flag,type:int"`                                     // Indicate whether the record is active or delete
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
