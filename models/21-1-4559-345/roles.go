package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Roles struct {
	bun.BaseModel           `bun:"table:roles"`
	RoleUid                 int32     `bun:"role_uid,type:int,pk"`                                      // Unique identifier for the role.
	Role                    string    `bun:"role,type:varchar(128)"`                                    // The name you define for roles within the system.  This appears in drop-down menus elsewhere in the system.
	DeleteFlag              string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	MinimumMarginPercentage float64   `bun:"minimum_margin_percentage,type:decimal(19,9),nullzero"`     // Minimum acceptable margin percentage for users with this role
}
