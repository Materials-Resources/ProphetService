package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type BinType struct {
	bun.BaseModel    `bun:"table:bin_type"`
	BinTypeUid       int32     `bun:"bin_type_uid,type:int,pk"`                                     // Unique identifier for bin_type table
	BinType          string    `bun:"bin_type,type:varchar(255),unique"`                            // Name; specifies bin type
	BinTypeDesc      string    `bun:"bin_type_desc,type:varchar(255)"`                              // Description
	CompanyId        string    `bun:"company_id,type:varchar(8),unique"`                            // Company identifier
	PutableFlag      string    `bun:"putable_flag,type:char(1)"`                                    // Whether user can put quantity in bin
	PickableFlag     string    `bun:"pickable_flag,type:char(1)"`                                   // Whether user can pick quantity from bin for orders and transfers
	QuarantineFlag   string    `bun:"quarantine_flag,type:char(1)"`                                 // Whether bin is quarantined
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`                                     // Whether row is deleted or not
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	WeighStationFlag string    `bun:"weigh_station_flag,type:char(1),default:('N')"`                // Signifies if the bin is a weigh station (used to pick tags at).
	FrontCounterFlag *string   `bun:"front_counter_flag,type:char(1)"`                              // Indicate a bin as Front Counter bin type
	BackflushFlag    string    `bun:"backflush_flag,type:char(1),default:('N')"`                    // Indicate if the bin is a backflush type or not.
	FullSkidOnlyFlag *string   `bun:"full_skid_only_flag,type:char(1)"`                             // Indicates this type of bin contains only full skids
	ShippableFlag    *string   `bun:"shippable_flag,type:char(1)"`                                  // Indicates if user can take inventory from the bin to ship the orders.
}
