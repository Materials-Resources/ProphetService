package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PopupFieldBehavior struct {
	bun.BaseModel         `bun:"table:popup_field_behavior"`
	PopupFieldBehaviorUid int32     `bun:"popup_field_behavior_uid,type:int,autoincrement,identity,pk"`  // Unique identifier of popup_field_behavior
	PopupFieldUid         int32     `bun:"popup_field_uid,type:int"`                                     // Foreign key association to popup_field
	Condition             string    `bun:"condition,type:varchar(max)"`                                  // Condition to evaluate if the behavior should be used
	ConfigurationId       int32     `bun:"configuration_id,type:int,default:((0))"`                      // Defines a configuration_id for this behavior
	Sort                  bool      `bun:"sort,type:bit"`                                                // If the popup field should be used as a sort
	Filter                bool      `bun:"filter,type:bit"`                                              // If the popup field should be used as a filter
	Visible               bool      `bun:"visible,type:bit"`                                             // If the popup field should be visible in the grid
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	AsParameter           bool      `bun:"as_parameter,type:bit,default:((0))"`                          // Override AsParameter for a Specific Dynachange or condition.
	ParameterDefaultValue string    `bun:"parameter_default_value,type:varchar(255),nullzero"`           // Override Parameter Default Value for a Specific Dynachange or condition.
}
