package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type OeLineTax struct {
	bun.BaseModel    `bun:"table:oe_line_tax"`
	OrderNumber      string    `bun:"order_number,type:varchar(8),pk"`                               // What order is this for?
	LineNumber       float64   `bun:"line_number,type:decimal(19,0),pk"`                             // The line number on the order.
	JurisdictionId   string    `bun:"jurisdiction_id,type:varchar(10),pk"`                           // What is the unique identifier for this tax jurisdiction?
	Taxable          string    `bun:"taxable,type:char(1)"`                                          // Is this line item taxable?
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	LineType         *int32    `bun:"line_type,type:int"`                                            // Indicates any special type of oe_line sub-type, such as a part or labor line.
}
