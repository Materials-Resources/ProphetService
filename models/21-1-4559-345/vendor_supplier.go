package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type VendorSupplier struct {
	bun.BaseModel    `bun:"table:vendor_supplier"`
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`                                 // Unique code that identifies a company.
	VendorId         float64   `bun:"vendor_id,type:decimal(19,0),pk"`                               // What is the vendor ID?
	SupplierId       float64   `bun:"supplier_id,type:decimal(19,0),pk"`                             // What is the supplier ID?
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	PrimaryVendor    string    `bun:"primary_vendor,type:char(1),nullzero"`                          // Indicates whether this vendor is primary.
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
}
