package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdOrderHdr struct {
	bun.BaseModel            `bun:"table:prod_order_hdr"`
	ProdOrderNumber          float64   `bun:"prod_order_number,type:decimal(19,0),pk"`                       // Production order number.
	CompanyId                string    `bun:"company_id,type:varchar(8)"`                                    // Unique code that identifies a company.
	SourceLocationId         float64   `bun:"source_location_id,type:decimal(19,0)"`                         // Location ID
	OrderDate                time.Time `bun:"order_date,type:datetime"`                                      // When was this order taken?
	RequiredDate             time.Time `bun:"required_date,type:datetime"`                                   // When is this purchase order line item required by?
	ExpectedCompletionDate   time.Time `bun:"expected_completion_date,type:datetime"`                        // Date the production order should be completed.
	Approved                 string    `bun:"approved,type:char(1)"`                                         // Has this product order been approved?
	Complete                 string    `bun:"complete,type:char(1)"`                                         // Indicates whether the production order has been completed.
	DeleteFlag               string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	Cancel                   string    `bun:"cancel,type:char(1)"`                                           // Indicates whether the production order has been canceled.
	Comment                  string    `bun:"comment,type:varchar(255),nullzero"`                            // Production order comments.
	EnteredBy                string    `bun:"entered_by,type:varchar(30)"`                                   // Id of who entered the production order.
	Period                   float64   `bun:"period,type:decimal(3,0),nullzero"`                             // Which period does this quota apply to?
	YearForPeriod            float64   `bun:"year_for_period,type:decimal(4,0),nullzero"`                    // What year does the period belong to?
	AssembleDisassemble      string    `bun:"assemble_disassemble,type:char(1)"`                             // Determines if this production order is an assembly or disassembly.
	Printed                  string    `bun:"printed,type:char(1)"`                                          // Indicates whether a pick ticket, invoice, shipping paper, transfer form, or production order form has been printed for the shipment or release.
	DateCreated              time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	CreateAutoTransfer       string    `bun:"create_auto_transfer,type:char(1),default:('N')"`               // This option allows you to bypass Order Based Transfers on an individual Production Order basis.
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	RetrievedByWms           string    `bun:"retrieved_by_wms,type:char(1),default:('N')"`
	ExpediteDate             time.Time `bun:"expedite_date,type:datetime,nullzero"`                   // The date in which this production order needs to be expedited in order to get assembly out on time
	EstimatedFreightCost     float64   `bun:"estimated_freight_cost,type:decimal(19,9),nullzero"`     // Indicates estimated freight cost for the production order
	ActualFreightCost        float64   `bun:"actual_freight_cost,type:decimal(19,9),nullzero"`        // Indicates actual freight cost for the production order
	ManualFreightOverideFlag string    `bun:"manual_freight_overide_flag,type:char(1),default:('N')"` // This will be set to Y if the user changed freight.
	ReleaseDate              time.Time `bun:"release_date,type:datetime,nullzero"`                    // The date when the order is released to production
	DrawingsComplete         string    `bun:"drawings_complete,type:char(1),default:('N')"`           // Indicates that the drawings have been completed
	InventoryLocationId      float64   `bun:"inventory_location_id,type:decimal(19,0),nullzero"`      // Default Inventory Location for production order
}
