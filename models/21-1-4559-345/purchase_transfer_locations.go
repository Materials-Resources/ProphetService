package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchaseTransferLocations struct {
	bun.BaseModel           `bun:"table:purchase_transfer_locations"`
	PurchaseTransferGroupId string    `bun:"purchase_transfer_group_id,type:varchar(8),pk"`                 // Unique ID for the Purchase Transfer Group (User defined)
	CompanyId               string    `bun:"company_id,type:varchar(8),pk"`                                 // Unique code that identifies a company.
	LocationId              float64   `bun:"location_id,type:decimal(19,0),pk"`                             // System generated identified for the location.
	DeleteFlag              string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	SequenceNumber          float64   `bun:"sequence_number,type:decimal(3,0),nullzero"`                    // The order of the locations as they were entered.  Used for GPOR group purchasing processing.
}
