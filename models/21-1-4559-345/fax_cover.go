package model

type FaxCover struct {
	bun.BaseModel    `bun:"table:fax_cover"`
	FaxCoverUid      int32     `bun:"fax_cover_uid,type:int,pk,identity"`
	FaxCoverId       string    `bun:"fax_cover_id,type:varchar(255)"`
	FaxCoverDesc     string    `bun:"fax_cover_desc,type:varchar(255),nullzero"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
