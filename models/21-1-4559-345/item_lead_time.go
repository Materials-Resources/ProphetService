package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ItemLeadTime struct {
	bun.BaseModel       `bun:"table:item_lead_time"`
	ItemLeadTimeKey     int32     `bun:"item_lead_time_key,type:int,pk"` // Unique Identifier for the record - primary key of table.
	LocationId          *float64  `bun:"location_id,type:decimal(19,0)"` // Location for which we are tracking lead time
	SupplierId          *float64  `bun:"supplier_id,type:decimal(19,0)"` // Supplier for which we are tracking lead time
	ReceiptNumber       *float64  `bun:"receipt_number,type:decimal(19,0)"`
	ReceiptDate         time.Time `bun:"receipt_date,type:datetime"`                                    // The approximate date a transfer will be received at the destination location.
	PoNo                *float64  `bun:"po_no,type:decimal(19,0)"`                                      // System-generated number assigned to a purchase order.
	PoDate              time.Time `bun:"po_date,type:datetime"`                                         // The date this purchase order was generated.
	ExcludeFromLeadTime string    `bun:"exclude_from_lead_time,type:char(1)"`                           // should we calc lead time for this combination of item/loc/supplier?
	ItemOrSupplier      string    `bun:"item_or_supplier,type:char(1)"`                                 // Does the lead time get calculated at the item or supplier level?
	DateCreated         time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	InvMastUid          *int32    `bun:"inv_mast_uid,type:int"`                                         // Unique identifier for the item id.
	ActualLeadDays      int16     `bun:"actual_lead_days,type:smallint,default:(0)"`                    // The number of days between the date the purchase order was generated and the day it was received.
	EstimatedLeadDays   int16     `bun:"estimated_lead_days,type:smallint,default:(0)"`                 // The average lead time for an item at the time the purchase order is generated.
	Edited              string    `bun:"edited,type:char(1),default:('N')"`                             // Indicates if the lead days value was manually edited.
}
