package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PopupColumn struct {
	bun.BaseModel               `bun:"table:popup_column"`
	PopupColumnUid              int32     `bun:"popup_column_uid,type:int,autoincrement,pk"`                   // Unique identifier for popup_column
	ColumnName                  string    `bun:"column_name,type:varchar(255)"`                                // Defines the column name for the popup window
	DatawindowName              string    `bun:"datawindow_name,type:varchar(255),nullzero"`                   // Datawindow name for the popup
	AutoRetrieveFlag            string    `bun:"auto_retrieve_flag,type:char(1),nullzero"`                     // Enables auto retrieve of data
	AllowAdvancedSearchFlag     string    `bun:"allow_advanced_search_flag,type:char(1),nullzero"`             // Enables allow advanced search
	CompanySecurityFlag         string    `bun:"company_security_flag,type:char(1),nullzero"`                  // Enables company security
	CompanySecurityAdvancedFlag string    `bun:"company_security_advanced_flag,type:char(1),nullzero"`         // Enables advanced company security
	ColumnToReturn              string    `bun:"column_to_return,type:varchar(255),nullzero"`                  // Column to return
	RedirectToColumn            string    `bun:"redirect_to_column,type:varchar(255),nullzero"`                // Redirect to column
	DeveloperNotes              string    `bun:"developer_notes,type:varchar(8000),nullzero"`                  // Developer notes
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	DefaultSortColumn           string    `bun:"default_sort_column,type:varchar(255),nullzero"`               // Defaut Sort Column
	SortDescending              string    `bun:"sort_descending,type:char(1),nullzero"`                        // If set to Y, indicates that the popup should be sorted in descending order.
}
