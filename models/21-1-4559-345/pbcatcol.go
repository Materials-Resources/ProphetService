package model

type Pbcatcol struct {
	bun.BaseModel `bun:"table:pbcatcol"`
	PbcTnam       string `bun:"pbc_tnam,type:char(30),nullzero"`
	PbcTid        int32  `bun:"pbc_tid,type:int,pk"`
	PbcOwnr       string `bun:"pbc_ownr,type:char(30),nullzero"`
	PbcCnam       string `bun:"pbc_cnam,type:char(30),nullzero"`
	PbcCid        int16  `bun:"pbc_cid,type:smallint,pk"`
	PbcLabl       string `bun:"pbc_labl,type:varchar(254),nullzero"`
	PbcLpos       int16  `bun:"pbc_lpos,type:smallint,nullzero"`
	PbcHdr        string `bun:"pbc_hdr,type:varchar(254),nullzero"`
	PbcHpos       int16  `bun:"pbc_hpos,type:smallint,nullzero"`
	PbcJtfy       int16  `bun:"pbc_jtfy,type:smallint,nullzero"`
	PbcMask       string `bun:"pbc_mask,type:varchar(31),nullzero"`
	PbcCase       int16  `bun:"pbc_case,type:smallint,nullzero"`
	PbcHght       int16  `bun:"pbc_hght,type:smallint,nullzero"`
	PbcWdth       int16  `bun:"pbc_wdth,type:smallint,nullzero"`
	PbcPtrn       string `bun:"pbc_ptrn,type:varchar(31),nullzero"`
	PbcBmap       string `bun:"pbc_bmap,type:char,nullzero"`
	PbcInit       string `bun:"pbc_init,type:varchar(254),nullzero"`
	PbcCmnt       string `bun:"pbc_cmnt,type:varchar(254),nullzero"`
	PbcEdit       string `bun:"pbc_edit,type:varchar(31),nullzero"`
	PbcTag        string `bun:"pbc_tag,type:varchar(254),nullzero"`
}
