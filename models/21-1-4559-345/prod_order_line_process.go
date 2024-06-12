package model

type ProdOrderLineProcess struct {
	bun.BaseModel           `bun:"table:prod_order_line_process"`
	ProdOrderLineProcessUid int32     `bun:"prod_order_line_process_uid,type:int,pk"`
	ProdOrderLineUid        int32     `bun:"prod_order_line_uid,type:int"`
	SequenceNumber          int32     `bun:"sequence_number,type:int"`
	QtyCompleted            float64   `bun:"qty_completed,type:decimal(19,9)"`
	MaterialCost            float64   `bun:"material_cost,type:decimal(19,9)"`
	LaborCost               float64   `bun:"labor_cost,type:decimal(19,9)"`
	FreightCost             float64   `bun:"freight_cost,type:decimal(19,9)"`
	OtherChargeCost         float64   `bun:"other_charge_cost,type:decimal(19,9)"`
	ProcessCost             float64   `bun:"process_cost,type:decimal(19,9)"`
	RowStatusFlag           int32     `bun:"row_status_flag,type:int"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	AdditionalLabor         float64   `bun:"additional_labor,type:decimal(19,9),nullzero"`
	AdditionalFreight       float64   `bun:"additional_freight,type:decimal(19,9),nullzero"`
	AdditionalMaterial      float64   `bun:"additional_material,type:decimal(19,9),nullzero"`
	AdditionalOtherCharge   float64   `bun:"additional_other_charge,type:decimal(19,9),nullzero"`
	SourceTypeCd            int32     `bun:"source_type_cd,type:int,nullzero"`
	LaborCostIndirect       float64   `bun:"labor_cost_indirect,type:decimal(19,9)"`
	AdditionalLaborIndirect float64   `bun:"additional_labor_indirect,type:decimal(19,9),nullzero"`
}
