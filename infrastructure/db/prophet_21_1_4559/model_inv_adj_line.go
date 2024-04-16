package prophet_21_1_4559

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type InvAdjLine struct {
	bun.BaseModel          `bun:"table:inv_adj_line"`
	AdjustmentNumber       float64         `bun:"adjustment_number,type:decimal(19,0),pk"`
	LineNumber             int32           `bun:"line_number,type:int,pk"`
	UnitQuantity           float64         `bun:"unit_quantity,type:decimal(19,9)"`
	Cost                   float64         `bun:"cost,type:decimal(19,9)"`
	UnitOfMeasure          string          `bun:"unit_of_measure,type:varchar(8)"`
	DateCreated            time.Time       `bun:"date_created,type:datetime"`
	DateLastModified       time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy       string          `bun:"last_maintained_by,type:varchar(30)"`
	Quantity               float64         `bun:"quantity,type:decimal(19,9)"`
	InvMastUid             int32           `bun:"inv_mast_uid,type:int"`
	RowStatusFlag          int32           `bun:"row_status_flag,type:int,default:(701)"`
	QtyCountedFlag         sql.NullString  `bun:"qty_counted_flag,type:char,nullzero"`
	VendorId               sql.NullFloat64 `bun:"vendor_id,type:decimal(19,0),nullzero"`
	DeallocateQtyFlag      string          `bun:"deallocate_qty_flag,type:char,default:('N')"`
	WwmsDeallocateFailFlag string          `bun:"wwms_deallocate_fail_flag,type:char,default:('N')"`
	ParentTransLineNo      sql.NullInt32   `bun:"parent_trans_line_no,type:int,nullzero"`
	ActualPhysicalCountQty sql.NullFloat64 `bun:"actual_physical_count_qty,type:decimal(19,9),default:((0))"`
	RecountFlag            sql.NullString  `bun:"recount_flag,type:char,default:('N')"`
	TimesRecounted         sql.NullInt32   `bun:"times_recounted,type:int,default:((0))"`
	ReasonId               sql.NullString  `bun:"reason_id,type:varchar(5),nullzero"`
	Comment                sql.NullString  `bun:"comment,type:varchar(255),nullzero"`
	AdjustFoundItemFlag    string          `bun:"adjust_found_item_flag,type:char,default:('N')"`
}
