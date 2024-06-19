package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type BusinessRuleEventKey struct {
	bun.BaseModel           `bun:"table:business_rule_event_key"`
	BusinessRuleEventKeyUid int32     `bun:"business_rule_event_key_uid,type:int,autoincrement,pk"`        // UID for this table.
	BusinessRuleEventUid    int32     `bun:"business_rule_event_uid,type:int"`                             // Business rule event that this value relates to.
	KeyValue                string    `bun:"key_value,type:varchar(255)"`                                  // Internal value of the key for this event.
	DisplayValue            string    `bun:"display_value,type:varchar(255)"`                              // Value of the key for this event that will be shown to the user.
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
