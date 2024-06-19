package gen

import "github.com/uptrace/bun"

type WzdPrcsSesnStateP21 struct {
	bun.BaseModel   `bun:"table:wzd_prcs_sesn_state_p21"`
	WizardStateCode string `bun:"wizard_state_code,type:char(1),pk"`                  // State code
	WizardStateDesc string `bun:"wizard_state_desc,type:varchar(15)"`                 // State description
	CreatedOn       string `bun:"created_on,type:smalldatetime,default:(getdate())"`  // When was this document_line_bin created?
	ModifiedOn      string `bun:"modified_on,type:smalldatetime,default:(getdate())"` // Date on which this record was last modified.
	ModifiedBy      string `bun:"modified_by,type:varchar(30),default:(user_name())"` // Name of user who created or last modified this rec
}
