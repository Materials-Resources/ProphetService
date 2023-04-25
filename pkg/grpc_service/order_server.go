package grpc_service

import (
	"context"
	"materials-resources.com/msv/prophet/pkg/database"
	rpc "materials-resources.com/msv/prophet/proto/order/v1alpha0"
)

type OrderServer struct {
	rpc.UnimplementedOrderServiceServer
	OrderHandler *database.OrderHandler
}

func (s *OrderServer) GetOrder(ctx context.Context, req *rpc.GetOrderRequest) (*rpc.GetOrderResponse, error) {
	hRes, err := s.OrderHandler.SelectOrder(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &rpc.GetOrderResponse{
		DeliveryInstructions: hRes.DeliveryInstructions,
		ShippingAddress: &rpc.GetOrderResponse_ShippingAddress{
			Name:       hRes.ShipToName,
			LineOne:    hRes.ShipToAdd1,
			LineTwo:    hRes.ShipToAdd1,
			City:       hRes.ShipToCity,
			State:      hRes.ShipToState,
			PostalCode: hRes.ShipToZip,
			Country:    hRes.ShipToCountry,
		},
		PoNo:      hRes.PoNo,
		ContactId: hRes.ContactId,
	}, nil
}

func (s *OrderServer) GetOrderItem(ctx context.Context, req *rpc.GetOrderItemRequest) (*rpc.GetOrderItemResponse, error) {
	hRes, err := s.OrderHandler.SelectOrderItem(ctx, req.GetOrderId(), req.GetProductId())
	if err != nil {
		return nil, err
	}
	rpcRes := new(rpc.GetOrderItemResponse)

	for _, item := range *hRes {
		rpcRes.Item = append(rpcRes.Item, &rpc.GetOrderItemResponse_Item{
			ProductId:         item.InvMastUid,
			CustomerProductId: item.CustomerPartId,
		})
	}

	return rpcRes, nil

}
