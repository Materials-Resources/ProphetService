package model

type Counter struct {
	bun.BaseModel `bun:"table:counter"`
	Id            string `bun:"id,type:varchar(64),pk"`
	Description   string `bun:"description,type:varchar(255),nullzero"`
	CounterNum    int32  `bun:"counter_num,type:int"`
	SeqName       `bun:"seq_name,type:nvarchar(255),nullzero"`
	TableName     `bun:"table_name,type:nvarchar(255),nullzero"`
	ColumnName    `bun:"column_name,type:nvarchar(255),nullzero"`
}
