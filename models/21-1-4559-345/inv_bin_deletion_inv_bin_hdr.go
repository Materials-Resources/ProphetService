package model

import "github.com/uptrace/bun"

type InvBinDeletionInvBinHdr struct {
	bun.BaseModel `bun:"table:inv_bin_deletion_inv_bin_hdr"`
	InvMastUid    int32  `bun:"inv_mast_uid,type:int"`
	LocationId    int32  `bun:"location_id,type:int"`
	Bin           string `bun:"bin,type:varchar(255)"`
}
