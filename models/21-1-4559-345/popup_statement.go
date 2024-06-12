package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PopupStatement struct {
	bun.BaseModel           `bun:"table:popup_statement"`
	PopupStatementUid       int32     `bun:"popup_statement_uid,type:int,pk,identity"`
	Columns                 string    `bun:"columns,type:varchar"`
	FromJoin                string    `bun:"from_join,type:varchar"`
	Where                   string    `bun:"where,type:varchar"`
	OrderBy                 string    `bun:"order_by,type:varchar"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	PopupDetailUid          int32     `bun:"popup_detail_uid,type:int"`
	GroupBy                 string    `bun:"group_by,type:varchar,nullzero"`
	OverrideColumns         string    `bun:"override_columns,type:char(3),nullzero"`
	OverrideFromJoin        string    `bun:"override_from_join,type:char(3),nullzero"`
	OverrideWhere           string    `bun:"override_where,type:char(3),nullzero"`
	OverrideGroupBy         string    `bun:"override_group_by,type:char(3),nullzero"`
	PopupStatementUidParent int32     `bun:"popup_statement_uid_parent,type:int,nullzero"`
	OverrideOrderBy         string    `bun:"override_order_by,type:char(3),nullzero"`
	Option                  string    `bun:"option,type:varchar,default:(NULL)"`
	OverrideOption          bool      `bun:"override_option,type:bit,default:((0))"`
}
