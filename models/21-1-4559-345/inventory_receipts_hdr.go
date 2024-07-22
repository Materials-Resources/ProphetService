package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InventoryReceiptsHdr struct {
	bun.BaseModel           `bun:"table:inventory_receipts_hdr"`
	ReceiptNumber           float64    `bun:"receipt_number,type:decimal(19,0),pk"`
	PoNumber                float64    `bun:"po_number,type:decimal(19,0)"`                              // What is the purchase order that pertains to this inventory reciept?
	Approved                string     `bun:"approved,type:char(1)"`                                     // This indicates whether or not the voucher is approved
	DeleteFlag              string     `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated             time.Time  `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time  `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	VouchComplete           string     `bun:"vouch_complete,type:char(1)"`                               // Set to Y when the receipt has been completed converted to voucher.
	Period                  float64    `bun:"period,type:decimal(3,0)"`                                  // In which period did the inventory transfer occur?
	YearForPeriod           float64    `bun:"year_for_period,type:decimal(4,0)"`                         // What year does the period belong to?
	FreightAmount           *float64   `bun:"freight_amount,type:decimal(19,4)"`                         // The freight amount of the receipt.
	FreightAmountDisplay    *float64   `bun:"freight_amount_display,type:decimal(19,4)"`                 // The freight amount of the receipt in the appropriate currency.
	CurrencyId              float64    `bun:"currency_id,type:decimal(19,0)"`                            // Currency associated with the receipt.
	ReceiptType             string     `bun:"receipt_type,type:char(1)"`                                 // Indicates type of receipt (ie po receipt, transfer receipt etc)
	PackingSlipNumber       *string    `bun:"packing_slip_number,type:varchar(255)"`                     // Vendors packing slip number
	FreightCodeUid          *int32     `bun:"freight_code_uid,type:int"`                                 // Unique identifier of the Freight Code.
	ReceiverNumber          *string    `bun:"receiver_number,type:varchar(255)"`                         // Used for FASCOR integration
	ExternalReferenceNo     *string    `bun:"external_reference_no,type:varchar(255)"`                   // Vessel ID Number if material is been received by container.
	PoAsnHdrUid             *int32     `bun:"po_asn_hdr_uid,type:int"`                                   // Indicates that the receipt is created based on what ASN record
	ContainerReceiptsHdrUid *int32     `bun:"container_receipts_hdr_uid,type:int"`                       // Reference to the container receipt record that created this PO receipt record
	ShipmentDate            *time.Time `bun:"shipment_date,type:datetime"`                               // Date in which the supplier shipped the material
	SourceTypeCd            *int32     `bun:"source_type_cd,type:int"`                                   // Indicates where this receipt was created from (currently only populated for custom features)
	InvReceiptsClearingAcct *string    `bun:"inv_receipts_clearing_acct,type:varchar(32)"`               // The account that stores the inventory value of goods that have been received against a PO while you are still waiting for the invoice for those goods
	CreatedBy               string     `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	GroupPoHdrUid           *int32     `bun:"group_po_hdr_uid,type:int"` // Group po no for the po that this receipt is created under.
	RfnavTransNo            *int32     `bun:"rfnav_trans_no,type:int"`   // F73204: RF Navigator transaction number
}
