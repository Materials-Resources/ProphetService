package db_receiving

import (
	pb "github.com/materials-resources/s_prophet/proto/receiving/v1alpha0"
	"gorm.io/gorm"
	"strconv"
)

type Handler struct {
	db *gorm.DB
}

func NewReceivingHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetReceipt(request *pb.GetReceiptRequest) (*pb.GetReceiptResponse, error) {
	pbResponse := new(pb.GetReceiptResponse)

	var receipt resultGetReceipt
	var receiptItems []*resultGetReceiptItem
	var receiptItemsId []*int32
	var allocatedOrders []*resultGetReceiptItemAllocatedOrder

	h.db.Raw("select receipt_number FROM inventory_receipts_hdr where receipt_number = ?", request.GetId()).Scan(&receipt)

	h.db.Raw(`select inventory_receipts_line.inv_mast_uid, inventory_receipts_line.qty_received, inv_mast.item_id, inv_mast.item_desc, inv_loc.primary_bin
					FROM inventory_receipts_line
					INNER JOIN inv_mast on inv_mast.inv_mast_uid = inventory_receipts_line.inv_mast_uid
					INNER JOIN inv_loc on inv_mast.inv_mast_uid = inv_loc.inv_mast_uid
					where inventory_receipts_line.receipt_number = ?
					`, request.GetId()).Scan(&receiptItems)

	for _, receiptItem := range receiptItems {
		receiptItemsId = append(receiptItemsId, &receiptItem.ID)
	}

	h.db.Raw(`SELECT inv_tran.qty_allocated, inv_tran.document_no, inv_tran.inv_mast_uid, oe_hdr.ship2_name
					FROM inv_tran
					INNER JOIN oe_hdr on (cast(inv_tran.document_no AS varchar) = oe_hdr.order_no) 
					WHERE inv_tran.sub_document_no = ? AND inv_tran.trans_type = 'IRA' AND inv_tran.inv_mast_uid in (?)`, request.GetId(), receiptItemsId).Scan(&allocatedOrders)

	pbResponse.Id = receipt.Id

	for _, receiptItem := range receiptItems {

		item := &pb.GetReceiptResponse_Item{
			Id:          receiptItem.ID,
			Sn:          receiptItem.SN,
			Name:        receiptItem.Name,
			QtyReceived: receiptItem.QuantityReceived,
			PrimaryBin:  receiptItem.PrimaryBin.String,
		}

		for _, allocatedOrder := range allocatedOrders {
			if allocatedOrder.ProductId == receiptItem.ID {
				item.AllocatedOrders = append(item.AllocatedOrders, &pb.GetReceiptResponse_Item_Order{
					Id:   strconv.Itoa(int(allocatedOrder.OrderId)),
					Qty:  allocatedOrder.Qty,
					Name: allocatedOrder.OrderName,
				})
			}

		}

		pbResponse.Items = append(pbResponse.Items, item)

	}

	return pbResponse, nil
}
