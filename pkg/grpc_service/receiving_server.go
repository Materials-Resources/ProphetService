package grpc_service

import (
	"context"
	"github.com/materials-resources/s_prophet/pkg/database/db_receiving"
	rpc "github.com/materials-resources/s_prophet/proto/receiving/v1alpha0"
	"log"
)

type ReceivingServer struct {
	rpc.UnimplementedReceivingServiceServer
	DBHandler *db_receiving.Handler
}

func (s *ReceivingServer) GetReceipt(ctx context.Context, req *rpc.GetReceiptRequest) (*rpc.GetReceiptResponse, error) {
	res, err := s.DBHandler.GetReceipt(req)
	if err != nil {
		log.Fatal("failed to query")
	}

	return res, nil
}
