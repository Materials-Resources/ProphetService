package model

import (
	"github.com/uptrace/bun"
	"time"
)

type WzdProcessP21 struct {
	bun.BaseModel    `bun:"table:wzd_process_p21"`
	WizardProcessNo  int16     `bun:"wizard_process_no,type:smallint,pk"`
	WizardTypeNo     int16     `bun:"wizard_type_no,type:smallint"`
	WizardAppNo      int16     `bun:"wizard_app_no,type:smallint"`
	StepSeq          int16     `bun:"step_seq,type:smallint"`
	StepNm           string    `bun:"step_nm,type:varchar(60)"`
	StepStatusCode   string    `bun:"step_status_code,type:char,default:('A')"`
	WindowNm         string    `bun:"window_nm,type:varchar(60)"`
	DataobjectNm     string    `bun:"dataobject_nm,type:varchar(60)"`
	InstructTx       string    `bun:"instruct_tx,type:text(2147483647),nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
}
