package model

import "github.com/uptrace/bun"

type WzdTypeP21 struct {
	bun.BaseModel `bun:"table:wzd_type_p21"`
	WizardTypeNo  int16  `bun:"wizard_type_no,type:smallint,pk"`
	WizardNm      string `bun:"wizard_nm,type:varchar(40)"`
	WizardExe     string `bun:"wizard_exe,type:varchar(30)"`
	WizardPic     string `bun:"wizard_pic,type:varchar(30)"`
	WizardMode    string `bun:"wizard_mode,type:varchar(5)"`
	CreatedOn     string `bun:"created_on,type:smalldatetime,default:(getdate())"`
	ModifiedOn    string `bun:"modified_on,type:smalldatetime,default:(getdate())"`
	ModifiedBy    string `bun:"modified_by,type:varchar(30),default:(user_name())"`
}
