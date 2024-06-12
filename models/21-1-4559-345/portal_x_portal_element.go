package model

type PortalXPortalElement struct {
	bun.BaseModel           `bun:"table:portal_x_portal_element"`
	PortalXPortalElementUid int32     `bun:"portal_x_portal_element_uid,type:int,pk,identity"`
	PortalUid               int32     `bun:"portal_uid,type:int"`
	PortalElementUid        int32     `bun:"portal_element_uid,type:int"`
	SequenceNo              int32     `bun:"sequence_no,type:int"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
