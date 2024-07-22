package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Division struct {
	bun.BaseModel             `bun:"table:division"`
	SupplierId                float64   `bun:"supplier_id,type:decimal(19,0),pk"`                         // What supplier supplies material for this stage?
	DivisionId                float64   `bun:"division_id,type:decimal(19,0),pk"`                         // What is the unique identifier for this division?
	DeleteFlag                string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated               time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	DivisionName              *string   `bun:"division_name,type:varchar(255)"`                           // What is the name of this division?
	CreatedBy                 *string   `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	ReturnDivisionFlag        *string   `bun:"return_division_flag,type:char(1),default:('N')"`          // Determines if the division is an inventory return division
	DefaultCarrierId          *float64  `bun:"default_carrier_id,type:decimal(19,0)"`                    // Determines the default carrier for inventory returns for this division.
	DefaultFob                *string   `bun:"default_fob,type:varchar(20)"`                             // Defaults the FOB setting for the inventory return.
	ChargeFreightToVendorFlag *string   `bun:"charge_freight_to_vendor_flag,type:char(1),default:('N')"` // Determines default for the Charge Freight to Vendor checkbox on the Charges tab in Inventory Return Entry
	DefaultAuthorizationNo    *string   `bun:"default_authorization_no,type:varchar(255)"`               // Determines the default return authorization number for returns created for the division.
}
