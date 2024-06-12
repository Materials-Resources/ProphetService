package model

type NavigationIndex struct {
	bun.BaseModel      `bun:"table:navigation_index"`
	NavigationIndexUid int32     `bun:"navigation_index_uid,type:int,pk,identity"`
	UsersId            string    `bun:"users_id,type:varchar(30)"`
	MenuName           string    `bun:"menu_name,type:varchar(2000)"`
	MenuText           string    `bun:"menu_text,type:varchar(255)"`
	RowStatusFlag      int32     `bun:"row_status_flag,type:int,default:((704))"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	WindowName         string    `bun:"window_name,type:varchar(255),nullzero"`
}
