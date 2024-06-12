package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CompanyLostSales struct {
	bun.BaseModel       `bun:"table:company_lost_sales"`
	CompanyLostSalesUid int32     `bun:"company_lost_sales_uid,type:int,pk,identity"`
	CompanyId           string    `bun:"company_id,type:varchar(8)"`
	TransactionCodeNo   int32     `bun:"transaction_code_no,type:int"`
	ActionCodeNo        int32     `bun:"action_code_no,type:int"`
	LostSalesUid        int32     `bun:"lost_sales_uid,type:int,nullzero"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
