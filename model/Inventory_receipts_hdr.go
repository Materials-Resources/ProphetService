package model

import (
	"github.com/uptrace/bun"
)

type InventoryReceiptsHdr struct {
	bun.BaseModel `bun:"table:inventory_receipts_hdr"`
	ReceiptNumber float64 `bun:"receipt_number,type:decimal(19),pk"`

	InventoryReceiptsLineItems []InventoryReceiptsLine `bun:"rel:has-many,join:receipt_number=receipt_number"`
}
