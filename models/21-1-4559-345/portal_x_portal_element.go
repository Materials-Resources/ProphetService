package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PortalXPortalElement struct {
	bun.BaseModel           `bun:"table:portal_x_portal_element"`
	PortalXPortalElementUid int32     `bun:"portal_x_portal_element_uid,type:int,autoincrement,identity,pk"` // Unique ID for the portal/portal_element link
	PortalUid               int32     `bun:"portal_uid,type:int,unique"`                                     // Portal UID
	PortalElementUid        int32     `bun:"portal_element_uid,type:int"`                                    // Portal Element UID
	SequenceNo              int32     `bun:"sequence_no,type:int,unique"`                                    // The sequence of this element on the portal
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`                 // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`           // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`           // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`   // User who last changed the record
}
