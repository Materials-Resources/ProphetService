package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type SalesrepCommission struct {
	bun.BaseModel             `bun:"table:salesrep_commission"`
	CompanyId                 string    `bun:"company_id,type:varchar(8),unique"`                          // Unique code that identifies a company.
	SalesrepId                string    `bun:"salesrep_id,type:varchar(16),unique,nullzero"`               // What sales representative does this relationship refer to?
	CommissionPaidOn          string    `bun:"commission_paid_on,type:char(1)"`                            // Indicates where commissions are paid on Invoice Generation or Cash Receipts.
	PaidOnPartialPayments     string    `bun:"paid_on_partial_payments,type:char(1)"`                      // If commission are paid on Cash receipts, are commissions paid on partial payments?
	CommissionCutOff          string    `bun:"commission_cut_off,type:char(1)"`                            // Indicates whether commission is paid on payments received after the due date of the invoice.
	NumberOfDaysOverdue       float64   `bun:"number_of_days_overdue,type:decimal(3,0),nullzero"`          // Ddetermines how many days must pass after an invoice is overdue before commission is not paid.
	TrackQuotas               string    `bun:"track_quotas,type:char(1)"`                                  // Indicates if quotas are tracked for this salesrep.
	CommissionScheduleId      string    `bun:"commission_schedule_id,type:varchar(8),nullzero"`            // User-defined code that identifies a commission schedule.
	DeleteFlag                string    `bun:"delete_flag,type:char(1)"`                                   // Indicates whether this record is logically deleted
	DateCreated               time.Time `bun:"date_created,type:datetime"`                                 // Indicates the date/time this record was created.
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`                           // Indicates the date/time this record was last modified.
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`  // ID of the user who last maintained this record
	MfrRepCommissionPaidOnCd  int32     `bun:"mfr_rep_commission_paid_on_cd,type:int,default:((1893))"`    // Indicates where commissions are paid on for Manufacturer Rep Orders.
	MfrRepPaidOnPartialFlag   string    `bun:"mfr_rep_paid_on_partial_flag,type:char(1),default:('N')"`    // Determines whether or not commission for Cash Receipts are paid on partial payments.
	TotalProfitThreshold      float64   `bun:"total_profit_threshold,type:decimal(19,2),default:((0.00))"` // The aggregate profit over invoice lines that generate a commission needs to meet this value.
	TotalProfitThresholdType  string    `bun:"total_profit_threshold_type,type:char(1),nullzero"`          // The type of total_profit_threshold ($ or %).  Can be null if neither is selected.
	SalesrepCommissionUid     int32     `bun:"salesrep_commission_uid,type:int,autoincrement,pk"`          // Unique identifier on the salerep_commission table
	TerritoryUid              int32     `bun:"territory_uid,type:int,unique,nullzero"`                     // Custom This column will hold the territory used for commission calcs.
	HouseSplitPercentage      float64   `bun:"house_split_percentage,type:decimal(19,9),nullzero"`         // Decimal value indicating a split percentage that goes to the house when commissions are calculated for the salesrep.
	CalcDaysOverdueFromDateCd int32     `bun:"calc_days_overdue_from_date_cd,type:int,default:((1347))"`   // Indicates whether the days overdue should be calculated by due date or invoice date.
}
