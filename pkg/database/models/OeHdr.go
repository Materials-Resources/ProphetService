package models

import "database/sql"

func (OeHDR) TableName() string {
	return "oe_hdr"
}

type OeHDR struct {
	OrderNo              string         `gorm:"column:order_no;primaryKey;size:8;"`
	Ship2Name            sql.NullString `gorm:"column:ship2_name;size:50"`
	Ship2Add1            sql.NullString `gorm:"column:ship2_add1;size:50"`
	Ship2Add2            sql.NullString `gorm:"column:ship2_add2;size:50"`
	Ship2City            sql.NullString `gorm:"column:ship2_city;size:50"`
	Ship2State           sql.NullString `gorm:"column:ship2_state;size:50"`
	Ship2Zip             sql.NullString `gorm:"column:ship2_zip;size:10"`
	Ship2Country         sql.NullString `gorm:"column:ship2_country;size:50"`
	ContactId            sql.NullString `gorm:"column:contact_id;size:16"`
	DeliveryInstructions string         `gorm:"column:delivery_instructions;size:255"`

	OeLineItems []OeLine `gorm:"foreignKey:OrderNo;references:OrderNo"`
}
