package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerListTemp struct {
	bun.BaseModel       `bun:"table:customer_list_temp"`
	CustomerListTempUid int32     `bun:"customer_list_temp_uid,type:int,identity"`
	RunNumber           int32     `bun:"run_number,type:int"`
	CompanyId           string    `bun:"company_id,type:varchar(8)"`
	CustomerId          float64   `bun:"customer_id,type:decimal(19,0)"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
