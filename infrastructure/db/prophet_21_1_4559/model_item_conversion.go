package prophet_21_1_4559

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type ItemConversion struct {
	bun.BaseModel          `bun:"table:item_conversion"`
	FromUom                string          `bun:"from_uom,type:varchar(8),pk"`
	ToUom                  string          `bun:"to_uom,type:varchar(8),pk"`
	Round                  string          `bun:"round,type:char"`
	Accumulate             string          `bun:"accumulate,type:char"`
	PrintOnPickTicket      string          `bun:"print_on_pick_ticket,type:char"`
	PrintOnOrderAck        string          `bun:"print_on_order_ack,type:char"`
	PrintOnInvoice         string          `bun:"print_on_invoice,type:char"`
	PrintOnPo              string          `bun:"print_on_po,type:char"`
	ConvertAtOe            string          `bun:"convert_at_oe,type:char"`
	ConvertAtPo            string          `bun:"convert_at_po,type:char"`
	DeleteFlag             string          `bun:"delete_flag,type:char"`
	DateCreated            time.Time       `bun:"date_created,type:datetime"`
	DateLastModified       time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy       string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	SortOrder              float64         `bun:"sort_order,type:decimal(19,0),pk"`
	InvMastUid             int32           `bun:"inv_mast_uid,type:int,pk"`
	PrintOnProdOrder       string          `bun:"print_on_prod_order,type:char,default:('N')"`
	ConvertAtTransferFlag  string          `bun:"convert_at_transfer_flag,type:char,default:('N')"`
	MinOrderQty            sql.NullFloat64 `bun:"min_order_qty,type:decimal(19,9),nullzero"`
	ConvertAtProdOrderFlag string          `bun:"convert_at_prod_order_flag,type:char,default:('N')"`
}
