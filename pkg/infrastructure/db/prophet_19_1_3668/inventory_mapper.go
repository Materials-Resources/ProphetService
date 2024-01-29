package prophet_19_1_3668

import (
	"fmt"
	"strconv"

	"github.com/materials-resources/s_prophet/pkg/domain/entities"
	"github.com/materials-resources/s_prophet/pkg/infrastructure/db/prophet_19_1_3668/model"
)

func FromDBReceipt(dbInventoryReceiptsHDR *model.InventoryReceiptsHdr) (*entities.ValidatedInventoryReceipt, error) {
	inventoryReceipt := &entities.InventoryReceipt{
		ID: fmt.Sprintf("%.0f", dbInventoryReceiptsHDR.ReceiptNumber),
	}

	for _, item := range dbInventoryReceiptsHDR.InventoryReceiptsLineItems {
		ent := entities.InventoryReceiptItem{
			ProductId:        strconv.Itoa(int(item.InvMastUid)),
			ProductSn:        item.InvLoc.InvMast.ItemId,
			ProductName:      item.InvLoc.InvMast.ItemDesc,
			PrimaryBin:       item.InvLoc.PrimaryBin.String,
			ReceivedQuantity: item.QtyReceived,
			ReceivedUnit:     item.UnitOfMeasure,
			Orders:           nil,
		}

		validatedEnt, err := entities.NewValidatedInventoryReceiptItem(&ent)
		if err != nil {
			//TODO handle error
			fmt.Println(err)
		}
		inventoryReceipt.Items = append(inventoryReceipt.Items, *validatedEnt)
	}
	return entities.NewValidatedInventoryReceipt(inventoryReceipt)
}
