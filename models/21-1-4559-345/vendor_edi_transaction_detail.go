package model

type VendorEdiTransactionDetail struct {
	bun.BaseModel           `bun:"table:vendor_edi_transaction_detail"`
	VendorEdiTransDetailUid int32     `bun:"vendor_edi_trans_detail_uid,type:int,pk"`
	VendorEdiTransactionUid int32     `bun:"vendor_edi_transaction_uid,type:int"`
	Name                    string    `bun:"name,type:varchar(255)"`
	Value                   string    `bun:"value,type:varchar(255)"`
	DataTypeCd              int32     `bun:"data_type_cd,type:int"`
	DataTypeLength          int32     `bun:"data_type_length,type:int"`
	DataTypeScale           int32     `bun:"data_type_scale,type:int"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
