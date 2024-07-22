package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PcUserDef struct {
	bun.BaseModel `bun:"table:pc_user_def"`
	UserSkey      int32      `bun:"user_skey,type:int,pk"`           // Padlock table
	UserId        string     `bun:"user_id,type:varchar(30)"`        // Padlock table
	LocationSkey  *int32     `bun:"location_skey,type:int"`          // Padlock table
	Password      *string    `bun:"password,type:varchar(40)"`       // Padlock table
	FirstName     string     `bun:"first_name,type:varchar(20)"`     // Padlock table
	Mi            *string    `bun:"mi,type:varchar(2)"`              // Padlock table
	LastName      string     `bun:"last_name,type:varchar(24)"`      // Padlock table
	ProfileSkey   *int32     `bun:"profile_skey,type:int"`           // Padlock table
	Ssn           *string    `bun:"ssn,type:char(9)"`                // Padlock table
	HomePhone     *string    `bun:"home_phone,type:varchar(20)"`     // Padlock table
	WorkPhone     *string    `bun:"work_phone,type:varchar(20)"`     // Padlock table
	WorkExt       *string    `bun:"work_ext,type:char(5)"`           // Padlock table
	FaxNbr        *string    `bun:"fax_nbr,type:char(20)"`           // Padlock table
	EmailAddress  *string    `bun:"email_address,type:varchar(100)"` // Padlock table
	ActiveInd     *string    `bun:"active_ind,type:char(1)"`         // Padlock table
	CreateUserId  *string    `bun:"create_user_id,type:char(10)"`    // Padlock table
	CreateDate    *time.Time `bun:"create_date,type:datetime"`       // Padlock table
	MaintUserId   *string    `bun:"maint_user_id,type:char(10)"`     // Padlock table
	MaintDate     *time.Time `bun:"maint_date,type:datetime"`        // Padlock table
}
