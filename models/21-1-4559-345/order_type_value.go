package model

import (
	"github.com/uptrace/bun"
	"time"
)

type OrderTypeValue struct {
	bun.BaseModel     `bun:"table:order_type_value"`
	OrderTypeValueUid int32     `bun:"order_type_value_uid,type:int,pk,identity"`
	OrderTypeId       int32     `bun:"order_type_id,type:int"`
	RequiredFlag      string    `bun:"required_flag,type:char,default:('N')"`
	Value             string    `bun:"value,type:varchar(40)"`
	DefaultColumnFlag string    `bun:"default_column_flag,type:char,default:('N')"`
	RowStatusFlag     int32     `bun:"row_status_flag,type:int"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
