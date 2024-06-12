package model

type SysParamsP21 struct {
	bun.BaseModel       `bun:"table:sys_params_p21"`
	SysParamsUid        int32     `bun:"sys_params_uid,type:int,pk"`
	ConfigurationId     int32     `bun:"configuration_id,type:int"`
	Module              string    `bun:"module,type:varchar(32)"`
	SysParamName        string    `bun:"sys_param_name,type:varchar(32)"`
	SysParamValue       string    `bun:"sys_param_value,type:varchar(255)"`
	SysParamDataType    string    `bun:"sys_param_data_type,type:varchar(16)"`
	SysParamDescription string    `bun:"sys_param_description,type:varchar(255)"`
	DateCreated         time.Time `bun:"date_created,type:datetime"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30)"`
	CodeGroupNo         int32     `bun:"code_group_no,type:int,default:(0)"`
}
