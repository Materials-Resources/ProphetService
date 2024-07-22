package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Stop struct {
	bun.BaseModel      `bun:"table:stop"`
	StopUid            int32      `bun:"stop_uid,type:int,pk"`                 // Unique Identifier
	DeliveryUid        int32      `bun:"delivery_uid,type:int"`                // Pointer to the delviery that this stop is for.
	RecipientName      *string    `bun:"recipient_name,type:varchar(40)"`      // Name of the person that received the shipment.
	SignatureFilename  *string    `bun:"signature_filename,type:varchar(255)"` // Name of the bitmap file that contains the recipients signature.
	Notes              *string    `bun:"notes,type:varchar(255)"`              // Notes that the driver may have taken for the stop.
	BeginStopDate      *time.Time `bun:"begin_stop_date,type:datetime"`        // Date time that the driver signed on at the stop location.
	EndStopDate        *time.Time `bun:"end_stop_date,type:datetime"`          // Date time that the driver signed off at the stop location.
	DeliveryOrder      int32      `bun:"delivery_order,type:int"`              // Sequence number for the stop within the delivery.
	DeliveryStatus     int32      `bun:"delivery_status,type:int"`             // Flag indicating delivered or not delivered.
	DateCreated        time.Time  `bun:"date_created,type:datetime"`           // Indicates the date/time this record was created.
	DateLastModified   time.Time  `bun:"date_last_modified,type:datetime"`     // Indicates the date/time this record was last modified.
	LastMaintainedBy   string     `bun:"last_maintained_by,type:varchar(30)"`  // ID of the user who last maintained this record
	SignatureImageData *string    `bun:"signature_image_data,type:text"`       // Base64 data of the recipient signature.
	DeliveryPhotoData  *string    `bun:"delivery_photo_data,type:text"`        // Base64 data of the stop
	DownloadStatus     *int32     `bun:"download_status,type:int"`             // Status of the download.
}
