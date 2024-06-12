package model

import (
	"github.com/uptrace/bun"
	"time"
)

type OeHdrSalesrep struct {
	bun.BaseModel              `bun:"table:oe_hdr_salesrep"`
	OrderNumber                string    `bun:"order_number,type:varchar(8),pk"`
	SalesrepId                 string    `bun:"salesrep_id,type:varchar(16),pk"`
	CommissionSplit            float64   `bun:"commission_split,type:decimal(5,2)"`
	DeleteFlag                 string    `bun:"delete_flag,type:char"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	DatePaid                   time.Time `bun:"date_paid,type:datetime,nullzero"`
	PrimarySalesrep            string    `bun:"primary_salesrep,type:char,nullzero"`
	ExcludeSplitValidationFlag string    `bun:"exclude_split_validation_flag,type:char,default:('Y')"`
	CommissionOverridePercent  float64   `bun:"commission_override_percent,type:decimal(19,9),nullzero"`
}
