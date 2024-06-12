package model

type ScriptCommand struct {
	bun.BaseModel      `bun:"table:script_command"`
	ScriptCommandUid   int32 `bun:"script_command_uid,type:int,pk,identity"`
	CommandName        `bun:"command_name,type:nvarchar(50)"`
	CommandDescription `bun:"command_description,type:nvarchar(255),nullzero"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
