package model

type TempUdKeyUpdate struct {
	bun.BaseModel      `bun:"table:temp_ud_key_update"`
	TempUdKeyUpdateUid int32 `bun:"temp_ud_key_update_uid,type:int,identity"`
	TableName          `bun:"table_name,type:nvarchar(255)"`
	BaseTable          `bun:"base_table,type:nvarchar(255)"`
	UdKeyName          `bun:"ud_key_name,type:nvarchar(255)"`
}
