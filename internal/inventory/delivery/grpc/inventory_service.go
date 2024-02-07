package inventory

import (
	"context"

	"github.com/materials-resources/s_prophet/internal/inventory"
	rpc "github.com/materials-resources/s_prophet/proto/inventory/v1alpha0"
)

type inventoryService struct {
	repo inventory.Repository
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
