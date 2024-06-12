package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PcUserDef struct {
	bun.BaseModel `bun:"table:pc_user_def"`
	UserSkey      int32     `bun:"user_skey,type:int,pk"`
	UserId        string    `bun:"user_id,type:varchar(30)"`
	LocationSkey  int32     `bun:"location_skey,type:int,nullzero"`
	Password      string    `bun:"password,type:varchar(40),nullzero"`
	FirstName     string    `bun:"first_name,type:varchar(20)"`
	Mi            string    `bun:"mi,type:varchar(2),nullzero"`
	LastName      string    `bun:"last_name,type:varchar(24)"`
	ProfileSkey   int32     `bun:"profile_skey,type:int,nullzero"`
	Ssn           string    `bun:"ssn,type:char(9),nullzero"`
	HomePhone     string    `bun:"home_phone,type:varchar(20),nullzero"`
	WorkPhone     string    `bun:"work_phone,type:varchar(20),nullzero"`
	WorkExt       string    `bun:"work_ext,type:char(5),nullzero"`
	FaxNbr        string    `bun:"fax_nbr,type:char(20),nullzero"`
	EmailAddress  string    `bun:"email_address,type:varchar(100),nullzero"`
	ActiveInd     string    `bun:"active_ind,type:char,nullzero"`
	CreateUserId  string    `bun:"create_user_id,type:char(10),nullzero"`
	CreateDate    time.Time `bun:"create_date,type:datetime,nullzero"`
	MaintUserId   string    `bun:"maint_user_id,type:char(10),nullzero"`
	MaintDate     time.Time `bun:"maint_date,type:datetime,nullzero"`
}
