package model

import (
	"github.com/uptrace/bun"
	"time"
)

type VendorSupplier struct {
	bun.BaseModel    `bun:"table:vendor_supplier"`
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`
	VendorId         float64   `bun:"vendor_id,type:decimal(19,0),pk"`
	SupplierId       float64   `bun:"supplier_id,type:decimal(19,0),pk"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	PrimaryVendor    string    `bun:"primary_vendor,type:char,nullzero"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
}
