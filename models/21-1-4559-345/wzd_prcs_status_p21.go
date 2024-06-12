package model

type WzdPrcsStatusP21 struct {
	bun.BaseModel    `bun:"table:wzd_prcs_status_p21"`
	WizardStatusCode string `bun:"wizard_status_code,type:char,pk"`
	WizardStatusDesc string `bun:"wizard_status_desc,type:varchar(15)"`
	CreatedOn        string `bun:"created_on,type:smalldatetime,default:(getdate())"`
	ModifiedOn       string `bun:"modified_on,type:smalldatetime,default:(getdate())"`
	ModifiedBy       string `bun:"modified_by,type:varchar(30),default:(user_name())"`
}
