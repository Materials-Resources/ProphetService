package model

type LinkQuantity struct {
	bun.BaseModel    `bun:"table:link_quantity"`
	LinkQuantityUid  int32     `bun:"link_quantity_uid,type:int,pk,identity"`
	FromUid          int32     `bun:"from_uid,type:int"`
	FromTypeCd       int32     `bun:"from_type_cd,type:int"`
	ToUid            int32     `bun:"to_uid,type:int"`
	ToTypeCd         int32     `bun:"to_type_cd,type:int"`
	Quantity         float64   `bun:"quantity,type:decimal(19,9)"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
