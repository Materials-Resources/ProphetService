package model

type FrameMenuReporting struct {
	bun.BaseModel         `bun:"table:frame_menu_reporting"`
	FrameMenuReportingUid int32     `bun:"frame_menu_reporting_uid,type:int,pk,identity"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	FrameMenuUid          int32     `bun:"frame_menu_uid,type:int"`
	ReportSyntax          string    `bun:"report_syntax,type:varchar,nullzero"`
}
