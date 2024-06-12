package model

import "github.com/uptrace/bun"

type WzdAppP21 struct {
	bun.BaseModel `bun:"table:wzd_app_p21"`
	WizardAppNo   int16  `bun:"wizard_app_no,type:smallint,pk"`
	AppExe        string `bun:"app_exe,type:varchar(30)"`
	ModuleNm      string `bun:"module_nm,type:varchar(40)"`
	ModuleFrame   string `bun:"module_frame,type:varchar(60)"`
	CreatedOn     string `bun:"created_on,type:smalldatetime,default:(getdate())"`
	ModifiedOn    string `bun:"modified_on,type:smalldatetime,default:(getdate())"`
	ModifiedBy    string `bun:"modified_by,type:varchar(30),default:(user_name())"`
}
