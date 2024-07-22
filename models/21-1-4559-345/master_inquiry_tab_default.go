package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type MasterInquiryTabDefault struct {
	bun.BaseModel              `bun:"table:master_inquiry_tab_default"`
	MasterInquiryTabDefaultUid int32     `bun:"master_inquiry_tab_default_uid,type:int,autoincrement,identity,pk"` // Unique indentifier for the table
	MasterInquiryTypeCd        int32     `bun:"master_inquiry_type_cd,type:int,unique"`                            // Denotes the window i.e. Supplier or Customer Master Inquiry
	UserId                     string    `bun:"user_id,type:varchar(30),unique"`                                   // The user who specified the preference
	DefaultTabIndex            int16     `bun:"default_tab_index,type:smallint"`                                   // The number of the preferred tabpage
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`                    // Date and time the record was originally created
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`              // User who created the record
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`              // Date and time the record was modified
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`      // User who last changed the record
}
