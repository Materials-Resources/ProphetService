package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type SerialNumber struct {
	bun.BaseModel               `bun:"table:serial_number"`
	CompanyNo                   string     `bun:"company_no,type:varchar(8),pk"`                                 // Unique code that identifies a company.
	LocationId                  float64    `bun:"location_id,type:decimal(19,0),pk"`                             // Where was the used material located?
	SerialNumber                string     `bun:"serial_number,type:varchar(40),pk"`                             // A serial number that is allocated for the item.
	DeleteFlag                  string     `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated                 time.Time  `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified            time.Time  `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy            string     `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	DimensionTrackingKey        *int32     `bun:"dimension_tracking_key,type:int"`                               // What -  if any -  dimension is used to track this seri
	RowStatus                   int16      `bun:"row_status,type:tinyint"`                                       // Indicates current record status.
	InvMastUid                  int32      `bun:"inv_mast_uid,type:int,pk"`                                      // Unique identifier for the item id.
	SerialNumberUid             int32      `bun:"serial_number_uid,type:int,unique"`                             // Uniquer row ID for the table.
	CustomerId                  *float64   `bun:"customer_id,type:decimal(19,0)"`                                // track which customer owns this serial number
	BinId                       *string    `bun:"bin_id,type:varchar(10)"`                                       // The bin this serial number is stored in.
	LotUid                      *int32     `bun:"lot_uid,type:int"`                                              // lot uid when serial lot integration is used
	AllowanceAmount             *float64   `bun:"allowance_amount,type:decimal(19,4)"`                           // (Custom) Total Service Allowance Amount
	AllowanceAmountModifiedDate *time.Time `bun:"allowance_amount_modified_date,type:datetime"`                  // (Custom) Date that allowance_amount field was last modified
	TradeLayerUid               *int32     `bun:"trade_layer_uid,type:int"`                                      // Unique Key reference to trade_layer table
}
