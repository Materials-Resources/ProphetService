package model

type InvTranSerialDetail struct {
	bun.BaseModel          `bun:"table:inv_tran_serial_detail"`
	CompanyId              string    `bun:"company_id,type:varchar(8)"`
	LocationId             float64   `bun:"location_id,type:decimal(19,0)"`
	SerialNumber           string    `bun:"serial_number,type:varchar(40)"`
	TransactionNumber      int32     `bun:"transaction_number,type:int,default:((-1))"`
	DateCreated            time.Time `bun:"date_created,type:datetime"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30)"`
	DimensionTrackingKey   int32     `bun:"dimension_tracking_key,type:int,nullzero"`
	InvTranSerialDetailUid int32     `bun:"inv_tran_serial_detail_uid,type:int,pk"`
	RowStatus              int16     `bun:"row_status,type:tinyint"`
	DocumentLineSerialUid  int32     `bun:"document_line_serial_uid,type:int"`
	InvMastUid             int32     `bun:"inv_mast_uid,type:int"`
	LotUid                 int32     `bun:"lot_uid,type:int,nullzero"`
}
