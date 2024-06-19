package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceClass struct {
	bun.BaseModel    `bun:"table:invoice_class"`
	InvoiceClass     string    `bun:"invoice_class,type:varchar(8),pk"`                          // What is the unique identifier for this invoice class?
	InvoiceClassDesc string    `bun:"invoice_class_desc,type:varchar(30)"`                       // How would you describe this invoice class?
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
