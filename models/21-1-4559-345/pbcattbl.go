package prophet

import "github.com/uptrace/bun"

type Pbcattbl struct {
	bun.BaseModel `bun:"table:pbcattbl"`
	PbtTnam       *string `bun:"pbt_tnam,type:char(30)"`     // Powerbuilder catalog information
	PbtTid        int32   `bun:"pbt_tid,type:int,pk"`        // Powerbuilder catalog information
	PbtOwnr       *string `bun:"pbt_ownr,type:char(30)"`     // Powerbuilder catalog information
	PbdFhgt       *int16  `bun:"pbd_fhgt,type:smallint"`     // Powerbuilder catalog information
	PbdFwgt       *int16  `bun:"pbd_fwgt,type:smallint"`     // Powerbuilder catalog information
	PbdFitl       *string `bun:"pbd_fitl,type:char(1)"`      // Powerbuilder catalog information
	PbdFunl       *string `bun:"pbd_funl,type:char(1)"`      // Powerbuilder catalog information
	PbdFchr       *int16  `bun:"pbd_fchr,type:smallint"`     // Powerbuilder catalog information
	PbdFptc       *int16  `bun:"pbd_fptc,type:smallint"`     // Powerbuilder catalog information
	PbdFfce       *string `bun:"pbd_ffce,type:char(18)"`     // Powerbuilder catalog information
	PbhFhgt       *int16  `bun:"pbh_fhgt,type:smallint"`     // Powerbuilder catalog information
	PbhFwgt       *int16  `bun:"pbh_fwgt,type:smallint"`     // Powerbuilder catalog information
	PbhFitl       *string `bun:"pbh_fitl,type:char(1)"`      // Powerbuilder catalog information
	PbhFunl       *string `bun:"pbh_funl,type:char(1)"`      // Powerbuilder catalog information
	PbhFchr       *int16  `bun:"pbh_fchr,type:smallint"`     // Powerbuilder catalog information
	PbhFptc       *int16  `bun:"pbh_fptc,type:smallint"`     // Powerbuilder catalog information
	PbhFfce       *string `bun:"pbh_ffce,type:char(18)"`     // Powerbuilder catalog information
	PblFhgt       *int16  `bun:"pbl_fhgt,type:smallint"`     // Powerbuilder catalog information
	PblFwgt       *int16  `bun:"pbl_fwgt,type:smallint"`     // Powerbuilder catalog information
	PblFitl       *string `bun:"pbl_fitl,type:char(1)"`      // Powerbuilder catalog information
	PblFunl       *string `bun:"pbl_funl,type:char(1)"`      // Powerbuilder catalog information
	PblFchr       *int16  `bun:"pbl_fchr,type:smallint"`     // Powerbuilder catalog information
	PblFptc       *int16  `bun:"pbl_fptc,type:smallint"`     // Powerbuilder catalog information
	PblFfce       *string `bun:"pbl_ffce,type:char(18)"`     // Powerbuilder catalog information
	PbtCmnt       *string `bun:"pbt_cmnt,type:varchar(254)"` // Powerbuilder catalog information
}
