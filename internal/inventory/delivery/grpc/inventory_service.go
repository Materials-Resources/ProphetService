package inventory

import (
	"context"
	"fmt"

	"github.com/materials-resources/s_prophet/internal/inventory"
	rpc "github.com/materials-resources/s_prophet/proto/inventory/v1alpha0"
)

type inventoryService struct {
	repo inventory.Repository
}

func (s inventoryService) GetProductStock(
	ctx context.Context,
	request *rpc.GetProductStockRequest) (*rpc.GetProductStockResponse, error) {

	dProductStock, err := s.repo.SelectProductStockById(ctx, request.GetProductUid())
	if err != nil {
		return &rpc.GetProductStockResponse{}, fmt.Errorf("error getting product stock: %v", err)
	}
	return ToPbGetProductStockResponse(dProductStock)

}

func (s inventoryService) GetReceiptByID(
	ctx context.Context, request *rpc.GetReceiptByIDRequest,
) (*rpc.GetReceiptByIDResponse, error) {
	//TODO implement me

	inventoryReceipt, err := s.repo.FindReceiptById(request.GetId())
	if err != nil {
		return nil, err
	}
	return ToPBGetReceiptByIDResponse(inventoryReceipt)
}
