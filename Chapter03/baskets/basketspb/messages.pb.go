// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: basketspb/messages.proto

package basketspb

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

type Basket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Items []*Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Basket) Reset() {
	*x = Basket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Basket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Basket) ProtoMessage() {}

func (x *Basket) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Basket.ProtoReflect.Descriptor instead.
func (*Basket) Descriptor() ([]byte, []int) {
	return file_basketspb_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Basket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Basket) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId      string  `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	ProductId    string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	StoreName    string  `protobuf:"bytes,3,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	ProductName  string  `protobuf:"bytes,4,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	ProductPrice float64 `protobuf:"fixed64,5,opt,name=product_price,json=productPrice,proto3" json:"product_price,omitempty"`
	Quantity     int32   `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_basketspb_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Item) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *Item) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Item) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *Item) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Item) GetProductPrice() float64 {
	if x != nil {
		return x.ProductPrice
	}
	return 0
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_basketspb_messages_proto protoreflect.FileDescriptor

var file_basketspb_messages_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x73, 0x70, 0x62, 0x22, 0x3f, 0x0a, 0x06, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xc3, 0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0xa6, 0x01, 0x0a,
	0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x42, 0x0d,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x75, 0x73, 0x2f, 0x65, 0x64, 0x61, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x2d, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x63, 0x68, 0x34, 0x2f, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x2f,
	0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x73, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x42, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x73, 0x70, 0x62, 0xca, 0x02, 0x09, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70,
	0x62, 0xe2, 0x02, 0x15, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x42, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basketspb_messages_proto_rawDescOnce sync.Once
	file_basketspb_messages_proto_rawDescData = file_basketspb_messages_proto_rawDesc
)

func file_basketspb_messages_proto_rawDescGZIP() []byte {
	file_basketspb_messages_proto_rawDescOnce.Do(func() {
		file_basketspb_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_basketspb_messages_proto_rawDescData)
	})
	return file_basketspb_messages_proto_rawDescData
}

var file_basketspb_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_basketspb_messages_proto_goTypes = []interface{}{
	(*Basket)(nil), // 0: basketspb.Basket
	(*Item)(nil),   // 1: basketspb.Item
}
var file_basketspb_messages_proto_depIdxs = []int32{
	1, // 0: basketspb.Basket.items:type_name -> basketspb.Item
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_basketspb_messages_proto_init() }
func file_basketspb_messages_proto_init() {
	if File_basketspb_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_basketspb_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Basket); i {
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
		file_basketspb_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
			RawDescriptor: file_basketspb_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_basketspb_messages_proto_goTypes,
		DependencyIndexes: file_basketspb_messages_proto_depIdxs,
		MessageInfos:      file_basketspb_messages_proto_msgTypes,
	}.Build()
	File_basketspb_messages_proto = out.File
	file_basketspb_messages_proto_rawDesc = nil
	file_basketspb_messages_proto_goTypes = nil
	file_basketspb_messages_proto_depIdxs = nil
}
