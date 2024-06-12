package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CommissionDefaults struct {
	bun.BaseModel             `bun:"table:commission_defaults"`
	CompanyId                 string    `bun:"company_id,type:varchar(8),pk"`
	CommissionBasedOn         string    `bun:"commission_based_on,type:char(2)"`
	CommissionPaidOn          string    `bun:"commission_paid_on,type:char"`
	PaidOnPartialPayments     string    `bun:"paid_on_partial_payments,type:char"`
	CommissionCutOff          string    `bun:"commission_cut_off,type:char"`
	NumberOfDaysOverdue       float64   `bun:"number_of_days_overdue,type:decimal(3,0),nullzero"`
	IncludeOtherCharge        string    `bun:"include_other_charge,type:char"`
	CommissionScheduleId      string    `bun:"commission_schedule_id,type:varchar(8),nullzero"`
	DeleteFlag                string    `bun:"delete_flag,type:char"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	IncludeFreightOut         string    `bun:"include_freight_out,type:char,default:('Y')"`
	IncludeFreightIn          string    `bun:"include_freight_in,type:char,default:('Y')"`
	IncludeNoChargeInvoices   string    `bun:"include_no_charge_invoices,type:char,default:('N')"`
	ReduceCommByTermsFlag     string    `bun:"reduce_comm_by_terms_flag,type:char,default:('N')"`
	TotalProfitThreshold      float64   `bun:"total_profit_threshold,type:decimal(19,2),default:((0.00))"`
	TotalProfitThresholdType  string    `bun:"total_profit_threshold_type,type:char,nullzero"`
	CalcDaysOverdueFromDateCd int32     `bun:"calc_days_overdue_from_date_cd,type:int,default:((1347))"`
}
