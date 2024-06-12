package model

type ClassNameXFormCode struct {
	bun.BaseModel         `bun:"table:class_name_x_form_code"`
	ClassNameXFormCodeUid int32     `bun:"class_name_x_form_code_uid,type:int,pk,identity"`
	ClassName             string    `bun:"class_name,type:varchar(255)"`
	FormCode              int32     `bun:"form_code,type:int"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
