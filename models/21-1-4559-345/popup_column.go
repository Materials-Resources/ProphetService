package model

type PopupColumn struct {
	bun.BaseModel               `bun:"table:popup_column"`
	PopupColumnUid              int32     `bun:"popup_column_uid,type:int,pk,identity"`
	ColumnName                  string    `bun:"column_name,type:varchar(255)"`
	DatawindowName              string    `bun:"datawindow_name,type:varchar(255),nullzero"`
	AutoRetrieveFlag            string    `bun:"auto_retrieve_flag,type:char,nullzero"`
	AllowAdvancedSearchFlag     string    `bun:"allow_advanced_search_flag,type:char,nullzero"`
	CompanySecurityFlag         string    `bun:"company_security_flag,type:char,nullzero"`
	CompanySecurityAdvancedFlag string    `bun:"company_security_advanced_flag,type:char,nullzero"`
	ColumnToReturn              string    `bun:"column_to_return,type:varchar(255),nullzero"`
	RedirectToColumn            string    `bun:"redirect_to_column,type:varchar(255),nullzero"`
	DeveloperNotes              string    `bun:"developer_notes,type:varchar(8000),nullzero"`
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DefaultSortColumn           string    `bun:"default_sort_column,type:varchar(255),nullzero"`
	SortDescending              string    `bun:"sort_descending,type:char,nullzero"`
}
