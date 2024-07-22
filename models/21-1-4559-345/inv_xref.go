package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvXref struct {
	bun.BaseModel             `bun:"table:inv_xref"`
	CompanyId                 string     `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	CustomerId                float64    `bun:"customer_id,type:decimal(19,0),pk"`                         // Customer paying invoice -  remitter
	TheirItemId               string     `bun:"their_item_id,type:varchar(40),pk"`                         // Cross refs with this item at the customer.
	DateCreated               time.Time  `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified          time.Time  `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy          string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	DeleteFlag                string     `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	InvMastUid                int32      `bun:"inv_mast_uid,type:int"`                                     // Unique identifier for the item id.
	DefaultSourceLocationId   *float64   `bun:"default_source_location_id,type:decimal(19,0)"`             // Default Source Location ID
	ConsignmentProcessing     int32      `bun:"consignment_processing,type:int,default:(2)"`               // Consignment Processing
	InvXrefUid                int32      `bun:"inv_xref_uid,type:int"`                                     // Unique identifier for customer part numbers
	CustomerAccountNumber     *string    `bun:"customer_account_number,type:varchar(16)"`
	BaselinePrice             *float64   `bun:"baseline_price,type:decimal(19,9)"`
	Requestor                 *string    `bun:"requestor,type:varchar(25)"`
	Department                *string    `bun:"department,type:varchar(16)"`
	Category                  *string    `bun:"category,type:varchar(16)"`
	ContractNumber            *string    `bun:"contract_number,type:varchar(16)"`
	Miscellaneous1            *string    `bun:"miscellaneous1,type:varchar(16)"`
	Miscellaneous2            *string    `bun:"miscellaneous2,type:varchar(16)"`
	Miscellaneous3            *string    `bun:"miscellaneous3,type:varchar(16)"`
	CreatedBy                 *string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	TheirBinId                *string    `bun:"their_bin_id,type:varchar(10)"`              // A column to hold the customer bin number
	VmiFlag                   *string    `bun:"vmi_flag,type:char(1)"`                      // If checked, Item is indicated to be processed via VMI for the customer
	LastReceiptDate           *time.Time `bun:"last_receipt_date,type:datetime"`            // Indicates the last date the item was included on an 852 sent from the customer.
	LastPoNo                  *string    `bun:"last_po_no,type:varchar(255)"`               // Indicates the Customer PO Number indicated on the last 852 sent that included the item.
	OnHandQty                 *float64   `bun:"on_hand_qty,type:decimal(19,9)"`             // Upon import, the system will calculate an On Hand Quantity from the individual data elements sent in the 852.
	MinQty                    *float64   `bun:"min_qty,type:decimal(19,9)"`                 // Indicates the Minimum (Order Point) quantity that is analyzed when generating a sales order from 852 requirements.
	MaxQty                    *float64   `bun:"max_qty,type:decimal(19,9)"`                 // Indicates the Maximum (Order Quantity) quantity that is used to determine an Order Quantity for the item if added to a sales order when generated from 852 requirements.
	MinUpdateOn852ImportFlag  *string    `bun:"min_update_on_852_import_flag,type:char(1)"` // If checked when an 852 is imported, the system will update the Minimum on the Customer Part Number record with the “Minimum” value from the import file.
	MaxUpdateOn852ImportFlag  *string    `bun:"max_update_on_852_import_flag,type:char(1)"` // If checked when an 852 is imported, the system will update the Maximum field on the Customer Part Number record with the “Maximum” value from the import file.
	DirectShipDispositionFlag *string    `bun:"direct_ship_disposition_flag,type:char(1)"`  // Indicates that the related Customer Part Number record is marked as a Direct Ship.
	LastPricePaid             *float64   `bun:"last_price_paid,type:decimal(19,9)"`         // User maintained value of the last price paid for this item
	ExtendedDescription       *string    `bun:"extended_description,type:varchar(255)"`     // Customer Part Number Extended Description
	ItemRevisionUid           *int32     `bun:"item_revision_uid,type:int"`                 // Column holds the item revision uid for the customer part number
	AvgUnitSellPrice          *float64   `bun:"avg_unit_sell_price,type:decimal(19,9)"`     // Average selling price for a customer part number in terms of the default sales pricing unit size.
	ItemLabelUom              *string    `bun:"item_label_uom,type:varchar(8)"`             // Custom (F45127): used in conjunction with column item_label_unit_size to determine the number of item labels to print
	ItemLabelUnitSize         *float64   `bun:"item_label_unit_size,type:decimal(19,4)"`    // Custom (F45127): used in conjunction with column item_label_uom to determine the number of item labels to print
	ItemLabelActionFlag       *string    `bun:"item_label_action_flag,type:char(1)"`        // Custom (F45127): determines how item labels will be handled in the warehouse.  Valid values are A (print and apply) and P (place in package).
	ItemLabelText             *string    `bun:"item_label_text,type:varchar(255)"`          // Custom (F45127): free form data that will be printed as a barcode on item labels.
	ItemLabelCalcTypeFlag     *string    `bun:"item_label_calc_type_flag,type:char(1)"`     // Custom (F45127): determines the method used to calculate the number of labels to print.  Valid values are U (calc using item_label_unit_size) and M (calc using item_label_multiplier).
	ItemLabelMultiplier       *float64   `bun:"item_label_multiplier,type:decimal(19,9)"`   // Custom (F45127): used in conjunction with the shipment qty to determine the number of item labels to print
}
