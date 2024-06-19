package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type EnterpriseSearch struct {
	bun.BaseModel             `bun:"table:enterprise_search"`
	EnterpriseSearchUid       int32     `bun:"enterprise_search_uid,type:int,autoincrement,pk"`              // Unique ID for this table
	EntityType                string    `bun:"entity_type,type:varchar(255)"`                                // Entity type for this record
	TableName                 string    `bun:"table_name,type:varchar(255)"`                                 // Table to search for this record
	SearchExpression          string    `bun:"search_expression,type:varchar(8000)"`                         // Search expression for the table related to this record
	Id                        string    `bun:"id,type:varchar(255)"`                                         // Record identifier for this table related records
	PrimaryDesc               string    `bun:"primary_desc,type:varchar(255)"`                               // Primary description for records from the table related to thsi records
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	DeleteClause              string    `bun:"delete_clause,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`
	CompanyColumn             string    `bun:"company_column,type:varchar(255),nullzero"`               // Stores the name of the column on the searched table that holds the company.
	IncludeNullCompanyRecords string    `bun:"include_null_company_records,type:char(1),default:('N')"` // Determines whether search for a particular entity type should include records that have a NULL company value.
}
