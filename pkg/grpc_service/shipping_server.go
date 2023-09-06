package grpc_service

import (
	"context"
	"github.com/materials-resources/s_prophet/pkg/database/db_shipping"
	rpc "github.com/materials-resources/s_prophet/proto/shipping/v1alpha0"
	"log"
)

type ShippingServer struct {
	rpc.UnimplementedShippingServiceServer
	DBHandler *db_shipping.Handler
}

func (s *ShippingServer) GetPickTicket(ctx context.Context, req *rpc.GetPickTicketRequest) (*rpc.GetPickTicketResponse, error) {
	res, err := s.DBHandler.GetPickTicket(req)
	if err != nil {
		log.Fatal("failed to query")
	}
	return res, nil
}
