package model

type UserWindowPref struct {
	bun.BaseModel     `bun:"table:user_window_pref"`
	UserWindowPrefUid int32     `bun:"user_window_pref_uid,type:int,pk,identity"`
	UserId            string    `bun:"user_id,type:varchar(30)"`
	WindowCd          int32     `bun:"window_cd,type:int"`
	ObjectName        string    `bun:"object_name,type:varchar(255)"`
	ObjectProperty    string    `bun:"object_property,type:varchar(255)"`
	ObjectValue       string    `bun:"object_value,type:varchar(8000)"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	WindowName        string    `bun:"window_name,type:varchar(255),nullzero"`
}
