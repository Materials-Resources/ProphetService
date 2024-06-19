package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type OeHdrSalesrep struct {
	bun.BaseModel              `bun:"table:oe_hdr_salesrep"`
	OrderNumber                string    `bun:"order_number,type:varchar(8),pk"`                               // Order number
	SalesrepId                 string    `bun:"salesrep_id,type:varchar(16),pk"`                               // Salesrep ID
	CommissionSplit            float64   `bun:"commission_split,type:decimal(5,2)"`                            // Commission percentage for this sales rep / order
	DeleteFlag                 string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated                time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	DatePaid                   time.Time `bun:"date_paid,type:datetime,nullzero"`                              // Date the commission was paid.
	PrimarySalesrep            string    `bun:"primary_salesrep,type:char(1),nullzero"`                        // Indicates whether this is the primary salesrep for this order.
	ExcludeSplitValidationFlag string    `bun:"exclude_split_validation_flag,type:char(1),default:('Y')"`      // Excludes this salesrep from any validation that requires the salesreps' total commission_percentage to add up to 100%.
	CommissionOverridePercent  float64   `bun:"commission_override_percent,type:decimal(19,9),nullzero"`       // Percent of the sales amount to be used to override the normal commission calculation for the associated salesrep. If field is left NULL, commission will be calculated using normal commission rules.
}
