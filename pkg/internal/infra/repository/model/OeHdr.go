package model

import "database/sql"

func (OeHdr) TableName() string {
	return "oe_hdr"
}

type OeHdr struct {
	OrderNo              string         `gorm:"column:order_no;primaryKey;size:8;"`
	Ship2Name            sql.NullString `gorm:"column:ship2_name;size:50"`
	Ship2Add1            sql.NullString `gorm:"column:ship2_add1;size:50"`
	Ship2Add2            sql.NullString `gorm:"column:ship2_add2;size:50"`
	Ship2City            sql.NullString `gorm:"column:ship2_city;size:50"`
	Ship2State           sql.NullString `gorm:"column:ship2_state;size:50"`
	Ship2Zip             sql.NullString `gorm:"column:ship2_zip;size:10"`
	Ship2Country         sql.NullString `gorm:"column:ship2_country;size:50"`
	ContactId            sql.NullString `gorm:"column:contact_id;size:16"`
	DeliveryInstructions sql.NullString `gorm:"column:delivery_instructions;type:varchar(255)"`
	PoNo                 sql.NullString `gorm:"column:po_no;type:varchar(50)"`

	OeLineItems []OeLine `gorm:"foreignKey:OrderNo;references:OrderNo"`
	Contact     Contact  `gorm:"foreignKey:ContactId;references:Id"`
}
