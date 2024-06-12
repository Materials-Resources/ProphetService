package model

type DocumentLinkKey struct {
	bun.BaseModel      `bun:"table:document_link_key"`
	DocumentLinkKeyUid int32     `bun:"document_link_key_uid,type:int,pk"`
	AreaCd             int32     `bun:"area_cd,type:int"`
	Key1Cd             string    `bun:"key1_cd,type:varchar(255),nullzero"`
	Key1Type           string    `bun:"key1_type,type:varchar(5),nullzero"`
	Key2Cd             string    `bun:"key2_cd,type:varchar(255),nullzero"`
	Key2Type           string    `bun:"key2_type,type:varchar(5),nullzero"`
	Key3Cd             string    `bun:"key3_cd,type:varchar(255),nullzero"`
	Key3Type           string    `bun:"key3_type,type:varchar(5),nullzero"`
	RowStatusFlag      int32     `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
