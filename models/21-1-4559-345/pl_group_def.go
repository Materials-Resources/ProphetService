package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PlGroupDef struct {
	bun.BaseModel `bun:"table:pl_group_def"`
	GroupSkey     int32     `bun:"group_skey,type:int,pk"`
	Priority      int32     `bun:"priority,type:int,nullzero"`
	SecAdminInd   int32     `bun:"sec_admin_ind,type:int"`
	GroupName     string    `bun:"group_name,type:varchar(100)"`
	CreateUserId  string    `bun:"create_user_id,type:char(10),nullzero"`
	CreateDate    time.Time `bun:"create_date,type:datetime,nullzero"`
	MaintUserId   string    `bun:"maint_user_id,type:char(10),nullzero"`
	MaintDate     time.Time `bun:"maint_date,type:datetime,nullzero"`
}
