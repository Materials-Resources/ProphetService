package grpc_service

import (
	"context"
	"fmt"
	"github.com/materials-resources/ProphetService/pkg/database"
	rpc "github.com/materials-resources/ProphetService/proto/warehouse/v1alpha0"
	"strconv"
)

type WarehouseServer struct {
	rpc.UnimplementedWarehouseServiceServer
	WarehouseHandler *database.WarehouseHandler
}

func (s *WarehouseServer) GetPickTicket(ctx context.Context, req *rpc.GetPickTicketRequest) (*rpc.GetPickTicketResponse, error) {
	hRes, err := s.WarehouseHandler.SelectPickTicket(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &rpc.GetPickTicketResponse{
		OrderId: hRes.OrderNo,
	}, nil
}

func (s *WarehouseServer) GetReceivingReport(ctx context.Context, req *rpc.GetReceivingReportRequest) (*rpc.GetReceivingReportResponse, error) {
	hRes, err := s.WarehouseHandler.SelectReceivingReport(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	rpcRes := new(rpc.GetReceivingReportResponse)

	for _, product := range hRes.Items {
		rpcRes.Products = append(rpcRes.Products, &rpc.GetReceivingReportResponse_Product{
			Id:            product.InvMastUid,
			Qty:           product.Qty,
			UnitOfMeasure: product.Uom,
		})
	}
	return rpcRes, nil
}

func (s *WarehouseServer) GetAllocatedOrdersForReceiptItem(ctx context.Context, req *rpc.GetAllocatedOrdersForReceiptItemRequest) (*rpc.GetAllocatedOrdersForReceiptItemResponse, error) {
	fmt.Println(req.GetPrimaryBins())
	hRes, err := s.WarehouseHandler.SelectReceivingReportAllocatedOrders(ctx, req.GetReceiptId(), req.GetProductId(), req.GetPrimaryBins())
	if err != nil {
		return nil, err
	}

	rpcRes := new(rpc.GetAllocatedOrdersForReceiptItemResponse)

	for _, order := range hRes.OrderTransactions {
		rpcRes.AllocatedOrders = append(rpcRes.AllocatedOrders, &rpc.GetAllocatedOrdersForReceiptItemResponse_AllocatedOrder{
			Id:  strconv.Itoa(int(order.DocumentNo)),
			Qty: order.QtyAllocated,
		})
	}

	return rpcRes, nil
}
