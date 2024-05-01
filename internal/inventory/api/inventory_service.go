package api

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/inventory/service"

	rpc "github.com/materials-resources/s-prophet/proto/inventory/v1"
)

func NewInventoryApi(service service.Service) InventoryApi {
	return InventoryApi{service: service}
}

type InventoryApi struct {
	service service.Service
}

func (s InventoryApi) GetProductStock(
	ctx context.Context,
	request *rpc.GetProductStockRequest) (*rpc.GetProductStockResponse, error) {

	productStock, err := s.service.GetProductStock(ctx, request.GetProductUid())
	if err != nil {
		return nil, err
	}

	res := &rpc.GetProductStockResponse{}

	res.ProductStock = make([]*rpc.GetProductStockResponse_ProductStock, len(productStock))
	for i, stock := range productStock {
		res.ProductStock[i] = &rpc.GetProductStockResponse_ProductStock{
			ProductUid:        stock.ProductUid,
			QuantityAvailable: stock.QuantityAvailable,
		}
	}

	return res, nil

}

func (s InventoryApi) GetReceiptByID(
	ctx context.Context, request *rpc.GetReceiptByIDRequest,
) (*rpc.GetReceiptByIDResponse, error) {

	inventoryReceipt, err := s.service.GetInventoryReceiptById(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	res := &rpc.GetReceiptByIDResponse{
		Id: inventoryReceipt.ID,
	}

	res.Items = make([]*rpc.GetReceiptByIDResponse_Item, len(inventoryReceipt.Items))
	for i, line := range inventoryReceipt.Items {
		res.Items[i] = &rpc.GetReceiptByIDResponse_Item{
			ProductUid:        line.ProductUid,
			ReceivedQuantity:  line.ReceivedQuantity,
			ReceivedUnit:      line.ReceivedUnit,
			AllocatedOrders:   line.AllocatedOrders,
			ProductPrimaryBin: line.PrimaryBin,
		}
	}

	return res, nil
}
