package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ContactsXShipTo struct {
	bun.BaseModel      `bun:"table:contacts_x_ship_to"`
	ContactsXShipToUid int32     `bun:"contacts_x_ship_to_uid,type:int,pk"`
	ContactId          string    `bun:"contact_id,type:varchar(16)"`
	CompanyId          string    `bun:"company_id,type:varchar(8)"`
	ShipToId           float64   `bun:"ship_to_id,type:decimal(19,0)"`
	RowStatusFlag      int32     `bun:"row_status_flag,type:int"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	PedigreeContact    string    `bun:"pedigree_contact,type:char,default:('N')"`
}
