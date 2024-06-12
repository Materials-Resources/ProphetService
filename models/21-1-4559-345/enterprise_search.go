package model

type EnterpriseSearch struct {
	bun.BaseModel             `bun:"table:enterprise_search"`
	EnterpriseSearchUid       int32     `bun:"enterprise_search_uid,type:int,pk,identity"`
	EntityType                string    `bun:"entity_type,type:varchar(255)"`
	TableName                 string    `bun:"table_name,type:varchar(255)"`
	SearchExpression          string    `bun:"search_expression,type:varchar(8000)"`
	Id                        string    `bun:"id,type:varchar(255)"`
	PrimaryDesc               string    `bun:"primary_desc,type:varchar(255)"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DeleteClause              `bun:"delete_clause,type:nvarchar,nullzero"`
	CompanyColumn             string `bun:"company_column,type:varchar(255),nullzero"`
	IncludeNullCompanyRecords string `bun:"include_null_company_records,type:char,default:('N')"`
}
