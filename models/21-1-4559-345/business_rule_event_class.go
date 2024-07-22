package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type BusinessRuleEventClass struct {
	bun.BaseModel             `bun:"table:business_rule_event_class"`
	BusinessRuleEventClassUid int32     `bun:"business_rule_event_class_uid,type:int,autoincrement,identity,pk"` // UID for this table.
	BusinessRuleEventUid      int32     `bun:"business_rule_event_uid,type:int"`                                 // FK to the business rule event record that this rule class goes with.
	RuleClassName             string    `bun:"rule_class_name,type:varchar(255)"`                                // Name of the .NET class that will be triggered for the associated business rule event.
	RunTypeCd                 int32     `bun:"run_type_cd,type:int"`                                             // Determines if this rule will be executed synchronously or asynchronously.
	ConfigurationId           *int32    `bun:"configuration_id,type:int"`                                        // Configuration for which this rule should be executed. Baseline if null or 0.
	SequenceNo                int32     `bun:"sequence_no,type:int"`                                             // Sequence that the rules for this even should be executed.
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`                   // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`             // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`             // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`     // User who last changed the record
}
