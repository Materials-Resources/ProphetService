package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ReportKeyword struct {
	bun.BaseModel    `bun:"table:report_keyword"`
	ReportKeywordUid int32     `bun:"report_keyword_uid,type:int,autoincrement,identity,pk"`        // Primary key
	Keyword          string    `bun:"keyword,type:varchar(250)"`                                    // Name of the keyword
	Statement        string    `bun:"statement,type:varchar(250)"`                                  // column that thave the name of the column or function for the keyword
	Alias            string    `bun:"alias,type:varchar(250)"`                                      // name used to display in the dataset
	Description      string    `bun:"description,type:varchar(250)"`                                // decription of what the keyword contains
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
