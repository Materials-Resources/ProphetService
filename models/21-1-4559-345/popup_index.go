package model

type PopupIndex struct {
	bun.BaseModel    `bun:"table:popup_index"`
	PopupIndexUid    int32     `bun:"popup_index_uid,type:int,pk,identity"`
	Dwcontrol        string    `bun:"dwcontrol,type:varchar(50),nullzero"`
	Window           string    `bun:"window,type:varchar(50),nullzero"`
	Dwfield          string    `bun:"dwfield,type:varchar(50),nullzero"`
	Role             string    `bun:"role,type:varchar(50),nullzero"`
	Condition        string    `bun:"condition,type:varchar(255),nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	PopupDetailUid   int32     `bun:"popup_detail_uid,type:int"`
	UserRole         string    `bun:"user_role,type:varchar(50),nullzero"`
	Dataobject       string    `bun:"dataobject,type:varchar(255),nullzero"`
	ConfigurationId  int32     `bun:"configuration_id,type:int,default:((0))"`
	Action           string    `bun:"action,type:varchar(50),nullzero"`
	UserId           string    `bun:"user_id,type:varchar(30),nullzero"`
	Shortcut         string    `bun:"shortcut,type:varchar(30),nullzero"`
	CoreFlag         `bun:"core_flag,type:bit,default:((0))"`
	DynachangeId     float64 `bun:"dynachange_id,type:decimal(6,0),nullzero"`
}
