package repository

import (
	"fmt"
	"strconv"

	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_19_1_3668"
	"github.com/materials-resources/s_prophet/internal/inventory/domain"
)

func FromDBReceipt(dbInventoryReceiptsHDR *prophet_19_1_3668.InventoryReceiptsHdr) (*domain.ValidatedInventoryReceipt,
	error,
) {
	inventoryReceipt := &domain.InventoryReceipt{
		ID: fmt.Sprintf("%.0f", dbInventoryReceiptsHDR.ReceiptNumber),
	}

	for _, item := range dbInventoryReceiptsHDR.InventoryReceiptsLineItems {
		ent := domain.InventoryReceiptItem{
			ProductId:        strconv.Itoa(int(item.InvMastUid)),
			ProductSn:        item.InvLoc.InvMast.ItemId,
			ProductName:      item.InvLoc.InvMast.ItemDesc,
			PrimaryBin:       item.InvLoc.PrimaryBin.String,
			ReceivedQuantity: item.QtyReceived,
			ReceivedUnit:     item.UnitOfMeasure,
			Orders:           nil,
		}

		validatedEnt, err := domain.NewValidatedInventoryReceiptItem(&ent)
		if err != nil {
			//TODO handle error
			fmt.Println(err)
		}
		inventoryReceipt.Items = append(inventoryReceipt.Items, *validatedEnt)
	}
	return domain.NewValidatedInventoryReceipt(inventoryReceipt)
}
