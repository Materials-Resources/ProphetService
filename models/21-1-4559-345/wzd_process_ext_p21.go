package gen

import "github.com/uptrace/bun"

type WzdProcessExtP21 struct {
	bun.BaseModel   `bun:"table:wzd_process_ext_p21"`
	WizardProcessNo int16  `bun:"wizard_process_no,type:smallint,pk"`                 // System assigned unique key
	PanelTx         string `bun:"panel_tx,type:varchar(60),nullzero"`                 // Display text for Wizard panel override
	PanelPic        string `bun:"panel_pic,type:varchar(30),nullzero"`                // Display graphic for Wizard panel override
	PanelObj        string `bun:"panel_obj,type:varchar(60),nullzero"`                // Custom PowerBuilder User Object override
	AddtlTx         string `bun:"addtl_tx,type:text,nullzero"`                        // Additional user instructions
	CreatedOn       string `bun:"created_on,type:smalldatetime,default:(getdate())"`  // When was this document_line_bin created?
	ModifiedOn      string `bun:"modified_on,type:smalldatetime,default:(getdate())"` // Date on which this record was last modified.
	ModifiedBy      string `bun:"modified_by,type:varchar(30),default:(user_name())"` // User who made last modification
}
