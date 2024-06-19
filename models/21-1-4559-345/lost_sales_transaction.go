package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type LostSalesTransaction struct {
	bun.BaseModel           `bun:"table:lost_sales_transaction"`
	LostSalesTransactionUid int32     `bun:"lost_sales_transaction_uid,type:int,autoincrement,identity,pk"` // Unique identifier for lost_sales_transaction.
	LostSalesUid            int32     `bun:"lost_sales_uid,type:int"`                                       // Idenitifies the reason chosen by the user to cancel the qty (from the lost_sales table).
	AffectUsage             string    `bun:"affect_usage,type:char(1)"`                                     // Whether or not the cancelled quantity associated with this change should affect the usage of the item.  This comes from the lost_sales table but is stored here also for historical data purposes.
	TransactionCodeNo       int32     `bun:"transaction_code_no,type:int"`                                  // The action/transaction type which led to the generation of this record.
	TransactionNo           int32     `bun:"transaction_no,type:int"`                                       // The transaction this reason is associated with.
	LineNo                  int32     `bun:"line_no,type:int,nullzero"`                                     // The line of the transaction this reason is associated with.
	SubLineNo               int32     `bun:"sub_line_no,type:int,nullzero"`                                 // Row for detail records of the main transaction row (assembly components, releases, etc).
	SkuQtyChange            float64   `bun:"sku_qty_change,type:decimal(19,9),nullzero"`                    // The change in qty this reason is associated with.
	UsageProcessedFlag      string    `bun:"usage_processed_flag,type:char(1)"`                             // Whether or not this cancelled qty has been applied to usage yet (irrespective of how affect_usage is set.)  Generally this will be N until the transaction is saved approved and then this will become Y.
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`                // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`          // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`          // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`  // User who last changed the record
}
