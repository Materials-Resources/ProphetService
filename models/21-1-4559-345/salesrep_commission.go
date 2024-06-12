package model

import (
	"github.com/uptrace/bun"
	"time"
)

type SalesrepCommission struct {
	bun.BaseModel             `bun:"table:salesrep_commission"`
	CompanyId                 string    `bun:"company_id,type:varchar(8)"`
	SalesrepId                string    `bun:"salesrep_id,type:varchar(16),nullzero"`
	CommissionPaidOn          string    `bun:"commission_paid_on,type:char"`
	PaidOnPartialPayments     string    `bun:"paid_on_partial_payments,type:char"`
	CommissionCutOff          string    `bun:"commission_cut_off,type:char"`
	NumberOfDaysOverdue       float64   `bun:"number_of_days_overdue,type:decimal(3,0),nullzero"`
	TrackQuotas               string    `bun:"track_quotas,type:char"`
	CommissionScheduleId      string    `bun:"commission_schedule_id,type:varchar(8),nullzero"`
	DeleteFlag                string    `bun:"delete_flag,type:char"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	MfrRepCommissionPaidOnCd  int32     `bun:"mfr_rep_commission_paid_on_cd,type:int,default:((1893))"`
	MfrRepPaidOnPartialFlag   string    `bun:"mfr_rep_paid_on_partial_flag,type:char,default:('N')"`
	TotalProfitThreshold      float64   `bun:"total_profit_threshold,type:decimal(19,2),default:((0.00))"`
	TotalProfitThresholdType  string    `bun:"total_profit_threshold_type,type:char,nullzero"`
	SalesrepCommissionUid     int32     `bun:"salesrep_commission_uid,type:int,pk,identity"`
	TerritoryUid              int32     `bun:"territory_uid,type:int,nullzero"`
	HouseSplitPercentage      float64   `bun:"house_split_percentage,type:decimal(19,9),nullzero"`
	CalcDaysOverdueFromDateCd int32     `bun:"calc_days_overdue_from_date_cd,type:int,default:((1347))"`
}
