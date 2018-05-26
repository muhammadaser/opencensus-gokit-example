// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event/transport/pb/event.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	event/transport/pb/event.proto

It has these top-level messages:
	Event
	CreateRequest
	CreateResponse
	GetRequest
	GetResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	ListRequest
	ListResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	Id   []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateRequest struct {
	TenantId []byte `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Event    *Event `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateRequest) GetTenantId() []byte {
	if m != nil {
		return m.TenantId
	}
	return nil
}

func (m *CreateRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type CreateResponse struct {
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateResponse) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type GetRequest struct {
	TenantId []byte `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Id       []byte `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetRequest) GetTenantId() []byte {
	if m != nil {
		return m.TenantId
	}
	return nil
}

func (m *GetRequest) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type GetResponse struct {
	Event *Event `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type UpdateRequest struct {
	TenantId []byte `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Event    *Event `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateRequest) GetTenantId() []byte {
	if m != nil {
		return m.TenantId
	}
	return nil
}

func (m *UpdateRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type DeleteRequest struct {
	TenantId []byte `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Id       []byte `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteRequest) GetTenantId() []byte {
	if m != nil {
		return m.TenantId
	}
	return nil
}

func (m *DeleteRequest) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ListRequest struct {
	TenantId []byte `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListRequest) GetTenantId() []byte {
	if m != nil {
		return m.TenantId
	}
	return nil
}

type ListResponse struct {
	Event *Event `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "pb.event")
	proto.RegisterType((*CreateRequest)(nil), "pb.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "pb.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "pb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "pb.GetResponse")
	proto.RegisterType((*UpdateRequest)(nil), "pb.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "pb.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "pb.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "pb.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "pb.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "pb.ListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Event service

type EventClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type eventClient struct {
	cc *grpc.ClientConn
}

func NewEventClient(cc *grpc.ClientConn) EventClient {
	return &eventClient{cc}
}

func (c *eventClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/pb.Event/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/pb.Event/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/pb.Event/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/pb.Event/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/pb.Event/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Event service

type EventServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterEventServer(s *grpc.Server, srv EventServer) {
	s.RegisterService(&_Event_serviceDesc, srv)
}

func _Event_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Event/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Event/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Event/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Event/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Event/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Event_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Event",
	HandlerType: (*EventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Event_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Event_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Event_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Event_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Event_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event/transport/pb/event.proto",
}

func init() { proto.RegisterFile("event/transport/pb/event.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x41, 0x6b, 0xb3, 0x40,
	0x10, 0x8d, 0x7e, 0x49, 0xf8, 0x32, 0x49, 0x6c, 0x3a, 0xa7, 0x60, 0xa1, 0x95, 0x3d, 0x49, 0x03,
	0x4a, 0xd2, 0x53, 0xa1, 0xb7, 0xb6, 0x84, 0x42, 0x73, 0x11, 0x7a, 0x2e, 0x8a, 0x73, 0x10, 0xda,
	0x75, 0xab, 0xdb, 0xfe, 0xf5, 0x5e, 0x8b, 0x3b, 0x6a, 0xa2, 0x50, 0xf0, 0xd0, 0x9b, 0x3e, 0xe7,
	0xbd, 0x7d, 0xf3, 0x9e, 0x0b, 0x97, 0xf4, 0x45, 0x52, 0x87, 0xba, 0x88, 0x65, 0xa9, 0xf2, 0x42,
	0x87, 0x2a, 0x09, 0x0d, 0x14, 0xa8, 0x22, 0xd7, 0x39, 0xda, 0x2a, 0x11, 0x1b, 0x98, 0x18, 0x08,
	0x1d, 0xb0, 0xb3, 0x74, 0x6d, 0x79, 0x96, 0xbf, 0x88, 0xec, 0x2c, 0x45, 0x84, 0xb1, 0x8c, 0xdf,
	0x69, 0x6d, 0x7b, 0x96, 0x3f, 0x8b, 0xcc, 0xb3, 0x38, 0xc0, 0xf2, 0xbe, 0xa0, 0x58, 0x53, 0x44,
	0x1f, 0x9f, 0x54, 0x6a, 0xbc, 0x80, 0x99, 0x26, 0x19, 0x4b, 0xfd, 0xda, 0x72, 0xff, 0x33, 0xf0,
	0x94, 0xe2, 0x55, 0x2d, 0x6d, 0x24, 0xe6, 0xbb, 0x59, 0xa0, 0x92, 0xc0, 0x00, 0x11, 0xe3, 0xc2,
	0x03, 0xa7, 0x91, 0x2b, 0x55, 0x2e, 0x4b, 0xea, 0x9b, 0x10, 0xb7, 0x00, 0x7b, 0xd2, 0x83, 0x4e,
	0x63, 0xaa, 0xdd, 0x52, 0x03, 0x98, 0x1b, 0x6a, 0xad, 0xdc, 0x9a, 0xb1, 0x7e, 0x31, 0x73, 0x80,
	0xe5, 0x8b, 0x4a, 0xff, 0x6c, 0xb7, 0x15, 0x38, 0x8d, 0x1c, 0x3b, 0x10, 0x77, 0xb0, 0x7c, 0xa0,
	0x37, 0x1a, 0x78, 0x40, 0x7f, 0x9d, 0x15, 0x38, 0x0d, 0xbb, 0xd6, 0xbb, 0x86, 0xf9, 0x73, 0x56,
	0x0e, 0x0a, 0x47, 0x84, 0xb0, 0xe0, 0xd9, 0x81, 0x69, 0xec, 0xbe, 0x2d, 0x98, 0x3c, 0x9a, 0xff,
	0x62, 0x0b, 0x53, 0x2e, 0x09, 0xcf, 0xab, 0xa9, 0x4e, 0xff, 0x2e, 0x9e, 0x42, 0xb5, 0xaf, 0x11,
	0xfa, 0xf0, 0x6f, 0x4f, 0x1a, 0x9d, 0xea, 0xe3, 0xb1, 0x3e, 0xf7, 0xac, 0x7d, 0x6f, 0x27, 0xb7,
	0x30, 0xe5, 0x94, 0x58, 0xbc, 0x53, 0x00, 0x8b, 0xf7, 0x42, 0x34, 0x14, 0x0e, 0x82, 0x29, 0x9d,
	0x48, 0x99, 0xd2, 0xcb, 0x69, 0x84, 0x1b, 0x18, 0x57, 0xdb, 0xa3, 0x31, 0x70, 0x92, 0x99, 0xbb,
	0x3a, 0x02, 0xcd, 0x70, 0x32, 0x35, 0x77, 0xe3, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0xad, 0x22,
	0xf3, 0x61, 0x3d, 0x03, 0x00, 0x00,
}
