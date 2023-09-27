package order

import (
	"context"

	"github.com/materials-resources/s_prophet/pkg/internal/core/common"
	"github.com/materials-resources/s_prophet/pkg/internal/core/domain"
	"github.com/materials-resources/s_prophet/pkg/internal/core/port/repository"
	rpc "github.com/materials-resources/s_prophet/proto/order/v1alpha0"
	"google.golang.org/grpc"
)

type service struct {
	orderRepository repository.OrderRepository
	//db           *gorm.DB
}

func NewOrderServer(server *grpc.Server, orderRepository repository.OrderRepository) {
	orderServer := &service{orderRepository: orderRepository}
	rpc.RegisterOrderServiceServer(
		server,
		orderServer,
	)
}

func (s *service) GetOrder(ctx context.Context, req *rpc.GetOrderRequest) (*rpc.GetOrderResponse, error) {
	order, err := s.orderRepository.Read(req.GetId())
	if err != nil {
		return nil, common.GetRpcError(err)
	}

	response := domainToOrderResponse(order)

	return response, nil
}

func domainToOrderResponse(order *domain.Order) *rpc.GetOrderResponse {
	getOrderResponse := &rpc.GetOrderResponse{
		Id: order.ID,
		ShippingDetails: &rpc.GetOrderResponse_ShippingDetails{
			Name:                 order.ShippingAddress.Name,
			LineOne:              order.ShippingAddress.LineOne,
			LineTwo:              order.ShippingAddress.LineTwo,
			City:                 order.ShippingAddress.City,
			State:                order.ShippingAddress.State,
			DeliveryInstructions: order.ShippingAddress.DeliveryInstructions,
		},
	}

	for _, orderItem := range order.OrderItems {
		getOrderResponse.OrderItems = append(
			getOrderResponse.OrderItems,
			&rpc.GetOrderResponse_OrderItem{
				Id:            orderItem.ID,
				Sn:            orderItem.SN,
				Name:          orderItem.Name,
				UnitPurchased: orderItem.UnitPurchased,
				CostPerUnit:   orderItem.CostPerUnit,
				UnitLabel:     orderItem.UnitLabel,
			},
		)
	}
	return getOrderResponse
}
