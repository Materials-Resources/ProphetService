package model

type DocumentLinkDocstar struct {
	bun.BaseModel          `bun:"table:document_link_docstar"`
	DocumentLinkDocstarUid int32     `bun:"document_link_docstar_uid,type:int,pk,identity"`
	SourceAreaCd           int32     `bun:"source_area_cd,type:int"`
	DocumentAreaCd         string    `bun:"document_area_cd,type:varchar(255)"`
	DisplayAreaCd          int32     `bun:"display_area_cd,type:int"`
	Key1Cd                 string    `bun:"key1_cd,type:varchar(255)"`
	Key1ValueField         string    `bun:"key1_value_field,type:varchar(255)"`
	Key2Cd                 string    `bun:"key2_cd,type:varchar(255),nullzero"`
	Key2ValueField         string    `bun:"key2_value_field,type:varchar(255),nullzero"`
	Key3Cd                 string    `bun:"key3_cd,type:varchar(255),nullzero"`
	Key3ValueField         string    `bun:"key3_value_field,type:varchar(255),nullzero"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
