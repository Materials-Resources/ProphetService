package model

type SupportSql struct {
	bun.BaseModel        `bun:"table:support_sql"`
	SupportSqlUid        int32  `bun:"support_sql_uid,type:int,identity"`
	SupportQueryCd       int32  `bun:"support_query_cd,type:int,pk"`
	SupportQuerySequence int32  `bun:"support_query_sequence,type:int,pk"`
	ViewName             string `bun:"view_name,type:varchar(255),nullzero"`
	WhereClause          string `bun:"where_clause,type:varchar(4000),nullzero"`
	SupportSql           string `bun:"support_sql,type:varchar(4000),nullzero"`
	P21Baseline          string `bun:"p21_baseline,type:char,default:('N')"`
	OrderByClause        string `bun:"order_by_clause,type:varchar(4000),nullzero"`
}
