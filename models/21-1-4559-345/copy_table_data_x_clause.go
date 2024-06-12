package model

type CopyTableDataXClause struct {
	bun.BaseModel           `bun:"table:copy_table_data_x_clause"`
	CopyTableDataXClauseUid int32     `bun:"copy_table_data_x_clause_uid,type:int,pk,identity"`
	CopyTableDataXTableUid  int32     `bun:"copy_table_data_x_table_uid,type:int"`
	FromClause              string    `bun:"from_clause,type:varchar(8000)"`
	WhereClause             string    `bun:"where_clause,type:varchar(8000)"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
