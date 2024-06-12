package model

type EccGetTableSchemaTable struct {
	bun.BaseModel             `bun:"table:ecc_get_table_schema_table"`
	EccGetTableSchemaTableUid int32 `bun:"ecc_get_table_schema_table_uid,type:int,identity"`
	TableName                 `bun:"table_name,type:nvarchar"`
}
