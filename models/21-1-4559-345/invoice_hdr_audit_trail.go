package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceHdrAuditTrail struct {
	bun.BaseModel            `bun:"table:invoice_hdr_audit_trail"`
	InvoiceHdrAuditTrailUid  int32     `bun:"invoice_hdr_audit_trail_uid,type:int,autoincrement,identity"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,nullzero"`
	AuditDate                time.Time `bun:"audit_date,type:datetime,nullzero"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),nullzero"`
	AuditUserName            string    `bun:"audit_user_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`
	AuditLoginName           string    `bun:"audit_login_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`
	InvoiceNo                string    `bun:"invoice_no,type:varchar(10)"`
	AmountPaidOld            string    `bun:"amount_paid_old,type:varchar(255),nullzero"`
	AmountPaidNew            string    `bun:"amount_paid_new,type:varchar(255),nullzero"`
	ApprovedOld              string    `bun:"approved_old,type:varchar(255),nullzero"`
	ApprovedNew              string    `bun:"approved_new,type:varchar(255),nullzero"`
	BranchIdOld              string    `bun:"branch_id_old,type:varchar(255),nullzero"`
	BranchIdNew              string    `bun:"branch_id_new,type:varchar(255),nullzero"`
	CompanyNoOld             string    `bun:"company_no_old,type:varchar(255),nullzero"`
	CompanyNoNew             string    `bun:"company_no_new,type:varchar(255),nullzero"`
	ConsolidatedOld          string    `bun:"consolidated_old,type:varchar(255),nullzero"`
	ConsolidatedNew          string    `bun:"consolidated_new,type:varchar(255),nullzero"`
	CustomerIdOld            string    `bun:"customer_id_old,type:varchar(255),nullzero"`
	CustomerIdNew            string    `bun:"customer_id_new,type:varchar(255),nullzero"`
	DownpaymentAppliedOld    string    `bun:"downpayment_applied_old,type:varchar(255),nullzero"`
	DownpaymentAppliedNew    string    `bun:"downpayment_applied_new,type:varchar(255),nullzero"`
	FreightOld               string    `bun:"freight_old,type:varchar(255),nullzero"`
	FreightNew               string    `bun:"freight_new,type:varchar(255),nullzero"`
	FreightCodeUidOld        string    `bun:"freight_code_uid_old,type:varchar(255),nullzero"`
	FreightCodeUidNew        string    `bun:"freight_code_uid_new,type:varchar(255),nullzero"`
	InvoiceAdjustmentTypeOld string    `bun:"invoice_adjustment_type_old,type:varchar(255),nullzero"`
	InvoiceAdjustmentTypeNew string    `bun:"invoice_adjustment_type_new,type:varchar(255),nullzero"`
	InvoiceClassOld          string    `bun:"invoice_class_old,type:varchar(255),nullzero"`
	InvoiceClassNew          string    `bun:"invoice_class_new,type:varchar(255),nullzero"`
	InvoiceDateOld           string    `bun:"invoice_date_old,type:varchar(255),nullzero"`
	InvoiceDateNew           string    `bun:"invoice_date_new,type:varchar(255),nullzero"`
	OriginalDocumentTypeOld  string    `bun:"original_document_type_old,type:varchar(255),nullzero"`
	OriginalDocumentTypeNew  string    `bun:"original_document_type_new,type:varchar(255),nullzero"`
	OtherChargeAmountOld     string    `bun:"other_charge_amount_old,type:varchar(255),nullzero"`
	OtherChargeAmountNew     string    `bun:"other_charge_amount_new,type:varchar(255),nullzero"`
	ShippingCostOld          string    `bun:"shipping_cost_old,type:varchar(255),nullzero"`
	ShippingCostNew          string    `bun:"shipping_cost_new,type:varchar(255),nullzero"`
	TaxAmountOld             string    `bun:"tax_amount_old,type:varchar(255),nullzero"`
	TaxAmountNew             string    `bun:"tax_amount_new,type:varchar(255),nullzero"`
	TotalAmountOld           string    `bun:"total_amount_old,type:varchar(255),nullzero"`
	TotalAmountNew           string    `bun:"total_amount_new,type:varchar(255),nullzero"`
	ShipToIdOld              string    `bun:"ship_to_id_old,type:varchar(255),nullzero"`
	ShipToIdNew              string    `bun:"ship_to_id_new,type:varchar(255),nullzero"`
	InvoiceBatchUidOld       string    `bun:"invoice_batch_uid_old,type:varchar(255),nullzero"`
	InvoiceBatchUidNew       string    `bun:"invoice_batch_uid_new,type:varchar(255),nullzero"`
	ShippingRouteUidOld      string    `bun:"shipping_route_uid_old,type:varchar(255),nullzero"`
	ShippingRouteUidNew      string    `bun:"shipping_route_uid_new,type:varchar(255),nullzero"`
	PoNoOld                  string    `bun:"po_no_old,type:varchar(255),nullzero"`
	PoNoNew                  string    `bun:"po_no_new,type:varchar(255),nullzero"`
	PeriodOld                string    `bun:"period_old,type:varchar(255),nullzero"`
	PeriodNew                string    `bun:"period_new,type:varchar(255),nullzero"`
	YearForPeriodOld         string    `bun:"year_for_period_old,type:varchar(255),nullzero"`
	YearForPeriodNew         string    `bun:"year_for_period_new,type:varchar(255),nullzero"`
	BrokerageOld             string    `bun:"brokerage_old,type:varchar(255),nullzero"`
	BrokerageNew             string    `bun:"brokerage_new,type:varchar(255),nullzero"`
	TermsIdOld               string    `bun:"terms_id_old,type:varchar(255),nullzero"`
	TermsIdNew               string    `bun:"terms_id_new,type:varchar(255),nullzero"`
	CommissionCostOld        string    `bun:"commission_cost_old,type:varchar(255),nullzero"`
	CommissionCostNew        string    `bun:"commission_cost_new,type:varchar(255),nullzero"`
	InvoiceReferenceNoOld    string    `bun:"invoice_reference_no_old,type:varchar(255),nullzero"`
	InvoiceReferenceNoNew    string    `bun:"invoice_reference_no_new,type:varchar(255),nullzero"`
	InvoiceDescOld           string    `bun:"invoice_desc_old,type:varchar(255),nullzero"`
	InvoiceDescNew           string    `bun:"invoice_desc_new,type:varchar(255),nullzero"`
}
