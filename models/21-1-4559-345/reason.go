package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Reason struct {
	bun.BaseModel           `bun:"table:reason"`
	Id                      string    `bun:"id,type:varchar(5),pk"`                                         // This column is unused.
	Reason                  string    `bun:"reason,type:varchar(40)"`                                       // What is the human-readable text of this reason?
	DeleteFlag              string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	GlAccountNo             string    `bun:"gl_account_no,type:varchar(32),nullzero"`                       // General Ledger Account number
	CompanyNo               string    `bun:"company_no,type:varchar(8),nullzero"`                           // Unique code that identifies a company.
	AffectActualUsageFlag   string    `bun:"affect_actual_usage_flag,type:char(1),default:('N')"`           // Indicates whether inventory adjusments with this reason affect usage
	DeliveryReasonFlag      string    `bun:"delivery_reason_flag,type:char(1),default:('N')"`               // Indicates this reason is appropriate for explaining a delivery shortage
	UseBranchAccFlag        string    `bun:"use_branch_acc_flag,type:char(1),default:('Y')"`                // Use Branch in Account
	ArCashReasonFlag        string    `bun:"ar_cash_reason_flag,type:char(1),default:('N')"`                // Indicates this reason is approriate for explaining zero dollar invoices created.
	MacUpdateReasonFlag     string    `bun:"mac_update_reason_flag,type:char(1),default:('N')"`             // Indicates a reason code that will be used for MAC adjustments
	RmaReasonFlag           string    `bun:"rma_reason_flag,type:char(1),default:('N')"`                    // Custom: Indicates a reason code will be used in RMA Entry and RMA Receipts
	FieldDestroyReasonFlag  string    `bun:"field_destroy_reason_flag,type:char(1),default:('N')"`          // Custom: Indicates a reason code will be used for Field Destroy requests on B2B website.
	InvAdjItemLevelFlag     string    `bun:"inv_adj_item_level_flag,type:char(1),default:('N')"`            // Custom (F65493): determines if this reason appllies to inventory adjustment item level only
	InvAdjItemLevelEditFlag string    `bun:"inv_adj_item_level_edit_flag,type:char(1),default:('N')"`       // Custom (F65493): for physical counts, when 	this reason is specified as the adjustment header reason, determines if item level reasons can be specified
	ReturnToStockFlag       string    `bun:"return_to_stock_flag,type:char(1),default:('N')"`               // Determines if a different action should occur for returned items
	ReturnAction            int32     `bun:"return_action,type:int,nullzero"`                               // Indicates the action to perform for returned items 1- Allocate, 2-Cancel
	MafSurchargeReasonFlag  string    `bun:"maf_surcharge_reason_flag,type:char(1),default:('N')"`          // Use the field to enter records to be used in OE MAF field
}
