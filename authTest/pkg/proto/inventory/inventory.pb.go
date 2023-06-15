// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: pkg/proto/inventory/inventory.proto

package inventory

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

type CaseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemName        string `protobuf:"bytes,2,opt,name=itemName,proto3" json:"itemName,omitempty"`
	ItemDescription string `protobuf:"bytes,3,opt,name=itemDescription,proto3" json:"itemDescription,omitempty"`
	Type            string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Stars           int64  `protobuf:"varint,5,opt,name=stars,proto3" json:"stars,omitempty"`
	Image           []byte `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *CaseItem) Reset() {
	*x = CaseItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_inventory_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaseItem) ProtoMessage() {}

func (x *CaseItem) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_inventory_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaseItem.ProtoReflect.Descriptor instead.
func (*CaseItem) Descriptor() ([]byte, []int) {
	return file_pkg_proto_inventory_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *CaseItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CaseItem) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *CaseItem) GetItemDescription() string {
	if x != nil {
		return x.ItemDescription
	}
	return ""
}

func (x *CaseItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CaseItem) GetStars() int64 {
	if x != nil {
		return x.Stars
	}
	return 0
}

func (x *CaseItem) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type Inventory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Userid int64       `protobuf:"varint,2,opt,name=userid,proto3" json:"userid,omitempty"`
	Items  []*CaseItem `protobuf:"bytes,3,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *Inventory) Reset() {
	*x = Inventory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_inventory_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inventory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inventory) ProtoMessage() {}

func (x *Inventory) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_inventory_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inventory.ProtoReflect.Descriptor instead.
func (*Inventory) Descriptor() ([]byte, []int) {
	return file_pkg_proto_inventory_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *Inventory) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Inventory) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *Inventory) GetItems() []*CaseItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type Confirm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok      bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Confirm) Reset() {
	*x = Confirm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_inventory_inventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirm) ProtoMessage() {}

func (x *Confirm) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_inventory_inventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirm.ProtoReflect.Descriptor instead.
func (*Confirm) Descriptor() ([]byte, []int) {
	return file_pkg_proto_inventory_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *Confirm) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *Confirm) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type InventoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenValue string `protobuf:"bytes,1,opt,name=tokenValue,proto3" json:"tokenValue,omitempty"`
	Id         int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InventoryRequest) Reset() {
	*x = InventoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_inventory_inventory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryRequest) ProtoMessage() {}

func (x *InventoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_inventory_inventory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryRequest.ProtoReflect.Descriptor instead.
func (*InventoryRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_inventory_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *InventoryRequest) GetTokenValue() string {
	if x != nil {
		return x.TokenValue
	}
	return ""
}

func (x *InventoryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_pkg_proto_inventory_inventory_proto protoreflect.FileDescriptor

var file_pkg_proto_inventory_inventory_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x22, 0xa0, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x74, 0x65,
	0x6d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x69, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x22, 0x5e, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x33, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x42, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xdc, 0x01, 0x0a,
	0x10, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x40, 0x0a, 0x0b, 0x54, 0x6f, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x1b, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e,
	0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_proto_inventory_inventory_proto_rawDescOnce sync.Once
	file_pkg_proto_inventory_inventory_proto_rawDescData = file_pkg_proto_inventory_inventory_proto_rawDesc
)

func file_pkg_proto_inventory_inventory_proto_rawDescGZIP() []byte {
	file_pkg_proto_inventory_inventory_proto_rawDescOnce.Do(func() {
		file_pkg_proto_inventory_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_inventory_inventory_proto_rawDescData)
	})
	return file_pkg_proto_inventory_inventory_proto_rawDescData
}

var file_pkg_proto_inventory_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_proto_inventory_inventory_proto_goTypes = []interface{}{
	(*CaseItem)(nil),         // 0: inventory.CaseItem
	(*Inventory)(nil),        // 1: inventory.Inventory
	(*Confirm)(nil),          // 2: inventory.Confirm
	(*InventoryRequest)(nil), // 3: inventory.InventoryRequest
}
var file_pkg_proto_inventory_inventory_proto_depIdxs = []int32{
	0, // 0: inventory.Inventory.Items:type_name -> inventory.CaseItem
	3, // 1: inventory.InventoryService.ToInventory:input_type -> inventory.InventoryRequest
	3, // 2: inventory.InventoryService.NewInventory:input_type -> inventory.InventoryRequest
	3, // 3: inventory.InventoryService.GetInventory:input_type -> inventory.InventoryRequest
	2, // 4: inventory.InventoryService.ToInventory:output_type -> inventory.Confirm
	2, // 5: inventory.InventoryService.NewInventory:output_type -> inventory.Confirm
	1, // 6: inventory.InventoryService.GetInventory:output_type -> inventory.Inventory
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_inventory_inventory_proto_init() }
func file_pkg_proto_inventory_inventory_proto_init() {
	if File_pkg_proto_inventory_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_inventory_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaseItem); i {
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
		file_pkg_proto_inventory_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inventory); i {
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
		file_pkg_proto_inventory_inventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Confirm); i {
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
		file_pkg_proto_inventory_inventory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryRequest); i {
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
			RawDescriptor: file_pkg_proto_inventory_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_inventory_inventory_proto_goTypes,
		DependencyIndexes: file_pkg_proto_inventory_inventory_proto_depIdxs,
		MessageInfos:      file_pkg_proto_inventory_inventory_proto_msgTypes,
	}.Build()
	File_pkg_proto_inventory_inventory_proto = out.File
	file_pkg_proto_inventory_inventory_proto_rawDesc = nil
	file_pkg_proto_inventory_inventory_proto_goTypes = nil
	file_pkg_proto_inventory_inventory_proto_depIdxs = nil
}
