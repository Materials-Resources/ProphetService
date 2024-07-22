package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Remittances struct {
	bun.BaseModel    `bun:"table:remittances"`
	RemittanceNumber float64   `bun:"remittance_number,type:decimal(19,0),pk"`                       // What remittance does this remittance detail belong to?
	OrderNo          string    `bun:"order_no,type:varchar(8)"`                                      // What order does this invoice belong to?
	CashDrawerId     string    `bun:"cash_drawer_id,type:varchar(8)"`                                // Identifier for drawer
	CompanyId        string    `bun:"company_id,type:varchar(8)"`                                    // Unique code that identifies a company.
	PaymentNumber    float64   `bun:"payment_number,type:decimal(19,0)"`                             // Unique identifier - system generated
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	InCashDrawer     string    `bun:"in_cash_drawer,type:char(1)"`                                   // Is the remittance in the cash drawer?
	CashDrawerSeqNo  int32     `bun:"cash_drawer_seq_no,type:int,default:(0)"`                       // Cash Drawer Sequence Number
}
