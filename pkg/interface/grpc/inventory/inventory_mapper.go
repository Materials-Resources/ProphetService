package inventory

import (
	"github.com/materials-resources/s_prophet/pkg/domain/entities"
	rpc "github.com/materials-resources/s_prophet/proto/inventory/v1alpha0"
)

func ToPBGetReceiptByIDResponse(ir *entities.ValidatedInventoryReceipt) (*rpc.GetReceiptByIDResponse, error) {
	pb := &rpc.GetReceiptByIDResponse{Id: ir.ID}

	for _, item := range ir.Items {
		pb.Items = append(pb.Items, &rpc.GetReceiptByIDResponse_Item{
			ProductId:        item.ProductId,
			ProductSn:        item.ProductSn,
			ProductName:      item.ProductName,
			ReceivedQuantity: item.ReceivedQuantity,
			ReceivedUnit:     item.ReceivedUnit,
		},
		)
	}
	return pb, nil
}
