package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceTypes struct {
	bun.BaseModel          `bun:"table:invoice_types"`
	InvoiceTypeId          string    `bun:"invoice_type_id,type:varchar(2),pk"`                        // Identifier of the invoice type.
	InvoiceTypeDescription string    `bun:"invoice_type_description,type:varchar(30)"`                 // This column is unused.
	DeleteFlag             string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated            time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	Object                 string    `bun:"object,type:varchar(50),nullzero"`                          // This column is unused.
}
