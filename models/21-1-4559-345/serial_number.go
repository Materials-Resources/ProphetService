package model

type SerialNumber struct {
	bun.BaseModel               `bun:"table:serial_number"`
	CompanyNo                   string    `bun:"company_no,type:varchar(8),pk"`
	LocationId                  float64   `bun:"location_id,type:decimal(19,0),pk"`
	SerialNumber                string    `bun:"serial_number,type:varchar(40),pk"`
	DeleteFlag                  string    `bun:"delete_flag,type:char"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	DimensionTrackingKey        int32     `bun:"dimension_tracking_key,type:int,nullzero"`
	RowStatus                   int16     `bun:"row_status,type:tinyint"`
	InvMastUid                  int32     `bun:"inv_mast_uid,type:int,pk"`
	SerialNumberUid             int32     `bun:"serial_number_uid,type:int"`
	CustomerId                  float64   `bun:"customer_id,type:decimal(19,0),nullzero"`
	BinId                       string    `bun:"bin_id,type:varchar(10),nullzero"`
	LotUid                      int32     `bun:"lot_uid,type:int,nullzero"`
	AllowanceAmount             float64   `bun:"allowance_amount,type:decimal(19,4),nullzero"`
	AllowanceAmountModifiedDate time.Time `bun:"allowance_amount_modified_date,type:datetime,nullzero"`
	TradeLayerUid               int32     `bun:"trade_layer_uid,type:int,nullzero"`
}
