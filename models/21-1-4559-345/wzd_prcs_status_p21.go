package gen

import "github.com/uptrace/bun"

type WzdPrcsStatusP21 struct {
	bun.BaseModel    `bun:"table:wzd_prcs_status_p21"`
	WizardStatusCode string `bun:"wizard_status_code,type:char(1),pk"`                 // State code
	WizardStatusDesc string `bun:"wizard_status_desc,type:varchar(15)"`                // State description
	CreatedOn        string `bun:"created_on,type:smalldatetime,default:(getdate())"`  // When was this document_line_lot created?
	ModifiedOn       string `bun:"modified_on,type:smalldatetime,default:(getdate())"` // Date created timestamp
	ModifiedBy       string `bun:"modified_by,type:varchar(30),default:(user_name())"` // Name of user who created or last modified this rec
}
