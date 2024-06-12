package model

type AssemblyHdr struct {
	bun.BaseModel                `bun:"table:assembly_hdr"`
	Backflush                    string    `bun:"backflush,type:char"`
	AssemblyDays                 float64   `bun:"assembly_days,type:decimal(4,1)"`
	DisassemblyDays              float64   `bun:"disassembly_days,type:decimal(4,1)"`
	Labor                        float64   `bun:"labor,type:decimal(19,9)"`
	LaborCostMultiplier          float64   `bun:"labor_cost_multiplier,type:decimal(19,9)"`
	Class1                       string    `bun:"class1,type:varchar(255),nullzero"`
	Class2                       string    `bun:"class2,type:varchar(255),nullzero"`
	Class3                       string    `bun:"class3,type:varchar(255),nullzero"`
	Class4                       string    `bun:"class4,type:varchar(255),nullzero"`
	Class5                       string    `bun:"class5,type:varchar(255),nullzero"`
	ShipPartialQuantity          string    `bun:"ship_partial_quantity,type:char"`
	AllowOeAddComponents         string    `bun:"allow_oe_add_components,type:char"`
	PrintCompntsOnPickTicket     string    `bun:"print_compnts_on_pick_ticket,type:char"`
	PrintComponentsOnInvoice     string    `bun:"print_components_on_invoice,type:char"`
	DefaultDisposition           string    `bun:"default_disposition,type:char"`
	ProductionOrderProcessing    string    `bun:"production_order_processing,type:char"`
	OverallAssemblyLength        float64   `bun:"overall_assembly_length,type:decimal(19,4)"`
	DeleteFlag                   string    `bun:"delete_flag,type:char"`
	DateCreated                  time.Time `bun:"date_created,type:datetime"`
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	ShipIncompleteAssemblies     string    `bun:"ship_incomplete_assemblies,type:char"`
	AutoCreateProdOrder          string    `bun:"auto_create_prod_order,type:char,default:('Y')"`
	PricingOption                int16     `bun:"pricing_option,type:tinyint,default:(3)"`
	PrintIncompleteAssemblies    string    `bun:"print_incomplete_assemblies,type:char,default:('N')"`
	InvMastUid                   int32     `bun:"inv_mast_uid,type:int,pk"`
	AssemblyForStock             string    `bun:"assembly_for_stock,type:char,default:('N')"`
	AssemblyUsageAccumulation    string    `bun:"assembly_usage_accumulation,type:char,default:('A')"`
	CalcComponentServiceLevel    string    `bun:"calc_component_service_level,type:char,default:('N')"`
	ExtendedDesc                 string    `bun:"extended_desc,type:varchar(255),nullzero"`
	AllowDisassembly             string    `bun:"allow_disassembly,type:char,default:('N')"`
	CostOfDisassembly            float64   `bun:"cost_of_disassembly,type:decimal(19,9),default:(0)"`
	ProrateCostBy                string    `bun:"prorate_cost_by,type:char,default:('A')"`
	PrintReturnToStockComponents string    `bun:"print_return_to_stock_components,type:char,default:('Y')"`
	PriceReturnToStockComponents string    `bun:"price_return_to_stock_components,type:char,default:('Y')"`
	HoseAssemblyFlag             string    `bun:"hose_assembly_flag,type:char,default:('N')"`
	HoseOverallLength            float64   `bun:"hose_overall_length,type:decimal(19,9),default:(0)"`
	HoseUnitOfMeasure            string    `bun:"hose_unit_of_measure,type:varchar(8),nullzero"`
	CreatedBy                    string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	Type                         int32     `bun:"type,type:int,nullzero"`
	ProcessUid                   int32     `bun:"process_uid,type:int,nullzero"`
	PrintTravelerWithProdorder   string    `bun:"print_traveler_with_prodorder,type:varchar,default:('N')"`
	DoNotBulkEditFlag            string    `bun:"do_not_bulk_edit_flag,type:char,default:('N')"`
	AutoExpandFlag               string    `bun:"auto_expand_flag,type:char,nullzero"`
	BypassOeProdOrderProcessing  string    `bun:"bypass_oe_prod_order_processing,type:char,nullzero"`
	AssemblyClassUid             int32     `bun:"assembly_class_uid,type:int,nullzero"`
	PackBy                       string    `bun:"pack_by,type:char,default:('H')"`
	MinutesToMake                int32     `bun:"minutes_to_make,type:int"`
	IncludeComponentsOnEdi856    string    `bun:"include_components_on_edi_856,type:char,default:('N')"`
}
