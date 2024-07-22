package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Dynachange struct {
	bun.BaseModel     `bun:"table:dynachange"`
	DynachangeId      float64   `bun:"dynachange_id,type:decimal(6,0),pk"`                        // Which dynachange is this menu for?
	BaseClass         string    `bun:"base_class,type:varchar(255)"`                              // What is the name of the class that the change is based on?
	PersonalizedClass string    `bun:"personalized_class,type:varchar(255)"`                      // What is the name of the new class?
	DateCreated       time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
