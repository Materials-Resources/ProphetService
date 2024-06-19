package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type WzdProcessP21 struct {
	bun.BaseModel    `bun:"table:wzd_process_p21"`
	WizardProcessNo  int16     `bun:"wizard_process_no,type:smallint,pk"`                            // System assigned unique key
	WizardTypeNo     int16     `bun:"wizard_type_no,type:smallint"`                                  // System assigned unique key
	WizardAppNo      int16     `bun:"wizard_app_no,type:smallint"`                                   // Wizard application foreign key
	StepSeq          int16     `bun:"step_seq,type:smallint"`                                        // Task/step sequence
	StepNm           string    `bun:"step_nm,type:varchar(60)"`                                      // Task/step name
	StepStatusCode   string    `bun:"step_status_code,type:char(1),default:('A')"`                   // Task/step status
	WindowNm         string    `bun:"window_nm,type:varchar(60)"`                                    // Window PowerBuilder object name
	DataobjectNm     string    `bun:"dataobject_nm,type:varchar(60)"`                                // DataWindow PowerBuilder object name
	InstructTx       string    `bun:"instruct_tx,type:text,nullzero"`                                // Wizard instruction
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
}
