package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type SystemParameters struct {
	bun.BaseModel               `bun:"table:system_parameters"`
	SecurityActive              string    `bun:"security_active,type:char(1)"`
	CompSecActive               string    `bun:"comp_sec_active,type:char(1)"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastModifiedBy              string    `bun:"last_modified_by,type:varchar(30)"`
	HelpDirectory               string    `bun:"help_directory,type:varchar(255),nullzero"`
	NoOfRetries                 float64   `bun:"no_of_retries,type:decimal(19,0),nullzero"`
	SetUpMode                   string    `bun:"set_up_mode,type:char(1),nullzero"`
	InventoryModuleInstalled    string    `bun:"inventory_module_installed,type:char(1),nullzero"`
	UsedInUs                    string    `bun:"used_in_us,type:char(1),nullzero"`
	ForeignCurrency             string    `bun:"foreign_currency,type:char(1),nullzero"`
	RevenueByItem               string    `bun:"revenue_by_item,type:char(1)"`
	ApModuleInstalled           string    `bun:"ap_module_installed,type:char(1)"`
	ArModuleInstalled           string    `bun:"ar_module_installed,type:char(1)"`
	OeModuleInstalled           string    `bun:"oe_module_installed,type:char(1)"`
	PoModuleInstalled           string    `bun:"po_module_installed,type:char(1)"`
	ProdOrdModuleInstalled      string    `bun:"prod_ord_module_installed,type:char(1)"`
	JcModuleInstalled           string    `bun:"jc_module_installed,type:char(1)"`
	MultiLingual                string    `bun:"multi_lingual,type:char(1)"`
	CurrencyMask                string    `bun:"currency_mask,type:varchar(50),nullzero"`
	NumberOfSigs                string    `bun:"number_of_sigs,type:char(1)"`
	Sig1Path                    string    `bun:"sig1_path,type:varchar(255),nullzero"`
	Sig2Path                    string    `bun:"sig2_path,type:varchar(255),nullzero"`
	PrinterDpi                  string    `bun:"printer_dpi,type:varchar(3)"`
	DefaultPeriod               string    `bun:"default_period,type:char(1)"`
	DefaultYear                 string    `bun:"default_year,type:char(1)"`
	CoaMaskCompliance           string    `bun:"coa_mask_compliance,type:char(1)"`
	AmountReqTwoSigs            float64   `bun:"amount_req_two_sigs,type:decimal(19,4),nullzero"`
	ValidateCustContactsOnOe    string    `bun:"validate_cust_contacts_on_oe,type:char(1)"`
	DefaultBackorderQty         string    `bun:"default_backorder_qty,type:char(1),nullzero"`
	CalcBackorderInOe           string    `bun:"calc_backorder_in_oe,type:char(1)"`
	PrintDetailOnCheckstub      string    `bun:"print_detail_on_checkstub,type:char(1)"`
	GlIntercompanyFunctionality string    `bun:"gl_intercompany_functionality,type:char(1)"`
	ApIntercompanyFunctionality string    `bun:"ap_intercompany_functionality,type:char(1)"`
	DatabaseVersion             string    `bun:"database_version,type:varchar(8),nullzero"`
	FireGlTrigger               string    `bun:"fire_gl_trigger,type:char(1),nullzero"`
	LogoPath                    string    `bun:"logo_path,type:varchar(255),nullzero"`
	NoOfCycleDays               int32     `bun:"no_of_cycle_days,type:int,nullzero"`
	UseBillingAddress           string    `bun:"use_billing_address,type:char(1),nullzero"`
	UsePayableGroups            string    `bun:"use_payable_groups,type:char(1),nullzero"`
	UseReceivableGroups         string    `bun:"use_receivable_groups,type:char(1),nullzero"`
	CommissionModuleInstalled   string    `bun:"commission_module_installed,type:char(1),nullzero"`
	UseEnterKey                 string    `bun:"use_enter_key,type:char(1),nullzero"`
	CharacterSensitivePopups    string    `bun:"character_sensitive_popups,type:char(1),nullzero"`
	DefaultPreviousPeriodYear   string    `bun:"default_previous_period_year,type:char(1),nullzero"`
	DisplayUpdateWasSuccessful  string    `bun:"display_update_was_successful,type:char(1),nullzero"`
	BuyerOnlyToPlacePo          string    `bun:"buyer_only_to_place_po,type:char(1),nullzero"`
	UsingApprovals              string    `bun:"using_approvals,type:char(1),nullzero"`
	SalesPricingMethod          string    `bun:"sales_pricing_method,type:varchar(2)"`
	TabSecActive                string    `bun:"tab_sec_active,type:char(1),nullzero"`
	BranchSecActive             string    `bun:"branch_sec_active,type:char(1),nullzero"`
	Reserved1                   float64   `bun:"reserved1,type:decimal(1,0),nullzero"`
	JobIdMask                   string    `bun:"job_id_mask,type:varchar(64),nullzero"`
	CharacterSpecificPopups     string    `bun:"character_specific_popups,type:char(1),nullzero"`
	NumOfDecPlacesUnitPrice     float64   `bun:"num_of_dec_places_unit_price,type:decimal(1,0)"`
	NumOfDecPlacesUnitCost      float64   `bun:"num_of_dec_places_unit_cost,type:decimal(1,0)"`
	MinCharForCharSpecific      float64   `bun:"min_char_for_char_specific,type:decimal(2,0),nullzero"`
	KeystrokeForPopups          string    `bun:"keystroke_for_popups,type:varchar(10),nullzero"`
	KeystrokeForNewLineItem     string    `bun:"keystroke_for_new_line_item,type:varchar(10),nullzero"`
	MailModuleInstalled         string    `bun:"mail_module_installed,type:char(1),nullzero"`
	SecurityMethod              string    `bun:"security_method,type:char(1)"`
	SystemParameterUid          string    `bun:"system_parameter_uid,type:numeric(19,0),pk"`
	OePrintQty                  float64   `bun:"oe_print_qty,type:decimal(5,0)"`
	PickTicketPrintQty          float64   `bun:"pick_ticket_print_qty,type:decimal(5,0)"`
	InvoicePrintQty             float64   `bun:"invoice_print_qty,type:decimal(5,0)"`
	PoPrintQty                  float64   `bun:"po_print_qty,type:decimal(5,0)"`
	TransferPrintQty            float64   `bun:"transfer_print_qty,type:decimal(5,0)"`
	StatementPrintQty           float64   `bun:"statement_print_qty,type:decimal(5,0)"`
	QuotePrintQty               float64   `bun:"quote_print_qty,type:decimal(5,0)"`
	CrystalDirectory            string    `bun:"crystal_directory,type:varchar(255),nullzero"`
	EnableFax                   string    `bun:"enable_fax,type:char(1),default:('N')"`
	FaxFilePath                 string    `bun:"fax_file_path,type:varchar(50),nullzero"`
	VoucherEntryApprovals       string    `bun:"voucher_entry_approvals,type:char(1),default:('N')"`
	CrDrMemoEntryApprovals      string    `bun:"cr_dr_memo_entry_approvals,type:char(1),default:('N')"`
	InvoiceEntryApprovals       string    `bun:"invoice_entry_approvals,type:char(1),default:('N')"`
	MiscCashReceiptsApprovals   string    `bun:"misc_cash_receipts_approvals,type:char(1),default:('N')"`
	CashReceiptsApprovals       string    `bun:"cash_receipts_approvals,type:char(1),default:('N')"`
	JournalEntriesApprovals     string    `bun:"journal_entries_approvals,type:char(1),default:('N')"`
	InvTransferRecApprovals     string    `bun:"inv_transfer_rec_approvals,type:char(1),default:('N')"`
	InvAdjApprovals             string    `bun:"inv_adj_approvals,type:char(1),default:('N')"`
	TransferEntryApprovals      string    `bun:"transfer_entry_approvals,type:char(1),default:('N')"`
	PhysicalCountApprovals      string    `bun:"physical_count_approvals,type:char(1),default:('N')"`
	OrderEntryApprovals         string    `bun:"order_entry_approvals,type:char(1),default:('N')"`
	RmaEntryApprovals           string    `bun:"rma_entry_approvals,type:char(1),default:('N')"`
	ProdOrderEntryApprovals     string    `bun:"prod_order_entry_approvals,type:char(1),default:('N')"`
	ProcessProdOrderApprovals   string    `bun:"process_prod_order_approvals,type:char(1),default:('N')"`
	PurchaseOrderEntryApprovals string    `bun:"purchase_order_entry_approvals,type:char(1),default:('N')"`
	ConvertPoVoucherApprovals   string    `bun:"convert_po_voucher_approvals,type:char(1),default:('N')"`
	RepetitiveJrnEntryApprovals string    `bun:"repetitive_jrn_entry_approvals,type:char(1),default:('N')"`
	ConfigurationId             int32     `bun:"configuration_id,type:int,default:((-1))"`
	MessageLogFilePath          string    `bun:"message_log_file_path,type:varchar(255),default:('c:\\msglog.txt')"`
	ScriptPath                  string    `bun:"script_path,type:varchar(255),default:(' ')"`
}
