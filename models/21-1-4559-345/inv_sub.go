package model

type InvSub struct {
	bun.BaseModel    `bun:"table:inv_sub"`
	InvMastUid       int32     `bun:"inv_mast_uid,type:int,pk"`
	SubInvMastUid    int32     `bun:"sub_inv_mast_uid,type:int,pk"`
	SubDesc          string    `bun:"sub_desc,type:varchar(40)"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	Interchangeable  string    `bun:"interchangeable,type:char"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
}
