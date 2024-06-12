package model

type WeboeCode struct {
	bun.BaseModel `bun:"table:weboe_code"`
	WeboeCodeUid  int32  `bun:"weboe_code_uid,type:int"`
	WeboeCodeDesc string `bun:"weboe_code_desc,type:varchar(255)"`
}
