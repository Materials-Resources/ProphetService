package model

type WzdPrcsSesnStateP21 struct {
	bun.BaseModel   `bun:"table:wzd_prcs_sesn_state_p21"`
	WizardStateCode string `bun:"wizard_state_code,type:char,pk"`
	WizardStateDesc string `bun:"wizard_state_desc,type:varchar(15)"`
	CreatedOn       string `bun:"created_on,type:smalldatetime,default:(getdate())"`
	ModifiedOn      string `bun:"modified_on,type:smalldatetime,default:(getdate())"`
	ModifiedBy      string `bun:"modified_by,type:varchar(30),default:(user_name())"`
}
