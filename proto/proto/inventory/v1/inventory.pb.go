// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: proto/inventory/v1/inventory.proto

package inventory_v1alpha0

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialId       int32   `protobuf:"varint,1,opt,name=serial_id,json=serialId,proto3" json:"serial_id,omitempty"`
	SupplierId     int32   `protobuf:"varint,2,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	DivisionId     int32   `protobuf:"varint,3,opt,name=division_id,json=divisionId,proto3" json:"division_id,omitempty"`
	LeadTime       int32   `protobuf:"varint,4,opt,name=lead_time,json=leadTime,proto3" json:"lead_time,omitempty"`
	ListPrice      float32 `protobuf:"fixed32,5,opt,name=list_price,json=listPrice,proto3" json:"list_price,omitempty"`
	Cost           float32 `protobuf:"fixed32,6,opt,name=cost,proto3" json:"cost,omitempty"`
	SupplierPartNo string  `protobuf:"bytes,7,opt,name=supplier_part_no,json=supplierPartNo,proto3" json:"supplier_part_no,omitempty"`
}

func (x *AddSupplierRequest) Reset() {
	*x = AddSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inventory_v1_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSupplierRequest) ProtoMessage() {}

func (x *AddSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_v1_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSupplierRequest.ProtoReflect.Descriptor instead.
func (*AddSupplierRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_v1_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *AddSupplierRequest) GetSerialId() int32 {
	if x != nil {
		return x.SerialId
	}
	return 0
}

func (x *AddSupplierRequest) GetSupplierId() int32 {
	if x != nil {
		return x.SupplierId
	}
	return 0
}

func (x *AddSupplierRequest) GetDivisionId() int32 {
	if x != nil {
		return x.DivisionId
	}
	return 0
}

func (x *AddSupplierRequest) GetLeadTime() int32 {
	if x != nil {
		return x.LeadTime
	}
	return 0
}

func (x *AddSupplierRequest) GetListPrice() float32 {
	if x != nil {
		return x.ListPrice
	}
	return 0
}

func (x *AddSupplierRequest) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *AddSupplierRequest) GetSupplierPartNo() string {
	if x != nil {
		return x.SupplierPartNo
	}
	return ""
}

type GetProductStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductUid []int32 `protobuf:"varint,1,rep,packed,name=product_uid,json=productUid,proto3" json:"product_uid,omitempty"`
}

func (x *GetProductStockRequest) Reset() {
	*x = GetProductStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inventory_v1_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductStockRequest) ProtoMessage() {}

func (x *GetProductStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_v1_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductStockRequest.ProtoReflect.Descriptor instead.
func (*GetProductStockRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_v1_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductStockRequest) GetProductUid() []int32 {
	if x != nil {
		return x.ProductUid
	}
	return nil
}

type GetProductStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductStock []*GetProductStockResponse_ProductStock `protobuf:"bytes,1,rep,name=product_stock,json=productStock,proto3" json:"product_stock,omitempty"`
}

func (x *GetProductStockResponse) Reset() {
	*x = GetProductStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inventory_v1_inventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductStockResponse) ProtoMessage() {}

func (x *GetProductStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_v1_inventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductStockResponse.ProtoReflect.Descriptor instead.
func (*GetProductStockResponse) Descriptor() ([]byte, []int) {
	return file_proto_inventory_v1_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductStockResponse) GetProductStock() []*GetProductStockResponse_ProductStock {
	if x != nil {
		return x.ProductStock
	}
	return nil
}

type AddSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddSupplierResponse) Reset() {
	*x = AddSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inventory_v1_inventory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSupplierResponse) ProtoMessage() {}

func (x *AddSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_v1_inventory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSupplierResponse.ProtoReflect.Descriptor instead.
func (*AddSupplierResponse) Descriptor() ([]byte, []int) {
	return file_proto_inventory_v1_inventory_proto_rawDescGZIP(), []int{3}
}

type DeleteSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSupplierRequest) Reset() {
	*x = DeleteSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inventory_v1_inventory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSupplierRequest) ProtoMessage() {}

func (x *DeleteSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_v1_inventory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSupplierRequest.ProtoReflect.Descriptor instead.
func (*DeleteSupplierRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_v1_inventory_proto_rawDescGZIP(), []int{4}
}

type DeleteSupplierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSupplierResponse) Reset() {
	*x = DeleteSupplierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inventory_v1_inventory_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSupplierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSupplierResponse) ProtoMessage() {}

func (x *DeleteSupplierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_v1_inventory_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSupplierResponse.ProtoReflect.Descriptor instead.
func (*DeleteSupplierResponse) Descriptor() ([]byte, []int) {
	return file_proto_inventory_v1_inventory_proto_rawDescGZIP(), []int{5}
}

type GetReceiptByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id float64 `protobuf:"fixed64,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetReceiptByIDRequest) Reset() {
	*x = GetReceiptByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inventory_v1_inventory_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptByIDRequest) ProtoMessage() {}

func (x *GetReceiptByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_v1_inventory_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptByIDRequest.ProtoReflect.Descriptor instead.
func (*GetReceiptByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_v1_inventory_proto_rawDescGZIP(), []int{6}
}

func (x *GetReceiptByIDRequest) GetId() float64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetReceiptByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique ID of receipt
	Id float64 `protobuf:"fixed64,1,opt,name=id,proto3" json:"id,omitempty"`
	// Items that were received for this receipt
	Items []*GetReceiptByIDResponse_Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetReceiptByIDResponse) Reset() {
	*x = GetReceiptByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inventory_v1_inventory_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptByIDResponse) ProtoMessage() {}

func (x *GetReceiptByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_v1_inventory_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptByIDResponse.ProtoReflect.Descriptor instead.
func (*GetReceiptByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_inventory_v1_inventory_proto_rawDescGZIP(), []int{7}
}

func (x *GetReceiptByIDResponse) GetId() float64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetReceiptByIDResponse) GetItems() []*GetReceiptByIDResponse_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetProductStockResponse_ProductStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductUid        int32   `protobuf:"varint,1,opt,name=product_uid,json=productUid,proto3" json:"product_uid,omitempty"`
	QuantityAvailable float64 `protobuf:"fixed64,2,opt,name=quantity_available,json=quantityAvailable,proto3" json:"quantity_available,omitempty"`
}

func (x *GetProductStockResponse_ProductStock) Reset() {
	*x = GetProductStockResponse_ProductStock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inventory_v1_inventory_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductStockResponse_ProductStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductStockResponse_ProductStock) ProtoMessage() {}

func (x *GetProductStockResponse_ProductStock) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_v1_inventory_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductStockResponse_ProductStock.ProtoReflect.Descriptor instead.
func (*GetProductStockResponse_ProductStock) Descriptor() ([]byte, []int) {
	return file_proto_inventory_v1_inventory_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetProductStockResponse_ProductStock) GetProductUid() int32 {
	if x != nil {
		return x.ProductUid
	}
	return 0
}

func (x *GetProductStockResponse_ProductStock) GetQuantityAvailable() float64 {
	if x != nil {
		return x.QuantityAvailable
	}
	return 0
}

type GetReceiptByIDResponse_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique ID of product
	ProductUid int32 `protobuf:"varint,1,opt,name=product_uid,json=productUid,proto3" json:"product_uid,omitempty"`
	// Primary bin where product is stored
	ProductPrimaryBin string `protobuf:"bytes,2,opt,name=product_primary_bin,json=productPrimaryBin,proto3" json:"product_primary_bin,omitempty"`
	// Quantity of product received
	ReceivedQuantity float64 `protobuf:"fixed64,4,opt,name=received_quantity,json=receivedQuantity,proto3" json:"received_quantity,omitempty"`
	// Unit associated with received_quantity
	ReceivedUnit string `protobuf:"bytes,5,opt,name=received_unit,json=receivedUnit,proto3" json:"received_unit,omitempty"`
	// Orders that were allocated to
	AllocatedOrders []string `protobuf:"bytes,6,rep,name=allocated_orders,json=allocatedOrders,proto3" json:"allocated_orders,omitempty"`
}

func (x *GetReceiptByIDResponse_Item) Reset() {
	*x = GetReceiptByIDResponse_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inventory_v1_inventory_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptByIDResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptByIDResponse_Item) ProtoMessage() {}

func (x *GetReceiptByIDResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_v1_inventory_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptByIDResponse_Item.ProtoReflect.Descriptor instead.
func (*GetReceiptByIDResponse_Item) Descriptor() ([]byte, []int) {
	return file_proto_inventory_v1_inventory_proto_rawDescGZIP(), []int{7, 0}
}

func (x *GetReceiptByIDResponse_Item) GetProductUid() int32 {
	if x != nil {
		return x.ProductUid
	}
	return 0
}

func (x *GetReceiptByIDResponse_Item) GetProductPrimaryBin() string {
	if x != nil {
		return x.ProductPrimaryBin
	}
	return ""
}

func (x *GetReceiptByIDResponse_Item) GetReceivedQuantity() float64 {
	if x != nil {
		return x.ReceivedQuantity
	}
	return 0
}

func (x *GetReceiptByIDResponse_Item) GetReceivedUnit() string {
	if x != nil {
		return x.ReceivedUnit
	}
	return ""
}

func (x *GetReceiptByIDResponse_Item) GetAllocatedOrders() []string {
	if x != nil {
		return x.AllocatedOrders
	}
	return nil
}

type GetReceiptByIDResponse_Item_Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique ID of order
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Customer Name
	CustomerName string `protobuf:"bytes,2,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	// Quantity of product allocated
	AllocatedQuantity float64 `protobuf:"fixed64,3,opt,name=allocated_quantity,json=allocatedQuantity,proto3" json:"allocated_quantity,omitempty"`
	// Unit associated with allocated_quantity
	AllocatedUnit string `protobuf:"bytes,4,opt,name=allocated_unit,json=allocatedUnit,proto3" json:"allocated_unit,omitempty"`
}

func (x *GetReceiptByIDResponse_Item_Order) Reset() {
	*x = GetReceiptByIDResponse_Item_Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inventory_v1_inventory_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptByIDResponse_Item_Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptByIDResponse_Item_Order) ProtoMessage() {}

func (x *GetReceiptByIDResponse_Item_Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_v1_inventory_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptByIDResponse_Item_Order.ProtoReflect.Descriptor instead.
func (*GetReceiptByIDResponse_Item_Order) Descriptor() ([]byte, []int) {
	return file_proto_inventory_v1_inventory_proto_rawDescGZIP(), []int{7, 0, 0}
}

func (x *GetReceiptByIDResponse_Item_Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetReceiptByIDResponse_Item_Order) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *GetReceiptByIDResponse_Item_Order) GetAllocatedQuantity() float64 {
	if x != nil {
		return x.AllocatedQuantity
	}
	return 0
}

func (x *GetReceiptByIDResponse_Item_Order) GetAllocatedUnit() string {
	if x != nil {
		return x.AllocatedUnit
	}
	return ""
}

var File_proto_inventory_v1_inventory_proto protoreflect.FileDescriptor

var file_proto_inventory_v1_inventory_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x22, 0xed, 0x01, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x69,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x65, 0x61,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x6e, 0x6f, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74,
	0x4e, 0x6f, 0x22, 0x39, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x69, 0x64, 0x22, 0xd2, 0x01,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0d, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x1a, 0x5e, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x55, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x11, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd5, 0x03, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x3f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x1a, 0xe9, 0x02, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x62,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x92, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x32, 0xcf, 0x01,
	0x0a, 0x10, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x24, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x23, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_inventory_v1_inventory_proto_rawDescOnce sync.Once
	file_proto_inventory_v1_inventory_proto_rawDescData = file_proto_inventory_v1_inventory_proto_rawDesc
)

func file_proto_inventory_v1_inventory_proto_rawDescGZIP() []byte {
	file_proto_inventory_v1_inventory_proto_rawDescOnce.Do(func() {
		file_proto_inventory_v1_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_inventory_v1_inventory_proto_rawDescData)
	})
	return file_proto_inventory_v1_inventory_proto_rawDescData
}

var file_proto_inventory_v1_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_inventory_v1_inventory_proto_goTypes = []interface{}{
	(*AddSupplierRequest)(nil),                   // 0: inventory.v1.AddSupplierRequest
	(*GetProductStockRequest)(nil),               // 1: inventory.v1.GetProductStockRequest
	(*GetProductStockResponse)(nil),              // 2: inventory.v1.GetProductStockResponse
	(*AddSupplierResponse)(nil),                  // 3: inventory.v1.AddSupplierResponse
	(*DeleteSupplierRequest)(nil),                // 4: inventory.v1.DeleteSupplierRequest
	(*DeleteSupplierResponse)(nil),               // 5: inventory.v1.DeleteSupplierResponse
	(*GetReceiptByIDRequest)(nil),                // 6: inventory.v1.GetReceiptByIDRequest
	(*GetReceiptByIDResponse)(nil),               // 7: inventory.v1.GetReceiptByIDResponse
	(*GetProductStockResponse_ProductStock)(nil), // 8: inventory.v1.GetProductStockResponse.ProductStock
	(*GetReceiptByIDResponse_Item)(nil),          // 9: inventory.v1.GetReceiptByIDResponse.Item
	(*GetReceiptByIDResponse_Item_Order)(nil),    // 10: inventory.v1.GetReceiptByIDResponse.Item.Order
}
var file_proto_inventory_v1_inventory_proto_depIdxs = []int32{
	8, // 0: inventory.v1.GetProductStockResponse.product_stock:type_name -> inventory.v1.GetProductStockResponse.ProductStock
	9, // 1: inventory.v1.GetReceiptByIDResponse.items:type_name -> inventory.v1.GetReceiptByIDResponse.Item
	1, // 2: inventory.v1.InventoryService.GetProductStock:input_type -> inventory.v1.GetProductStockRequest
	6, // 3: inventory.v1.InventoryService.GetReceiptByID:input_type -> inventory.v1.GetReceiptByIDRequest
	2, // 4: inventory.v1.InventoryService.GetProductStock:output_type -> inventory.v1.GetProductStockResponse
	7, // 5: inventory.v1.InventoryService.GetReceiptByID:output_type -> inventory.v1.GetReceiptByIDResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_inventory_v1_inventory_proto_init() }
func file_proto_inventory_v1_inventory_proto_init() {
	if File_proto_inventory_v1_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_inventory_v1_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSupplierRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_inventory_v1_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductStockRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_inventory_v1_inventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductStockResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_inventory_v1_inventory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSupplierResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_inventory_v1_inventory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSupplierRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_inventory_v1_inventory_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSupplierResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_inventory_v1_inventory_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiptByIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_inventory_v1_inventory_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiptByIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_inventory_v1_inventory_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductStockResponse_ProductStock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_inventory_v1_inventory_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiptByIDResponse_Item); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_inventory_v1_inventory_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiptByIDResponse_Item_Order); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_inventory_v1_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_inventory_v1_inventory_proto_goTypes,
		DependencyIndexes: file_proto_inventory_v1_inventory_proto_depIdxs,
		MessageInfos:      file_proto_inventory_v1_inventory_proto_msgTypes,
	}.Build()
	File_proto_inventory_v1_inventory_proto = out.File
	file_proto_inventory_v1_inventory_proto_rawDesc = nil
	file_proto_inventory_v1_inventory_proto_goTypes = nil
	file_proto_inventory_v1_inventory_proto_depIdxs = nil
}
