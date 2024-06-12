package model

import (
	"github.com/uptrace/bun"
	"time"
)

type OeSystemParameters struct {
	bun.BaseModel                 `bun:"table:oe_system_parameters"`
	UseShippingGroups             string    `bun:"use_shipping_groups,type:char"`
	DateCreated                   time.Time `bun:"date_created,type:datetime"`
	DateLastModified              time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy              string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	CommissionCostFlag            string    `bun:"commission_cost_flag,type:char"`
	OtherCost                     string    `bun:"other_cost,type:char"`
	OtherCostSource               string    `bun:"other_cost_source,type:varchar(4),nullzero"`
	DownpaymentOption             string    `bun:"downpayment_option,type:varchar(4),nullzero"`
	RecalculatePrice              string    `bun:"recalculate_price,type:char,nullzero"`
	PricingUomIsOrderUom          string    `bun:"pricing_uom_is_order_uom,type:varchar,nullzero"`
	PrintPickTicket               string    `bun:"print_pick_ticket,type:char"`
	AllOn1stPickHold              string    `bun:"all_on_1st_pick_hold,type:char,nullzero"`
	SuppressPickAndHold           string    `bun:"suppress_pick_and_hold,type:char"`
	OeSystemParameterUid          string    `bun:"oe_system_parameter_uid,type:numeric,pk"`
	AllowSourceShipToDiffer       string    `bun:"allow_source_ship_to_differ,type:char"`
	ChangeSourceLocViaPopup       string    `bun:"change_source_loc_via_popup,type:char"`
	SuppressBackorderOnPickticket string    `bun:"suppress_backorder_on_pickticket,type:char,default:('N')"`
}
