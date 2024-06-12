package model

type FrameMenu struct {
	bun.BaseModel             `bun:"table:frame_menu"`
	FrameMenuUid              int32     `bun:"frame_menu_uid,type:int,pk,identity"`
	Handle                    int32     `bun:"handle,type:int"`
	ParentHandle              int32     `bun:"parent_handle,type:int,nullzero"`
	ParentClassName           string    `bun:"parent_class_name,type:varchar(255),nullzero"`
	ModuleHandle              int32     `bun:"module_handle,type:int"`
	ClassName                 string    `bun:"class_name,type:varchar(255)"`
	FrameName                 string    `bun:"frame_name,type:varchar(255)"`
	Tag                       string    `bun:"tag,type:varchar(255),nullzero"`
	TextLabel                 string    `bun:"text_label,type:varchar(255)"`
	Enabled                   string    `bun:"enabled,type:char,default:('Y')"`
	ToolbarItemImage          string    `bun:"toolbar_item_image,type:varchar(255),nullzero"`
	ToolbarItemText           string    `bun:"toolbar_item_text,type:varchar(255),nullzero"`
	MenuImage                 string    `bun:"menu_image,type:varchar(255),nullzero"`
	Microhelp                 string    `bun:"microhelp,type:varchar(8000),nullzero"`
	ToolbarItemBarIndex       int16     `bun:"toolbar_item_bar_index,type:smallint,nullzero"`
	ToolbarItemOrder          int16     `bun:"toolbar_item_order,type:smallint,nullzero"`
	Stringparm                string    `bun:"stringparm,type:varchar(255),nullzero"`
	StringparmAddtl           string    `bun:"stringparm_addtl,type:varchar(255),nullzero"`
	DynachangeId              int32     `bun:"dynachange_id,type:int,nullzero"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	UseClickedService         string    `bun:"use_clicked_service,type:char,default:('N')"`
	WindowRole                int32     `bun:"window_role,type:int,nullzero"`
	UseCompanyRequiredMessage string    `bun:"use_company_required_message,type:char,default:('N')"`
	UserDefined               string    `bun:"user_defined,type:char,default:('N')"`
	TitleExpression           string    `bun:"title_expression,type:varchar(8000),nullzero"`
	ServiceName               string    `bun:"service_name,type:varchar(255),nullzero"`
	FasteditType              string    `bun:"fastedit_type,type:char,nullzero"`
	MenuContext               int32     `bun:"menu_context,type:int,default:((3))"`
	StringparmUrl             string    `bun:"stringparm_url,type:varchar(8000),nullzero"`
	ServiceKey                string    `bun:"service_key,type:varchar(8000),nullzero"`
	SchedulerEnabled          string    `bun:"scheduler_enabled,type:char,nullzero"`
	ReportMetadataUid         int32     `bun:"report_metadata_uid,type:int,nullzero"`
}
