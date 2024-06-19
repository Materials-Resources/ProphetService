package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ContactsXSupplier struct {
	bun.BaseModel    `bun:"table:contacts_x_supplier"`
	ContactId        string    `bun:"contact_id,type:varchar(16),pk"`              // Unique code to identify a Contact
	SupplierId       float64   `bun:"supplier_id,type:decimal(19,0),pk"`           // Unique code to identify a Supplier
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`                    // Indicates current record status
	DateCreated      time.Time `bun:"date_created,type:datetime"`                  // Date/time this record was created
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`            // Date/time this record was last modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30)"`         // ID of the user that last modified this record
	PedigreeContact  string    `bun:"pedigree_contact,type:char(1),default:('N')"` // Serves as pedigree relationship point of contact
}
