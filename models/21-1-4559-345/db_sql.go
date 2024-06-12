package model

type DbSql struct {
	bun.BaseModel   `bun:"table:db_sql"`
	LastSqlExecuted string    `bun:"last_sql_executed,type:varchar(50)"`
	DateSqlExecuted time.Time `bun:"date_sql_executed,type:datetime"`
	SqlDescription  string    `bun:"sql_description,type:varchar(255)"`
	DbSqlUid        int32     `bun:"db_sql_uid,type:int,pk,identity"`
}
