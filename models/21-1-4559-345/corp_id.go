package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CorpId struct {
	bun.BaseModel    `bun:"table:corp_id"`
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	AddressId        float64   `bun:"address_id,type:decimal(19,0),pk"`                          // What is the address of this contact?
	CreditLimit      float64   `bun:"credit_limit,type:decimal(19,4)"`                           // Credit limit available for this corp
	CreditLimitUsed  float64   `bun:"credit_limit_used,type:decimal(19,4)"`                      // Credit limit used for this corp
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	AddressName      *string   `bun:"address_name,type:varchar(255)"`                            // Name of the corp
}
