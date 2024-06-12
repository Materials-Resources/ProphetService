package model

import "github.com/uptrace/bun"

type WebImages struct {
	bun.BaseModel `bun:"table:web_images"`
	InvMastUid    int32  `bun:"inv_mast_uid,type:int,pk"`
	ImageId       string `bun:"image_id,type:varchar(60),nullzero"`
	ImageUrl      string `bun:"image_url,type:varchar(120),nullzero"`
}
