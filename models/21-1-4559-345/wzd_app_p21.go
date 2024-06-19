package gen

import "github.com/uptrace/bun"

type WzdAppP21 struct {
	bun.BaseModel `bun:"table:wzd_app_p21"`
	WizardAppNo   int16  `bun:"wizard_app_no,type:smallint,pk"`                     // System assigned unique key
	AppExe        string `bun:"app_exe,type:varchar(30)"`                           // File name
	ModuleNm      string `bun:"module_nm,type:varchar(40)"`                         // MDI Frame Title
	ModuleFrame   string `bun:"module_frame,type:varchar(60)"`                      // MDI Frame PowerBuilder object name
	CreatedOn     string `bun:"created_on,type:smalldatetime,default:(getdate())"`  // When was this document_line_lot created?
	ModifiedOn    string `bun:"modified_on,type:smalldatetime,default:(getdate())"` // Date created timestamp
	ModifiedBy    string `bun:"modified_by,type:varchar(30),default:(user_name())"` // Name of user who created or last modified this rec
}
