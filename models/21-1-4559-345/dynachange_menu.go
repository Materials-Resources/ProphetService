package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DynachangeMenu struct {
	bun.BaseModel         `bun:"table:dynachange_menu"`
	DynachangeId          float64   `bun:"dynachange_id,type:decimal(6,0),pk"`                        // What is the unique identifier for this dynachange?
	BaseClassItem         string    `bun:"base_class_item,type:varchar(255),pk"`                      // What is the base class for this dynachange menu?
	PersonalizedClassItem string    `bun:"personalized_class_item,type:varchar(255),pk"`              // What is the new class item?
	DateCreated           time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
