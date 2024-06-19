package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CustPartNoGroupLine struct {
	bun.BaseModel          `bun:"table:cust_part_no_group_line"`
	CustPartNoGroupLineUid int32     `bun:"cust_part_no_group_line_uid,type:int,pk"`                      // Sys generated, uniquely distinguishes each row.
	CustPartNoGroupHdrUid  int32     `bun:"cust_part_no_group_hdr_uid,type:int"`                          // Link to Hdr table
	CustomerId             float64   `bun:"customer_id,type:decimal(19,0)"`                               // customer who shares the common CPN
	RowStatusFlag          int32     `bun:"row_status_flag,type:int"`                                     // row status flag
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
