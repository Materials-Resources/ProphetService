package receiving

import (
	"context"

	"github.com/materials-resources/s_prophet/pkg/internal/core/port/repository"
	rpc "github.com/materials-resources/s_prophet/proto/receiving/v1alpha0"
	"google.golang.org/grpc"
)

type service struct {
	receivingRepository repository.ReceivingRepository
}

func (s service) GetReceipt(ctx context.Context, request *rpc.GetReceiptRequest) (*rpc.GetReceiptResponse, error) {
	err := s.receivingRepository.ReadReceivingReceipt()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func NewReceivingServer(server *grpc.Server, receivingRepository repository.ReceivingRepository) {
	receivingServer := &service{receivingRepository: receivingRepository}
	rpc.RegisterReceivingServiceServer(
		server,
		receivingServer,
	)
}
