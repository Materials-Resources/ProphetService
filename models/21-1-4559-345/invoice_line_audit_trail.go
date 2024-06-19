package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceLineAuditTrail struct {
	bun.BaseModel            `bun:"table:invoice_line_audit_trail"`
	InvoiceLineAuditTrailUid int32     `bun:"invoice_line_audit_trail_uid,type:int,autoincrement,identity"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,nullzero"`
	AuditDate                time.Time `bun:"audit_date,type:datetime,nullzero"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),nullzero"`
	AuditUserName            string    `bun:"audit_user_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`
	AuditLoginName           string    `bun:"audit_login_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`
	InvoiceLineUid           int32     `bun:"invoice_line_uid,type:int"`
	OrderNoOld               string    `bun:"order_no_old,type:varchar(255),nullzero"`
	OrderNoNew               string    `bun:"order_no_new,type:varchar(255),nullzero"`
	OeLineNumberOld          string    `bun:"oe_line_number_old,type:varchar(255),nullzero"`
	OeLineNumberNew          string    `bun:"oe_line_number_new,type:varchar(255),nullzero"`
	InvoiceLineTypeOld       string    `bun:"invoice_line_type_old,type:varchar(255),nullzero"`
	InvoiceLineTypeNew       string    `bun:"invoice_line_type_new,type:varchar(255),nullzero"`
	TaxItemOld               string    `bun:"tax_item_old,type:varchar(255),nullzero"`
	TaxItemNew               string    `bun:"tax_item_new,type:varchar(255),nullzero"`
	CogsAmountOld            string    `bun:"cogs_amount_old,type:varchar(255),nullzero"`
	CogsAmountNew            string    `bun:"cogs_amount_new,type:varchar(255),nullzero"`
	ExtendedPriceOld         string    `bun:"extended_price_old,type:varchar(255),nullzero"`
	ExtendedPriceNew         string    `bun:"extended_price_new,type:varchar(255),nullzero"`
	ItemIdOld                string    `bun:"item_id_old,type:varchar(255),nullzero"`
	ItemIdNew                string    `bun:"item_id_new,type:varchar(255),nullzero"`
	UnitPriceOld             string    `bun:"unit_price_old,type:varchar(255),nullzero"`
	UnitPriceNew             string    `bun:"unit_price_new,type:varchar(255),nullzero"`
	InvoiceLineUidParentOld  string    `bun:"invoice_line_uid_parent_old,type:varchar(255),nullzero"`
	InvoiceLineUidParentNew  string    `bun:"invoice_line_uid_parent_new,type:varchar(255),nullzero"`
	OtherCostOld             string    `bun:"other_cost_old,type:varchar(255),nullzero"`
	OtherCostNew             string    `bun:"other_cost_new,type:varchar(255),nullzero"`
	QtyShippedOld            string    `bun:"qty_shipped_old,type:varchar(255),nullzero"`
	QtyShippedNew            string    `bun:"qty_shipped_new,type:varchar(255),nullzero"`
	CommissionCostOld        string    `bun:"commission_cost_old,type:varchar(255),nullzero"`
	CommissionCostNew        string    `bun:"commission_cost_new,type:varchar(255),nullzero"`
	CompanyIdOld             string    `bun:"company_id_old,type:varchar(255),nullzero"`
	CompanyIdNew             string    `bun:"company_id_new,type:varchar(255),nullzero"`
	ItemDescOld              string    `bun:"item_desc_old,type:varchar(255),nullzero"`
	ItemDescNew              string    `bun:"item_desc_new,type:varchar(255),nullzero"`
	GlRevenueAccountNoOld    string    `bun:"gl_revenue_account_no_old,type:varchar(255),nullzero"`
	GlRevenueAccountNoNew    string    `bun:"gl_revenue_account_no_new,type:varchar(255),nullzero"`
	JobIdOld                 string    `bun:"job_id_old,type:varchar(255),nullzero"`
	JobIdNew                 string    `bun:"job_id_new,type:varchar(255),nullzero"`
	QtyRequestedOld          string    `bun:"qty_requested_old,type:varchar(255),nullzero"`
	QtyRequestedNew          string    `bun:"qty_requested_new,type:varchar(255),nullzero"`
	UnitOfMeasureOld         string    `bun:"unit_of_measure_old,type:varchar(255),nullzero"`
	UnitOfMeasureNew         string    `bun:"unit_of_measure_new,type:varchar(255),nullzero"`
	InvoiceNo                string    `bun:"invoice_no,type:varchar(255),nullzero"`
	LineNo                   float64   `bun:"line_no,type:decimal(19,0),nullzero"`
	GlDimenTypeUidOld        string    `bun:"gl_dimen_type_uid_old,type:varchar(255),nullzero"`
	GlDimenTypeUidNew        string    `bun:"gl_dimen_type_uid_new,type:varchar(255),nullzero"`
	GlDimensionIdOld         string    `bun:"gl_dimension_id_old,type:varchar(255),nullzero"`
	GlDimensionIdNew         string    `bun:"gl_dimension_id_new,type:varchar(255),nullzero"`
}
