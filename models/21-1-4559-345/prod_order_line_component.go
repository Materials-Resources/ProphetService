package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdOrderLineComponent struct {
	bun.BaseModel             `bun:"table:prod_order_line_component"`
	ProdOrderLineComponentUid int32     `bun:"prod_order_line_component_uid,type:int,pk,default:(0)"`
	ProdOrderNumber           float64   `bun:"prod_order_number,type:decimal(19,0)"`
	LineNumber                float64   `bun:"line_number,type:decimal(19,0)"`
	ComponentNumber           float64   `bun:"component_number,type:decimal(19,0)"`
	CompanyId                 string    `bun:"company_id,type:varchar(8)"`
	SourceLocationId          float64   `bun:"source_location_id,type:decimal(19,0)"`
	ExtendedDesc              string    `bun:"extended_desc,type:varchar(255),nullzero"`
	QtyRequested              float64   `bun:"qty_requested,type:decimal(19,9)"`
	QtyAllocated              float64   `bun:"qty_allocated,type:decimal(19,9)"`
	QtyCanceled               float64   `bun:"qty_canceled,type:decimal(19,9),nullzero"`
	QtyUsed                   float64   `bun:"qty_used,type:decimal(19,9),nullzero"`
	UnitOfMeasure             string    `bun:"unit_of_measure,type:varchar(8)"`
	UnitSize                  float64   `bun:"unit_size,type:decimal(19,4)"`
	Disposition               string    `bun:"disposition,type:char,nullzero"`
	SupplierId                float64   `bun:"supplier_id,type:decimal(19,0)"`
	DivisionId                float64   `bun:"division_id,type:decimal(19,0)"`
	Taxable                   string    `bun:"taxable,type:char"`
	OtherCharge               string    `bun:"other_charge,type:char"`
	UnitCost                  float64   `bun:"unit_cost,type:decimal(19,9)"`
	PoCost                    float64   `bun:"po_cost,type:decimal(19,9),nullzero"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30)"`
	UnitsRequested            float64   `bun:"units_requested,type:decimal(19,9)"`
	SubAssembly               string    `bun:"sub_assembly,type:char"`
	QtyPerAssembly            float64   `bun:"qty_per_assembly,type:decimal(19,9)"`
	Complete                  string    `bun:"complete,type:char"`
	InvMastUid                int32     `bun:"inv_mast_uid,type:int"`
	CompExtendedDesc          string    `bun:"comp_extended_desc,type:varchar(255),nullzero"`
	ComponentCutLength        float64   `bun:"component_cut_length,type:decimal(19,9),default:(0)"`
	CutLengthEditedFlag       string    `bun:"cut_length_edited_flag,type:char,default:('N')"`
	QtyNeeded                 float64   `bun:"qty_needed,type:decimal(19,9),default:(0)"`
	UsedSpecificCostFlag      string    `bun:"used_specific_cost_flag,type:char,default:('N')"`
	QtyOnPickTickets          float64   `bun:"qty_on_pick_tickets,type:decimal(19,9),default:((0))"`
	QtyConfirmed              float64   `bun:"qty_confirmed,type:decimal(19,9),default:((0))"`
	RequiredDate              time.Time `bun:"required_date,type:datetime,nullzero"`
	ExpediteDate              time.Time `bun:"expedite_date,type:datetime,nullzero"`
	ServiceLaborUid           int32     `bun:"service_labor_uid,type:int,nullzero"`
	OperationUid              int32     `bun:"operation_uid,type:int,nullzero"`
	InventoryCost             float64   `bun:"inventory_cost,type:decimal(19,9),nullzero"`
	NewCost                   float64   `bun:"new_cost,type:decimal(19,9),nullzero"`
	RetrievedByWms            string    `bun:"retrieved_by_wms,type:char,default:('N')"`
	OriginalCompletionDate    time.Time `bun:"original_completion_date,type:datetime,nullzero"`
	CompletionDate            time.Time `bun:"completion_date,type:datetime,nullzero"`
	BackflushFlag             string    `bun:"backflush_flag,type:char,default:('N')"`
	QtyScrapped               float64   `bun:"qty_scrapped,type:decimal(19,9),nullzero"`
	CountryOfOrigin           string    `bun:"country_of_origin,type:varchar(255),nullzero"`
	BeltingLength             float64   `bun:"belting_length,type:decimal(19,9),nullzero"`
	BeltingWidth              float64   `bun:"belting_width,type:decimal(19,9),nullzero"`
	UomConvertedFlag          string    `bun:"uom_converted_flag,type:char,nullzero"`
	LooseShipFlag             string    `bun:"loose_ship_flag,type:char,nullzero"`
	ExtendedItemFlag          string    `bun:"extended_item_flag,type:char,nullzero"`
	UnitCostEdited            string    `bun:"unit_cost_edited,type:char,default:('N')"`
	LaborTypeCd               int32     `bun:"labor_type_cd,type:int,nullzero"`
}
