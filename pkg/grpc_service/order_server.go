package grpc_service

import (
	"context"
	"github.com/materials-resources/s_prophet/pkg/database"
	"github.com/materials-resources/s_prophet/pkg/database/models"
	rpc "github.com/materials-resources/s_prophet/proto/order/v1alpha0"
	"gorm.io/gorm"
)

type OrderServer struct {
	rpc.UnimplementedOrderServiceServer
	OrderHandler *database.OrderHandler
	db           *gorm.DB
}

func NewOrderServer(db *gorm.DB, orderHandler *database.OrderHandler) *OrderServer {
	return &OrderServer{
		UnimplementedOrderServiceServer: rpc.UnimplementedOrderServiceServer{},
		OrderHandler:                    orderHandler,
		db:                              db,
	}
}

func (s *OrderServer) GetOrder(ctx context.Context, req *rpc.GetOrderRequest) (*rpc.GetOrderResponse, error) {
	var order models.OeHDR
	err := s.db.Model(&models.OeHDR{}).Preload("OeLineItems").First(&order, req.GetId()).Error
	if err != nil {
		return nil, err
	}

	return &rpc.GetOrderResponse{
		DeliveryInstructions: order.DeliveryInstructions,
		ShippingAddress: &rpc.GetOrderResponse_ShippingAddress{
			Name:       order.Ship2Name.String,
			LineOne:    order.Ship2Add1.String,
			LineTwo:    order.Ship2Add1.String,
			City:       order.Ship2City.String,
			State:      order.Ship2State.String,
			PostalCode: order.Ship2Zip.String,
			Country:    order.Ship2Country.String,
		},
		ContactId: order.ContactId.String,
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
			QuantityOrdered:   item.QtyOrdered,
			UnitPrice:         item.UnitPrice,
			UnitMeasurement:   item.UnitOfMeasurement,
			TotalPrice:        item.ExtendedPrice,
			CustomerProductId: item.CustomerPartId,
		})
	}

	return rpcRes, nil

}
