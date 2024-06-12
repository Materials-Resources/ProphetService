package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ContactsXSupplier struct {
	bun.BaseModel    `bun:"table:contacts_x_supplier"`
	ContactId        string    `bun:"contact_id,type:varchar(16),pk"`
	SupplierId       float64   `bun:"supplier_id,type:decimal(19,0),pk"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30)"`
	PedigreeContact  string    `bun:"pedigree_contact,type:char,default:('N')"`
}
