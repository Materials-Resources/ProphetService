package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PortalParamDef struct {
	bun.BaseModel     `bun:"table:portal_param_def"`
	PortalParamDefUid int32     `bun:"portal_param_def_uid,type:int,pk"`                             // Unique ID for a portal element's parameter
	PortalElementUid  int32     `bun:"portal_element_uid,type:int,unique"`                           // Unique ID for the portal element
	SequenceNo        int32     `bun:"sequence_no,type:int,unique"`                                  // Sequence of parameter for the portal
	ParameterName     string    `bun:"parameter_name,type:varchar(255)"`                             // Name of the parameter (English)
	ParameterDesc     string    `bun:"parameter_desc,type:varchar(255)"`                             // Portal description
	DatatypeCd        int32     `bun:"datatype_cd,type:int"`                                         // Datatype of the portal element parameter
	DefaultValue      string    `bun:"default_value,type:varchar(255)"`                              // Default value for any user
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
