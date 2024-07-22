package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ContactsXShipTo struct {
	bun.BaseModel      `bun:"table:contacts_x_ship_to"`
	ContactsXShipToUid int32     `bun:"contacts_x_ship_to_uid,type:int,pk"`                           // Unique record identifier
	ContactId          string    `bun:"contact_id,type:varchar(16)"`                                  // Unique code to identify a Contact
	CompanyId          string    `bun:"company_id,type:varchar(8)"`                                   // Unique code that identifies a company
	ShipToId           float64   `bun:"ship_to_id,type:decimal(19,0)"`                                // Unique code to identify a Ship to
	RowStatusFlag      int32     `bun:"row_status_flag,type:int"`                                     // Indicates current record status
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	PedigreeContact    string    `bun:"pedigree_contact,type:char(1),default:('N')"`                  // Serves as pedigree relationship point of contact
}
