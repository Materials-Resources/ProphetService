package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type JobPriceLine struct {
	bun.BaseModel               `bun:"table:job_price_line"`
	JobPriceLineUid             int32      `bun:"job_price_line_uid,type:int,pk"`                             // System generated. Uniquely distinguishes each row.
	JobPriceHdrUid              int32      `bun:"job_price_hdr_uid,type:int"`                                 // Link to system generated identifier for job_price_hdr
	InvMastUid                  *int32     `bun:"inv_mast_uid,type:int"`                                      // Unique identifier for the item id.
	Uom                         *string    `bun:"uom,type:varchar(8)"`                                        // Unit of measure for the qty ordered
	UnitSize                    *float64   `bun:"unit_size,type:decimal(19,4)"`                               // Unit size for the qty ordered
	PricingMethod               *int16     `bun:"pricing_method,type:smallint"`                               // Could be source, price Or none
	SourcePrice                 *int16     `bun:"source_price,type:smallint"`                                 // Could be price 1, Price 2, supplier list price… multiplier Number by which source price should be multiplied
	Price                       *float64   `bun:"price,type:decimal(19,9)"`                                   // Indicates the price of the item.
	QtyOrdered                  *float64   `bun:"qty_ordered,type:decimal(19,9)"`                             // Number of qty ordered
	QtyMaximum                  *float64   `bun:"qty_maximum,type:decimal(19,9)"`                             // Quantity Cap
	OtherCostTypeCd             *int16     `bun:"other_cost_type_cd,type:smallint"`                           // Could be Order, Source, value Or None
	OtherCostValue              *float64   `bun:"other_cost_value,type:decimal(19,9)"`                        // Indicate the Value for the other cost
	OtherCostSourceCd           *int16     `bun:"other_cost_source_cd,type:smallint"`                         // Source for the other cost (price 1, Pr 2…)
	OtherCostCalcMethodCd       *int16     `bun:"other_cost_calc_method_cd,type:smallint"`                    // Could be by difference, mark up, Multiplier Or %
	OtherCostCalcValue          *float64   `bun:"other_cost_calc_value,type:decimal(19,9)"`                   // Multiplication factor.
	CommissionCostTypeCd        *int16     `bun:"commission_cost_type_cd,type:smallint"`                      // Could be Order, Source, value Or None
	CommissionCostValue         *float64   `bun:"commission_cost_value,type:decimal(19,9)"`                   // Indicate the Value for the Comm cost
	CommissionCostSourceCd      *int16     `bun:"commission_cost_source_cd,type:smallint"`                    // Source for the comm cost (price 1, Pr 2…)
	CommissionCostCalcMethodCd  *int16     `bun:"commission_cost_calc_method_cd,type:smallint"`               // Could be by difference, mark up, Multiplier Or %
	CommissionCostCalcValue     *float64   `bun:"commission_cost_calc_value,type:decimal(19,9)"`              // Multiplication factor.
	RowStatusFlag               int32      `bun:"row_status_flag,type:int"`                                   // Indicates current record status.
	DateLastModified            time.Time  `bun:"date_last_modified,type:datetime"`                           // Indicates the date/time this record was last modified.
	DateCreated                 time.Time  `bun:"date_created,type:datetime"`                                 // Indicates the date/time this record was created.
	LastMaintainedBy            string     `bun:"last_maintained_by,type:varchar(30)"`                        // ID of the user who last maintained this record
	Multiplier                  *float64   `bun:"multiplier,type:decimal(19,9)"`                              // Determines what number to multiply the cost by.
	LineNo                      int32      `bun:"line_no,type:int"`                                           // Unique Line identifier for this Contract Item.
	SourceLocationId            *float64   `bun:"source_location_id,type:decimal(19,0)"`                      // Indicates where this Contract Item is Sourced from.
	CustomerPartNo              *string    `bun:"customer_part_no,type:varchar(40)"`                          // Indicates Item on the Contract Line is AKA.
	ExpirationDate              *time.Time `bun:"expiration_date,type:datetime"`                              // The date when this item expires from this contract
	VendorAuthNo                *string    `bun:"vendor_auth_no,type:varchar(255)"`                           // vendor authorization number for this contract item
	PoCost                      *float64   `bun:"po_cost,type:decimal(19,9)"`                                 // The PO cost for the item if we have a DS or Special during order entry
	DiscountGroupId             *string    `bun:"discount_group_id,type:varchar(8)"`                          // Discount group associated with the contract line.
	CustPoNo                    *string    `bun:"cust_po_no,type:varchar(255)"`                               // Customer Purchase Order number for this item record
	ItemCategoryUid             *int32     `bun:"item_category_uid,type:int"`                                 // Links the item to a category.
	SubCategoryUid              *int32     `bun:"sub_category_uid,type:int"`                                  // Links the item to a sub-category.
	ProductGroupId              *string    `bun:"product_group_id,type:varchar(8)"`                           // Custom - F23038: FK to product_group.product_group_id.  Link to associated product group instance.
	CurrencyId                  *float64   `bun:"currency_id,type:decimal(19,0)"`                             // Job/Contract Item currency identifier for the table.
	TerminalId                  *float64   `bun:"terminal_id,type:decimal(19,0)"`                             // The terminal ID used to retrieve this price for this item
	ContractLineCostPageUid     *int32     `bun:"contract_line_cost_page_uid,type:int"`                       // Identifies the associated cost page for the contract line.
	ContractLinePricePageUid    *int32     `bun:"contract_line_price_page_uid,type:int"`                      // Identifies the associated price page for the contract line.
	BudgetCd                    *string    `bun:"budget_cd,type:varchar(255)"`                                // Custom (F34176): User defined budget code.
	ItemRevisionUid             *int32     `bun:"item_revision_uid,type:int"`                                 // Column holds item revision uid of item.
	PickFee                     *float64   `bun:"pick_fee,type:decimal(19,9)"`                                // Custom column to specify the addictional fee for the item
	LineStartDate               *time.Time `bun:"line_start_date,type:datetime"`                              // Start date at the line level for the item on the contract.
	InitialCommitmentAmount     *float64   `bun:"initial_commitment_amount,type:decimal(19,9),default:((0))"` // Inital amount committed to the contract line
	CommitmentAmount            *float64   `bun:"commitment_amount,type:decimal(19,9),default:((0))"`         // Amount committed to the contract line
	TotalCommitmentAmount       *float64   `bun:"total_commitment_amount,type:decimal(19,9),default:((0))"`   // Total amount committed to the contract line
	CadPurchaseCost             *float64   `bun:"cad_purchase_cost,type:decimal(19,9)"`                       // Determines the purchase price at which the item will be purchased in PORG or Purchase Order Entry if purchased under the contract.
	StartingVirtualInventoryQty *float64   `bun:"starting_virtual_inventory_qty,type:decimal(19,9)"`          // Specifies the starting balance for virtual inventory for the item on the CAD or CHA contract, so that contracts imported from a legacy system will have their current virtual inventory represented accurately.
	SnapshotCost                *float64   `bun:"snapshot_cost,type:decimal(19,9)"`                           // (Custom) Commission cost for this line item from the last review of this contract.
	MinimumMccCode              *string    `bun:"minimum_mcc_code,type:varchar(255)"`                         // Minimum Material Classification Codes for the item on this contract line.
	CommissionClassId           *string    `bun:"commission_class_id,type:varchar(8)"`                        // Unique Identifier for the commission class
	CommissionOverridePercent   *float64   `bun:"commission_override_percent,type:decimal(19,9)"`             // Commission override percent to be set on oe_line_salesrep for any sales order line for this item that is priced by this contract.
	AllDiscountGroupsFlag       string     `bun:"all_discount_groups_flag,type:char(1),default:('N')"`        // Custom (F63764): determines if this record applies to all sales discount groups
	PocostingMethod             *int16     `bun:"pocosting_method,type:smallint"`                             // PO Cost Multiplier Tab po costing method
	PocostingSourcePoCostCd     *int16     `bun:"pocosting_source_po_cost_cd,type:smallint"`                  // PO Cost Multiplier Tab po costing po cost source cd
	PocostingMultiplier         *float64   `bun:"pocosting_multiplier,type:decimal(19,9)"`                    // PO Cost Multiplier Tab po costing po po cost multiplier
	PocostingPoCost             *float64   `bun:"pocosting_po_cost,type:decimal(19,9)"`                       // PO Cost Multiplier Tab po cost
	PocostingPoContractNumber   *string    `bun:"pocosting_po_contract_number,type:varchar(40)"`              // PO Cost Multiplier Tab po contract no
	DefaultDisposition          *string    `bun:"default_disposition,type:char(1)"`                           // Custom column to indicates the what Default Disposition should be used when the material is out of stock for the line
}
