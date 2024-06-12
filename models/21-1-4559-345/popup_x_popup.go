package model

type PopupXPopup struct {
	bun.BaseModel       `bun:"table:popup_x_popup"`
	PopupXPopupUid      int32     `bun:"popup_x_popup_uid,type:int,pk,identity"`
	PopupIndexUid       int32     `bun:"popup_index_uid,type:int"`
	LinkedPopupIndexUid int32     `bun:"linked_popup_index_uid,type:int"`
	Order               int32     `bun:"order,type:int"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
