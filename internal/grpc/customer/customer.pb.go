// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: customer.proto

package customer

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

type Branch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CustomerId string `protobuf:"bytes,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *Branch) Reset() {
	*x = Branch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Branch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch) ProtoMessage() {}

func (x *Branch) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Branch.ProtoReflect.Descriptor instead.
func (*Branch) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Branch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Branch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Branch) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type BranchSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BranchSummary) Reset() {
	*x = BranchSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchSummary) ProtoMessage() {}

func (x *BranchSummary) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchSummary.ProtoReflect.Descriptor instead.
func (*BranchSummary) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{1}
}

func (x *BranchSummary) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BranchSummary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetBranchesForContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContactId string `protobuf:"bytes,1,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
}

func (x *GetBranchesForContactRequest) Reset() {
	*x = GetBranchesForContactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBranchesForContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranchesForContactRequest) ProtoMessage() {}

func (x *GetBranchesForContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBranchesForContactRequest.ProtoReflect.Descriptor instead.
func (*GetBranchesForContactRequest) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{2}
}

func (x *GetBranchesForContactRequest) GetContactId() string {
	if x != nil {
		return x.ContactId
	}
	return ""
}

type GetBranchesForContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Branches []*BranchSummary `protobuf:"bytes,1,rep,name=branches,proto3" json:"branches,omitempty"`
}

func (x *GetBranchesForContactResponse) Reset() {
	*x = GetBranchesForContactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBranchesForContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranchesForContactResponse) ProtoMessage() {}

func (x *GetBranchesForContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBranchesForContactResponse.ProtoReflect.Descriptor instead.
func (*GetBranchesForContactResponse) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{3}
}

func (x *GetBranchesForContactResponse) GetBranches() []*BranchSummary {
	if x != nil {
		return x.Branches
	}
	return nil
}

type GetBranchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBranchRequest) Reset() {
	*x = GetBranchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranchRequest) ProtoMessage() {}

func (x *GetBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBranchRequest.ProtoReflect.Descriptor instead.
func (*GetBranchRequest) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{4}
}

func (x *GetBranchRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBranchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Branch *Branch `protobuf:"bytes,1,opt,name=branch,proto3" json:"branch,omitempty"`
}

func (x *GetBranchResponse) Reset() {
	*x = GetBranchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBranchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranchResponse) ProtoMessage() {}

func (x *GetBranchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBranchResponse.ProtoReflect.Descriptor instead.
func (*GetBranchResponse) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{5}
}

func (x *GetBranchResponse) GetBranch() *Branch {
	if x != nil {
		return x.Branch
	}
	return nil
}

var File_customer_proto protoreflect.FileDescriptor

var file_customer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x4d, 0x0a,
	0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x0d,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x3d, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73,
	0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x57, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x46,
	0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52,
	0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x32,
	0xcd, 0x01, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x65, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x29, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65,
	0x73, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x12, 0x1d, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customer_proto_rawDescOnce sync.Once
	file_customer_proto_rawDescData = file_customer_proto_rawDesc
)

func file_customer_proto_rawDescGZIP() []byte {
	file_customer_proto_rawDescOnce.Do(func() {
		file_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_customer_proto_rawDescData)
	})
	return file_customer_proto_rawDescData
}

var file_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_customer_proto_goTypes = []any{
	(*Branch)(nil),                        // 0: customer.v1.Branch
	(*BranchSummary)(nil),                 // 1: customer.v1.BranchSummary
	(*GetBranchesForContactRequest)(nil),  // 2: customer.v1.GetBranchesForContactRequest
	(*GetBranchesForContactResponse)(nil), // 3: customer.v1.GetBranchesForContactResponse
	(*GetBranchRequest)(nil),              // 4: customer.v1.GetBranchRequest
	(*GetBranchResponse)(nil),             // 5: customer.v1.GetBranchResponse
}
var file_customer_proto_depIdxs = []int32{
	1, // 0: customer.v1.GetBranchesForContactResponse.branches:type_name -> customer.v1.BranchSummary
	0, // 1: customer.v1.GetBranchResponse.branch:type_name -> customer.v1.Branch
	2, // 2: customer.v1.CustomerService.GetBranchesForContact:input_type -> customer.v1.GetBranchesForContactRequest
	4, // 3: customer.v1.CustomerService.GetBranch:input_type -> customer.v1.GetBranchRequest
	3, // 4: customer.v1.CustomerService.GetBranchesForContact:output_type -> customer.v1.GetBranchesForContactResponse
	5, // 5: customer.v1.CustomerService.GetBranch:output_type -> customer.v1.GetBranchResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_customer_proto_init() }
func file_customer_proto_init() {
	if File_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Branch); i {
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
		file_customer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BranchSummary); i {
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
		file_customer_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetBranchesForContactRequest); i {
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
		file_customer_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetBranchesForContactResponse); i {
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
		file_customer_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetBranchRequest); i {
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
		file_customer_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetBranchResponse); i {
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
			RawDescriptor: file_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customer_proto_goTypes,
		DependencyIndexes: file_customer_proto_depIdxs,
		MessageInfos:      file_customer_proto_msgTypes,
	}.Build()
	File_customer_proto = out.File
	file_customer_proto_rawDesc = nil
	file_customer_proto_goTypes = nil
	file_customer_proto_depIdxs = nil
}
