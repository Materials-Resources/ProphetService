package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type GporDynamicLookAhead struct {
	bun.BaseModel        `bun:"table:gpor_dynamic_look_ahead"`
	GporRunHdrUid        int32     `bun:"gpor_run_hdr_uid,type:int,unique"`                             // Unique ID for the gpor_run_hdr table
	InvMastUid           int32     `bun:"inv_mast_uid,type:int,unique"`                                 // Unique ID for the item
	LocationId           float64   `bun:"location_id,type:decimal(19,0),unique"`                        // Unique ID for the location
	DynamicLookAheadDate time.Time `bun:"dynamic_look_ahead_date,type:datetime"`                        // Calculated dynamic look ahead date for the gpor run, item, location
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	DrpItemFlag          string    `bun:"drp_item_flag,type:char(1),default:('N')"`                     // Indicates if DRP forecasting/purchasing is performed for item
	TransactionTypeCd    int32     `bun:"transaction_type_cd,type:int,default:((1322)),unique"`
}
