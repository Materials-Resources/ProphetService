package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Delivery struct {
	bun.BaseModel            `bun:"table:delivery"`
	DeliveryUid              int32     `bun:"delivery_uid,type:int,pk"`
	DeliveryNo               int32     `bun:"delivery_no,type:int"`
	BeginShipmentDate        time.Time `bun:"begin_shipment_date,type:datetime,nullzero"`
	EndShipmentDate          time.Time `bun:"end_shipment_date,type:datetime,nullzero"`
	DriverId                 string    `bun:"driver_id,type:varchar(16)"`
	DownloadStatus           int32     `bun:"download_status,type:int,nullzero"`
	RowStatusFlag            int16     `bun:"row_status_flag,type:smallint"`
	DeliveryListDate         time.Time `bun:"delivery_list_date,type:datetime"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30)"`
	ActiveFlag               string    `bun:"active_flag,type:char,default:('N')"`
	GpsData                  string    `bun:"gps_data,type:text(2147483647),nullzero"`
	DeliveryOutputCd         int32     `bun:"delivery_output_cd,type:int,default:((2845))"`
	DelvChargeInvCreatedFlag string    `bun:"delv_charge_inv_created_flag,type:char,default:('N')"`
}
