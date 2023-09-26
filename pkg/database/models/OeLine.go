package models

import "database/sql"

func (OeLine) TableName() string {
	return "oe_line"
}

type OeLine struct {
	OrderNo       string          `gorm:"column:order_no;type:varchar(8);primaryKey"`
	LineNo        float64         `gorm:"column:line_no;type:decimal(19);primaryKey"`
	UnitPrice     sql.NullFloat64 `gorm:"column:unit_price;type:decimal(19, 9)'"`
	QtyOrdered    sql.NullFloat64 `gorm:"column:qty_ordered;type:decimal(19, 9)'"`
	ExtendedPrice sql.NullFloat64 `gorm:"column:extended_price;type:decimal(19, 4)'"`
	UnitOfMeasure sql.NullString  `gorm:"column:unit_of_measure;type:varchar(8);"`
}
