package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InventoryReturnHdr struct {
	bun.BaseModel               `bun:"table:inventory_return_hdr"`
	InventoryReturnHdrUid       int32     `bun:"inventory_return_hdr_uid,type:int,pk"`                         // Unique Identifier for this record.
	ReturnNumber                int32     `bun:"return_number,type:int,unique"`                                // Return number - used by user to access record.
	CompanyId                   string    `bun:"company_id,type:varchar(8)"`                                   // Unique code that identifies a company.
	VendorId                    float64   `bun:"vendor_id,type:decimal(19,0)"`                                 // What is the unique vendor identifier for this row?
	SupplierId                  float64   `bun:"supplier_id,type:decimal(19,0)"`                               // What supplier supplies material for this stage?
	DivisionId                  float64   `bun:"division_id,type:decimal(19,0)"`                               // What is the unique identifier for this division?
	ReturnDate                  time.Time `bun:"return_date,type:datetime,default:(getdate())"`                // Date inventory return was entered
	ExpectedDate                time.Time `bun:"expected_date,type:datetime,default:(getdate())"`              // What is the expected date that this relatinship between a process and a transaction should be completed?
	FreightAmount               float64   `bun:"freight_amount,type:decimal(19,4),default:(0)"`                // Freight amount in standard currency.
	FreightAmountDisplay        float64   `bun:"freight_amount_display,type:decimal(19,4),default:(0)"`        // Holds the display freight amount. It is necessary
	RestockingFeeAmount         float64   `bun:"restocking_fee_amount,type:decimal(19,4),default:(0)"`         // Restocking fee amount in standard currency
	RestockingFeeAmountDisplay  float64   `bun:"restocking_fee_amount_display,type:decimal(19,4),default:(0)"` // Restocking fee amount in foreign currency
	RmaNumber                   string    `bun:"rma_number,type:varchar(255),nullzero"`                        // Used by some users to track vendors authorization number.
	PoReferenceNumber           float64   `bun:"po_reference_number,type:decimal(19,0),nullzero"`              // PO Number for the items being returned
	Ship2Name                   string    `bun:"ship2_name,type:varchar(255)"`                                 // Ship to name.
	Ship2Address1               string    `bun:"ship2_address1,type:varchar(255),nullzero"`                    // What is the first line of the ship to address?
	Ship2Address2               string    `bun:"ship2_address2,type:varchar(255),nullzero"`                    // What is the second  line of the ship to address?
	Ship2City                   string    `bun:"ship2_city,type:varchar(255),nullzero"`                        // What city should this order be shipped to?
	Ship2State                  string    `bun:"ship2_state,type:varchar(255),nullzero"`                       // What is the state  for the ship to address?
	Ship2Zip                    string    `bun:"ship2_zip,type:varchar(255),nullzero"`                         // What postal code should this order be shipped to?
	Ship2Country                string    `bun:"ship2_country,type:varchar(255),nullzero"`                     // What country should this order be shipped to?
	CurrencyId                  float64   `bun:"currency_id,type:decimal(19,0),default:(1)"`                   // What is the unique currency identifier for this ro
	RowStatusFlag               int32     `bun:"row_status_flag,type:int"`                                     // Indicates current record status.
	DateCreated                 time.Time `bun:"date_created,type:datetime"`                                   // Indicates the date/time this record was created.
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`                             // Indicates the date/time this record was last modified.
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`    // ID of the user who last maintained this record
	LocationId                  float64   `bun:"location_id,type:decimal(19,0)"`                               // What is the unique location identifier for this ro
	BuyerId                     string    `bun:"buyer_id,type:varchar(16)"`                                    // What buyer is responsible for this transfer?
	Ship2Phone                  string    `bun:"ship2_phone,type:varchar(20),nullzero"`                        // Ship to phone number
	CarrierId                   float64   `bun:"carrier_id,type:decimal(19,0),nullzero"`                       // What is the id of this carrier (if any)?
	CarrierTrackingNo           string    `bun:"carrier_tracking_no,type:varchar(255),nullzero"`               // The tracking number assigned by the carrier
	ChargeFreightToVendor       string    `bun:"charge_freight_to_vendor,type:char(1),default:('Y')"`
	ContactId                   string    `bun:"contact_id,type:varchar(16),nullzero"`                 // Contact associated with the inventory return
	RetrievedByWms              string    `bun:"retrieved_by_wms,type:char(1),default:('N')"`          // Determines whether or not this record has been exported.
	CurrencyLineUid             int32     `bun:"currency_line_uid,type:int,nullzero"`                  // Used to identify currency info for the Inventory Return, namely the exchange rate for a foreign currency transaction.
	InvReceiptsClearingAcct     string    `bun:"inv_receipts_clearing_acct,type:varchar(32),nullzero"` // The account that stores the inventory value of goods that have been received against a PO while you are still waiting for the invoice for those goods
	Ship2Address3               string    `bun:"ship2_address3,type:varchar(255),nullzero"`            // Ship to address line 3
	BlindShipFlag               string    `bun:"blind_ship_flag,type:char(1),nullzero"`                // Custom column to set an inventory return as a blind shipment so it does not appear to come from the distributor.
	Fob                         string    `bun:"fob,type:varchar(20),nullzero"`                        // custom column to default from division maint
	TrackingNoList              string    `bun:"tracking_no_list,type:varchar(1000),nullzero"`         // Custom column to hold multiple comma-delimited tracking no list
	ExpediteNotes               string    `bun:"expedite_notes,type:varchar(8000),nullzero"`           // Notes which help expedite inventory returns.
	WarrantyInventoryReturnFlag string    `bun:"warranty_inventory_return_flag,type:char(1),nullzero"` // Custom column to indicate the inventory return is a warranty inventory return
}
