package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DocumentTypes struct {
	bun.BaseModel           `bun:"table:document_types"`
	DocumentTypeId          string    `bun:"document_type_id,type:varchar(2),pk"`
	DocumentTypeDescription string    `bun:"document_type_description,type:varchar(30)"`
	DocumentObject          string    `bun:"document_object,type:varchar(50)"`
	DocumentId              string    `bun:"document_id,type:varchar(3)"`
	DeleteFlag              string    `bun:"delete_flag,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
