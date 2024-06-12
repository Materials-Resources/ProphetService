package model

type DatasourceXUsers struct {
	bun.BaseModel       `bun:"table:datasource_x_users"`
	DatasourceXUsersUid int32     `bun:"datasource_x_users_uid,type:int,pk,identity"`
	DatasourceUid       int32     `bun:"datasource_uid,type:int"`
	UsersId             string    `bun:"users_id,type:varchar(30)"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
