package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type JobPriceCustomerShipto struct {
	bun.BaseModel                `bun:"table:job_price_customer_shipto"`
	JobPriceCustShiptoUid        int32      `bun:"job_price_cust_shipto_uid,type:int,pk"`                        // Uniquely Identifier for each record in the table.
	JobPriceHdrUid               int32      `bun:"job_price_hdr_uid,type:int,unique"`                            // Unique ID for associated Job/Contract.
	CompanyId                    string     `bun:"company_id,type:varchar(8),unique"`                            // Identifies the company.
	CustomerId                   float64    `bun:"customer_id,type:decimal(19,0),unique"`                        // The customer identifier.
	ShipToId                     float64    `bun:"ship_to_id,type:decimal(19,0),unique"`                         // The ship to identifier.
	PreferredLocationId          *float64   `bun:"preferred_location_id,type:decimal(19,0)"`                     // Location from which inventory should be shipped.
	TermsId                      *string    `bun:"terms_id,type:varchar(2)"`                                     // The terms you are normally granting when contracting material.
	ConsolidatedInvoicing        *string    `bun:"consolidated_invoicing,type:char(1)"`                          // Yes or No value field.
	AcceptableWaitTime           *float64   `bun:"acceptable_wait_time,type:decimal(19,0)"`                      // Indicates acceptable wait time for this customer.
	DefaultDisposition           *string    `bun:"default_disposition,type:char(1)"`                             // Indicates the what Default Disposition should be used when the material is out of stock
	DefaultCarrierId             *float64   `bun:"default_carrier_id,type:decimal(19,0)"`                        // Indicates the what carrier should normally be used while shipping orders.
	FreightCdUid                 *int32     `bun:"freight_cd_uid,type:int"`                                      // Indicates the freight code associated with shipto.
	ContractDollarLimit          *float64   `bun:"contract_dollar_limit,type:decimal(19,9)"`                     // Optional $ value associated with this contract.
	ShipToContactId              *string    `bun:"ship_to_contact_id,type:varchar(16)"`                          // Indicates the contact_id of the ship tos
	DateCreated                  time.Time  `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                    string     `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified             time.Time  `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy             string     `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	RowStatusFlag                int32      `bun:"row_status_flag,type:int"`                                     // Status of row (e.g., active, inactive, deleted)
	LocationId                   *float64   `bun:"location_id,type:decimal(19,0)"`                               // Indicates the default sales location for the ship tos
	ContractDollarLimitUsed      float64    `bun:"contract_dollar_limit_used,type:decimal(19,9),default:(0)"`    // Stores the dollar amount used against the contract
	CustPoNo                     *string    `bun:"cust_po_no,type:varchar(255)"`                                 // Customer / ShipTo Purchase Order number specific to this Job Pricing Customer / ShipTo record.
	ContractOrderLimit           *int32     `bun:"contract_order_limit,type:int"`                                // Custom - F23038: maximum number of orders that can be associated w/this contract.
	ContractOrderLimitUsed       *int32     `bun:"contract_order_limit_used,type:int"`                           // Custom - F23038: number of orders that have been associated w/this contract
	ActivationDate               *time.Time `bun:"activation_date,type:datetime"`                                // Date of activation for the record
	ExpirationDate               *time.Time `bun:"expiration_date,type:datetime"`                                // Date of expiration for the record.
	AutoAssignedFlag             string     `bun:"auto_assigned_flag,type:char(1),default:('N')"`                // To determine a record was manually added or automatically added to the system.
	ExcludeFromFreightFactorFlag string     `bun:"exclude_from_freight_factor_flag,type:char(1),default:('N')"`  // Indicates if customer/ship-to should be excluded from freight factor calculations for the price pages linked to the lines on this contract
	Multiplier                   *float64   `bun:"multiplier,type:decimal(19,9)"`                                // Customer column to specify what number to multiply at Customer/ShipTo level
}
