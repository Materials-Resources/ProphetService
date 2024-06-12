package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InvAdjHdr struct {
	bun.BaseModel         `bun:"table:inv_adj_hdr"`
	AdjustmentNumber      float64   `bun:"adjustment_number,type:decimal(19,0),pk"`
	CompanyId             string    `bun:"company_id,type:varchar(8)"`
	LocationId            float64   `bun:"location_id,type:decimal(19,0)"`
	ReasonId              string    `bun:"reason_id,type:varchar(5)"`
	Period                float64   `bun:"period,type:decimal(3,0),nullzero"`
	YearForPeriod         float64   `bun:"year_for_period,type:decimal(4,0),nullzero"`
	Approved              string    `bun:"approved,type:char"`
	DeleteFlag            string    `bun:"delete_flag,type:char"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	InvAdjDescription     string    `bun:"inv_adj_description,type:varchar(255),nullzero"`
	ParentTransNo         int32     `bun:"parent_trans_no,type:int,nullzero"`
	ParentTransTypeCd     int32     `bun:"parent_trans_type_cd,type:int,nullzero"`
	RfLastLineNo          int32     `bun:"rf_last_line_no,type:int,nullzero"`
	RfCountInProcessFlag  string    `bun:"rf_count_in_process_flag,type:char,default:('N')"`
	PaperlessCountFlag    string    `bun:"paperless_count_flag,type:char,default:('N')"`
	SortOrder             string    `bun:"sort_order,type:char,nullzero"`
	AffectActualUsageFlag string    `bun:"affect_actual_usage_flag,type:char,nullzero"`
	OverrideDecPrecFlag   string    `bun:"override_dec_prec_flag,type:char,default:('N')"`
	ShowAllItemsFlag      string    `bun:"show_all_items_flag,type:char,default:('N')"`
	ShowExpectedQtyFlag   string    `bun:"show_expected_qty_flag,type:char,default:('N')"`
	DisplayFoundItemsFlag string    `bun:"display_found_items_flag,type:char,default:('N')"`
}
