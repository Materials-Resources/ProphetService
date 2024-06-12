package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ProcessPoShipmentHdr struct {
	bun.BaseModel           `bun:"table:process_po_shipment_hdr"`
	ProcessPoShipmentHdrUid int32     `bun:"process_po_shipment_hdr_uid,type:int,pk,identity"`
	PoNo                    float64   `bun:"po_no,type:decimal(19,0)"`
	ConfirmedFlag           string    `bun:"confirmed_flag,type:char,default:('N')"`
	ShipDate                time.Time `bun:"ship_date,type:datetime,nullzero"`
	CarrierId               float64   `bun:"carrier_id,type:decimal(19,0),nullzero"`
	TrackingNumber          string    `bun:"tracking_number,type:varchar(255),nullzero"`
	Instructions            string    `bun:"instructions,type:varchar(8000),nullzero"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
