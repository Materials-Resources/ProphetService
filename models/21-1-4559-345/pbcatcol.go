package gen

import "github.com/uptrace/bun"

type Pbcatcol struct {
	bun.BaseModel `bun:"table:pbcatcol"`
	PbcTnam       string `bun:"pbc_tnam,type:char(30),nullzero"`     // Powerbuilder catalog information
	PbcTid        int32  `bun:"pbc_tid,type:int,pk"`                 // Powerbuilder catalog information
	PbcOwnr       string `bun:"pbc_ownr,type:char(30),nullzero"`     // Powerbuilder catalog information
	PbcCnam       string `bun:"pbc_cnam,type:char(30),nullzero"`     // Powerbuilder catalog information
	PbcCid        int16  `bun:"pbc_cid,type:smallint,pk"`            // Powerbuilder catalog information
	PbcLabl       string `bun:"pbc_labl,type:varchar(254),nullzero"` // Powerbuilder catalog information
	PbcLpos       int16  `bun:"pbc_lpos,type:smallint,nullzero"`     // Powerbuilder catalog information
	PbcHdr        string `bun:"pbc_hdr,type:varchar(254),nullzero"`  // Powerbuilder catalog information
	PbcHpos       int16  `bun:"pbc_hpos,type:smallint,nullzero"`     // Powerbuilder catalog information
	PbcJtfy       int16  `bun:"pbc_jtfy,type:smallint,nullzero"`     // Powerbuilder catalog information
	PbcMask       string `bun:"pbc_mask,type:varchar(31),nullzero"`  // Powerbuilder catalog information
	PbcCase       int16  `bun:"pbc_case,type:smallint,nullzero"`     // Powerbuilder catalog information
	PbcHght       int16  `bun:"pbc_hght,type:smallint,nullzero"`     // Powerbuilder catalog information
	PbcWdth       int16  `bun:"pbc_wdth,type:smallint,nullzero"`     // Powerbuilder catalog information
	PbcPtrn       string `bun:"pbc_ptrn,type:varchar(31),nullzero"`  // Powerbuilder catalog information
	PbcBmap       string `bun:"pbc_bmap,type:char(1),nullzero"`      // Powerbuilder catalog information
	PbcInit       string `bun:"pbc_init,type:varchar(254),nullzero"` // Powerbuilder catalog information
	PbcCmnt       string `bun:"pbc_cmnt,type:varchar(254),nullzero"` // Powerbuilder catalog information
	PbcEdit       string `bun:"pbc_edit,type:varchar(31),nullzero"`  // Powerbuilder catalog information
	PbcTag        string `bun:"pbc_tag,type:varchar(254),nullzero"`  // Powerbuilder catalog information
}
