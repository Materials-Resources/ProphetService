package model

import (
	"database/sql"
)

func (InvMast) TableName() string {
	return "inv_mast"
}

type InvMast struct {
	InvMastUid   int32           `gorm:"column:inv_mast_uid;type:int;primaryKey"`
	ItemId       string          `gorm:"column:item_id;type:varchar(40);unique"`
	ItemDesc     string          `gorm:"column:item_desc;type:varchar(40)"`
	ExtendedDesc sql.NullString  `gorm:"column:extended_desc;type:varchar(255)"`
	Price1       sql.NullFloat64 `gorm:"column:price1;type:decimal(19, 9)"`
}
