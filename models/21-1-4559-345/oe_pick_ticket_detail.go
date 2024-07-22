package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type OePickTicketDetail struct {
	bun.BaseModel               `bun:"table:oe_pick_ticket_detail"`
	CompanyId                   string    `bun:"company_id,type:varchar(8)"`                                    // Unique code that identifies a company.
	PickTicketNo                float64   `bun:"pick_ticket_no,type:decimal(19,0),pk"`                          // Pick ticket number.
	LineNumber                  float64   `bun:"line_number,type:decimal(19,0),pk"`                             // Invoice line number
	PrintQuantity               *float64  `bun:"print_quantity,type:decimal(19,9)"`                             // Quantity printed.
	ShipQuantity                *float64  `bun:"ship_quantity,type:decimal(19,9)"`                              // Quantity shipped.
	FreightIn                   *float64  `bun:"freight_in,type:decimal(19,4)"`                                 // Freight cost incurred to ship the item in to a warehouse
	DateCreated                 time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	UnitOfMeasure               string    `bun:"unit_of_measure,type:varchar(8)"`                               // What is the unit of measure for this row?
	UnitSize                    float64   `bun:"unit_size,type:decimal(19,4)"`                                  // The size of the sales unit of measure.
	UnitQuantity                float64   `bun:"unit_quantity,type:decimal(19,9)"`                              // The quantity shipped in terms of sales units.
	OeLineNo                    float64   `bun:"oe_line_no,type:decimal(19,0)"`                                 // Order line number.
	QtyRequested                *float64  `bun:"qty_requested,type:decimal(19,9)"`                              // How many/much of the item was requested on the order?
	QtyToPick                   *float64  `bun:"qty_to_pick,type:decimal(19,9)"`                                // Contains the quantity that needs to be picked (not
	Staged                      *string   `bun:"staged,type:char(1)"`                                           // Indicated if this material has been placed in a holding area for later shipment
	ReleaseNo                   *float64  `bun:"release_no,type:decimal(6,0)"`                                  // To properly update qty_invoiced on scheduled relea
	InvMastUid                  int32     `bun:"inv_mast_uid,type:int"`                                         // Unique identifier for the item id.
	InvoiceLineUid              *int32    `bun:"invoice_line_uid,type:int"`                                     // Unique identifier for the invoice line.
	FreightCodeUid              *int32    `bun:"freight_code_uid,type:int"`                                     // Freight code for this line.
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`          // User who created the record
	UnitQuantityEdited          *string   `bun:"unit_quantity_edited,type:char(1),default:('N')"`               // Indicates if the item had its quantity edited
	EnvironmentalFeeFlag        *string   `bun:"environmental_fee_flag,type:char(1)"`                           // Indicates if the customer will be charged an environmental fee.
	AdminFeeFlag                *string   `bun:"admin_fee_flag,type:char(1)"`                                   // Indicates if this customer should be charged a handling/admin fee
	InspectedFlag               *string   `bun:"inspected_flag,type:char(1),default:('N')"`                     // Indicates that the line item has been manually inspected.
	OriginalFreightIn           *float64  `bun:"original_freight_in,type:decimal(19,9)"`                        // Freight value orginally entered for a line if strategic pricing is enabled.
	QtyScanned                  *float64  `bun:"qty_scanned,type:decimal(19,9)"`                                // Custom  (Shipping window item  scan feature): holds the qty shipped when shipment is saved unconfirmed.
	FreightInMinAppliedFlag     *string   `bun:"freight_in_min_applied_flag,type:char(1)"`                      // Custom (F46747): determines if a supplier in-freight minimum was applied to this line.
	BoxNumber                   *string   `bun:"box_number,type:varchar(255)"`                                  // Custom column that indicates a Box Number per line item within the Shipping window.  The user will be able to specify an alphanumeric box number at the line item level; this data will not be required.
	OriginalQtyToPick           *float64  `bun:"original_qty_to_pick,type:decimal(19,9)"`                       // Keeps track of the original qty to pick to determine an under/over-shipment occurred.
	CountryOfOrigin             *string   `bun:"country_of_origin,type:varchar(8)"`                             // Country of origin
	SubPickTicketNo             *int32    `bun:"sub_pick_ticket_no,type:int"`                                   // Sub pick ticket no for this line. (Custom column)
	ShippingCountryCodeUid      *int32    `bun:"shipping_country_code_uid,type:int"`                            // Stores the country code value
	BoxLpn                      *string   `bun:"box_lpn,type:varchar(255)"`                                     // Box line package number for warehouse shipping
	PrintPartNo                 *string   `bun:"print_part_no,type:varchar(40)"`
	NewJobPriceShipControlNoUid *int32    `bun:"new_job_price_ship_control_no_uid,type:int"` // Custom F68631: Stores the contract ship control number which is new created when a pick ticket is generated for a contract item.
	ReturnedCylinderQuantity    *int32    `bun:"returned_cylinder_quantity,type:int"`        // The returned cylinder quantity associated with a trackabout empty cylinder
	RfnavPickStatusCd           *int32    `bun:"rfnav_pick_status_cd,type:int"`              // For the RF Navigator WMS integration: specifies the current pick status.  Valid values contained in code group 2299.
}
