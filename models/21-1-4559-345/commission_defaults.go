package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CommissionDefaults struct {
	bun.BaseModel             `bun:"table:commission_defaults"`
	CompanyId                 string    `bun:"company_id,type:varchar(8),pk"`                              // Unique code that identifies a company.
	CommissionBasedOn         string    `bun:"commission_based_on,type:char(2)"`                           // Indicates default basis for computing commissions
	CommissionPaidOn          string    `bun:"commission_paid_on,type:char(1)"`                            // Indicates default basis for paying commissions~r~n
	PaidOnPartialPayments     string    `bun:"paid_on_partial_payments,type:char(1)"`                      // Indicates whether commission can be paid on partia
	CommissionCutOff          string    `bun:"commission_cut_off,type:char(1)"`                            // Indicates whether commission is not paid on overdu
	NumberOfDaysOverdue       float64   `bun:"number_of_days_overdue,type:decimal(3,0),nullzero"`          // Number of days overdue before an invoice is not el
	IncludeOtherCharge        string    `bun:"include_other_charge,type:char(1)"`                          // Indicates if commission is paid on Other Charge It
	CommissionScheduleId      string    `bun:"commission_schedule_id,type:varchar(8),nullzero"`            // Commission Schedule for this salesrep (used only w
	DeleteFlag                string    `bun:"delete_flag,type:char(1)"`                                   // Indicates whether this record is logically deleted
	DateCreated               time.Time `bun:"date_created,type:datetime"`                                 // Indicates the date/time this record was created.
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`                           // Indicates the date/time this record was last modified.
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`  // ID of the user who last maintained this record
	IncludeFreightOut         string    `bun:"include_freight_out,type:char(1),default:('Y')"`             // Indicates if commission if paid on out freight billed to customer
	IncludeFreightIn          string    `bun:"include_freight_in,type:char(1),default:('Y')"`              // Indicates if commission if paid on in freight billed to customer
	IncludeNoChargeInvoices   string    `bun:"include_no_charge_invoices,type:char(1),default:('N')"`      // Determines whether to include invoices with a 0 price in comm calculations
	ReduceCommByTermsFlag     string    `bun:"reduce_comm_by_terms_flag,type:char(1),default:('N')"`       // Indicates whether the amount used to calculate the commission should be reduced by terms taken.
	TotalProfitThreshold      float64   `bun:"total_profit_threshold,type:decimal(19,2),default:((0.00))"` // The aggregate profit over invoice lines that generate a commission needs to meet this value.
	TotalProfitThresholdType  string    `bun:"total_profit_threshold_type,type:char(1),nullzero"`          // The type of total_profit_threshold ($ or %).  Can be null if neither is selected.
	CalcDaysOverdueFromDateCd int32     `bun:"calc_days_overdue_from_date_cd,type:int,default:((1347))"`   // Indicates whether the days overdue should be calculated by due date or invoice date.
}
