package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvAccessory struct {
	bun.BaseModel         `bun:"table:inv_accessory"`
	InvAccessoryUid       int32     `bun:"inv_accessory_uid,type:int,pk"`                             // Unique identifier for the record.
	ParentInvMastUid      int32     `bun:"parent_inv_mast_uid,type:int"`                              // Unique Identifier for the parent item.
	ChildInvMastUid       int32     `bun:"child_inv_mast_uid,type:int"`                               // Unique Identifier for the child item.
	AutoPopulateQuantity  string    `bun:"auto_populate_quantity,type:char(1),default:('N')"`         // This checkbox is used to indicate whether the system should automatically fill in an Order Quantity for the Accessory Item when the Parent Item is ordered.
	ScaleToParentQuantity string    `bun:"scale_to_parent_quantity,type:char(1),default:('Y')"`       // This checkbox indicates if the recommended Order Quantity for this Accessory Item will be scaled at Order Entry time to correspond to the Order Quantity of the Parent Item.
	ChildUnitQuantity     float64   `bun:"child_unit_quantity,type:decimal(19,9),default:(0)"`        // nullThis is the base quantity for the child item that relates to the quantity and UOM fields for the accessory item.
	ChildUomCodeNo        int32     `bun:"child_uom_code_no,type:int"`                                // This is the unit of measure related to the child quantity column.
	ParentUnitQuantity    float64   `bun:"parent_unit_quantity,type:decimal(19,9),default:(0)"`       // This is the base quantity for the parent item that relates to the quantity and UOM fields for the accessory item.  This column defaults to 1.
	ParentUomCodeNo       int32     `bun:"parent_uom_code_no,type:int"`                               // This is the unit of measure related to the parent quantity column.
	DateCreated           time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	LinkQtyToParentFlag   string    `bun:"link_qty_to_parent_flag,type:char(1),default:('N')"`        // This column is set in Item Maintenace on the Accessory Items tab.  This column will be used to insure that the Accessory Item only is released to a Pick Ticket when the
	ShipTogetherFlag      string    `bun:"ship_together_flag,type:char(1),default:('N')"`             // Determines if the item's quantity should never be shipped if the parent is not also shipped
	MandatoryFlag         string    `bun:"mandatory_flag,type:char(1),default:('N')"`                 // Determines if item can be de-selected in the Accessory popup in OE. If no popup, item is added automatically
}
