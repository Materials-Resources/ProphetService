package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PredefinedCoa struct {
	bun.BaseModel      `bun:"table:predefined_coa"`
	PredefinedCoaUid   int32     `bun:"predefined_coa_uid,type:int,pk,identity"`
	AccountNumber      string    `bun:"account_number,type:varchar(10)"`
	AccountDescription string    `bun:"account_description,type:varchar(255)"`
	AccountType        string    `bun:"account_type,type:char"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
