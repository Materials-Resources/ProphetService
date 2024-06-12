package model

import "github.com/uptrace/bun"

type EccColumnLookup struct {
	bun.BaseModel      `bun:"table:ecc_column_lookup"`
	EccColumnLookupUid int32  `bun:"ecc_column_lookup_uid,type:int,pk,identity"`
	TransactionType    string `bun:"transaction_type,type:nvarchar(255)"`
	ParameterTag       string `bun:"parameter_tag,type:nvarchar(255)"`
	TableName          string `bun:"table_name,type:nvarchar(255)"`
	ColumnName         string `bun:"column_name,type:nvarchar(4000)"`
	Datatype           string `bun:"datatype,type:nvarchar(255)"`
}
