package model

import (
	"github.com/uptrace/bun"
	"time"
)

type AssemblyLine struct {
	bun.BaseModel            `bun:"table:assembly_line"`
	Quantity                 float64   `bun:"quantity,type:decimal(19,9)"`
	DeleteFlag               string    `bun:"delete_flag,type:char"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	SequenceNumber           float64   `bun:"sequence_number,type:decimal(5,0),pk"`
	SubAssembly              string    `bun:"sub_assembly,type:char"`
	Taxable                  string    `bun:"taxable,type:char"`
	UnitOfMeasure            string    `bun:"unit_of_measure,type:varchar(8),nullzero"`
	OtherChargeItem          string    `bun:"other_charge_item,type:char"`
	InvMastUid               int32     `bun:"inv_mast_uid,type:int,pk"`
	ComponentInvMastUid      int32     `bun:"component_inv_mast_uid,type:int"`
	ExtendedDesc             string    `bun:"extended_desc,type:varchar(255),nullzero"`
	ComponentCutLength       float64   `bun:"component_cut_length,type:decimal(19,9),default:(0)"`
	CutLengthEditedFlag      string    `bun:"cut_length_edited_flag,type:char,default:('N')"`
	QtyNeeded                float64   `bun:"qty_needed,type:decimal(19,9),default:(0)"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	ServiceLaborUid          int32     `bun:"service_labor_uid,type:int,nullzero"`
	OperationUid             int32     `bun:"operation_uid,type:int,nullzero"`
	DateOffsetDays           int32     `bun:"date_offset_days,type:int,nullzero"`
	AssemblyItemRevisionUid  int32     `bun:"assembly_item_revision_uid,type:int,nullzero"`
	ComponentItemRevisionUid int32     `bun:"component_item_revision_uid,type:int,nullzero"`
	AssemblyLineUid          int32     `bun:"assembly_line_uid,type:int,pk,identity"`
	BackflushFlag            string    `bun:"backflush_flag,type:char,default:('N')"`
	RefDesignatorLocator     string    `bun:"ref_designator_locator,type:varchar(255),nullzero"`
	UserComponentNumber      int32     `bun:"user_component_number,type:int,nullzero"`
	BeltingNumCuts           int32     `bun:"belting_num_cuts,type:int,nullzero"`
	BeltingLength            float64   `bun:"belting_length,type:decimal(19,9),nullzero"`
	BeltingWidth             float64   `bun:"belting_width,type:decimal(19,9),nullzero"`
	LooseShipFlag            string    `bun:"loose_ship_flag,type:char,nullzero"`
	MinimumMccCode           string    `bun:"minimum_mcc_code,type:varchar(255),nullzero"`
	ExtendedItemFlag         string    `bun:"extended_item_flag,type:char,nullzero"`
}
