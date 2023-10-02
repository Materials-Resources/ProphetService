package controller

import (
	"context"

	"github.com/materials-resources/s_prophet/core/domain"
	"github.com/materials-resources/s_prophet/core/port/repository"
	rpc "github.com/materials-resources/s_prophet/proto/receiving/v1alpha0"
	"google.golang.org/grpc"
)

type receivingService struct {
	receivingRepository repository.ReceivingRepository
}

func (s receivingService) GetReceipt(
	ctx context.Context, request *rpc.GetReceiptRequest,
) (*rpc.GetReceiptResponse, error) {
	receivingReceipt, err := s.receivingRepository.ReadReceipt(
		ctx,
		request.GetId(),
	)
	if err != nil {
		return nil, err
	}
	return receiptResponseFromDomain(*receivingReceipt), nil
}

func receiptResponseFromDomain(receipt domain.ReceivingReceipt) *rpc.GetReceiptResponse {
	allocatedOrders := func(orders []*domain.ReceivingReceiptAllocatedOrder) []*rpc.GetReceiptResponse_Product_Order {
		var res []*rpc.GetReceiptResponse_Product_Order
		for _, order := range orders {
			res = append(
				res,
				&rpc.GetReceiptResponse_Product_Order{
					Id:             order.Id,
					Name:           order.Name,
					UnitsAllocated: order.UnitsAllocated,
					UnitLabel:      order.UnitLabel,
				},
			)
		}
		return res
	}

	rpcRes := new(rpc.GetReceiptResponse)

	rpcRes.Id = receipt.Id

	for _, product := range receipt.Products {
		rpcRes.Products = append(
			rpcRes.Products,
			&rpc.GetReceiptResponse_Product{
				Id:              product.Id,
				Sn:              product.Sn,
				Name:            product.Name,
				UnitsReceived:   product.UnitsReceived,
				UnitLabel:       product.UnitLabel,
				PrimaryBin:      product.PrimaryBin,
				AllocatedOrders: allocatedOrders(product.AllocatedOrders),
			},
		)
	}

	return rpcRes

}

func NewReceivingService(server *grpc.Server, receivingRepository repository.ReceivingRepository) {
	receivingServer := &receivingService{receivingRepository: receivingRepository}
	rpc.RegisterReceivingServiceServer(
		server,
		receivingServer,
	)
}
