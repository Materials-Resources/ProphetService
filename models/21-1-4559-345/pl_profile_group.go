package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PlProfileGroup struct {
	bun.BaseModel `bun:"table:pl_profile_group"`
	GroupSkey     int32     `bun:"group_skey,type:int,pk"`
	ProfileSkey   int32     `bun:"profile_skey,type:int,pk"`
	CreateUserId  string    `bun:"create_user_id,type:char(10),nullzero"`
	CreateDate    time.Time `bun:"create_date,type:datetime,nullzero"`
	MaintUserId   string    `bun:"maint_user_id,type:char(10),nullzero"`
	MaintDate     time.Time `bun:"maint_date,type:datetime,nullzero"`
}
