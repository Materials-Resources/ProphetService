package model

type RfKeysDefault struct {
	bun.BaseModel    `bun:"table:rf_keys_default"`
	KeyhelpUid       int32     `bun:"keyhelp_uid,type:int,pk,identity"`
	KeyhelpCode      int32     `bun:"keyhelp_code,type:int"`
	WwmsWindowCode   int32     `bun:"wwms_window_code,type:int"`
	WwmsScreenCode   int32     `bun:"wwms_screen_code,type:int"`
	Keystroke        string    `bun:"keystroke,type:varchar(255)"`
	Description      string    `bun:"description,type:varchar(255)"`
	EventName        string    `bun:"event_name,type:varchar(255),nullzero"`
	KeystrokeDesc    string    `bun:"keystroke_desc,type:varchar(255),nullzero"`
	KeystrokeMod     string    `bun:"keystroke_mod,type:varchar(255),nullzero"`
	ImageFile        string    `bun:"image_file,type:varchar(255)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:('P21_DBA')"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:('P21_DBA')"`
}
