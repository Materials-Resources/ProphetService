package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type AssemblyLine struct {
	bun.BaseModel            `bun:"table:assembly_line"`
	Quantity                 float64   `bun:"quantity,type:decimal(19,9)"`                               // Component quantity per assembly.
	DeleteFlag               string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated              time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	SequenceNumber           float64   `bun:"sequence_number,type:decimal(5,0),pk"`                      // Indicates the sequence in which to process the loc
	SubAssembly              string    `bun:"sub_assembly,type:char(1)"`                                 // Indicates whether this component is an assembly itself.
	Taxable                  string    `bun:"taxable,type:char(1)"`                                      // Is this line item taxable?
	UnitOfMeasure            string    `bun:"unit_of_measure,type:varchar(8),nullzero"`                  // What is the unit of measure for this row?
	OtherChargeItem          string    `bun:"other_charge_item,type:char(1)"`                            // Indicates whether the component is a charge rather than material.
	InvMastUid               int32     `bun:"inv_mast_uid,type:int,pk"`                                  // Unique identifier for the item id.
	ComponentInvMastUid      int32     `bun:"component_inv_mast_uid,type:int"`                           // Unique identifier for the component item id.
	ExtendedDesc             string    `bun:"extended_desc,type:varchar(255),nullzero"`                  // Additional description for the assembly component.
	ComponentCutLength       float64   `bun:"component_cut_length,type:decimal(19,9),default:(0)"`       // Length of a hose to be cut
	CutLengthEditedFlag      string    `bun:"cut_length_edited_flag,type:char(1),default:('N')"`         // Indicate if the hose_cut_length column is edited
	QtyNeeded                float64   `bun:"qty_needed,type:decimal(19,9),default:(0)"`                 // The quantity per aaembly for hose and hose sleeve type components.
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	ServiceLaborUid          int32     `bun:"service_labor_uid,type:int,nullzero"`           // Unique identifier of labor associated with this assembly.
	OperationUid             int32     `bun:"operation_uid,type:int,nullzero"`               // Indicates if the line is for material (code: 1578) or for labor (code: 1577).
	DateOffsetDays           int32     `bun:"date_offset_days,type:int,nullzero"`            // Number of days later a component is required than the header Required Date
	AssemblyItemRevisionUid  int32     `bun:"assembly_item_revision_uid,type:int,nullzero"`  // Column holds the revision uid of the assembly item.
	ComponentItemRevisionUid int32     `bun:"component_item_revision_uid,type:int,nullzero"` // Column holds the revision uid of the component item.
	AssemblyLineUid          int32     `bun:"assembly_line_uid,type:int,autoincrement,pk"`
	BackflushFlag            string    `bun:"backflush_flag,type:char(1),default:('N')"`         // Indicate if the component item is a backflush component.
	RefDesignatorLocator     string    `bun:"ref_designator_locator,type:varchar(255),nullzero"` // Custom column at the component line level for extra information
	UserComponentNumber      int32     `bun:"user_component_number,type:int,nullzero"`           // Custom: User defined line number for component
	BeltingNumCuts           int32     `bun:"belting_num_cuts,type:int,nullzero"`                // Number of cuts needed for belting component.
	BeltingLength            float64   `bun:"belting_length,type:decimal(19,9),nullzero"`        // Length of the cut for belting component.
	BeltingWidth             float64   `bun:"belting_width,type:decimal(19,9),nullzero"`         // Width of the cut for belting component.
	LooseShipFlag            string    `bun:"loose_ship_flag,type:char(1),nullzero"`             // Flag to indicate whether an assembly component is a loose ship item.
	MinimumMccCode           string    `bun:"minimum_mcc_code,type:varchar(255),nullzero"`       // Minimum Material Classification Codes for this assembly component.
	ExtendedItemFlag         string    `bun:"extended_item_flag,type:char(1),nullzero"`          // This flag marks if an qty_ordered of the assembly component, should be affected when the assembly item qty is changed
}
