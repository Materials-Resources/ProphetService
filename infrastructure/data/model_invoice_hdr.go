package data

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"time"
)

type InvoiceHdr struct {
	bun.BaseModel              `bun:"table:invoice_hdr"`
	InvoiceNo                  string          `bun:"invoice_no,type:varchar(10),pk"`
	OrderNo                    sql.NullString  `bun:"order_no,type:varchar(8),nullzero"`
	OrderDate                  sql.NullTime    `bun:"order_date,type:datetime,nullzero"`
	InvoiceDate                time.Time       `bun:"invoice_date,type:datetime"`
	CustomerId                 string          `bun:"customer_id,type:varchar(12)"`
	Bill2Name                  sql.NullString  `bun:"bill2_name,type:varchar(50),nullzero"`
	Bill2Contact               sql.NullString  `bun:"bill2_contact,type:varchar(30),nullzero"`
	Bill2Address1              sql.NullString  `bun:"bill2_address1,type:varchar(50),nullzero"`
	Bill2Address2              sql.NullString  `bun:"bill2_address2,type:varchar(50),nullzero"`
	Bill2City                  sql.NullString  `bun:"bill2_city,type:varchar(50),nullzero"`
	Bill2State                 sql.NullString  `bun:"bill2_state,type:varchar(50),nullzero"`
	Bill2PostalCode            sql.NullString  `bun:"bill2_postal_code,type:varchar(10),nullzero"`
	Ship2Name                  sql.NullString  `bun:"ship2_name,type:varchar(50),nullzero"`
	Ship2Contact               sql.NullString  `bun:"ship2_contact,type:varchar(30),nullzero"`
	Ship2Address1              sql.NullString  `bun:"ship2_address1,type:varchar(50),nullzero"`
	Ship2Address2              sql.NullString  `bun:"ship2_address2,type:varchar(50),nullzero"`
	Ship2City                  sql.NullString  `bun:"ship2_city,type:varchar(50),nullzero"`
	Ship2State                 sql.NullString  `bun:"ship2_state,type:varchar(50),nullzero"`
	Ship2PostalCode            sql.NullString  `bun:"ship2_postal_code,type:varchar(10),nullzero"`
	CarrierName                sql.NullString  `bun:"carrier_name,type:varchar(50),nullzero"`
	Fob                        sql.NullString  `bun:"fob,type:varchar(20),nullzero"`
	TermsDesc                  sql.NullString  `bun:"terms_desc,type:varchar(20),nullzero"`
	PoNo                       sql.NullString  `bun:"po_no,type:varchar(255),nullzero"`
	SalesrepId                 sql.NullString  `bun:"salesrep_id,type:varchar(16),nullzero"`
	SalesrepName               sql.NullString  `bun:"salesrep_name,type:varchar(40),nullzero"`
	Brokerage                  sql.NullFloat64 `bun:"brokerage,type:decimal(19,4),nullzero"`
	Freight                    sql.NullFloat64 `bun:"freight,type:decimal(19,4),nullzero"`
	ArAccountNo                sql.NullString  `bun:"ar_account_no,type:varchar(32),nullzero"`
	GlFreightAccountNo         sql.NullString  `bun:"gl_freight_account_no,type:varchar(32),nullzero"`
	GlBrokerageAccountNo       sql.NullString  `bun:"gl_brokerage_account_no,type:varchar(32),nullzero"`
	BrokerageAmt               sql.NullFloat64 `bun:"brokerage_amt,type:decimal(19,4),nullzero"`
	Period                     sql.NullFloat64 `bun:"period,type:decimal(3,0),nullzero"`
	YearForPeriod              sql.NullFloat64 `bun:"year_for_period,type:decimal(4,0),nullzero"`
	StoreNo                    sql.NullString  `bun:"store_no,type:varchar(5),nullzero"`
	InvoiceType                sql.NullString  `bun:"invoice_type,type:varchar(2),nullzero"`
	ShipToId                   sql.NullFloat64 `bun:"ship_to_id,type:decimal(19,0),nullzero"`
	ShipDate                   sql.NullTime    `bun:"ship_date,type:datetime,nullzero"`
	TotalAmount                float64         `bun:"total_amount,type:decimal(19,4),default:(0)"`
	AmountPaid                 float64         `bun:"amount_paid,type:decimal(19,4),default:(0)"`
	TermsTaken                 float64         `bun:"terms_taken,type:decimal(19,4),default:(0)"`
	Allowed                    float64         `bun:"allowed,type:decimal(19,4),default:(0)"`
	PaidInFullFlag             sql.NullString  `bun:"paid_in_full_flag,type:char,default:('N')"`
	PaidByCheckNo              sql.NullFloat64 `bun:"paid_by_check_no,type:decimal(19,4),nullzero"`
	DatePaid                   sql.NullTime    `bun:"date_paid,type:datetime,nullzero"`
	PrintFlag                  sql.NullString  `bun:"print_flag,type:char,nullzero"`
	PrintDate                  sql.NullTime    `bun:"print_date,type:datetime,nullzero"`
	CompanyNo                  string          `bun:"company_no,type:varchar(8)"`
	CustomerIdNumber           float64         `bun:"customer_id_number,type:decimal(19,0)"`
	DateCreated                time.Time       `bun:"date_created,type:datetime"`
	DateLastModified           time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	Printed                    sql.NullString  `bun:"printed,type:char,nullzero"`
	PrintedDate                sql.NullTime    `bun:"printed_date,type:datetime,nullzero"`
	CorpAddressId              sql.NullFloat64 `bun:"corp_address_id,type:decimal(19,0),nullzero"`
	ShippingCost               sql.NullFloat64 `bun:"shipping_cost,type:decimal(19,4),nullzero"`
	Bill2Country               sql.NullString  `bun:"bill2_country,type:varchar(50),nullzero"`
	Ship2Country               sql.NullString  `bun:"ship2_country,type:varchar(50),nullzero"`
	InvoiceReferenceNo         string          `bun:"invoice_reference_no,type:varchar(10)"`
	InvoiceAdjustmentType      sql.NullString  `bun:"invoice_adjustment_type,type:char,nullzero"`
	InvoiceDesc                sql.NullString  `bun:"invoice_desc,type:varchar(255),nullzero"`
	MemoAmount                 float64         `bun:"memo_amount,type:decimal(19,4),default:(0)"`
	BadDebtAmount              float64         `bun:"bad_debt_amount,type:decimal(19,4),default:(0)"`
	InvoiceClass               sql.NullString  `bun:"invoice_class,type:varchar(8),nullzero"`
	PeriodFullyPaid            sql.NullFloat64 `bun:"period_fully_paid,type:decimal(3,0),nullzero"`
	YearFullyPaid              sql.NullFloat64 `bun:"year_fully_paid,type:decimal(4,0),nullzero"`
	Approved                   sql.NullString  `bun:"approved,type:char,nullzero"`
	FcThruDate                 sql.NullTime    `bun:"fc_thru_date,type:datetime,nullzero"`
	CumulativeFc               sql.NullFloat64 `bun:"cumulative_fc,type:decimal(19,4),nullzero"`
	NetDueDate                 sql.NullTime    `bun:"net_due_date,type:datetime,nullzero"`
	TermsDueDate               sql.NullTime    `bun:"terms_due_date,type:datetime,nullzero"`
	TermsId                    sql.NullString  `bun:"terms_id,type:varchar(2),nullzero"`
	BranchId                   string          `bun:"branch_id,type:varchar(8)"`
	DisputedFlag               sql.NullString  `bun:"disputed_flag,type:char,nullzero"`
	StatementPeriod            sql.NullFloat64 `bun:"statement_period,type:decimal(3,0),nullzero"`
	StatementYear              sql.NullFloat64 `bun:"statement_year,type:decimal(4,0),nullzero"`
	OtherChargeAmount          float64         `bun:"other_charge_amount,type:decimal(19,4),default:(0)"`
	TaxAmount                  float64         `bun:"tax_amount,type:decimal(19,4),default:(0)"`
	OriginalDocumentType       string          `bun:"original_document_type,type:char"`
	Consolidated               string          `bun:"consolidated,type:char"`
	SoldToAhUid                int32           `bun:"sold_to_ah_uid,type:int"`
	SoldToCustomerId           float64         `bun:"sold_to_customer_id,type:decimal(19,0)"`
	ShipToPhone                sql.NullString  `bun:"ship_to_phone,type:varchar(20),nullzero"`
	InvoiceBatchUid            int32           `bun:"invoice_batch_uid,type:int,default:(1)"`
	FreightCodeUid             sql.NullInt32   `bun:"freight_code_uid,type:int,nullzero"`
	ShippingRouteUid           sql.NullInt32   `bun:"shipping_route_uid,type:int,nullzero"`
	TransmissionMethod         sql.NullInt32   `bun:"transmission_method,type:int,nullzero"`
	TermsAmount                float64         `bun:"terms_amount,type:decimal(19,4),default:(0)"`
	SalesLocationId            sql.NullFloat64 `bun:"sales_location_id,type:decimal(19,0),nullzero"`
	SourceTypeCd               sql.NullInt32   `bun:"source_type_cd,type:int,nullzero"`
	Ship2EmailAddress          sql.NullString  `bun:"ship2_email_address,type:varchar(255),nullzero"`
	CurrencyLineUid            sql.NullInt32   `bun:"currency_line_uid,type:int,nullzero"`
	CreatedBy                  sql.NullString  `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	TaxTermsTaken              float64         `bun:"tax_terms_taken,type:decimal(19,9),default:((0.00))"`
	CarrierInsuranceAmt        sql.NullFloat64 `bun:"carrier_insurance_amt,type:decimal(19,9),nullzero"`
	InvNoDisplay               sql.NullString  `bun:"inv_no_display,type:varchar(255),nullzero"`
	IvaExemptionNumber         sql.NullString  `bun:"iva_exemption_number,type:varchar(30),nullzero"`
	IvaTaxableFlag             sql.NullString  `bun:"iva_taxable_flag,type:char,nullzero"`
	JobId                      sql.NullString  `bun:"job_id,type:varchar(32),nullzero"`
	InvoicePaidPeriod          sql.NullInt32   `bun:"invoice_paid_period,type:int,nullzero"`
	InvoicePeriod              sql.NullInt32   `bun:"invoice_period,type:int,nullzero"`
	CreditMemoForTermsFlag     sql.NullString  `bun:"credit_memo_for_terms_flag,type:char,nullzero"`
	RecordTypeCd               sql.NullInt32   `bun:"record_type_cd,type:int,nullzero"`
	CommissionCost             sql.NullFloat64 `bun:"commission_cost,type:decimal(19,9),default:((0))"`
	ReceiverName               sql.NullString  `bun:"receiver_name,type:varchar(255),nullzero"`
	CartonCount                sql.NullInt32   `bun:"carton_count,type:int,default:(NULL)"`
	StrategicFreightIn         sql.NullFloat64 `bun:"strategic_freight_in,type:decimal(19,9),nullzero"`
	StrategicFreightOut        sql.NullFloat64 `bun:"strategic_freight_out,type:decimal(19,9),nullzero"`
	AffiliatedTrainingCenter   sql.NullFloat64 `bun:"affiliated_training_center,type:decimal(19,0),nullzero"`
	ExternalReferenceNo        sql.NullString  `bun:"external_reference_no,type:varchar(255),nullzero"`
	TotalFreightchargeWeight   sql.NullFloat64 `bun:"total_freightcharge_weight,type:decimal(19,4),nullzero"`
	CardlockConsInvoiceFlag    sql.NullString  `bun:"cardlock_cons_invoice_flag,type:char,nullzero"`
	RemoveFromOpenDefRevWindow sql.NullString  `bun:"remove_from_open_def_rev_window,type:char,nullzero"`
	RecordTypeReferenceNo      sql.NullString  `bun:"record_type_reference_no,type:varchar(10),nullzero"`
	RebillInvoiceReasonUid     sql.NullInt32   `bun:"rebill_invoice_reason_uid,type:int,nullzero"`
	EdiOrderPrintedFlag        sql.NullString  `bun:"edi_order_printed_flag,type:char,nullzero"`
	DownpaymentApplied         sql.NullFloat64 `bun:"downpayment_applied,type:decimal(19,2),nullzero"`
	ReverseRedemptionFlag      sql.NullString  `bun:"reverse_redemption_flag,type:char,default:('N')"`
	SentToCarrierDate          sql.NullTime    `bun:"sent_to_carrier_date,type:datetime,nullzero"`
	TaxTermsAmt                float64         `bun:"tax_terms_amt,type:decimal(19,9),default:((0))"`
	TaxAmountPaid              float64         `bun:"tax_amount_paid,type:decimal(19,9),default:((0))"`
	ReasonCreditMemoCodeUid    sql.NullInt32   `bun:"reason_credit_memo_code_uid,type:int,nullzero"`
	SourceCreditMemoCodeUid    sql.NullInt32   `bun:"source_credit_memo_code_uid,type:int,nullzero"`
	ExternalClaimId            sql.NullString  `bun:"external_claim_id,type:varchar(255),nullzero"`
	Bill2Address3              sql.NullString  `bun:"bill2_address3,type:varchar(50),nullzero"`
	Ship2Address3              sql.NullString  `bun:"ship2_address3,type:varchar(50),nullzero"`
	EftExportedDate            sql.NullTime    `bun:"eft_exported_date,type:datetime,nullzero"`
	CourtesyInvoiceFlag        sql.NullString  `bun:"courtesy_invoice_flag,type:char,default:('N')"`
	IvaWithheldAmount          float64         `bun:"iva_withheld_amount,type:decimal(19,2),default:((0))"`
	CarrierId                  sql.NullFloat64 `bun:"carrier_id,type:decimal(19,0),nullzero"`
	FreeOutFreightMinMetFlag   sql.NullString  `bun:"free_out_freight_min_met_flag,type:char,default:('N')"`
	RegNumOfPackages           sql.NullString  `bun:"reg_num_of_packages,type:varchar(255),nullzero"`
	Ship2Latitude              sql.NullString  `bun:"ship2_latitude,type:varchar(20),nullzero"`
	Ship2Longitude             sql.NullString  `bun:"ship2_longitude,type:varchar(20),nullzero"`
	DeliveryNo                 sql.NullInt32   `bun:"delivery_no,type:int,nullzero"`
	AllowedAmtReasonId         sql.NullString  `bun:"allowed_amt_reason_id,type:varchar(255),nullzero"`
	TotalServiceAmountRedeemed sql.NullFloat64 `bun:"total_service_amount_redeemed,type:decimal(19,9),nullzero"`
	CustomerClaimNo            sql.NullString  `bun:"customer_claim_no,type:varchar(255),nullzero"`
	ClaimType                  sql.NullString  `bun:"claim_type,type:varchar(255),nullzero"`
	ClaimCoopFlag              sql.NullString  `bun:"claim_coop_flag,type:char,nullzero"`
	BillToSupplierFlag         sql.NullString  `bun:"bill_to_supplier_flag,type:char,nullzero"`
	CustPriceProtVarAcctNo     sql.NullString  `bun:"cust_price_prot_var_acct_no,type:varchar(32),nullzero"`
	ConsolidatedTypeFlag       sql.NullString  `bun:"consolidated_type_flag,type:char,nullzero"`
	MerchandiseInvoiceFlag     sql.NullString  `bun:"merchandise_invoice_flag,type:char,nullzero"`
	SalesMarketGroupUid        sql.NullInt32   `bun:"sales_market_group_uid,type:int,nullzero"`
	ActualFreightCost          sql.NullFloat64 `bun:"actual_freight_cost,type:decimal(19,9),nullzero"`
	EcoFeeAmt                  sql.NullFloat64 `bun:"eco_fee_amt,type:decimal(19,9),nullzero"`
	FrghtDiscTaken             float64         `bun:"frght_disc_taken,type:decimal(19,4),default:((0.00))"`
	RentalInvoiceNo            sql.NullString  `bun:"rental_invoice_no,type:varchar(10),nullzero"`
}

type InvoiceHdrModel struct {
	bun bun.IDB
}

func (m *InvoiceHdrModel) Get(invoiceNo string) (*InvoiceHdr, error) {
	return nil, nil
}

type InvoiceHdrGetAllParams struct {
	OrderNo *[]string
}

func (m *InvoiceHdrModel) GetAll(ctx context.Context, params InvoiceHdrGetAllParams) ([]*InvoiceHdr, error) {
	var invoiceHdrRecords []*InvoiceHdr
	q := m.bun.NewSelect().Model(&invoiceHdrRecords)
	if params.OrderNo != nil {
		q = q.Where("order_no IN (?)", bun.In(*params.OrderNo))
	}

	err := q.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return invoiceHdrRecords, nil
}
