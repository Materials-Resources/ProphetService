package model

type DcNavDrill struct {
	bun.BaseModel        `bun:"table:dc_nav_drill"`
	DcNavDrillUid        int32     `bun:"dc_nav_drill_uid,type:int,pk,identity"`
	SourceWindow         string    `bun:"source_window,type:varchar(255)"`
	SourceDatawindow     string    `bun:"source_datawindow,type:varchar(255),nullzero"`
	SourceField          string    `bun:"source_field,type:varchar(255)"`
	DestWindow           string    `bun:"dest_window,type:varchar(255)"`
	DestDatawindow       string    `bun:"dest_datawindow,type:varchar(255),nullzero"`
	DestField            string    `bun:"dest_field,type:varchar(255)"`
	RowStatusFlag        int32     `bun:"row_status_flag,type:int,default:((704))"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	TypeCd               int32     `bun:"type_cd,type:int,default:((1418))"`
	ApplyDrillToAllUsers string    `bun:"apply_drill_to_all_users,type:char,default:('Y')"`
	DestWindowName       string    `bun:"dest_window_name,type:varchar(255),nullzero"`
	SourceDataField      string    `bun:"source_data_field,type:varchar(255),nullzero"`
	NavigationType       int32     `bun:"navigation_type,type:int,default:((3590))"`
}
