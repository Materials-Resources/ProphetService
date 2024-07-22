package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceHdrAuditTrail struct {
	bun.BaseModel            `bun:"table:invoice_hdr_audit_trail"`
	InvoiceHdrAuditTrailUid  int32      `bun:"invoice_hdr_audit_trail_uid,type:int,autoincrement,identity"`
	DateLastModified         *time.Time `bun:"date_last_modified,type:datetime"`
	AuditDate                *time.Time `bun:"audit_date,type:datetime"`
	LastMaintainedBy         *string    `bun:"last_maintained_by,type:varchar(255)"`
	AuditUserName            *string    `bun:"audit_user_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	AuditLoginName           *string    `bun:"audit_login_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	InvoiceNo                string     `bun:"invoice_no,type:varchar(10)"`
	AmountPaidOld            *string    `bun:"amount_paid_old,type:varchar(255)"`
	AmountPaidNew            *string    `bun:"amount_paid_new,type:varchar(255)"`
	ApprovedOld              *string    `bun:"approved_old,type:varchar(255)"`
	ApprovedNew              *string    `bun:"approved_new,type:varchar(255)"`
	BranchIdOld              *string    `bun:"branch_id_old,type:varchar(255)"`
	BranchIdNew              *string    `bun:"branch_id_new,type:varchar(255)"`
	CompanyNoOld             *string    `bun:"company_no_old,type:varchar(255)"`
	CompanyNoNew             *string    `bun:"company_no_new,type:varchar(255)"`
	ConsolidatedOld          *string    `bun:"consolidated_old,type:varchar(255)"`
	ConsolidatedNew          *string    `bun:"consolidated_new,type:varchar(255)"`
	CustomerIdOld            *string    `bun:"customer_id_old,type:varchar(255)"`
	CustomerIdNew            *string    `bun:"customer_id_new,type:varchar(255)"`
	DownpaymentAppliedOld    *string    `bun:"downpayment_applied_old,type:varchar(255)"`
	DownpaymentAppliedNew    *string    `bun:"downpayment_applied_new,type:varchar(255)"`
	FreightOld               *string    `bun:"freight_old,type:varchar(255)"`
	FreightNew               *string    `bun:"freight_new,type:varchar(255)"`
	FreightCodeUidOld        *string    `bun:"freight_code_uid_old,type:varchar(255)"`
	FreightCodeUidNew        *string    `bun:"freight_code_uid_new,type:varchar(255)"`
	InvoiceAdjustmentTypeOld *string    `bun:"invoice_adjustment_type_old,type:varchar(255)"`
	InvoiceAdjustmentTypeNew *string    `bun:"invoice_adjustment_type_new,type:varchar(255)"`
	InvoiceClassOld          *string    `bun:"invoice_class_old,type:varchar(255)"`
	InvoiceClassNew          *string    `bun:"invoice_class_new,type:varchar(255)"`
	InvoiceDateOld           *string    `bun:"invoice_date_old,type:varchar(255)"`
	InvoiceDateNew           *string    `bun:"invoice_date_new,type:varchar(255)"`
	OriginalDocumentTypeOld  *string    `bun:"original_document_type_old,type:varchar(255)"`
	OriginalDocumentTypeNew  *string    `bun:"original_document_type_new,type:varchar(255)"`
	OtherChargeAmountOld     *string    `bun:"other_charge_amount_old,type:varchar(255)"`
	OtherChargeAmountNew     *string    `bun:"other_charge_amount_new,type:varchar(255)"`
	ShippingCostOld          *string    `bun:"shipping_cost_old,type:varchar(255)"`
	ShippingCostNew          *string    `bun:"shipping_cost_new,type:varchar(255)"`
	TaxAmountOld             *string    `bun:"tax_amount_old,type:varchar(255)"`
	TaxAmountNew             *string    `bun:"tax_amount_new,type:varchar(255)"`
	TotalAmountOld           *string    `bun:"total_amount_old,type:varchar(255)"`
	TotalAmountNew           *string    `bun:"total_amount_new,type:varchar(255)"`
	ShipToIdOld              *string    `bun:"ship_to_id_old,type:varchar(255)"`
	ShipToIdNew              *string    `bun:"ship_to_id_new,type:varchar(255)"`
	InvoiceBatchUidOld       *string    `bun:"invoice_batch_uid_old,type:varchar(255)"`
	InvoiceBatchUidNew       *string    `bun:"invoice_batch_uid_new,type:varchar(255)"`
	ShippingRouteUidOld      *string    `bun:"shipping_route_uid_old,type:varchar(255)"`
	ShippingRouteUidNew      *string    `bun:"shipping_route_uid_new,type:varchar(255)"`
	PoNoOld                  *string    `bun:"po_no_old,type:varchar(255)"`
	PoNoNew                  *string    `bun:"po_no_new,type:varchar(255)"`
	PeriodOld                *string    `bun:"period_old,type:varchar(255)"`
	PeriodNew                *string    `bun:"period_new,type:varchar(255)"`
	YearForPeriodOld         *string    `bun:"year_for_period_old,type:varchar(255)"`
	YearForPeriodNew         *string    `bun:"year_for_period_new,type:varchar(255)"`
	BrokerageOld             *string    `bun:"brokerage_old,type:varchar(255)"`
	BrokerageNew             *string    `bun:"brokerage_new,type:varchar(255)"`
	TermsIdOld               *string    `bun:"terms_id_old,type:varchar(255)"`
	TermsIdNew               *string    `bun:"terms_id_new,type:varchar(255)"`
	CommissionCostOld        *string    `bun:"commission_cost_old,type:varchar(255)"`
	CommissionCostNew        *string    `bun:"commission_cost_new,type:varchar(255)"`
	InvoiceReferenceNoOld    *string    `bun:"invoice_reference_no_old,type:varchar(255)"`
	InvoiceReferenceNoNew    *string    `bun:"invoice_reference_no_new,type:varchar(255)"`
	InvoiceDescOld           *string    `bun:"invoice_desc_old,type:varchar(255)"`
	InvoiceDescNew           *string    `bun:"invoice_desc_new,type:varchar(255)"`
}
