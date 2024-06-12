package model

type EccColumnLookup struct {
	bun.BaseModel      `bun:"table:ecc_column_lookup"`
	EccColumnLookupUid int32 `bun:"ecc_column_lookup_uid,type:int,pk,identity"`
	TransactionType    `bun:"transaction_type,type:nvarchar(255)"`
	ParameterTag       `bun:"parameter_tag,type:nvarchar(255)"`
	TableName          `bun:"table_name,type:nvarchar(255)"`
	ColumnName         `bun:"column_name,type:nvarchar(4000)"`
	Datatype           `bun:"datatype,type:nvarchar(255)"`
}
