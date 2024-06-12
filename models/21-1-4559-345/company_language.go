package model

type CompanyLanguage struct {
	bun.BaseModel      `bun:"table:company_language"`
	CompanyLanguageUid int32     `bun:"company_language_uid,type:int,pk,identity"`
	CompanyId          string    `bun:"company_id,type:varchar(8)"`
	LanguageId         string    `bun:"language_id,type:varchar(8),nullzero"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
