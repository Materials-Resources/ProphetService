package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ProcessPoShipmentHdr struct {
	bun.BaseModel           `bun:"table:process_po_shipment_hdr"`
	ProcessPoShipmentHdrUid int32     `bun:"process_po_shipment_hdr_uid,type:int,autoincrement,identity,pk"` // Unique identifier for this shipment record.
	PoNo                    float64   `bun:"po_no,type:decimal(19,0)"`                                       // The PO number for the Process PO this shipment record is associated with.
	ConfirmedFlag           string    `bun:"confirmed_flag,type:char(1),default:('N')"`                      // Indicates that the associated Process PO was shipped.
	ShipDate                time.Time `bun:"ship_date,type:datetime,nullzero"`                               // The date that the inventory for the associated Process PO was shipped.
	CarrierId               float64   `bun:"carrier_id,type:decimal(19,0),nullzero"`                         // The carrier for this shipment.
	TrackingNumber          string    `bun:"tracking_number,type:varchar(255),nullzero"`                     // The tracking number for this shipment.
	Instructions            string    `bun:"instructions,type:varchar(8000),nullzero"`                       // A note field for any shipping instructions related to this shipment
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`                 // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`           // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`           // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`   // User who last changed the record
}
