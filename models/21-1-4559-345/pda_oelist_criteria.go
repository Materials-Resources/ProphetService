package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PdaOelistCriteria struct {
	bun.BaseModel         `bun:"table:pda_oelist_criteria"`
	PdaOelistCriteriaUid  int32     `bun:"pda_oelist_criteria_uid,type:int,pk"`
	PdaOelistCriteriaId   string    `bun:"pda_oelist_criteria_id,type:varchar(255)"`
	PdaOelistCriteriaDesc string    `bun:"pda_oelist_criteria_desc,type:varchar(255)"`
	ItemOption            string    `bun:"item_option,type:char"`
	TopXSearchOption      string    `bun:"top_x_search_option,type:char,nullzero"`
	TopXCount             int32     `bun:"top_x_count,type:int,nullzero"`
	InvoiceDays           int32     `bun:"invoice_days,type:int,nullzero"`
	BegSupplierId         float64   `bun:"beg_supplier_id,type:decimal(19,0)"`
	EndSupplierId         float64   `bun:"end_supplier_id,type:decimal(19,0)"`
	BegProductGroupId     string    `bun:"beg_product_group_id,type:varchar(8)"`
	EndProductGroupId     string    `bun:"end_product_group_id,type:varchar(8)"`
	BegItemId             string    `bun:"beg_item_id,type:varchar(40)"`
	EndItemId             string    `bun:"end_item_id,type:varchar(40)"`
	RowStatusFlag         int32     `bun:"row_status_flag,type:int"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`
	UseAdditionalCriteria string    `bun:"use_additional_criteria,type:char,default:('N')"`
}
