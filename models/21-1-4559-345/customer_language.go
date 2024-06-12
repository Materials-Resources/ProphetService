package model

type CustomerLanguage struct {
	bun.BaseModel       `bun:"table:customer_language"`
	CustomerLanguageUid int32     `bun:"customer_language_uid,type:int,pk,identity"`
	CompanyId           string    `bun:"company_id,type:varchar(8)"`
	CustomerId          float64   `bun:"customer_id,type:decimal(19,0)"`
	LanguageId          string    `bun:"language_id,type:varchar(8)"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
