package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type AuditTrailSupport struct {
	bun.BaseModel         `bun:"table:audit_trail_support"`
	AuditTrailSupportUid  int32     `bun:"audit_trail_support_uid,type:int,pk"`                          // Unique identifier for each record
	SourceAreaCd          int32     `bun:"source_area_cd,type:int"`                                      // Code representing transaction or maintenance record where audit_trail record was established
	SourceTableName       string    `bun:"source_table_name,type:varchar(255)"`                          // Used during purge, identifies table to join when additional contraints are appropriate
	CompanyColumnName     string    `bun:"company_column_name,type:varchar(255),nullzero"`               // Used during purge, identifies company column name in source table
	LocationColumnName    string    `bun:"location_column_name,type:varchar(255),nullzero"`              // Used during purge, identifies location column name in source table
	CompletedColumnName   string    `bun:"completed_column_name,type:varchar(255),nullzero"`             // Used during purge, identifies completed column name in source table
	LineNoDisplayFlag     string    `bun:"line_no_display_flag,type:char(1),default:('N')"`              // Used during display, identifies whether to display audit_trail.line_no
	InvMastUidDisplayFlag string    `bun:"inv_mast_uid_display_flag,type:char(1),default:('N')"`         // Used during display, identifies whether to display audit_trail.inv_mast_uid
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date record is created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // Identifies the person who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Last date the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // Identifies the person who last modified the record
}
