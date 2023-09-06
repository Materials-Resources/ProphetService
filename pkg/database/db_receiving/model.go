package db_receiving

import "database/sql"

type resultGetReceipt struct {
	Id int32 `gorm:"column:receipt_number"`
}

type resultGetReceiptItem struct {
	ID               int32          `gorm:"column:inv_mast_uid"`
	SN               string         `gorm:"column:item_id"`
	Name             string         `gorm:"column:item_desc"`
	PrimaryBin       sql.NullString `gorm:"primary_bin"`
	QuantityReceived float32        `gorm:"column:qty_received"`
	QuantityUOM      string         `gorm:"column:unit_of_measure"`
}

type resultGetReceiptItemAllocatedOrder struct {
	OrderId   int32   `gorm:"column:document_no"`
	OrderName string  `gorm:"column:ship2_name"`
	Qty       float32 `gorm:"column:qty_allocated"`
	ProductId int32   `gorm:"column:inv_mast_uid"`
}
