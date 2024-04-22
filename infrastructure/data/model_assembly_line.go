package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type AssemblyLine struct {
	bun.BaseModel            `bun:"table:assembly_line"`
	Quantity                 float64         `bun:"quantity,type:decimal(19,9)"`
	DeleteFlag               string          `bun:"delete_flag,type:char"`
	DateCreated              time.Time       `bun:"date_created,type:datetime"`
	DateLastModified         time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	SequenceNumber           float64         `bun:"sequence_number,type:decimal(5,0),pk"`
	SubAssembly              string          `bun:"sub_assembly,type:char"`
	Taxable                  string          `bun:"taxable,type:char"`
	UnitOfMeasure            sql.NullString  `bun:"unit_of_measure,type:varchar(8),nullzero"`
	OtherChargeItem          string          `bun:"other_charge_item,type:char"`
	InvMastUid               int32           `bun:"inv_mast_uid,type:int,pk"`
	ComponentInvMastUid      int32           `bun:"component_inv_mast_uid,type:int"`
	ExtendedDesc             sql.NullString  `bun:"extended_desc,type:varchar(255),nullzero"`
	ComponentCutLength       float64         `bun:"component_cut_length,type:decimal(19,9),default:(0)"`
	CutLengthEditedFlag      string          `bun:"cut_length_edited_flag,type:char,default:('N')"`
	QtyNeeded                float64         `bun:"qty_needed,type:decimal(19,9),default:(0)"`
	CreatedBy                sql.NullString  `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	ServiceLaborUid          sql.NullInt32   `bun:"service_labor_uid,type:int,nullzero"`
	OperationUid             sql.NullInt32   `bun:"operation_uid,type:int,nullzero"`
	DateOffsetDays           sql.NullInt32   `bun:"date_offset_days,type:int,nullzero"`
	AssemblyItemRevisionUid  sql.NullInt32   `bun:"assembly_item_revision_uid,type:int,nullzero"`
	ComponentItemRevisionUid sql.NullInt32   `bun:"component_item_revision_uid,type:int,nullzero"`
	AssemblyLineUid          int32           `bun:"assembly_line_uid,type:int,pk,identity"`
	BackflushFlag            string          `bun:"backflush_flag,type:char,default:('N')"`
	RefDesignatorLocator     sql.NullString  `bun:"ref_designator_locator,type:varchar(255),nullzero"`
	UserComponentNumber      sql.NullInt32   `bun:"user_component_number,type:int,nullzero"`
	BeltingNumCuts           sql.NullInt32   `bun:"belting_num_cuts,type:int,nullzero"`
	BeltingLength            sql.NullFloat64 `bun:"belting_length,type:decimal(19,9),nullzero"`
	BeltingWidth             sql.NullFloat64 `bun:"belting_width,type:decimal(19,9),nullzero"`
	LooseShipFlag            sql.NullString  `bun:"loose_ship_flag,type:char,nullzero"`
	MinimumMccCode           sql.NullString  `bun:"minimum_mcc_code,type:varchar(255),nullzero"`
	ExtendedItemFlag         sql.NullString  `bun:"extended_item_flag,type:char,nullzero"`
}

type AssemblyLineModel struct {
	bun bun.IDB
}

// GetByInvMastUid returns a slice of AssemblyLine by the given InvMastUid
func (m AssemblyLineModel) GetByInvMastUid(ctx context.Context, invMastUid int32) ([]*AssemblyLine, error) {
	var assemblyLines []*AssemblyLine
	err := m.bun.NewSelect().Model(&assemblyLines).Where("inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return assemblyLines, nil
}

// Delete removes the record from the database.
func (m AssemblyLineModel) Delete(ctx context.Context, assemblyLine *AssemblyLine) error {
	_, err := m.bun.NewDelete().Model(assemblyLine).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
