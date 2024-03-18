package inventory

import (
	"github.com/materials-resources/s_prophet/internal/inventory/domain"
	rpc "github.com/materials-resources/s_prophet/proto/inventory/v1alpha0"
)

func ToPbGetProductStockResponse(
	dPS []*domain.ProductStock,
) (*rpc.GetProductStockResponse, error) {
	pb := &rpc.GetProductStockResponse{}
	for _, ps := range dPS {
		pb.ProductStock = append(pb.ProductStock, &rpc.GetProductStockResponseProductStock{
			ProductUid:        ps.ProductUID,
			QuantityAvailable: ps.QuantityAvailable,
		},
		)
	}
	return pb, nil
}

func ToPBGetReceiptByIDResponse(ir *domain.ValidatedInventoryReceipt) (*rpc.GetReceiptByIDResponse, error) {
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
