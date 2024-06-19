package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CrmContactInformation struct {
	bun.BaseModel            `bun:"table:crm_contact_information"`
	CrmContactInformationUid int32     `bun:"crm_contact_information_uid,type:int,autoincrement,pk"`        // Unique identifier for the record.
	CompanyId                string    `bun:"company_id,type:varchar(8),unique,nullzero"`                   // The identifier of the company the contains this record.
	EntityLinkIdChar         string    `bun:"entity_link_id_char,type:varchar(255),unique,nullzero"`        // The character identifier for the entity.
	EntityLinkIdDec          float64   `bun:"entity_link_id_dec,type:decimal(19,0),unique,nullzero"`        // The decimal identifier for the entity.
	EntityTypeCd             int32     `bun:"entity_type_cd,type:int,unique"`                               // Determine the type of this record.
	LastHardTouchDate        time.Time `bun:"last_hard_touch_date,type:datetime"`                           // The last hard-touch date for this records.
	ActivityTransNo          string    `bun:"activity_trans_no,type:varchar(10),nullzero"`                  // The corresponding task that affected this record.
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
