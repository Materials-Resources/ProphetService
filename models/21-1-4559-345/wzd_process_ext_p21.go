package model

type WzdProcessExtP21 struct {
	bun.BaseModel   `bun:"table:wzd_process_ext_p21"`
	WizardProcessNo int16  `bun:"wizard_process_no,type:smallint,pk"`
	PanelTx         string `bun:"panel_tx,type:varchar(60),nullzero"`
	PanelPic        string `bun:"panel_pic,type:varchar(30),nullzero"`
	PanelObj        string `bun:"panel_obj,type:varchar(60),nullzero"`
	AddtlTx         string `bun:"addtl_tx,type:text(2147483647),nullzero"`
	CreatedOn       string `bun:"created_on,type:smalldatetime,default:(getdate())"`
	ModifiedOn      string `bun:"modified_on,type:smalldatetime,default:(getdate())"`
	ModifiedBy      string `bun:"modified_by,type:varchar(30),default:(user_name())"`
}
