package model

type Token struct {
	bun.BaseModel    `bun:"table:token"`
	TokenUid         int32     `bun:"token_uid,type:int,pk"`
	Name             string    `bun:"name,type:varchar(40)"`
	Description      string    `bun:"description,type:varchar(80)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	DataTypeCd       int32     `bun:"data_type_cd,type:int,nullzero"`
	AvailableAreas   int32     `bun:"available_areas,type:int,nullzero"`
	CodeGroupNo      int16     `bun:"code_group_no,type:smallint,nullzero"`
	UserLookupFlag   string    `bun:"user_lookup_flag,type:char,nullzero"`
}
