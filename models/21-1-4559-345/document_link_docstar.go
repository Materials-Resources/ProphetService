package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DocumentLinkDocstar struct {
	bun.BaseModel          `bun:"table:document_link_docstar"`
	DocumentLinkDocstarUid int32     `bun:"document_link_docstar_uid,type:int,autoincrement,identity,pk"` // Unique identifier for table document_link_docstar
	SourceAreaCd           int32     `bun:"source_area_cd,type:int"`                                      // Document link source area for this document link ID
	DocumentAreaCd         string    `bun:"document_area_cd,type:varchar(255)"`                           // form document linking area this record relates to
	DisplayAreaCd          int32     `bun:"display_area_cd,type:int"`                                     // Document link display area for which we want to store this document information
	Key1Cd                 string    `bun:"key1_cd,type:varchar(255)"`                                    // Key 1 cd for transaction records related to this type of transaction
	Key1ValueField         string    `bun:"key1_value_field,type:varchar(255)"`                           // form field to get the key1 value from
	Key2Cd                 *string   `bun:"key2_cd,type:varchar(255)"`                                    // Key 2 cd for transaction records related to this type of transaction
	Key2ValueField         *string   `bun:"key2_value_field,type:varchar(255)"`                           // form field to get the key2 value from
	Key3Cd                 *string   `bun:"key3_cd,type:varchar(255)"`                                    // Key 3 cd for transaction records related to this type of transaction
	Key3ValueField         *string   `bun:"key3_value_field,type:varchar(255)"`                           // form field to get the key3 value from
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
