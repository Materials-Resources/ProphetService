package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DocumentTypes struct {
	bun.BaseModel           `bun:"table:document_types"`
	DocumentTypeId          string    `bun:"document_type_id,type:varchar(2),pk"`                       // Unique identifier of document_type
	DocumentTypeDescription string    `bun:"document_type_description,type:varchar(30)"`                // Description of document type.
	DocumentObject          string    `bun:"document_object,type:varchar(50)"`                          // What is the report object in the system?
	DocumentId              string    `bun:"document_id,type:varchar(3)"`                               // What is the unique identifer for this document?
	DeleteFlag              string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
