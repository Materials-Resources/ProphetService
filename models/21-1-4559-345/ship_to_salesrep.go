package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ShipToSalesrep struct {
	bun.BaseModel        `bun:"table:ship_to_salesrep"`
	CompanyId            string    `bun:"company_id,type:varchar(8),pk"`                                 // Unique code that identifies a company.
	ShipToId             float64   `bun:"ship_to_id,type:decimal(19,0),pk"`                              // What
	SalesrepId           string    `bun:"salesrep_id,type:varchar(16),pk"`                               // Who is the actual sales representative for this intersection?
	DeleteFlag           string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated          time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	PrimarySalesrep      *string   `bun:"primary_salesrep,type:char(1)"`
	CommissionPercentage float64   `bun:"commission_percentage,type:decimal(19,4),default:((100))"` // Percentage of commission split among the inside salesreps
	PrimaryServiceRep    string    `bun:"primary_service_rep,type:char(1),default:('N')"`           // Primary rep for service orders
}
