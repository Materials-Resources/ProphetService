package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerSalesrep struct {
	bun.BaseModel        `bun:"table:customer_salesrep"`
	CustomerSalesrepUid  int32     `bun:"customer_salesrep_uid,type:int,pk"`
	CompanyId            string    `bun:"company_id,type:varchar(8)"`
	CustomerId           float64   `bun:"customer_id,type:decimal(19,0)"`
	SalesrepId           string    `bun:"salesrep_id,type:varchar(16)"`
	CommissionPercentage float64   `bun:"commission_percentage,type:decimal(19,4)"`
	PrimarySalesrepFlag  string    `bun:"primary_salesrep_flag,type:char,default:('N')"`
	RowStatusFlag        int32     `bun:"row_status_flag,type:int"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
