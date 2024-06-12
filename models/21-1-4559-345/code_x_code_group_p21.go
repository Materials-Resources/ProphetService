package model

type CodeXCodeGroupP21 struct {
	bun.BaseModel    `bun:"table:code_x_code_group_p21"`
	CodeGroupNo      int16     `bun:"code_group_no,type:smallint,pk"`
	CodeNo           int32     `bun:"code_no,type:int,pk"`
	SubGroupNo       int16     `bun:"sub_group_no,type:smallint,nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
