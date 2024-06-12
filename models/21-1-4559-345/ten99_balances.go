package model

type Ten99Balances struct {
	bun.BaseModel           `bun:"table:ten99_balances"`
	CompanyId               string    `bun:"company_id,type:varchar(8)"`
	VendorId                float64   `bun:"vendor_id,type:decimal(18,0)"`
	Year                    float64   `bun:"year,type:decimal(4,0)"`
	Rents                   float64   `bun:"rents,type:decimal(19,4)"`
	Royalties               float64   `bun:"royalties,type:decimal(19,4)"`
	OtherIncome             float64   `bun:"other_income,type:decimal(19,4)"`
	FitWh                   float64   `bun:"fit_wh,type:decimal(19,4)"`
	FishingBoatProceeds     float64   `bun:"fishing_boat_proceeds,type:decimal(19,4)"`
	MedicalAndHealthcare    float64   `bun:"medical_and_healthcare,type:decimal(19,4)"`
	NonEmployeeCompensation float64   `bun:"non_employee_compensation,type:decimal(19,4)"`
	CropInsuranceProceeds   float64   `bun:"crop_insurance_proceeds,type:decimal(19,4)"`
	SubstitutePayments      float64   `bun:"substitute_payments,type:decimal(19,4)"`
	ExcessGoldenParachute   float64   `bun:"excess_golden_parachute,type:decimal(19,4)"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	Ten99BalancesUid        int32     `bun:"ten99_balances_uid,type:int,pk"`
	CashLiquidations        float64   `bun:"cash_liquidations,type:decimal(19,9),default:(0)"`
	OrdindaryDividends      float64   `bun:"ordindary_dividends,type:decimal(19,9),default:(0)"`
	AttorneyProceeds        float64   `bun:"attorney_proceeds,type:decimal(19,4),default:(0)"`
	InterestIncome          float64   `bun:"interest_income,type:decimal(19,9),default:((0))"`
	IsrTaxWithheld          float64   `bun:"isr_tax_withheld,type:decimal(19,9),nullzero"`
	QualifiedDividends      float64   `bun:"qualified_dividends,type:decimal(19,9),default:((0))"`
}
