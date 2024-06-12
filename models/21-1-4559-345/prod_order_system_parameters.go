package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdOrderSystemParameters struct {
	bun.BaseModel             `bun:"table:prod_order_system_parameters"`
	AllowMultipleAssemblies   string    `bun:"allow_multiple_assemblies,type:char"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	LaborCostMultiplier       float64   `bun:"labor_cost_multiplier,type:decimal(19,4),nullzero"`
	ShipPartialQuantity       string    `bun:"ship_partial_quantity,type:char"`
	ShipIncompleteAssemblies  string    `bun:"ship_incomplete_assemblies,type:char"`
	PrintCompntsOnPickTicket  string    `bun:"print_compnts_on_pick_ticket,type:char"`
	PrintComponentsOnInvoice  string    `bun:"print_components_on_invoice,type:char"`
	ProductionOrderProcessing string    `bun:"production_order_processing,type:char"`
	DefaultDisposition        string    `bun:"default_disposition,type:char"`
	OverallAssemblyLength     float64   `bun:"overall_assembly_length,type:decimal(19,4),nullzero"`
	ProdOrderSystemParamUid   float64   `bun:"prod_order_system_param_uid,type:decimal(19,0),pk"`
	AutoCreateProdOrders      string    `bun:"auto_create_prod_orders,type:char"`
	AppendToProdOrders        string    `bun:"append_to_prod_orders,type:char"`
	AllowOeAddComponents      string    `bun:"allow_oe_add_components,type:char"`
	PricingOption             int16     `bun:"pricing_option,type:tinyint,default:(3)"`
}
