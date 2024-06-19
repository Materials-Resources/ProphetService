package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InvAdjHdr struct {
	bun.BaseModel         `bun:"table:inv_adj_hdr"`
	AdjustmentNumber      float64   `bun:"adjustment_number,type:decimal(19,0),pk"`                   // System defined Id for the adjustment.
	CompanyId             string    `bun:"company_id,type:varchar(8)"`                                // Unique code that identifies a company.
	LocationId            float64   `bun:"location_id,type:decimal(19,0)"`                            // Location ID of the adjustment.
	ReasonId              string    `bun:"reason_id,type:varchar(5)"`                                 // Reason ID for the adjustment, decode from reason table.
	Period                float64   `bun:"period,type:decimal(3,0),nullzero"`                         // In which period did the inventory transfer occur?
	YearForPeriod         float64   `bun:"year_for_period,type:decimal(4,0),nullzero"`                // What year does the period belong to?
	Approved              string    `bun:"approved,type:char(1)"`                                     // This indicates whether or not the voucher is approved
	DeleteFlag            string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated           time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	InvAdjDescription     string    `bun:"inv_adj_description,type:varchar(255),nullzero"`            // Text description of the cause for the adjustment.
	ParentTransNo         int32     `bun:"parent_trans_no,type:int,nullzero"`                         // Source of the adjustment, decode from code_p21.
	ParentTransTypeCd     int32     `bun:"parent_trans_type_cd,type:int,nullzero"`                    // If a document is associated with an adjustment, the doc number may be stored here.
	RfLastLineNo          int32     `bun:"rf_last_line_no,type:int,nullzero"`                         // Corresponds to last line number counted on inv_adj_line (for saving place on RF Counting).
	RfCountInProcessFlag  string    `bun:"rf_count_in_process_flag,type:char(1),default:('N')"`       // Specifies if a RF user is currently counting this physical count.
	PaperlessCountFlag    string    `bun:"paperless_count_flag,type:char(1),default:('N')"`           // Denotes if adjustment is a paperless count
	SortOrder             string    `bun:"sort_order,type:char(1),nullzero"`                          // Determines whether user was grouping by Bin or Item in Physical Count.
	AffectActualUsageFlag string    `bun:"affect_actual_usage_flag,type:char(1),nullzero"`            // Affect Actual Usage Flag
	OverrideDecPrecFlag   string    `bun:"override_dec_prec_flag,type:char(1),default:('N')"`         // overrides decimal precision for quantities in line, lot, bin and lot/bin tabs
	ShowAllItemsFlag      string    `bun:"show_all_items_flag,type:char(1),default:('N')"`            // Show all items on this physical count even if they have 0 quantity on hand
	ShowExpectedQtyFlag   string    `bun:"show_expected_qty_flag,type:char(1),default:('N')"`         // Indicates whether to show the expected qty for an item being counted in WWMS.
	DisplayFoundItemsFlag string    `bun:"display_found_items_flag,type:char(1),default:('N')"`       // For adjustments associated with physical/cycle counts, determines if found items are included in the adjustment.
}
