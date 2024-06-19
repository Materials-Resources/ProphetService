package gen

import "github.com/uptrace/bun"

type WzdTypeP21 struct {
	bun.BaseModel `bun:"table:wzd_type_p21"`
	WizardTypeNo  int16  `bun:"wizard_type_no,type:smallint,pk"`                    // Wizard type foreign key
	WizardNm      string `bun:"wizard_nm,type:varchar(40)"`                         // Wizard Name
	WizardExe     string `bun:"wizard_exe,type:varchar(30)"`                        // Wizard application name
	WizardPic     string `bun:"wizard_pic,type:varchar(30)"`                        // A descriptive name for the Wizard used for display in the interface.
	WizardMode    string `bun:"wizard_mode,type:varchar(5)"`                        // Mode of operation
	CreatedOn     string `bun:"created_on,type:smalldatetime,default:(getdate())"`  // When was this document_line_bin created?
	ModifiedOn    string `bun:"modified_on,type:smalldatetime,default:(getdate())"` // Date on which this record was last modified.
	ModifiedBy    string `bun:"modified_by,type:varchar(30),default:(user_name())"` // User who made last modification
}
