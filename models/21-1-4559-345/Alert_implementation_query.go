package model

import (
	"github.com/uptrace/bun"
	"time"
)

type AlertImplementationQuery struct {
	bun.BaseModel               `bun:"table:Alert_implementation_query"`
	AlertImplementationQueryUid int32     `bun:"alert_implementation_query_uid,type:int,pk"`
	AlertImplementationUid      int32     `bun:"alert_implementation_uid,type:int"`
	ColumnId                    int32     `bun:"column_id,type:int"`
	OperatorCd                  int32     `bun:"operator_cd,type:int"`
	ColumnValue                 string    `bun:"column_value,type:varchar(255)"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30)"`
	RowStatusFlag               int32     `bun:"row_status_flag,type:int"`
	ColumnValueDescription      string    `bun:"column_value_description,type:varchar(255)"`
}
