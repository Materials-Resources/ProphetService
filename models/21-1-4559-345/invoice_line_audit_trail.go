package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceLineAuditTrail struct {
	bun.BaseModel            `bun:"table:invoice_line_audit_trail"`
	InvoiceLineAuditTrailUid int32      `bun:"invoice_line_audit_trail_uid,type:int,autoincrement,identity"`
	DateLastModified         *time.Time `bun:"date_last_modified,type:datetime"`
	AuditDate                *time.Time `bun:"audit_date,type:datetime"`
	LastMaintainedBy         *string    `bun:"last_maintained_by,type:varchar(255)"`
	AuditUserName            *string    `bun:"audit_user_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	AuditLoginName           *string    `bun:"audit_login_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	InvoiceLineUid           int32      `bun:"invoice_line_uid,type:int"`
	OrderNoOld               *string    `bun:"order_no_old,type:varchar(255)"`
	OrderNoNew               *string    `bun:"order_no_new,type:varchar(255)"`
	OeLineNumberOld          *string    `bun:"oe_line_number_old,type:varchar(255)"`
	OeLineNumberNew          *string    `bun:"oe_line_number_new,type:varchar(255)"`
	InvoiceLineTypeOld       *string    `bun:"invoice_line_type_old,type:varchar(255)"`
	InvoiceLineTypeNew       *string    `bun:"invoice_line_type_new,type:varchar(255)"`
	TaxItemOld               *string    `bun:"tax_item_old,type:varchar(255)"`
	TaxItemNew               *string    `bun:"tax_item_new,type:varchar(255)"`
	CogsAmountOld            *string    `bun:"cogs_amount_old,type:varchar(255)"`
	CogsAmountNew            *string    `bun:"cogs_amount_new,type:varchar(255)"`
	ExtendedPriceOld         *string    `bun:"extended_price_old,type:varchar(255)"`
	ExtendedPriceNew         *string    `bun:"extended_price_new,type:varchar(255)"`
	ItemIdOld                *string    `bun:"item_id_old,type:varchar(255)"`
	ItemIdNew                *string    `bun:"item_id_new,type:varchar(255)"`
	UnitPriceOld             *string    `bun:"unit_price_old,type:varchar(255)"`
	UnitPriceNew             *string    `bun:"unit_price_new,type:varchar(255)"`
	InvoiceLineUidParentOld  *string    `bun:"invoice_line_uid_parent_old,type:varchar(255)"`
	InvoiceLineUidParentNew  *string    `bun:"invoice_line_uid_parent_new,type:varchar(255)"`
	OtherCostOld             *string    `bun:"other_cost_old,type:varchar(255)"`
	OtherCostNew             *string    `bun:"other_cost_new,type:varchar(255)"`
	QtyShippedOld            *string    `bun:"qty_shipped_old,type:varchar(255)"`
	QtyShippedNew            *string    `bun:"qty_shipped_new,type:varchar(255)"`
	CommissionCostOld        *string    `bun:"commission_cost_old,type:varchar(255)"`
	CommissionCostNew        *string    `bun:"commission_cost_new,type:varchar(255)"`
	CompanyIdOld             *string    `bun:"company_id_old,type:varchar(255)"`
	CompanyIdNew             *string    `bun:"company_id_new,type:varchar(255)"`
	ItemDescOld              *string    `bun:"item_desc_old,type:varchar(255)"`
	ItemDescNew              *string    `bun:"item_desc_new,type:varchar(255)"`
	GlRevenueAccountNoOld    *string    `bun:"gl_revenue_account_no_old,type:varchar(255)"`
	GlRevenueAccountNoNew    *string    `bun:"gl_revenue_account_no_new,type:varchar(255)"`
	JobIdOld                 *string    `bun:"job_id_old,type:varchar(255)"`
	JobIdNew                 *string    `bun:"job_id_new,type:varchar(255)"`
	QtyRequestedOld          *string    `bun:"qty_requested_old,type:varchar(255)"`
	QtyRequestedNew          *string    `bun:"qty_requested_new,type:varchar(255)"`
	UnitOfMeasureOld         *string    `bun:"unit_of_measure_old,type:varchar(255)"`
	UnitOfMeasureNew         *string    `bun:"unit_of_measure_new,type:varchar(255)"`
	InvoiceNo                *string    `bun:"invoice_no,type:varchar(255)"`
	LineNo                   *float64   `bun:"line_no,type:decimal(19,0)"`
	GlDimenTypeUidOld        *string    `bun:"gl_dimen_type_uid_old,type:varchar(255)"`
	GlDimenTypeUidNew        *string    `bun:"gl_dimen_type_uid_new,type:varchar(255)"`
	GlDimensionIdOld         *string    `bun:"gl_dimension_id_old,type:varchar(255)"`
	GlDimensionIdNew         *string    `bun:"gl_dimension_id_new,type:varchar(255)"`
}
