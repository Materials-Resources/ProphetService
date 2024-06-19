package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ApinvLine struct {
	bun.BaseModel         `bun:"table:apinv_line"`
	VoucherNo             string    `bun:"voucher_no,type:varchar(10)"`                   // Unique identifier, system generated number for a voucher
	Quantity              float64   `bun:"quantity,type:decimal(19,9)"`                   // Quantity on the voucher
	ItemId                string    `bun:"item_id,type:varchar(40),nullzero"`             // What material is held in this bin?
	UnitPrice             float64   `bun:"unit_price,type:decimal(19,9)"`                 // What is the unit price for this line item?
	PurchaseAmount        float64   `bun:"purchase_amount,type:decimal(19,4)"`            // Purchase amount per line item
	PurchaseAccount       string    `bun:"purchase_account,type:varchar(32)"`             // Enter the purchase account
	Description           string    `bun:"description,type:varchar(30),nullzero"`         // How would you describe this repeating journal entry?
	TypeId                string    `bun:"type_id,type:varchar(8),nullzero"`              // Enter the type ID
	Ten99Type             float64   `bun:"ten99_type,type:decimal(19,0)"`                 // Ten99 Type Amount
	DateCreated           time.Time `bun:"date_created,type:datetime"`                    // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`              // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`           // ID of the user who last maintained this record
	CompanyNo             string    `bun:"company_no,type:varchar(8)"`                    // Unique code that identifies a company.
	HomeCurrencyAmt       float64   `bun:"home_currency_amt,type:decimal(19,4),nullzero"` // Please Enter The Home Currency Amount.
	ProjectId             string    `bun:"project_id,type:varchar(32),nullzero"`          // Enter the project ID
	JobId                 string    `bun:"job_id,type:varchar(32),nullzero"`              // Job associated with the line item
	BranchId              string    `bun:"branch_id,type:varchar(8),nullzero"`            // What branch issued this Purchase Order?
	UnitPriceDisplay      float64   `bun:"unit_price_display,type:decimal(19,9)"`         // Holds the values that are displayed on the window.
	PurchaseAmountDisplay float64   `bun:"purchase_amount_display,type:decimal(19,4)"`    // Holds the values that are displayed on the window.
	ApinvLineUid          int32     `bun:"apinv_line_uid,type:int,default:(0),pk"`        // UID for the table
	InvMastUid            int32     `bun:"inv_mast_uid,type:int,nullzero"`                // Unique identifier for valid inventory item
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DisputedFlag          string    `bun:"disputed_flag,type:char(1),default:('N')"`           // Indicates whether the voucher is in dispute status.
	DisputedAmt           float64   `bun:"disputed_amt,type:decimal(19,4),default:(0)"`        // The amount disputed on the voucher line
	OriginalDisputedAmt   float64   `bun:"original_disputed_amt,type:decimal(19,4),nullzero"`  // Original disputed value of the voucher line
	ApinvLineTypeCd       int32     `bun:"apinv_line_type_cd,type:int,nullzero"`               // Type code for the apinv_line record - Regular, Cost Variance, Quantity Variance, and Freight Variance.
	ParentApinvLineUid    int32     `bun:"parent_apinv_line_uid,type:int,nullzero"`            // Groups apinv line records together (groups variance records with the apinv_line record they are associated with)
	PoLineUid             int32     `bun:"po_line_uid,type:int,nullzero"`                      // surrogate key to table po_line
	SourceNo              float64   `bun:"source_no,type:decimal(19,9),nullzero"`              // Source record from which this line was created.
	SequenceNumber        int32     `bun:"sequence_number,type:int,nullzero"`                  // Sequence Number
	PayableSourceUid      int32     `bun:"payable_source_uid,type:int,nullzero"`               // Transaction Detail Associated with the Voucher Line (Landed Cost Driver, Inventory Receipt, etc...)
	FreightAmount         float64   `bun:"freight_amount,type:decimal(19,4),nullzero"`         // Freight amount for the voucher (Home)
	FreightAmountDisplay  float64   `bun:"freight_amount_display,type:decimal(19,4),nullzero"` // Freight amount for the voucher (Foreign)
	GlDimenTypeUid        int32     `bun:"gl_dimen_type_uid,type:int,nullzero"`                // UID for GL dimension type
	GlDimensionId         string    `bun:"gl_dimension_id,type:varchar(255),nullzero"`         // Value for dimension type
	GlDimensionDesc       string    `bun:"gl_dimension_desc,type:varchar(255),nullzero"`       // Desc of dimension type
	PercentAllocated      float64   `bun:"percent_allocated,type:decimal(19,4),nullzero"`      // percentage of total amount for the line
}
