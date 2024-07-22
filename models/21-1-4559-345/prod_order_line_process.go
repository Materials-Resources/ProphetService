package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdOrderLineProcess struct {
	bun.BaseModel           `bun:"table:prod_order_line_process"`
	ProdOrderLineProcessUid int32     `bun:"prod_order_line_process_uid,type:int,pk"`                      // Unique Identifier for the table.
	ProdOrderLineUid        int32     `bun:"prod_order_line_uid,type:int"`                                 // Associated Production Order Line record.
	SequenceNumber          int32     `bun:"sequence_number,type:int"`                                     // Number of times the Assembly have been completed.
	QtyCompleted            float64   `bun:"qty_completed,type:decimal(19,9)"`                             // Quantity Completed for the Assembly.
	MaterialCost            float64   `bun:"material_cost,type:decimal(19,9)"`                             // Total of the Material Cost for the Completed Assemblies.
	LaborCost               float64   `bun:"labor_cost,type:decimal(19,9)"`                                // Total of the Labor Cost for the Completed Assemblies.
	FreightCost             float64   `bun:"freight_cost,type:decimal(19,9)"`                              // Total of the Freight Cost for the Completed Assemblies.
	OtherChargeCost         float64   `bun:"other_charge_cost,type:decimal(19,9)"`                         // Total of the Other Charge Cost for the Completed Assemblies.
	ProcessCost             float64   `bun:"process_cost,type:decimal(19,9)"`                              // Total of the Secondary Processing Cost for the Completed Assemblies.
	RowStatusFlag           int32     `bun:"row_status_flag,type:int"`                                     // Reflects the status of the record.
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created.
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record.
	AdditionalLabor         *float64  `bun:"additional_labor,type:decimal(19,9)"`                          // Labor Cost specified after the Completion of the Production Order.
	AdditionalFreight       *float64  `bun:"additional_freight,type:decimal(19,9)"`                        // Freight Cost specified after the Completion of the Production Order.
	AdditionalMaterial      *float64  `bun:"additional_material,type:decimal(19,9)"`                       // Material Cost specified after the Completion of the Production Order.
	AdditionalOtherCharge   *float64  `bun:"additional_other_charge,type:decimal(19,9)"`                   // Other Charge Cost specified after the Completion of the Production Order.
	SourceTypeCd            *int32    `bun:"source_type_cd,type:int"`                                      // Code (from code_p21 table) which indicates where the record was created
	LaborCostIndirect       float64   `bun:"labor_cost_indirect,type:decimal(19,9)"`                       // Total Labor (indirect)
	AdditionalLaborIndirect *float64  `bun:"additional_labor_indirect,type:decimal(19,9)"`                 // Additional total labor (Indirect)
}
