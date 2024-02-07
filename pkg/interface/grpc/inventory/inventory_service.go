package inventory

import (
	"context"

	"github.com/materials-resources/s_prophet/pkg/domain/repositories"
	rpc "github.com/materials-resources/s_prophet/proto/inventory/v1alpha0"
	"google.golang.org/grpc"
)

func NewInventoryService(server *grpc.Server, repo repositories.InventoryRepository) {
	svc := &inventoryService{repo: repo}
	rpc.RegisterInventoryServiceServer(server, svc)
}

type inventoryService struct {
	repo repositories.InventoryRepository
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
