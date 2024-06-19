package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PlGroupDef struct {
	bun.BaseModel `bun:"table:pl_group_def"`
	GroupSkey     int32     `bun:"group_skey,type:int,pk"`                // Padlock table.
	Priority      int32     `bun:"priority,type:int,nullzero"`            // Padlock table.
	SecAdminInd   int32     `bun:"sec_admin_ind,type:int"`                // Padlock table.
	GroupName     string    `bun:"group_name,type:varchar(100),unique"`   // Padlock table.
	CreateUserId  string    `bun:"create_user_id,type:char(10),nullzero"` // Padlock table.
	CreateDate    time.Time `bun:"create_date,type:datetime,nullzero"`    // Padlock table.
	MaintUserId   string    `bun:"maint_user_id,type:char(10),nullzero"`  // Padlock table.
	MaintDate     time.Time `bun:"maint_date,type:datetime,nullzero"`     // This column is unused.
}
