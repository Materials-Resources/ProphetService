package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Delivery struct {
	bun.BaseModel            `bun:"table:delivery"`
	DeliveryUid              int32     `bun:"delivery_uid,type:int,pk"`                                // Unique Identifier
	DeliveryNo               int32     `bun:"delivery_no,type:int,unique"`                             // Delivery number associated with the set of pick tickets that will be delivered.
	BeginShipmentDate        time.Time `bun:"begin_shipment_date,type:datetime,nullzero"`              // Date Time that the delivery started. Time that the driver clocked in at the first stop.
	EndShipmentDate          time.Time `bun:"end_shipment_date,type:datetime,nullzero"`                // Date Time that the delivery ended. Time that the driver clocked out at the last stop.
	DriverId                 string    `bun:"driver_id,type:varchar(16)"`                              // ID of the driver, from the contacts table.
	DownloadStatus           int32     `bun:"download_status,type:int,nullzero"`                       // Not downloaded, downloaded or uploaded.
	RowStatusFlag            int16     `bun:"row_status_flag,type:smallint"`                           // Indicates current record status.
	DeliveryListDate         time.Time `bun:"delivery_list_date,type:datetime"`                        // Date the the list was last edited.
	DateCreated              time.Time `bun:"date_created,type:datetime"`                              // Indicates the date/time this record was created.
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`                        // Indicates the date/time this record was last modified.
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30)"`                     // ID of the user who last maintained this record
	ActiveFlag               string    `bun:"active_flag,type:char(1),default:('N')"`                  // User by OTA POD, which list is active for a driver for the OTA device to pull.
	GpsData                  string    `bun:"gps_data,type:text,nullzero"`                             // GPS Tracking data of the delivery
	DeliveryOutputCd         int32     `bun:"delivery_output_cd,type:int,default:((2845))"`            // The type of device the output data was created for.
	DelvChargeInvCreatedFlag string    `bun:"delv_charge_inv_created_flag,type:char(1),default:('N')"` // Determines if delivery charge invoices have been created for this list.
}
