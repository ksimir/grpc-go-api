// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inventory.proto

package inventory

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for adding/deleting/getting an Item
type ItemRequest struct {
	// API versioning
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Pid                  string   `protobuf:"bytes,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Quantity             int64    `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemRequest) Reset()         { *m = ItemRequest{} }
func (m *ItemRequest) String() string { return proto.CompactTextString(m) }
func (*ItemRequest) ProtoMessage()    {}
func (*ItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7173caedb7c6ae96, []int{0}
}

func (m *ItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemRequest.Unmarshal(m, b)
}
func (m *ItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemRequest.Marshal(b, m, deterministic)
}
func (m *ItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemRequest.Merge(m, src)
}
func (m *ItemRequest) XXX_Size() int {
	return xxx_messageInfo_ItemRequest.Size(m)
}
func (m *ItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ItemRequest proto.InternalMessageInfo

func (m *ItemRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ItemRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemRequest) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *ItemRequest) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type ItemResponse struct {
	// API versioning
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemResponse) Reset()         { *m = ItemResponse{} }
func (m *ItemResponse) String() string { return proto.CompactTextString(m) }
func (*ItemResponse) ProtoMessage()    {}
func (*ItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7173caedb7c6ae96, []int{1}
}

func (m *ItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemResponse.Unmarshal(m, b)
}
func (m *ItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemResponse.Marshal(b, m, deterministic)
}
func (m *ItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemResponse.Merge(m, src)
}
func (m *ItemResponse) XXX_Size() int {
	return xxx_messageInfo_ItemResponse.Size(m)
}
func (m *ItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ItemResponse proto.InternalMessageInfo

func (m *ItemResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ItemResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ItemResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*ItemRequest)(nil), "inventory.ItemRequest")
	proto.RegisterType((*ItemResponse)(nil), "inventory.ItemResponse")
}

func init() { proto.RegisterFile("inventory.proto", fileDescriptor_7173caedb7c6ae96) }

var fileDescriptor_7173caedb7c6ae96 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0x12, 0x35, 0xc9, 0x53, 0x54, 0x06, 0xd1, 0xa5, 0xa7, 0x92, 0x53, 0x4f, 0x45, 0xf4,
	0xa8, 0x58, 0x04, 0x41, 0x7a, 0x93, 0x80, 0x47, 0x0f, 0xb1, 0x3b, 0x87, 0x05, 0xcd, 0x6e, 0x77,
	0x27, 0x42, 0xff, 0xa1, 0x3f, 0x4b, 0x92, 0xb4, 0x55, 0xd0, 0x1e, 0xec, 0x6d, 0xdf, 0xe3, 0x7d,
	0x0d, 0x2c, 0x4e, 0x4c, 0xfd, 0xc1, 0xb5, 0x58, 0xbf, 0x9c, 0x38, 0x6f, 0xc5, 0x52, 0xbe, 0x21,
	0x8a, 0x17, 0x1c, 0xce, 0x84, 0xdf, 0x4b, 0x5e, 0x34, 0x1c, 0x84, 0x4e, 0x91, 0x54, 0xce, 0xa8,
	0x68, 0x14, 0x8d, 0xf3, 0xb2, 0x7d, 0xd2, 0x31, 0x62, 0xa3, 0x55, 0x3c, 0x8a, 0xc6, 0x49, 0x19,
	0x1b, 0xdd, 0x2a, 0x9c, 0xd1, 0x2a, 0xe9, 0x15, 0xce, 0x68, 0x1a, 0x22, 0x5b, 0x34, 0x55, 0x2d,
	0x46, 0x96, 0x6a, 0xaf, 0xd3, 0x6d, 0x70, 0xf1, 0x84, 0xa3, 0x3e, 0x3e, 0x38, 0x5b, 0x07, 0xfe,
	0x23, 0x5f, 0x21, 0x0d, 0xcd, 0x7c, 0xce, 0x21, 0x74, 0x25, 0x59, 0xb9, 0x86, 0x74, 0x86, 0x7d,
	0xf6, 0xde, 0xfa, 0x55, 0x57, 0x0f, 0xae, 0x3e, 0x63, 0xe4, 0xb3, 0xf5, 0x7c, 0xba, 0x45, 0x7a,
	0xaf, 0x75, 0x5b, 0x41, 0xe7, 0x93, 0xef, 0x33, 0x7f, 0x9c, 0x34, 0xbc, 0xf8, 0xc5, 0xf7, 0x5b,
	0x8a, 0x01, 0xdd, 0x21, 0x7b, 0x64, 0x69, 0xc9, 0xb0, 0xd5, 0xbe, 0x85, 0x2f, 0x06, 0x97, 0x11,
	0xdd, 0x20, 0x5d, 0xf9, 0xff, 0x6f, 0xa7, 0x29, 0xf0, 0xec, 0x74, 0x25, 0xbc, 0xeb, 0xfa, 0x29,
	0xf0, 0xc0, 0x6f, 0xbc, 0x73, 0xc0, 0xeb, 0x41, 0xf7, 0x1b, 0xae, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xd5, 0x77, 0x07, 0x07, 0x20, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// InventoryClient is the client API for Inventory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InventoryClient interface {
	// Add a new Item in the inventory
	AddItem(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemResponse, error)
	// Get all items from a specific player
	GetItems(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (Inventory_GetItemsClient, error)
	// Return a specific item based on its id and player id
	GetItem(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemRequest, error)
	// Update an item
	UpdateItem(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemResponse, error)
	// Delete an item
	DeleteItem(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemResponse, error)
}

type inventoryClient struct {
	cc *grpc.ClientConn
}

func NewInventoryClient(cc *grpc.ClientConn) InventoryClient {
	return &inventoryClient{cc}
}

func (c *inventoryClient) AddItem(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemResponse, error) {
	out := new(ItemResponse)
	err := c.cc.Invoke(ctx, "/inventory.Inventory/AddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) GetItems(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (Inventory_GetItemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Inventory_serviceDesc.Streams[0], "/inventory.Inventory/GetItems", opts...)
	if err != nil {
		return nil, err
	}
	x := &inventoryGetItemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Inventory_GetItemsClient interface {
	Recv() (*ItemRequest, error)
	grpc.ClientStream
}

type inventoryGetItemsClient struct {
	grpc.ClientStream
}

func (x *inventoryGetItemsClient) Recv() (*ItemRequest, error) {
	m := new(ItemRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *inventoryClient) GetItem(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemRequest, error) {
	out := new(ItemRequest)
	err := c.cc.Invoke(ctx, "/inventory.Inventory/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) UpdateItem(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemResponse, error) {
	out := new(ItemResponse)
	err := c.cc.Invoke(ctx, "/inventory.Inventory/UpdateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) DeleteItem(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemResponse, error) {
	out := new(ItemResponse)
	err := c.cc.Invoke(ctx, "/inventory.Inventory/DeleteItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServer is the server API for Inventory service.
type InventoryServer interface {
	// Add a new Item in the inventory
	AddItem(context.Context, *ItemRequest) (*ItemResponse, error)
	// Get all items from a specific player
	GetItems(*ItemRequest, Inventory_GetItemsServer) error
	// Return a specific item based on its id and player id
	GetItem(context.Context, *ItemRequest) (*ItemRequest, error)
	// Update an item
	UpdateItem(context.Context, *ItemRequest) (*ItemResponse, error)
	// Delete an item
	DeleteItem(context.Context, *ItemRequest) (*ItemResponse, error)
}

func RegisterInventoryServer(s *grpc.Server, srv InventoryServer) {
	s.RegisterService(&_Inventory_serviceDesc, srv)
}

func _Inventory_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Inventory/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).AddItem(ctx, req.(*ItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_GetItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ItemRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InventoryServer).GetItems(m, &inventoryGetItemsServer{stream})
}

type Inventory_GetItemsServer interface {
	Send(*ItemRequest) error
	grpc.ServerStream
}

type inventoryGetItemsServer struct {
	grpc.ServerStream
}

func (x *inventoryGetItemsServer) Send(m *ItemRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _Inventory_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Inventory/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).GetItem(ctx, req.(*ItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Inventory/UpdateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).UpdateItem(ctx, req.(*ItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Inventory/DeleteItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).DeleteItem(ctx, req.(*ItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Inventory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.Inventory",
	HandlerType: (*InventoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItem",
			Handler:    _Inventory_AddItem_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _Inventory_GetItem_Handler,
		},
		{
			MethodName: "UpdateItem",
			Handler:    _Inventory_UpdateItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _Inventory_DeleteItem_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetItems",
			Handler:       _Inventory_GetItems_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "inventory.proto",
}