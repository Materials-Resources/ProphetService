package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PcProfileDef struct {
	bun.BaseModel `bun:"table:pc_profile_def"`
	ProfileSkey   int32     `bun:"profile_skey,type:int,pk"`
	SysAdminInd   int32     `bun:"sys_admin_ind,type:int,default:(0)"`
	ProfileName   string    `bun:"profile_name,type:varchar(100)"`
	CreateUserId  string    `bun:"create_user_id,type:char(10),nullzero"`
	CreateDate    time.Time `bun:"create_date,type:datetime,nullzero"`
	MaintUserId   string    `bun:"maint_user_id,type:char(10),nullzero"`
	MaintDate     time.Time `bun:"maint_date,type:datetime,nullzero"`
}
