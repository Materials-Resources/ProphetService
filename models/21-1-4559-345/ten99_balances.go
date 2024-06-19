package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Ten99Balances struct {
	bun.BaseModel           `bun:"table:ten99_balances"`
	CompanyId               string    `bun:"company_id,type:varchar(8)"`                                // Unique code that identifies a company.
	VendorId                float64   `bun:"vendor_id,type:decimal(18,0)"`                              // What is the vendor for this stage?
	Year                    float64   `bun:"year,type:decimal(4,0)"`                                    // Year of the 1099 balances
	Rents                   float64   `bun:"rents,type:decimal(19,4)"`                                  // Rents  year balance
	Royalties               float64   `bun:"royalties,type:decimal(19,4)"`                              // Royalties year balance
	OtherIncome             float64   `bun:"other_income,type:decimal(19,4)"`                           // Other Income year balance
	FitWh                   float64   `bun:"fit_wh,type:decimal(19,4)"`                                 // Federal income tax withheld year balance
	FishingBoatProceeds     float64   `bun:"fishing_boat_proceeds,type:decimal(19,4)"`                  // Fishing Boat Proceeds year balance
	MedicalAndHealthcare    float64   `bun:"medical_and_healthcare,type:decimal(19,4)"`                 // Medical/healthcare payments year balance
	NonEmployeeCompensation float64   `bun:"non_employee_compensation,type:decimal(19,4)"`              // Nonemployee compensation year balance
	CropInsuranceProceeds   float64   `bun:"crop_insurance_proceeds,type:decimal(19,4)"`                // Crop Insurance Proceeds year balance
	SubstitutePayments      float64   `bun:"substitute_payments,type:decimal(19,4)"`                    // Substitute payments in lieu  of dividends or interest year balance
	ExcessGoldenParachute   float64   `bun:"excess_golden_parachute,type:decimal(19,4)"`                // Excess Golden Parachute Payments year balance
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	Ten99BalancesUid        int32     `bun:"ten99_balances_uid,type:int,pk"`                            // Unique identifier for the table
	CashLiquidations        float64   `bun:"cash_liquidations,type:decimal(19,9),default:(0)"`          // Cash liquidation year balance
	OrdindaryDividends      float64   `bun:"ordindary_dividends,type:decimal(19,9),default:(0)"`        // Ordindary distributions year balance
	AttorneyProceeds        float64   `bun:"attorney_proceeds,type:decimal(19,4),default:(0)"`          // Gross proceeds paid to an attorney for 1099 tax form
	InterestIncome          float64   `bun:"interest_income,type:decimal(19,9),default:((0))"`          // Interest income yearly balance reported on 1099-INT
	IsrTaxWithheld          float64   `bun:"isr_tax_withheld,type:decimal(19,9),nullzero"`              // Custom (F30860): ISR tax amount withheld from vendor invoices
	QualifiedDividends      float64   `bun:"qualified_dividends,type:decimal(19,9),default:((0))"`      // Qualified Dividends for 1099-DIV
}
