package model

type UsersXApplicationSecurity struct {
	bun.BaseModel          `bun:"table:users_x_application_security"`
	UsersXAppSecurityUid   int32     `bun:"users_x_app_security_uid,type:int,pk,identity"`
	UsersId                string    `bun:"users_id,type:varchar(30)"`
	ApplicationSecurityUid int32     `bun:"application_security_uid,type:int"`
	CodeValue              int32     `bun:"code_value,type:int,nullzero"`
	DecimalValue           float64   `bun:"decimal_value,type:decimal(19,9),nullzero"`
	StringValue            string    `bun:"string_value,type:varchar(255),nullzero"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
