package model

type AiaElement struct {
	bun.BaseModel    `bun:"table:aia_element"`
	AiaElementUid    int32     `bun:"aia_element_uid,type:int,pk,identity"`
	Description      string    `bun:"description,type:varchar(255)"`
	PointValue       float64   `bun:"point_value,type:decimal(19,9),default:((0))"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
