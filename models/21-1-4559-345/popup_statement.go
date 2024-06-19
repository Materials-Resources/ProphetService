package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PopupStatement struct {
	bun.BaseModel           `bun:"table:popup_statement"`
	PopupStatementUid       int32     `bun:"popup_statement_uid,type:int,autoincrement,scanonly,pk"`       // Popup Statement Id
	Columns                 string    `bun:"columns,type:varchar(max)"`                                    // Select statement for popup
	FromJoin                string    `bun:"from_join,type:varchar(max)"`                                  // From-Join statement for popup
	Where                   string    `bun:"where,type:varchar(max)"`                                      // Where statement for popup
	OrderBy                 string    `bun:"order_by,type:varchar(max)"`                                   // Order By statement for popup
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	PopupDetailUid          int32     `bun:"popup_detail_uid,type:int"`                                    // Popup detail uid
	GroupBy                 string    `bun:"group_by,type:varchar(max),nullzero"`                          // Statement's group by clause
	OverrideColumns         string    `bun:"override_columns,type:char(3),nullzero"`                       // flag to override columns for Dynachange
	OverrideFromJoin        string    `bun:"override_from_join,type:char(3),nullzero"`                     // flag to override from join for Dynachange
	OverrideWhere           string    `bun:"override_where,type:char(3),nullzero"`                         // flag to override where for Dynachange
	OverrideGroupBy         string    `bun:"override_group_by,type:char(3),nullzero"`                      // flag to override order by for Dynachange
	PopupStatementUidParent int32     `bun:"popup_statement_uid_parent,type:int,nullzero"`                 // popup statement uid to manage the aditive for the dynachange
	OverrideOrderBy         string    `bun:"override_order_by,type:char(3),nullzero"`                      // flag to override order by for Dynachange
	Option                  string    `bun:"option,type:varchar(max),default:(NULL)"`                      // Store an option clause
	OverrideOption          bool      `bun:"override_option,type:bit,default:((0))"`                       // Identify if the popup should load the option clause from its child
}
