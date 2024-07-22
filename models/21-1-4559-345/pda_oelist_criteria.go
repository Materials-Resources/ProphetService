package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PdaOelistCriteria struct {
	bun.BaseModel         `bun:"table:pda_oelist_criteria"`
	PdaOelistCriteriaUid  int32     `bun:"pda_oelist_criteria_uid,type:int,pk"`                // Unique ID. Primary Key.
	PdaOelistCriteriaId   string    `bun:"pda_oelist_criteria_id,type:varchar(255),unique"`    // User defined ID for the set of criteria
	PdaOelistCriteriaDesc string    `bun:"pda_oelist_criteria_desc,type:varchar(255)"`         // User defined description for the criteria.
	ItemOption            string    `bun:"item_option,type:char(1)"`                           // Retrieving all items or top x items.
	TopXSearchOption      *string   `bun:"top_x_search_option,type:char(1)"`                   // Retrieving top x items by pieces sold or order lines.
	TopXCount             *int32    `bun:"top_x_count,type:int"`                               // Number of items to return for top X.
	InvoiceDays           *int32    `bun:"invoice_days,type:int"`                              // How many days in the past to look for invoices.
	BegSupplierId         float64   `bun:"beg_supplier_id,type:decimal(19,0)"`                 // Beginning range to filter items by supplier.
	EndSupplierId         float64   `bun:"end_supplier_id,type:decimal(19,0)"`                 // End range to filter items by supplier.
	BegProductGroupId     string    `bun:"beg_product_group_id,type:varchar(8)"`               // Beginning range to filter items by product group.
	EndProductGroupId     string    `bun:"end_product_group_id,type:varchar(8)"`               // Ending range to filter items by product group
	BegItemId             string    `bun:"beg_item_id,type:varchar(40)"`                       // Beginning range to filter items by item id.
	EndItemId             string    `bun:"end_item_id,type:varchar(40)"`                       // Ending range to filter items by item id.
	RowStatusFlag         int32     `bun:"row_status_flag,type:int"`                           // Indicates current record status.
	DateCreated           time.Time `bun:"date_created,type:datetime"`                         // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`                   // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`                // ID of the user who last maintained this record
	UseAdditionalCriteria string    `bun:"use_additional_criteria,type:char(1),default:('N')"` // Whether to use the extra criteria for retrieving items
}
