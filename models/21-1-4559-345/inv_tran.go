package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvTran struct {
	bun.BaseModel          `bun:"table:inv_tran"`
	TransactionNumber      float64   `bun:"transaction_number,type:decimal(19,0),pk"`                  // Unique indentifier for the inv_tran record.
	LocationId             float64   `bun:"location_id,type:decimal(19,0)"`                            // Where was the used material located?
	Period                 *float64  `bun:"period,type:decimal(3,0)"`                                  // Fiscal period associated with the change to quantity on-hand.
	YearForPeriod          *float64  `bun:"year_for_period,type:decimal(4,0)"`                         // Fiscal year associated with the change to quantity on-hand.
	TransType              string    `bun:"trans_type,type:varchar(5)"`                                // Type of transaction. ie. order, receipt, etc.
	Quantity               float64   `bun:"quantity,type:decimal(19,9)"`                               // Change to the quantity on-hand.
	UnitCostAmt            float64   `bun:"unit_cost_amt,type:decimal(19,9)"`                          // Cost if quantity on-hand is being added or relieved.
	Freight                float64   `bun:"freight,type:decimal(19,4)"`                                // Prorated freight amount for this item.
	ReservedBeforeTrans    *float64  `bun:"reserved_before_trans,type:decimal(19,9),default:(0)"`      // Qty reserved/due-in before the transaction.
	DocumentNo             float64   `bun:"document_no,type:decimal(19,0)"`                            // References the originating transaction. ie. order no, receipt no, etc.
	OnHandBeforeTrans      float64   `bun:"on_hand_before_trans,type:decimal(19,9)"`                   // Quantity on-hand before the transaction.
	AllocatedBeforeTrans   float64   `bun:"allocated_before_trans,type:decimal(19,9)"`                 // Quantity allocated before the transaction.
	BackorderedBeforeTrans float64   `bun:"backordered_before_trans,type:decimal(19,9)"`               // Quantity backordered before the transaction.
	OnPoBeforeTrans        float64   `bun:"on_po_before_trans,type:decimal(19,9)"`                     // Quantity on-po before the transaction.
	InTransitBeforeTrans   float64   `bun:"in_transit_before_trans,type:decimal(19,9)"`                // Quantity in-transit before the transaction.
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	CurrencyId             *float64  `bun:"currency_id,type:decimal(19,0)"`                            // What is the unique currency identifier for this ro
	SubDocumentNo          *float64  `bun:"sub_document_no,type:decimal(19,0)"`                        // Used on a per transaction basis to provide additional transaction information.
	QtyAllocated           float64   `bun:"qty_allocated,type:decimal(19,9)"`                          // Change to quantity allocated.
	QtyOnBo                float64   `bun:"qty_on_bo,type:decimal(19,9)"`                              // Change to quantity backordered.
	QtyOnPo                float64   `bun:"qty_on_po,type:decimal(19,9)"`                              // Change to quantity on-po.
	QtyInTransit           float64   `bun:"qty_in_transit,type:decimal(19,9)"`                         // Change to quantity in-transit.
	LineNo                 *float64  `bun:"line_no,type:decimal(19,0)"`                                // Line number associated with the originating transaction. ie. order line no
	UnitOfMeasure          string    `bun:"unit_of_measure,type:varchar(8)"`                           // What is the unit of measure for this row? - DO NOT USE
	UnitSize               float64   `bun:"unit_size,type:decimal(19,4)"`                              // Unit size associated with the UOM. - DO NOT USE
	QtyReservedDueIn       *float64  `bun:"qty_reserved_due_in,type:decimal(19,9),default:(0)"`        // Change to quantity reserved/due-in.
	InvMastUid             int32     `bun:"inv_mast_uid,type:int"`                                     // Unique identifier for the item id.
	InProcessBeforeTrans   float64   `bun:"in_process_before_trans,type:decimal(19,9),default:(0)"`    // store inv_loc.qty_in_process before this inv_tran record
	QtyInProcess           float64   `bun:"qty_in_process,type:decimal(19,9),default:(0)"`             // store secondary process finished item how much need to make
	UsedSpecificCostFlag   *string   `bun:"used_specific_cost_flag,type:char(1)"`                      // column used to specify if special costing was used or not.
}
