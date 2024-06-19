package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type BusinessRuleLog struct {
	bun.BaseModel      `bun:"table:business_rule_log"`
	BusinessRuleLogUid int32     `bun:"business_rule_log_uid,type:int,autoincrement,pk"`              // Unique identifier for record
	UserId             string    `bun:"user_id,type:varchar(255)"`                                    // User ID of current user
	LogAction          string    `bun:"log_action,type:varchar(255)"`                                 // Action being  recorded:  Invoke, Return, Update
	RuleName           string    `bun:"rule_name,type:varchar(255),nullzero"`                         // Name of rule being executed
	RuleManagerName    string    `bun:"rule_manager_name,type:varchar(255),nullzero"`                 // Name of rule manager DLL being executed
	RuleAssemblyName   string    `bun:"rule_assembly_name,type:varchar(255),nullzero"`                // Name of rule assembly DLL being executed
	RunType            string    `bun:"run_type,type:varchar(255),nullzero"`                          // Type of execution: Synchronous or Asynchronous
	Xml                string    `bun:"xml,type:text,nullzero"`                                       // XML being sent to rule or returned from rule
	ReturnValue        string    `bun:"return_value,type:varchar(255),nullzero"`                      // Rule return: Success or Fail
	ReturnMessage      string    `bun:"return_message,type:varchar(8000),nullzero"`                   // Message returned by rule
	UpdateClassName    string    `bun:"update_class_name,type:varchar(255),nullzero"`                 // Class name being updated
	UpdateRowId        int32     `bun:"update_row_id,type:int,nullzero"`                              // Row ID being updated
	UpdateRowNumber    int32     `bun:"update_row_number,type:int,nullzero"`                          // Row Number being updated
	UpdateFieldName    string    `bun:"update_field_name,type:varchar(255),nullzero"`                 // Field Name being updated
	UpdateFieldValue   string    `bun:"update_field_value,type:varchar(8000),nullzero"`               // Value being updated to field
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
