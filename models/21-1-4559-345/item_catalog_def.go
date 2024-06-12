package model

type ItemCatalogDef struct {
	bun.BaseModel     `bun:"table:item_catalog_def"`
	ItemCatalogDefUid int32     `bun:"item_catalog_def_uid,type:int,pk,identity"`
	ColumnId          string    `bun:"column_id,type:varchar(255)"`
	ColumnName        string    `bun:"column_name,type:varchar(255),nullzero"`
	ItemSearchFlag    string    `bun:"item_search_flag,type:char,default:('N')"`
	DisplayColumnId   string    `bun:"display_column_id,type:varchar(255)"`
	RowStatusFlag     int32     `bun:"row_status_flag,type:int,default:((704))"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
