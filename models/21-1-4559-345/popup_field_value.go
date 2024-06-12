package model

type PopupFieldValue struct {
	bun.BaseModel      `bun:"table:popup_field_value"`
	PopupFieldValueUid int32     `bun:"popup_field_value_uid,type:int,pk,identity"`
	PopupFieldUid      int32     `bun:"popup_field_uid,type:int"`
	Value              string    `bun:"value,type:varchar(255),nullzero"`
	Condition          string    `bun:"condition,type:varchar(5000),nullzero"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
