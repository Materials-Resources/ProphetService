package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Stop struct {
	bun.BaseModel      `bun:"table:stop"`
	StopUid            int32     `bun:"stop_uid,type:int,pk"`
	DeliveryUid        int32     `bun:"delivery_uid,type:int"`
	RecipientName      string    `bun:"recipient_name,type:varchar(40),nullzero"`
	SignatureFilename  string    `bun:"signature_filename,type:varchar(255),nullzero"`
	Notes              string    `bun:"notes,type:varchar(255),nullzero"`
	BeginStopDate      time.Time `bun:"begin_stop_date,type:datetime,nullzero"`
	EndStopDate        time.Time `bun:"end_stop_date,type:datetime,nullzero"`
	DeliveryOrder      int32     `bun:"delivery_order,type:int"`
	DeliveryStatus     int32     `bun:"delivery_status,type:int"`
	DateCreated        time.Time `bun:"date_created,type:datetime"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30)"`
	SignatureImageData string    `bun:"signature_image_data,type:text(2147483647),nullzero"`
	DeliveryPhotoData  string    `bun:"delivery_photo_data,type:text(2147483647),nullzero"`
	DownloadStatus     int32     `bun:"download_status,type:int,nullzero"`
}
