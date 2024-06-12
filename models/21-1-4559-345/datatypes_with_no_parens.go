package model

type DatatypesWithNoParens struct {
	bun.BaseModel            `bun:"table:datatypes_with_no_parens"`
	DatatypesWithNoParensUid int32  `bun:"datatypes_with_no_parens_uid,type:int,identity"`
	ColumnDatatype           string `bun:"column_datatype,type:varchar(50),pk"`
}
