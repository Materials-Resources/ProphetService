package model

type PopupFieldBehavior struct {
	bun.BaseModel         `bun:"table:popup_field_behavior"`
	PopupFieldBehaviorUid int32  `bun:"popup_field_behavior_uid,type:int,pk,identity"`
	PopupFieldUid         int32  `bun:"popup_field_uid,type:int"`
	Condition             string `bun:"condition,type:varchar"`
	ConfigurationId       int32  `bun:"configuration_id,type:int,default:((0))"`
	Sort                  `bun:"sort,type:bit"`
	Filter                `bun:"filter,type:bit"`
	Visible               `bun:"visible,type:bit"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	AsParameter           `bun:"as_parameter,type:bit,default:((0))"`
	ParameterDefaultValue string `bun:"parameter_default_value,type:varchar(255),nullzero"`
}
