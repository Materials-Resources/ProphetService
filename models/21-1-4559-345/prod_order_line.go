package model

type ProdOrderLine struct {
	bun.BaseModel                `bun:"table:prod_order_line"`
	ProdOrderNumber              float64   `bun:"prod_order_number,type:decimal(19,0),pk"`
	LineNumber                   float64   `bun:"line_number,type:decimal(19,0),pk"`
	QtyToMake                    float64   `bun:"qty_to_make,type:decimal(19,9)"`
	QtyCompleted                 float64   `bun:"qty_completed,type:decimal(19,9)"`
	Completed                    string    `bun:"completed,type:char"`
	Cancel                       string    `bun:"cancel,type:char"`
	Labor                        float64   `bun:"labor,type:decimal(19,2),nullzero"`
	ShipPartialQuantity          string    `bun:"ship_partial_quantity,type:char"`
	ShipIncompleteAssemblies     string    `bun:"ship_incomplete_assemblies,type:char"`
	PrintCompntsOnPickTicket     string    `bun:"print_compnts_on_pick_ticket,type:char"`
	PrintComponentsOnInvoice     string    `bun:"print_components_on_invoice,type:char"`
	DefaultDisposition           string    `bun:"default_disposition,type:char"`
	OverallAssemblyLength        float64   `bun:"overall_assembly_length,type:decimal(19,9),nullzero"`
	DateCreated                  time.Time `bun:"date_created,type:datetime"`
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	UnitOfMeasure                string    `bun:"unit_of_measure,type:varchar(8),nullzero"`
	UnitSize                     float64   `bun:"unit_size,type:decimal(19,4)"`
	FreeForm                     string    `bun:"free_form,type:char"`
	CumAssemblyCost              float64   `bun:"cum_assembly_cost,type:decimal(19,9),nullzero"`
	PrintIncompleteAssemblies    string    `bun:"print_incomplete_assemblies,type:char,default:('N')"`
	InvMastUid                   int32     `bun:"inv_mast_uid,type:int"`
	AssemblyForStock             string    `bun:"assembly_for_stock,type:char,default:('N')"`
	ExtendedDesc                 string    `bun:"extended_desc,type:varchar(255),nullzero"`
	CostToDisassemble            float64   `bun:"cost_to_disassemble,type:decimal(19,9),default:(0)"`
	PrintReturnToStockComponents string    `bun:"print_return_to_stock_components,type:char,default:('Y')"`
	PriceReturnToStockComponents string    `bun:"price_return_to_stock_components,type:char,default:('Y')"`
	HoseAssemblyFlag             string    `bun:"hose_assembly_flag,type:char,default:('N')"`
	HoseOverallLength            float64   `bun:"hose_overall_length,type:decimal(19,9),default:(0)"`
	HoseUom                      string    `bun:"hose_uom,type:varchar(8),nullzero"`
	CreatedBy                    string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	ExtdCompCost                 float64   `bun:"extd_comp_cost,type:decimal(19,4),nullzero"`
	ExtdCompCostEditedFlag       string    `bun:"extd_comp_cost_edited_flag,type:char,nullzero"`
	DepositBin                   string    `bun:"deposit_bin,type:varchar(10),nullzero"`
	QtyOnPickTickets             float64   `bun:"qty_on_pick_tickets,type:decimal(19,9),default:((0))"`
	QtyConfirmed                 float64   `bun:"qty_confirmed,type:decimal(19,9),default:((0))"`
	RequiredDate                 time.Time `bun:"required_date,type:datetime,nullzero"`
	ExpediteDate                 time.Time `bun:"expedite_date,type:datetime,nullzero"`
	EstimatedFreightCost         float64   `bun:"estimated_freight_cost,type:decimal(19,9),default:((0))"`
	ActualFreightCost            float64   `bun:"actual_freight_cost,type:decimal(19,9),default:((0))"`
	ManualFreightOverideFlag     string    `bun:"manual_freight_overide_flag,type:char,default:('N')"`
	CommittedStartDate           time.Time `bun:"committed_start_date,type:datetime,nullzero"`
	StartDate                    time.Time `bun:"start_date,type:datetime,nullzero"`
	EndDate                      time.Time `bun:"end_date,type:datetime,nullzero"`
	CommentDate                  time.Time `bun:"comment_date,type:datetime,nullzero"`
	ProdOrderLineUid             int32     `bun:"prod_order_line_uid,type:int"`
	DesignCompleteFlag           string    `bun:"design_complete_flag,type:char,nullzero"`
	EstimatedProcessCost         float64   `bun:"estimated_process_cost,type:decimal(19,9),nullzero"`
	CountryOfOrigin              string    `bun:"country_of_origin,type:varchar(255),nullzero"`
}
