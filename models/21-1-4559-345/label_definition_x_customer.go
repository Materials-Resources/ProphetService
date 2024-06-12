package model

import (
	"github.com/uptrace/bun"
	"time"
)

type LabelDefinitionXCustomer struct {
	bun.BaseModel           `bun:"table:label_definition_x_customer"`
	LabelDefinitionXCustUid int32     `bun:"label_definition_x_cust_uid,type:int,pk"`
	LabelDefinitionUid      int32     `bun:"label_definition_uid,type:int"`
	CompanyId               string    `bun:"company_id,type:varchar(8)"`
	CustomerId              float64   `bun:"customer_id,type:decimal(19,0)"`
	RowStatusFlag           int32     `bun:"row_status_flag,type:int"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
