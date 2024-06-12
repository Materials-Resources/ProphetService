package model

type CcPaymentTypeXProcessor struct {
	bun.BaseModel          `bun:"table:cc_payment_type_x_processor"`
	CcPaymentXProcessorUid int32     `bun:"cc_payment_x_processor_uid,type:int,pk"`
	PaymentTypeId          float64   `bun:"payment_type_id,type:decimal(19,0)"`
	CreditcardProcessorUid int32     `bun:"creditcard_processor_uid,type:int"`
	OrderSourceTypeCd      int32     `bun:"order_source_type_cd,type:int"`
	RowStatusFlag          int32     `bun:"row_status_flag,type:int"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	LocationId             float64   `bun:"location_id,type:decimal(19,0)"`
}
