package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerEmailDefaults struct {
	bun.BaseModel            `bun:"table:customer_email_defaults"`
	CustomerEmailDefaultsUid int32     `bun:"customer_email_defaults_uid,type:int,autoincrement,identity,pk"` // UID
	CompanyId                string    `bun:"company_id,type:varchar(8)"`                                     // Company ID
	CustomerId               float64   `bun:"customer_id,type:decimal(19,0)"`                                 // Customer ID
	FormTypeCd               int32     `bun:"form_type_cd,type:int"`                                          // Form type
	Subject                  string    `bun:"subject,type:varchar(255),nullzero"`                             // Email subject
	Copy                     string    `bun:"copy,type:varchar(255),nullzero"`                                // Email copy list
	BlindCopy                string    `bun:"blind_copy,type:varchar(255),nullzero"`                          // Email blind copy list
	Memo                     string    `bun:"memo,type:varchar(255),nullzero"`                                // Email memo
	CcUser                   string    `bun:"cc_user,type:char(1),nullzero"`                                  // Copy user
	CcSalesrep               string    `bun:"cc_salesrep,type:char(1),nullzero"`                              // Copy salesrep
	CcTaker                  string    `bun:"cc_taker,type:char(1),nullzero"`                                 // Copy taker
	BcUser                   string    `bun:"bc_user,type:char(1),nullzero"`                                  // Blind copy user
	BcSalesrep               string    `bun:"bc_salesrep,type:char(1),nullzero"`                              // Blind copy salesrep
	BcTaker                  string    `bun:"bc_taker,type:char(1),nullzero"`                                 // Blind copy taker
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`                 // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`           // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`           // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`   // User who last changed the record
}
