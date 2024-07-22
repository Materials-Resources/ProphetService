package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PlProfileGroup struct {
	bun.BaseModel `bun:"table:pl_profile_group"`
	GroupSkey     int32      `bun:"group_skey,type:int,pk"`       // Padlock table.
	ProfileSkey   int32      `bun:"profile_skey,type:int,pk"`     // Padlock table.
	CreateUserId  *string    `bun:"create_user_id,type:char(10)"` // Padlock table.
	CreateDate    *time.Time `bun:"create_date,type:datetime"`    // Padlock table.
	MaintUserId   *string    `bun:"maint_user_id,type:char(10)"`  // This column is unused.
	MaintDate     *time.Time `bun:"maint_date,type:datetime"`     // This column is unused.
}
