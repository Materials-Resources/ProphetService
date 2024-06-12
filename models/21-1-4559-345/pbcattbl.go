package model

type Pbcattbl struct {
	bun.BaseModel `bun:"table:pbcattbl"`
	PbtTnam       string `bun:"pbt_tnam,type:char(30),nullzero"`
	PbtTid        int32  `bun:"pbt_tid,type:int,pk"`
	PbtOwnr       string `bun:"pbt_ownr,type:char(30),nullzero"`
	PbdFhgt       int16  `bun:"pbd_fhgt,type:smallint,nullzero"`
	PbdFwgt       int16  `bun:"pbd_fwgt,type:smallint,nullzero"`
	PbdFitl       string `bun:"pbd_fitl,type:char,nullzero"`
	PbdFunl       string `bun:"pbd_funl,type:char,nullzero"`
	PbdFchr       int16  `bun:"pbd_fchr,type:smallint,nullzero"`
	PbdFptc       int16  `bun:"pbd_fptc,type:smallint,nullzero"`
	PbdFfce       string `bun:"pbd_ffce,type:char(18),nullzero"`
	PbhFhgt       int16  `bun:"pbh_fhgt,type:smallint,nullzero"`
	PbhFwgt       int16  `bun:"pbh_fwgt,type:smallint,nullzero"`
	PbhFitl       string `bun:"pbh_fitl,type:char,nullzero"`
	PbhFunl       string `bun:"pbh_funl,type:char,nullzero"`
	PbhFchr       int16  `bun:"pbh_fchr,type:smallint,nullzero"`
	PbhFptc       int16  `bun:"pbh_fptc,type:smallint,nullzero"`
	PbhFfce       string `bun:"pbh_ffce,type:char(18),nullzero"`
	PblFhgt       int16  `bun:"pbl_fhgt,type:smallint,nullzero"`
	PblFwgt       int16  `bun:"pbl_fwgt,type:smallint,nullzero"`
	PblFitl       string `bun:"pbl_fitl,type:char,nullzero"`
	PblFunl       string `bun:"pbl_funl,type:char,nullzero"`
	PblFchr       int16  `bun:"pbl_fchr,type:smallint,nullzero"`
	PblFptc       int16  `bun:"pbl_fptc,type:smallint,nullzero"`
	PblFfce       string `bun:"pbl_ffce,type:char(18),nullzero"`
	PbtCmnt       string `bun:"pbt_cmnt,type:varchar(254),nullzero"`
}
