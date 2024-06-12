package model

import (
	"github.com/uptrace/bun"
	"time"
)

type AccntGroupMx struct {
	bun.BaseModel    `bun:"table:accnt_group_mx"`
	AccntGroupMxUid  int32     `bun:"accnt_group_mx_uid,type:int,pk,identity"`
	AccntGroupMxId   string    `bun:"accnt_group_mx_id,type:varchar(255)"`
	GroupDescription string    `bun:"group_description,type:varchar(255),nullzero"`
	SequenceNo       int32     `bun:"sequence_no,type:int,default:((0))"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:((704))"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	Level            int32     `bun:"level,type:int,nullzero"`
}
