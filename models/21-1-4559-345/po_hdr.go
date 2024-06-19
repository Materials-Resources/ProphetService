package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PoHdr struct {
	bun.BaseModel                 `bun:"table:po_hdr"`
	PoNo                          float64   `bun:"po_no,type:decimal(19,0),pk"`                       // Purchase Order Number
	VendorId                      float64   `bun:"vendor_id,type:decimal(19,0)"`                      // What is the unique vendor identifier for this row?
	OrderDate                     time.Time `bun:"order_date,type:datetime,nullzero"`                 // When was this order taken?
	CompanyNo                     string    `bun:"company_no,type:varchar(8)"`                        // Unique code that identifies a company.
	Ship2Name                     string    `bun:"ship2_name,type:varchar(50),nullzero"`              // Ship To information located on ship to tab on PO Entry window, usually for Direct Ships. - Name
	PackingSlipNumber             string    `bun:"packing_slip_number,type:varchar(50),nullzero"`     // Ship To information located on ship to tab on PO Entry window, usually for Direct Ships. - Attention Line
	Ship2Add1                     string    `bun:"ship2_add1,type:varchar(50),nullzero"`              // Ship To information located on ship to tab on PO Entry window, usually for Direct Ships. - Address
	Ship2Add2                     string    `bun:"ship2_add2,type:varchar(50),nullzero"`              // Ship To information located on ship to tab on PO Entry window, usually for Direct Ships. - Address Line 2
	Ship2City                     string    `bun:"ship2_city,type:varchar(50),nullzero"`              // Ship To information located on ship to tab on PO Entry window, usually for Direct Ships. - City
	Ship2State                    string    `bun:"ship2_state,type:varchar(50),nullzero"`             // Ship To information located on ship to tab on PO Entry window, usually for Direct Ships. - State
	Ship2Country                  string    `bun:"ship2_country,type:varchar(50),nullzero"`           // Ship To information located on ship to tab on PO Entry window, usually for Direct Ships. - Country
	RequestedBy                   string    `bun:"requested_by,type:varchar(20),nullzero"`            // Buyer ID for this PO.
	Fob                           string    `bun:"fob,type:varchar(20),nullzero"`                     // Free On Board -- point in the delivery process at which freight costs and liability become the responsibility of the customer.
	DateDue                       time.Time `bun:"date_due,type:datetime,nullzero"`                   // Required date for the PO, the date the Items should be delivered by.
	ReceiptDate                   time.Time `bun:"receipt_date,type:datetime,nullzero"`               // Date of the last receipt of items on this PO.
	PoDesc                        string    `bun:"po_desc,type:varchar(255),nullzero"`                // Used for shipping instructions on the supplier tab.
	Ship2Zip                      string    `bun:"ship2_zip,type:varchar(10),nullzero"`               // Ship To information located on ship to tab on PO Entry window, usually for Direct Ships. - Zip
	DeleteFlag                    string    `bun:"delete_flag,type:char(1)"`                          // Indicates whether this record is logically deleted
	Complete                      string    `bun:"complete,type:char(1),nullzero"`                    // Indicates whether this PO is complete, do we expect any more shipments on this PO?
	DateCreated                   time.Time `bun:"date_created,type:datetime"`                        // Indicates the date/time this record was created.
	DateLastModified              time.Time `bun:"date_last_modified,type:datetime"`                  // Indicates the date/time this record was last modified.
	LastMaintainedBy              string    `bun:"last_maintained_by,type:varchar(30)"`               // ID of the user who last maintained this record
	LocationId                    float64   `bun:"location_id,type:decimal(19,0)"`                    // Where was the used material located?
	Terms                         string    `bun:"terms,type:varchar(20),nullzero"`                   // Terms ID – the terms you are normally granted when purchasing material from a particular vendor.
	CarrierId                     string    `bun:"carrier_id,type:varchar(8),nullzero"`               // What is the id of this carrier (if any)?
	VouchFlag                     string    `bun:"vouch_flag,type:char(1),nullzero"`                  // Indicates whether the PO has been vouched.
	Period                        float64   `bun:"period,type:decimal(3,0),nullzero"`                 // Which period does this PO apply to?
	YearForPeriod                 float64   `bun:"year_for_period,type:decimal(4,0),nullzero"`        // Which year does this PO apply to?
	CancelFlag                    string    `bun:"cancel_flag,type:char(1),nullzero"`                 // This column is unused.
	CurrencyId                    float64   `bun:"currency_id,type:decimal(19,0),nullzero"`           // What is the unique currency identifier for this PO
	Printed                       string    `bun:"printed,type:char(1),nullzero"`                     // Has this Po been printed?
	AckDate                       time.Time `bun:"ack_date,type:datetime,nullzero"`                   // This should be updated when an 855 *EDI Po Acknowledgement* comes in for this PO, currently not.
	AckCode                       string    `bun:"ack_code,type:varchar(2),nullzero"`                 // For EDI interfaces.
	ClosedFlag                    string    `bun:"closed_flag,type:char(1),nullzero"`                 // When this is "Y" the line item cannot be vouched
	SupplierId                    float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`           // What supplier supplies material for this PO?
	DivisionId                    float64   `bun:"division_id,type:decimal(19,0),nullzero"`           // What is the unique identifier for the division for the supplier?
	BranchId                      string    `bun:"branch_id,type:varchar(8)"`                         // What branch issued this Purchase Order?
	Approved                      string    `bun:"approved,type:char(1)"`                             // Has this Purchase order been approved?
	PurchaseGroupId               string    `bun:"purchase_group_id,type:varchar(8),nullzero"`        // Purchase/Transfer Group ID – a code that identifies a group of locations.  These locations are grouped together to make the processes of purchasing and transferring material more efficient.
	PoClass1                      string    `bun:"po_class1,type:varchar(255),nullzero"`              // What is the class for this purchase order?
	PoClass2                      string    `bun:"po_class2,type:varchar(255),nullzero"`              // What is the class for this purchase order?
	PoClass3                      string    `bun:"po_class3,type:varchar(255),nullzero"`              // What is the class for this purchase order?
	PoClass4                      string    `bun:"po_class4,type:varchar(255),nullzero"`              // What is the class for this purchase order?
	PoClass5                      string    `bun:"po_class5,type:varchar(255),nullzero"`              // What is the class for this purchase order?
	SalesOrderNumber              string    `bun:"sales_order_number,type:varchar(8),nullzero"`       // Maintains the sales order # when a purch order is created from a Sales Order.
	PoType                        string    `bun:"po_type,type:char(1)"`                              // Shows which requirement type the purchasing session should examine:  direct ship, regular stock, nonstock only, special order, or requisition.
	ExchangeRate                  float64   `bun:"exchange_rate,type:decimal(19,6),nullzero"`         // Exchange rate if foreign currency is being used.
	ExternalPoNo                  string    `bun:"external_po_no,type:varchar(40),nullzero"`          // The External PO Number is a PO number that was not generated by the system.
	ExcludeFromLeadTime           string    `bun:"exclude_from_lead_time,type:char(1),default:('N')"` // In some circumstances, a PO should not be a factor in the lead time calculation. This is the flag to indicate this.
	SourceType                    int32     `bun:"source_type,type:int,nullzero"`                     // How was the PO generated?  Order Entry, Import, EDI, Quote or PO Entry.
	TransmissionMethod            int32     `bun:"transmission_method,type:int,nullzero"`             // Indicates the EDI transmission method
	PoHdrUid                      int32     `bun:"po_hdr_uid,type:int,unique"`                        // Unique identifier of this PO.
	EdiEditCount                  int16     `bun:"edi_edit_count,type:smallint,default:(0)"`          // Number of times edi po has been edited
	RevisedPo                     string    `bun:"revised_po,type:char(1),nullzero"`                  // Indicates that a PO that had been previously printed has been changed and needs to be printed as a revised PO. This should not be consider a re-print since the PO changed. In a reprint the PO did not change.
	ContactId                     string    `bun:"contact_id,type:varchar(16),nullzero"`              // Contact associated with this purchase order
	CreatedBy                     string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	TotalIvaTaxAmt                float64   `bun:"total_iva_tax_amt,type:decimal(19,4),nullzero"`            // The column is used to display the total iva tax amount on the total tab in Purchase order Entry window
	ExpediteFlag                  string    `bun:"expedite_flag,type:char(1),nullzero"`                      // Indicates whether this PO should be included on the expedite request.
	VendorPaysFreightFlag         string    `bun:"vendor_pays_freight_flag,type:char(1),nullzero"`           // Used to indicate that the vendor will not charge the distributor any freight.
	RetrievedByWms                string    `bun:"retrieved_by_wms,type:char(1),default:('N')"`              // Determines whether or not this record has been exported.
	SupplierReleaseNo             string    `bun:"supplier_release_no,type:varchar(255),nullzero"`           // Custom: Release number for the PO on the supplier system
	HistoricalCompleteFlag        string    `bun:"historical_complete_flag,type:char(1),default:('N')"`      // Historical Complete Flag
	ExpectedDate                  time.Time `bun:"expected_date,type:datetime,nullzero"`                     // The date the items are expected to be received from the supplier.
	SentToCarrierDate             time.Time `bun:"sent_to_carrier_date,type:datetime,nullzero"`              // Date that this PO was exported to Carrier.
	PrintCanadianB3FormsFlag      string    `bun:"print_canadian_b3_forms_flag,type:char(1),default:('N')"`  // Determines if Canadian B3 customs forms should be avalable to print for this purchase order.
	CanadianB3FormsDate           time.Time `bun:"canadian_b3_forms_date,type:datetime,nullzero"`            // The date to employ when retrieving PO's in the Customs Forms Selection window.
	CadChaContractCd              int32     `bun:"cad_cha_contract_cd,type:int,nullzero"`                    // CAD/CHA Contract type
	CadChaQuoteNo                 string    `bun:"cad_cha_quote_no,type:varchar(255),nullzero"`              // CAD/CHA Quote number.
	TrackingNo                    string    `bun:"tracking_no,type:char(40),nullzero"`                       // Tracking No
	NetBillingFlag                string    `bun:"net_billing_flag,type:char(1),default:('N')"`              // Enable Net Billing custom functionality
	Ship2Add3                     string    `bun:"ship2_add3,type:varchar(50),nullzero"`                     // Ship to address line 3 for specials and direct ships
	DoNotExportCarrierPoFlag      string    `bun:"do_not_export_carrier_po_flag,type:char(1),default:('N')"` // Determines if a po can not be exported to Carrier.
	PackingBasis                  string    `bun:"packing_basis,type:varchar(255),nullzero"`                 // The packing basis for this purchase order, to be transmitted to the vendor.
	PackingBasisViolationCd       int32     `bun:"packing_basis_violation_cd,type:int,nullzero"`             // If the vendor ignores the PO Packing Basis, this column stores a code value that can be used to trigger an alert.
	TransmitPrint                 string    `bun:"transmit_print,type:char(1),nullzero"`                     // Set if print was selected as transmit method
	TransmitFax                   string    `bun:"transmit_fax,type:char(1),nullzero"`                       // Set if fax was selected as transmit method
	TransmitEmail                 string    `bun:"transmit_email,type:char(1),nullzero"`                     // Set if email was selected as transmit method
	TransmitEdi                   string    `bun:"transmit_edi,type:char(1),nullzero"`                       // Set if edi was selected as transmit method
	DfltListPriceMultiplier       float64   `bun:"dflt_list_price_multiplier,type:decimal(19,9),nullzero"`   // Default mutliplier for lines on the PO.
	ShipViaDesc                   string    `bun:"ship_via_desc,type:varchar(255),nullzero"`                 // Carrier Ship Via description. May store carrier account numbers.
	BlindShipFlag                 string    `bun:"blind_ship_flag,type:char(1),nullzero"`                    // Custom column to set an purchase order as a blind ship PO so that so that information can be added to the purchase order form / EDI export.
	PoQtyChangeCount              int32     `bun:"po_qty_change_count,type:int,nullzero"`                    // Count how many times qty on po has been changed.
	BulkDiscountFlag              string    `bun:"bulk_discount_flag,type:char(1),nullzero"`
	EstimatedValue                float64   `bun:"estimated_value,type:decimal(19,9),nullzero"`                    // Custom: Estimated value of this PO; Usually sum of line items, but can be edited.
	EstimatedValueEditFlag        string    `bun:"estimated_value_edit_flag,type:char(1),nullzero"`                // Custom: Indicates if the estimated value was edited by a user.
	BidNo                         string    `bun:"bid_no,type:varchar(255),nullzero"`                              // Custom: Tracks the bid/purchase order number.
	ExpectedShipDate              time.Time `bun:"expected_ship_date,type:datetime,nullzero"`                      // The date the items are expected to ship from the supplier
	AutoVouchExceptFlag           string    `bun:"auto_vouch_except_flag,type:char(1),default:('N')"`              // Determines if vendor invoices, associated with this PO, will be prevented from auto-vouching.
	SubmittedToVendorDate         time.Time `bun:"submitted_to_vendor_date,type:datetime,nullzero"`                // Custom column to record the date when PO has been submitted to vendor
	ExclSendLinkedInfoToRfnavFlag string    `bun:"excl_send_linked_info_to_rfnav_flag,type:char(1),default:('N')"` // Determines whether or not send linked po info to RF Navigator.
}
