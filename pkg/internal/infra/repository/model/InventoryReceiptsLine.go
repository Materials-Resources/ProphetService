package model

import (
	"github.com/uptrace/bun"
)

type InventoryReceiptsLine struct {
	bun.BaseModel `bun:"table:inventory_receipts_line"`
	ReceiptNumber float64 `bun:"receipt_number,type:decimal(19),pk"`
	LineNumber    int32   `bun:"line_number,type:int,pk"`
	InvMastUid    int32   `bun:"inv_mast_uid,type:int"`
	UnitQuantity  float64 `bun:"unit_quantity,type:decimal(19, 9)"`
	UnitOfMeasure string  `bun:"unit_of_measure,type:varchar(8)"`

	InvMast InvMast `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`

	//	Ira []InvTran `bun:"rel:has-many,
	//join:receipt_number=sub_document_no,join:type=trans_type,join:inv_mast_uid=inv_mast_uid,polymorphic:IRA"`
}
