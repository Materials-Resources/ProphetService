package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PoLine struct {
	bun.BaseModel            `bun:"table:po_line"`
	PoNo                     float64    `bun:"po_no,type:decimal(19,0),pk"`                                   // Purchase Order Number from po_hdr
	QtyOrdered               float64    `bun:"qty_ordered,type:decimal(19,9)"`                                // the quantity ordered in terms of SKUs.
	QtyReceived              float64    `bun:"qty_received,type:decimal(19,9)"`                               // Qty that has been received already.
	ReceivedDate             *time.Time `bun:"received_date,type:datetime"`                                   // Date the item was last received for this PO.
	UnitPrice                float64    `bun:"unit_price,type:decimal(19,6)"`                                 // What is the unit price for this line item?
	CompanyNo                string     `bun:"company_no,type:varchar(8)"`                                    // Unique code that identifies a company.
	MfgPartNo                *string    `bun:"mfg_part_no,type:varchar(20)"`                                  // Manufacturing Class ID for this item.
	DeleteFlag               string     `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateDue                  *time.Time `bun:"date_due,type:datetime"`                                        // When is this item due from the supplier?
	DateCreated              time.Time  `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified         time.Time  `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy         string     `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	NextDueInPoCost          *float64   `bun:"next_due_in_po_cost,type:decimal(19,9)"`                        // Item Cost of the next expected receipt of this item.
	Complete                 string     `bun:"complete,type:char(1)"`                                         // Has this po line been fully received?
	VouchCompleted           string     `bun:"vouch_completed,type:char(1)"`                                  // Has a voucher been completed?
	CancelFlag               *string    `bun:"cancel_flag,type:char(1)"`                                      // Was this PO line cancelled?
	InBoundCurryId           *float64   `bun:"in_bound_curry_id,type:decimal(19,0)"`                          // This goes with the in_bound_freight field.
	AccountNo                *string    `bun:"account_no,type:varchar(32)"`                                   // Enter the account number
	QtyToVouch               *float64   `bun:"qty_to_vouch,type:decimal(19,9)"`                               // Quantity that has been received and will be vouched.
	ClosedFlag               *string    `bun:"closed_flag,type:char(1)"`                                      // When this is "Y" the line item cannot be vouched
	ItemDescription          *string    `bun:"item_description,type:varchar(40)"`                             // What is the description of this item?
	UnitOfMeasure            *string    `bun:"unit_of_measure,type:varchar(8)"`                               // What is the unit of measure for this row?
	UnitSize                 float64    `bun:"unit_size,type:decimal(19,4)"`                                  // the quantity of the UOM.
	UnitQuantity             float64    `bun:"unit_quantity,type:decimal(19,9)"`                              // the quantity ordered in terms of UOMs.
	LineNo                   float64    `bun:"line_no,type:decimal(19,0),pk"`                                 // What line is this row?
	PricingBookId            *string    `bun:"pricing_book_id,type:varchar(8)"`                               // Unique identifier for pricing book used by this item.
	PricingBookItemId        *string    `bun:"pricing_book_item_id,type:varchar(40)"`                         // Unique identifier for item id pricing book used by this item.
	PricingBookSupplierId    *float64   `bun:"pricing_book_supplier_id,type:decimal(19,0)"`                   // Unique identifier for supplier id pricing book used by this item.
	PricingBookDiscGrpId     *string    `bun:"pricing_book_disc_grp_id,type:varchar(8)"`                      // Unique identifier for discount group id pricing book used by this item.
	PricingBookEffectiveDate *time.Time `bun:"pricing_book_effective_date,type:datetime"`                     // Date on which the pricing book takes effect.
	Combinable               *string    `bun:"combinable,type:char(1)"`                                       // Should this item be combined with like items for pricing purposes?
	CalcType                 *string    `bun:"calc_type,type:varchar(10)"`                                    // Pricing Calculation type (Multiplier, etc).
	CalcValue                *float64   `bun:"calc_value,type:decimal(19,9)"`                                 // Calculation Value in terms of calc_type.
	RequiredDate             *time.Time `bun:"required_date,type:datetime"`                                   // When is this order required by?
	NextBreak                *float64   `bun:"next_break,type:decimal(19,9)"`                                 // Quantity needed to achieve the next price break.
	NextUtPrice              *float64   `bun:"next_ut_price,type:decimal(19,9)"`                              // Unit price at the next price break.
	BaseUtPrice              float64    `bun:"base_ut_price,type:decimal(19,4),default:(0)"`                  // Either the source price or the actual price depending on the supplier’s pricing setup.
	PriceEdit                *string    `bun:"price_edit,type:char(1)"`                                       // Has the price been edited (Y), or was it defaulted from pricing service?
	NewItem                  *string    `bun:"new_item,type:char(1)"`                                         // Allows the purchase order to track items that have been added since original entry.
	QuantityChanged          *string    `bun:"quantity_changed,type:char(1)"`                                 // Allows the purchase order to track any items whose qty has been changed since original entry
	PricingUnit              *string    `bun:"pricing_unit,type:varchar(8)"`                                  // Maintains the pricing unit for the oe_line.
	PricingUnitSize          *float64   `bun:"pricing_unit_size,type:decimal(19,4)"`                          // Maintains the pricing unit size.
	ExtendedDesc             *string    `bun:"extended_desc,type:varchar(255)"`                               // Extended desc of the item from inv_mast
	UnitPriceDisplay         float64    `bun:"unit_price_display,type:decimal(19,6)"`                         // Holds the unit price that is displayed on the window
	InvMastUid               int32      `bun:"inv_mast_uid,type:int"`                                         // Unique identifier for the item id.
	ExcludeFromLeadTime      string     `bun:"exclude_from_lead_time,type:char(1),default:('N')"`             // Should this po line be excluded from the lead time calculation?
	SourceType               *int32     `bun:"source_type,type:int"`                                          // Was this PO entered from GPOR, PO Entry, Import, or OE?
	ExpDateUpdates           *int32     `bun:"exp_date_updates,type:int"`                                     // This column is unused.
	PoLineUid                int32      `bun:"po_line_uid,type:int,unique"`                                   // unique UID for this table.
	EdiNewStatus             string     `bun:"edi_new_status,type:char(1),default:('N')"`                     // Indicates this is a new edi PO
	LineType                 *string    `bun:"line_type,type:char(1)"`                                        // The po_line.line_type column is an override value for po_hdr.po_type.  In the code if po_line.line_type is NULL the po_hdr.po_type column will be used to deternine the value.  This column tells the source of the PO line and the values are: B - Backorder, D - Direct Ship, E Service PO, N - Non Stock, P - Special, Q - Vendor RFQ, R - Requisition, S - Stock, X - Secondary Processing
	ContractNumber           *string    `bun:"contract_number,type:varchar(40)"`                              // contract number assigned at po_line level
	CreatedBy                *string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	ParentPoLineNo           *int32     `bun:"parent_po_line_no,type:int"`                        // Associate po_line records
	SupplierShipDate         *time.Time `bun:"supplier_ship_date,type:datetime"`                  // Date supplier will ship the item
	EnteredAsCode            *string    `bun:"entered_as_code,type:varchar(40)"`                  // Stores the alternate code value used to enter the item.
	GporRunUid               *int32     `bun:"gpor_run_uid,type:int"`                             // GPOR requirement UID used to generate this po line
	PurchasePricingPageUid   *int32     `bun:"purchase_pricing_page_uid,type:int"`                // Unique UID for the purchase_pricing_page table
	ExpediteFlag             *string    `bun:"expedite_flag,type:char(1)"`                        // Indicates whether this PO line should be included on the expedite request.
	OriginalUnitPriceDisplay *float64   `bun:"original_unit_price_display,type:decimal(19,9)"`    // Custom feature 28086.  This column will hold the original unit price before it is modified by the user.
	RetrievedByWms           *string    `bun:"retrieved_by_wms,type:char(1),default:('N')"`       // column to indicate if the production order component was retrieved by wms
	ExpediteNotes            *string    `bun:"expedite_notes,type:varchar(8000)"`                 // Notes which help expedite POs.
	ExpediteFollowupFlag     string     `bun:"expedite_followup_flag,type:char(1),default:('N')"` // Provides an ability to help expedite a PO line by marking it for followup.
	DesiredReceiptLocationId *float64   `bun:"desired_receipt_location_id,type:decimal(19,0)"`    // The location ID at which receipt of the PO line item is desired.
	CountryOfOrigin          *string    `bun:"country_of_origin,type:char(8)"`                    // Indicates country of origin of the line item
	AcknowledgedDate         *time.Time `bun:"acknowledged_date,type:datetime"`                   // Date the supplier has acknowledged the fact that the PO needs to shipped to meet a required date
	B3Qty                    float64    `bun:"b3_qty,type:decimal(19,9),default:((0))"`           // Quantity to print on Canadian B3 customs forms.
	BulkBuyFlag              *string    `bun:"bulk_buy_flag,type:char(1)"`                        // Determines if this purchase order line will be for a bulk buy quantity.
	CadPurchaseCost          *float64   `bun:"cad_purchase_cost,type:decimal(19,9)"`              // CAD/CHA purchase cost
	QtyReady                 *float64   `bun:"qty_ready,type:decimal(19,9)"`                      // SKU quantity ready
	QtyReadyUnitSize         *float64   `bun:"qty_ready_unit_size,type:decimal(19,9)"`            // describes unit_qty_ready
	QtyReadyUom              *string    `bun:"qty_ready_uom,type:varchar(8)"`                     // describes unit_qty_ready
	UnitQtyReady             *float64   `bun:"unit_qty_ready,type:decimal(19,9)"`                 // quantity ready in entered unit terms
	ListPriceMultiplier      *float64   `bun:"list_price_multiplier,type:decimal(19,9)"`          // Stores the manual multiplier used for pricing based on list price.
	CarrierStatus            *string    `bun:"carrier_status,type:varchar(20)"`                   // Carrier status
	ExpectedShipDate         *time.Time `bun:"expected_ship_date,type:datetime"`                  // The date the items are expected to ship from the supplier
	DateDueLastModified      *time.Time `bun:"date_due_last_modified,type:datetime"`
	Acknowledged             *string    `bun:"acknowledged,type:char(1)"` // Indicate whether PO is acknowledged or not
}
