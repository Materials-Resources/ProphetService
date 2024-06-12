package model

type EscBaseViewAlias struct {
	bun.BaseModel       `bun:"table:esc_base_view_alias"`
	EscBaseViewAliasUid int32     `bun:"esc_base_view_alias_uid,type:int,pk,identity"`
	EscBaseViewName     string    `bun:"esc_base_view_name,type:varchar(255)"`
	BaseTable           string    `bun:"base_table,type:varchar(255)"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
