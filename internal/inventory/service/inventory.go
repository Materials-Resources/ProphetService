package service

import (
	"context"
	"fmt"
	"github.com/materials-resources/s_prophet/infrastructure/data"
	"github.com/materials-resources/s_prophet/internal/inventory/domain"
)

type InventoryService struct {
	models data.Models
}

func (s *InventoryService) GetProductStock(ctx context.Context, int32s []int32) ([]domain.ProductStock, error) {
	invLocs, err := s.models.InvLoc.GetByInvMastUids(ctx, int32s)
	if err != nil {
		return nil, err

	}
	productStocks := make([]domain.ProductStock, len(invLocs))

	for i, invLoc := range invLocs {
		productStocks[i] = domain.ProductStock{
			ProductUid:        invLoc.InvMastUid,
			QuantityAvailable: invLoc.QtyOnHand.Float64,
		}

	}
	// TODO implement me
	panic("implement me")
}

func NewInventoryService(models data.Models) *InventoryService {
	return &InventoryService{models: models}
}

func (s *InventoryService) GetInventoryReceiptById(ctx context.Context, id float64) (domain.InventoryReceipt, error) {

	inventoryReceiptsHdr, err := s.models.InventoryReceiptsHdr.GetByReceiptNumberWithLines(ctx, id)
	if err != nil {
		return domain.InventoryReceipt{}, err
	}

	invTrans, err := s.models.InvTran.GetBySubDocumentNo(ctx, id)
	if err != nil {
		return domain.InventoryReceipt{}, err
	}

	invTransMap := make(map[int32][]*data.InvTran)
	for _, invTran := range invTrans {
		invTransMap[invTran.InvMastUid] = append(invTransMap[invTran.InvMastUid], invTran)
	}

	invMastUids := make([]int32, len(inventoryReceiptsHdr.InventoryReceiptsLines))
	for i, line := range inventoryReceiptsHdr.InventoryReceiptsLines {
		invMastUids[i] = line.InvMastUid
	}

	invLocs, err := s.models.InvLoc.GetByInvMastUids(ctx, invMastUids)
	if err != nil {
		return domain.InventoryReceipt{}, err
	}
	invLocsMap := make(map[int32]*data.InvLoc)
	for _, invLoc := range invLocs {
		invLocsMap[invLoc.InvMastUid] = invLoc
	}

	items := make([]domain.InventoryReceiptItem, len(inventoryReceiptsHdr.InventoryReceiptsLines))
	for i, line := range inventoryReceiptsHdr.InventoryReceiptsLines {

		allocatedOrders := make([]string, len(invTransMap[line.InvMastUid]))

		for j, tran := range invTransMap[line.InvMastUid] {
			allocatedOrders[j] = fmt.Sprintf("%.0f", tran.DocumentNo)
		}

		items[i] = domain.InventoryReceiptItem{
			ProductUid:       line.InvMastUid,
			ReceivedQuantity: line.QtyReceived,
			ReceivedUnit:     line.UnitOfMeasure,
			AllocatedOrders:  allocatedOrders,
			PrimaryBin:       invLocsMap[line.InvMastUid].PrimaryBin.String,
		}
	}
	return domain.InventoryReceipt{
		ID:    inventoryReceiptsHdr.ReceiptNumber,
		Items: items,
	}, nil

}
