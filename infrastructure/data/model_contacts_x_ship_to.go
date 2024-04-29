package data

import (
	"context"
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

type ContactsXShipToModel struct {
	bun bun.IDB
}

// GetByContactId retrieves all contacts_x_ship_to records by contact_id
func (m *ContactsXShipToModel) GetByContactId(ctx context.Context, companyId, contactId string) (
	[]*ContactsXShipTo, error) {
	var contactsXShipTo []*ContactsXShipTo
	err := m.bun.NewSelect().Model(&contactsXShipTo).Where(
		"contact_id= ? AND company_id = ?", contactId, companyId).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return contactsXShipTo, nil
}

// GetByShipToId retrieves all contacts_x_ship_to records by ship_to_id
func (m *ContactsXShipToModel) GetByShipToId(ctx context.Context, companyId string, shipToId float64) (
	[]*ContactsXShipTo, error) {
	var contactsXShipTo []*ContactsXShipTo
	err := m.bun.NewSelect().Model(&contactsXShipTo).Where(
		"ship_to_id = ? AND company_id = ?", shipToId, companyId).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return contactsXShipTo, nil
}
