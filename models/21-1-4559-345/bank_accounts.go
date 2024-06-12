package model

type BankAccounts struct {
	bun.BaseModel             `bun:"table:bank_accounts"`
	CompanyNo                 string    `bun:"company_no,type:varchar(8),pk"`
	BankNo                    float64   `bun:"bank_no,type:decimal(19,0),pk"`
	TransitCode               string    `bun:"transit_code,type:varchar(15)"`
	AccountNo                 string    `bun:"account_no,type:varchar(18)"`
	BankCode                  float64   `bun:"bank_code,type:decimal(18,0)"`
	FedResCode                string    `bun:"fed_res_code,type:varchar(10)"`
	BankName                  string    `bun:"bank_name,type:varchar(30),nullzero"`
	BankBranch                string    `bun:"bank_branch,type:varchar(30),nullzero"`
	BankAddress1              string    `bun:"bank_address1,type:varchar(30),nullzero"`
	BankAddress2              string    `bun:"bank_address2,type:varchar(30),nullzero"`
	LastCheckNo               float64   `bun:"last_check_no,type:decimal(18,0)"`
	GlAccountNo               string    `bun:"gl_account_no,type:varchar(32)"`
	Balance                   float64   `bun:"balance,type:decimal(19,4),default:(0)"`
	Active                    string    `bun:"active,type:char"`
	DeleteFlag                string    `bun:"delete_flag,type:char"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	CurrencyId                float64   `bun:"currency_id,type:decimal(19,0)"`
	CheckFormName             string    `bun:"check_form_name,type:varchar(255),nullzero"`
	BeginReconcileDate        time.Time `bun:"begin_reconcile_date,type:datetime,nullzero"`
	ReconcileFlag             string    `bun:"reconcile_flag,type:char,default:('N')"`
	BeginReconciledThruPeriod int16     `bun:"begin_reconciled_thru_period,type:smallint,nullzero"`
	BeginReconciledThruYear   int16     `bun:"begin_reconciled_thru_year,type:smallint,nullzero"`
	AddressTypeCd             int32     `bun:"address_type_cd,type:int"`
	UsePreprintedChecksFlag   string    `bun:"use_preprinted_checks_flag,type:char,nullzero"`
	NumberOfSigs              string    `bun:"number_of_sigs,type:char,nullzero"`
	Sig1Path                  string    `bun:"sig1_path,type:varchar(255),nullzero"`
	Sig2Path                  string    `bun:"sig2_path,type:varchar(255),nullzero"`
	AmountReqTwoSigs          float64   `bun:"amount_req_two_sigs,type:decimal(19,9),nullzero"`
	RoutingNumber             string    `bun:"routing_number,type:varchar(255),nullzero"`
	CheckDigit                string    `bun:"check_digit,type:char,nullzero"`
	AchFileIdModifier         string    `bun:"ach_file_id_modifier,type:char,nullzero"`
	AchFileIdModifierDate     time.Time `bun:"ach_file_id_modifier_date,type:datetime,nullzero"`
	CheckLanguageCd           int32     `bun:"check_language_cd,type:int,nullzero"`
	ExportPath                string    `bun:"export_path,type:varchar(255),nullzero"`
	AchCompanyId              string    `bun:"ach_company_id,type:varchar(40),nullzero"`
	CeoCompany                string    `bun:"ceo_company,type:varchar(40),nullzero"`
	CommercialCardNo          string    `bun:"commercial_card_no,type:varchar(40),nullzero"`
	EddBillerId               string    `bun:"edd_biller_id,type:varchar(255),nullzero"`
	DocumentTemplateNo        string    `bun:"document_template_no,type:varchar(40),nullzero"`
	BankCity                  string    `bun:"bank_city,type:varchar(40),nullzero"`
	BankState                 string    `bun:"bank_state,type:varchar(2),nullzero"`
	BankPostalCode            string    `bun:"bank_postal_code,type:varchar(10),nullzero"`
	RemoteId                  string    `bun:"remote_id,type:varchar(8),nullzero"`
	ApplicationId             string    `bun:"application_id,type:varchar(8),nullzero"`
	BankCountry               string    `bun:"bank_country,type:char(2),nullzero"`
	BankCompanyIdNo           string    `bun:"bank_company_id_no,type:varchar(255),nullzero"`
}
