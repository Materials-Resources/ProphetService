package model

type InvoiceHdrSalesrep struct {
	bun.BaseModel              `bun:"table:invoice_hdr_salesrep"`
	InvoiceNumber              string    `bun:"invoice_number,type:varchar(10)"`
	SalesrepId                 string    `bun:"salesrep_id,type:varchar(16)"`
	CompanyId                  string    `bun:"company_id,type:varchar(8)"`
	VendorId                   float64   `bun:"vendor_id,type:decimal(19,0),nullzero"`
	CommissionAmount           float64   `bun:"commission_amount,type:decimal(19,4)"`
	DeleteFlag                 string    `bun:"delete_flag,type:char"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	DatePaid                   time.Time `bun:"date_paid,type:datetime,nullzero"`
	CommissionPaid             float64   `bun:"commission_paid,type:decimal(19,4),nullzero"`
	PrimarySalesrep            string    `bun:"primary_salesrep,type:char,default:('N')"`
	CommissionPercentage       float64   `bun:"commission_percentage,type:decimal(19,4),default:((100))"`
	SalesrepSourceCd           int32     `bun:"salesrep_source_cd,type:int,nullzero"`
	EditedCommissionFlag       string    `bun:"edited_commission_flag,type:char,nullzero"`
	PriorCommissionAmount      float64   `bun:"prior_commission_amount,type:decimal(19,2),nullzero"`
	SplitFlag                  string    `bun:"split_flag,type:char,nullzero"`
	InvoiceHdrSalesrepUid      int32     `bun:"invoice_hdr_salesrep_uid,type:int,pk,identity"`
	CommissionScheduleUid      int32     `bun:"commission_schedule_uid,type:int,nullzero"`
	CommissionCost             float64   `bun:"commission_cost,type:decimal(19,4),nullzero"`
	ExtendedPrice              float64   `bun:"extended_price,type:decimal(19,4),nullzero"`
	CalculationFromSplitFlag   string    `bun:"calculation_from_split_flag,type:char,default:('N')"`
	ExcludeSplitValidationFlag string    `bun:"exclude_split_validation_flag,type:char,default:('Y')"`
	PartialInvoiceOnlyFlag     string    `bun:"partial_invoice_only_flag,type:char,nullzero"`
	ForfeitedAmount            float64   `bun:"forfeited_amount,type:decimal(19,9),nullzero"`
	LinePercentageEditedFlag   string    `bun:"line_percentage_edited_flag,type:char,nullzero"`
	CommissionOverridePercent  float64   `bun:"commission_override_percent,type:decimal(19,9),nullzero"`
	ExtendedPriceNoCndAdjust   float64   `bun:"extended_price_no_cnd_adjust,type:decimal(19,4),nullzero"`
}
