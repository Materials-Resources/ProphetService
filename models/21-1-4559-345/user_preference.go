package model

type UserPreference struct {
	bun.BaseModel     `bun:"table:user_preference"`
	UserPreferenceUid int32     `bun:"user_preference_uid,type:int,pk,identity"`
	UserId            string    `bun:"user_id,type:varchar(30)"`
	PreferenceUid     int32     `bun:"preference_uid,type:int"`
	Value             string    `bun:"value,type:varchar"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
