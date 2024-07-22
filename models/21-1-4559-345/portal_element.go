package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PortalElement struct {
	bun.BaseModel     `bun:"table:portal_element"`
	PortalElementUid  int32     `bun:"portal_element_uid,type:int,pk"`                               // Unique ID for the portal element
	PortalElementName string    `bun:"portal_element_name,type:varchar(255),unique"`                 // Name of the portal element
	Classname         string    `bun:"classname,type:varchar(255)"`                                  // Class name (e.g. DW name) for the portal element
	RowStatusFlag     int32     `bun:"row_status_flag,type:int,default:((704))"`                     // Indicates the status of the record - Active/Delete
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	IconName          *string   `bun:"icon_name,type:varchar(255)"`                                  // Icon name the portal element
	PortalCd          *int32    `bun:"portal_cd,type:int"`                                           // Code from code_p21 for portal.
	ReportMetadataUid *int32    `bun:"report_metadata_uid,type:int"`                                 // Unique Identifier for Report Metadata
}
