package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InventorySupplierXLoc struct {
	bun.BaseModel              `bun:"table:inventory_supplier_x_loc"`
	InventorySupplierXLocUid   int32      `bun:"inventory_supplier_x_loc_uid,type:int,pk"`                  // Unique identifier for record.
	InventorySupplierUid       int32      `bun:"inventory_supplier_uid,type:int,unique"`                    // Unique identifier from inventory_supplier table.
	LocationId                 float64    `bun:"location_id,type:decimal(19,0),unique"`                     // Identifier for location id.
	PrimarySupplier            string     `bun:"primary_supplier,type:char(1),default:('N')"`               // The supplier that is denoted as being used most often when ordering an item.
	AverageLeadTime            int16      `bun:"average_lead_time,type:smallint,default:(0)"`               // The average of the last x lead times for a supplier or an item where x = the number of lead times defined in System Settings Inventory Management General.
	RowStatusFlag              int32      `bun:"row_status_flag,type:int"`                                  // Indicates current record status.
	DateCreated                time.Time  `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified           time.Time  `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy           string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	LocListPrice               *float64   `bun:"loc_list_price,type:decimal(19,9),default:(0)"`             // Location Specific Supplier List Price
	LocCost                    *float64   `bun:"loc_cost,type:decimal(19,9),default:(0)"`                   // Location Specific Supplier Cost
	LocContractNumber          *string    `bun:"loc_contract_number,type:varchar(40)"`                      // contract number assigned at item/supplier level
	FixedLeadTime              *int16     `bun:"fixed_lead_time,type:smallint,default:(0)"`                 // Indicates what is the fixed lead time for the combination of location, supplier and item.
	DefaultDisposition         *string    `bun:"default_disposition,type:char(1)"`                          // Default Disposition for consignment
	VmiStatus                  *int32     `bun:"vmi_status,type:int"`                                       // Define a Primary Supplier/Location/Item as a Vendor Managed Inventory item
	CreatedBy                  *string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	OverrideBegDate            *time.Time `bun:"override_beg_date,type:datetime"`                           // The beginning date for the vmi status override
	OverrideEndDate            *time.Time `bun:"override_end_date,type:datetime"`                           // The ending date for the vmi status override
	VmiLastSendDate            *time.Time `bun:"vmi_last_send_date,type:datetime"`                          // The last date that vmi data was exported for the supplier/location
	OverrideVmiStatus          *int32     `bun:"override_vmi_status,type:int"`                              // Indicates the vmi status (manually done by customer) that overrides the system vmi status.
	OverrideVmiStatusFlag      string     `bun:"override_vmi_status_flag,type:char(1),default:('N')"`       // Override system generated vmi status
	KeyVmiIndicatorChangedFlag string     `bun:"key_vmi_indicator_changed_flag,type:char(1),default:('N')"` // Determines if one of the key vmi indicators has changed.  Reset during 852 export.
	FutureCost                 *float64   `bun:"future_cost,type:decimal(19,9)"`                            // Future cost of an item
	EffectiveDate              *time.Time `bun:"effective_date,type:datetime"`                              // Date future cost will be effective
	ManualLeadTime             *int32     `bun:"manual_lead_time,type:int"`                                 // Override the Average Lead Time if set.
	FutureListPrice            *float64   `bun:"future_list_price,type:decimal(19,9)"`                      // Custom Column for Future List Price of an item
}
