package model

type Language struct {
	bun.BaseModel        `bun:"table:language"`
	LanguageId           string    `bun:"language_id,type:varchar(8),pk"`
	LanguageDescription  string    `bun:"language_description,type:varchar(30)"`
	DeleteFlag           string    `bun:"delete_flag,type:char"`
	DateCreated          time.Time `bun:"date_created,type:datetime"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	LanguageUid          int32     `bun:"language_uid,type:int,identity"`
	CrystalReportsFolder string    `bun:"crystal_reports_folder,type:char(255),nullzero"`
	IsoCode              string    `bun:"iso_code,type:varchar(5),nullzero"`
}
