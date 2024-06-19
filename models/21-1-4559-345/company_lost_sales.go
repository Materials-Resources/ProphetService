package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CompanyLostSales struct {
	bun.BaseModel       `bun:"table:company_lost_sales"`
	CompanyLostSalesUid int32     `bun:"company_lost_sales_uid,type:int,autoincrement,pk"`             // Unique identifier for company lost sales table
	CompanyId           string    `bun:"company_id,type:varchar(8),unique"`                            // Identifies the company for lost sales settings
	TransactionCodeNo   int32     `bun:"transaction_code_no,type:int,unique"`                          // Indicates what type of transaction was the quantity cancelled in. Order - Cancelled, Order Edited, Shipment Cancelled, Shipped Short etc.
	ActionCodeNo        int32     `bun:"action_code_no,type:int"`                                      // Action taken at transaction entry
	LostSalesUid        int32     `bun:"lost_sales_uid,type:int,nullzero"`                             // Reference to the lost sales UID for defaulting
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
