package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CustPartNoGroupHdr struct {
	bun.BaseModel         `bun:"table:cust_part_no_group_hdr"`
	CustPartNoGroupHdrUid int32     `bun:"cust_part_no_group_hdr_uid,type:int,pk"`                       // System generated. Uniquely distinguishes each row.
	CustPartNoGroupNo     string    `bun:"cust_part_no_group_no,type:varchar(255)"`                      // Unique & User defined
	CustPartNoGroupDesc   string    `bun:"cust_part_no_group_desc,type:varchar(255)"`                    // User defined group desc.
	CompanyId             string    `bun:"company_id,type:varchar(8)"`                                   // Company ID
	RowStatusFlag         int32     `bun:"row_status_flag,type:int"`                                     // Row status flag
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
