package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type InvAdjHdr struct {
	bun.BaseModel         `bun:"table:inv_adj_hdr"`
	AdjustmentNumber      float64         `bun:"adjustment_number,type:decimal(19,0),pk"`
	CompanyId             string          `bun:"company_id,type:varchar(8)"`
	LocationId            float64         `bun:"location_id,type:decimal(19,0)"`
	ReasonId              string          `bun:"reason_id,type:varchar(5)"`
	Period                sql.NullFloat64 `bun:"period,type:decimal(3,0),nullzero"`
	YearForPeriod         sql.NullFloat64 `bun:"year_for_period,type:decimal(4,0),nullzero"`
	Approved              string          `bun:"approved,type:char"`
	DeleteFlag            string          `bun:"delete_flag,type:char"`
	DateCreated           time.Time       `bun:"date_created,type:datetime"`
	DateLastModified      time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	InvAdjDescription     sql.NullString  `bun:"inv_adj_description,type:varchar(255),nullzero"`
	ParentTransNo         sql.NullInt32   `bun:"parent_trans_no,type:int,nullzero"`
	ParentTransTypeCd     sql.NullInt32   `bun:"parent_trans_type_cd,type:int,nullzero"`
	RfLastLineNo          sql.NullInt32   `bun:"rf_last_line_no,type:int,nullzero"`
	RfCountInProcessFlag  string          `bun:"rf_count_in_process_flag,type:char,default:('N')"`
	PaperlessCountFlag    string          `bun:"paperless_count_flag,type:char,default:('N')"`
	SortOrder             sql.NullString  `bun:"sort_order,type:char,nullzero"`
	AffectActualUsageFlag sql.NullString  `bun:"affect_actual_usage_flag,type:char,nullzero"`
	OverrideDecPrecFlag   string          `bun:"override_dec_prec_flag,type:char,default:('N')"`
	ShowAllItemsFlag      string          `bun:"show_all_items_flag,type:char,default:('N')"`
	ShowExpectedQtyFlag   sql.NullString  `bun:"show_expected_qty_flag,type:char,default:('N')"`
	DisplayFoundItemsFlag string          `bun:"display_found_items_flag,type:char,default:('N')"`
}

type InvAdjHdrModel struct {
	bun bun.IDB
}

func (m InvAdjHdrModel) Get(ctx context.Context, adjustmentNumber float64) (*InvAdjHdr, error) {
	var invAdjHdr *InvAdjHdr
	err := m.bun.NewSelect().Model(&invAdjHdr).Where("adjustment_number = ?", adjustmentNumber).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return invAdjHdr, nil
}
