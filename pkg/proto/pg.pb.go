// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pg.proto

package proto

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

type PGResponse struct {
	Succeed              bool     `protobuf:"varint,1,opt,name=succeed,proto3" json:"succeed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PGResponse) Reset()         { *m = PGResponse{} }
func (m *PGResponse) String() string { return proto.CompactTextString(m) }
func (*PGResponse) ProtoMessage()    {}
func (*PGResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pg_00c96af96ecf53ee, []int{0}
}
func (m *PGResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PGResponse.Unmarshal(m, b)
}
func (m *PGResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PGResponse.Marshal(b, m, deterministic)
}
func (dst *PGResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PGResponse.Merge(dst, src)
}
func (m *PGResponse) XXX_Size() int {
	return xxx_messageInfo_PGResponse.Size(m)
}
func (m *PGResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PGResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PGResponse proto.InternalMessageInfo

func (m *PGResponse) GetSucceed() bool {
	if m != nil {
		return m.Succeed
	}
	return false
}

type UpdatePostgresqlConfRequest struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	IsMaster             bool     `protobuf:"varint,2,opt,name=is_master,json=isMaster,proto3" json:"is_master,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePostgresqlConfRequest) Reset()         { *m = UpdatePostgresqlConfRequest{} }
func (m *UpdatePostgresqlConfRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePostgresqlConfRequest) ProtoMessage()    {}
func (*UpdatePostgresqlConfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pg_00c96af96ecf53ee, []int{1}
}
func (m *UpdatePostgresqlConfRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePostgresqlConfRequest.Unmarshal(m, b)
}
func (m *UpdatePostgresqlConfRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePostgresqlConfRequest.Marshal(b, m, deterministic)
}
func (dst *UpdatePostgresqlConfRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePostgresqlConfRequest.Merge(dst, src)
}
func (m *UpdatePostgresqlConfRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePostgresqlConfRequest.Size(m)
}
func (m *UpdatePostgresqlConfRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePostgresqlConfRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePostgresqlConfRequest proto.InternalMessageInfo

func (m *UpdatePostgresqlConfRequest) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *UpdatePostgresqlConfRequest) GetIsMaster() bool {
	if m != nil {
		return m.IsMaster
	}
	return false
}

type UpdatePGHBAConfRequest struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	AnotherIp            string   `protobuf:"bytes,2,opt,name=another_ip,json=anotherIp,proto3" json:"another_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePGHBAConfRequest) Reset()         { *m = UpdatePGHBAConfRequest{} }
func (m *UpdatePGHBAConfRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePGHBAConfRequest) ProtoMessage()    {}
func (*UpdatePGHBAConfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pg_00c96af96ecf53ee, []int{2}
}
func (m *UpdatePGHBAConfRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePGHBAConfRequest.Unmarshal(m, b)
}
func (m *UpdatePGHBAConfRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePGHBAConfRequest.Marshal(b, m, deterministic)
}
func (dst *UpdatePGHBAConfRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePGHBAConfRequest.Merge(dst, src)
}
func (m *UpdatePGHBAConfRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePGHBAConfRequest.Size(m)
}
func (m *UpdatePGHBAConfRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePGHBAConfRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePGHBAConfRequest proto.InternalMessageInfo

func (m *UpdatePGHBAConfRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UpdatePGHBAConfRequest) GetAnotherIp() string {
	if m != nil {
		return m.AnotherIp
	}
	return ""
}

type CreateRecoveryConfRequest struct {
	AnotherIp            string   `protobuf:"bytes,1,opt,name=another_ip,json=anotherIp,proto3" json:"another_ip,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRecoveryConfRequest) Reset()         { *m = CreateRecoveryConfRequest{} }
func (m *CreateRecoveryConfRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRecoveryConfRequest) ProtoMessage()    {}
func (*CreateRecoveryConfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pg_00c96af96ecf53ee, []int{3}
}
func (m *CreateRecoveryConfRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRecoveryConfRequest.Unmarshal(m, b)
}
func (m *CreateRecoveryConfRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRecoveryConfRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRecoveryConfRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRecoveryConfRequest.Merge(dst, src)
}
func (m *CreateRecoveryConfRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRecoveryConfRequest.Size(m)
}
func (m *CreateRecoveryConfRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRecoveryConfRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRecoveryConfRequest proto.InternalMessageInfo

func (m *CreateRecoveryConfRequest) GetAnotherIp() string {
	if m != nil {
		return m.AnotherIp
	}
	return ""
}

func (m *CreateRecoveryConfRequest) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type DeleteRecoveryConfRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRecoveryConfRequest) Reset()         { *m = DeleteRecoveryConfRequest{} }
func (m *DeleteRecoveryConfRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRecoveryConfRequest) ProtoMessage()    {}
func (*DeleteRecoveryConfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pg_00c96af96ecf53ee, []int{4}
}
func (m *DeleteRecoveryConfRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRecoveryConfRequest.Unmarshal(m, b)
}
func (m *DeleteRecoveryConfRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRecoveryConfRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRecoveryConfRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRecoveryConfRequest.Merge(dst, src)
}
func (m *DeleteRecoveryConfRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRecoveryConfRequest.Size(m)
}
func (m *DeleteRecoveryConfRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRecoveryConfRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRecoveryConfRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PGResponse)(nil), "proto.PGResponse")
	proto.RegisterType((*UpdatePostgresqlConfRequest)(nil), "proto.UpdatePostgresqlConfRequest")
	proto.RegisterType((*UpdatePGHBAConfRequest)(nil), "proto.UpdatePGHBAConfRequest")
	proto.RegisterType((*CreateRecoveryConfRequest)(nil), "proto.CreateRecoveryConfRequest")
	proto.RegisterType((*DeleteRecoveryConfRequest)(nil), "proto.DeleteRecoveryConfRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PGManagerClient is the client API for PGManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PGManagerClient interface {
	UpdatePostgresqlConf(ctx context.Context, in *UpdatePostgresqlConfRequest, opts ...grpc.CallOption) (*PGResponse, error)
	UpdatePGHBAConf(ctx context.Context, in *UpdatePGHBAConfRequest, opts ...grpc.CallOption) (*PGResponse, error)
	CreateRecoveryConf(ctx context.Context, in *CreateRecoveryConfRequest, opts ...grpc.CallOption) (*PGResponse, error)
	DeleteRecoveryConf(ctx context.Context, in *DeleteRecoveryConfRequest, opts ...grpc.CallOption) (*PGResponse, error)
}

type pGManagerClient struct {
	cc *grpc.ClientConn
}

func NewPGManagerClient(cc *grpc.ClientConn) PGManagerClient {
	return &pGManagerClient{cc}
}

func (c *pGManagerClient) UpdatePostgresqlConf(ctx context.Context, in *UpdatePostgresqlConfRequest, opts ...grpc.CallOption) (*PGResponse, error) {
	out := new(PGResponse)
	err := c.cc.Invoke(ctx, "/proto.PGManager/UpdatePostgresqlConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pGManagerClient) UpdatePGHBAConf(ctx context.Context, in *UpdatePGHBAConfRequest, opts ...grpc.CallOption) (*PGResponse, error) {
	out := new(PGResponse)
	err := c.cc.Invoke(ctx, "/proto.PGManager/UpdatePGHBAConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pGManagerClient) CreateRecoveryConf(ctx context.Context, in *CreateRecoveryConfRequest, opts ...grpc.CallOption) (*PGResponse, error) {
	out := new(PGResponse)
	err := c.cc.Invoke(ctx, "/proto.PGManager/CreateRecoveryConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pGManagerClient) DeleteRecoveryConf(ctx context.Context, in *DeleteRecoveryConfRequest, opts ...grpc.CallOption) (*PGResponse, error) {
	out := new(PGResponse)
	err := c.cc.Invoke(ctx, "/proto.PGManager/DeleteRecoveryConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PGManagerServer is the server API for PGManager service.
type PGManagerServer interface {
	UpdatePostgresqlConf(context.Context, *UpdatePostgresqlConfRequest) (*PGResponse, error)
	UpdatePGHBAConf(context.Context, *UpdatePGHBAConfRequest) (*PGResponse, error)
	CreateRecoveryConf(context.Context, *CreateRecoveryConfRequest) (*PGResponse, error)
	DeleteRecoveryConf(context.Context, *DeleteRecoveryConfRequest) (*PGResponse, error)
}

func RegisterPGManagerServer(s *grpc.Server, srv PGManagerServer) {
	s.RegisterService(&_PGManager_serviceDesc, srv)
}

func _PGManager_UpdatePostgresqlConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostgresqlConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PGManagerServer).UpdatePostgresqlConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PGManager/UpdatePostgresqlConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PGManagerServer).UpdatePostgresqlConf(ctx, req.(*UpdatePostgresqlConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PGManager_UpdatePGHBAConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePGHBAConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PGManagerServer).UpdatePGHBAConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PGManager/UpdatePGHBAConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PGManagerServer).UpdatePGHBAConf(ctx, req.(*UpdatePGHBAConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PGManager_CreateRecoveryConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecoveryConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PGManagerServer).CreateRecoveryConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PGManager/CreateRecoveryConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PGManagerServer).CreateRecoveryConf(ctx, req.(*CreateRecoveryConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PGManager_DeleteRecoveryConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecoveryConfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PGManagerServer).DeleteRecoveryConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PGManager/DeleteRecoveryConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PGManagerServer).DeleteRecoveryConf(ctx, req.(*DeleteRecoveryConfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PGManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PGManager",
	HandlerType: (*PGManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdatePostgresqlConf",
			Handler:    _PGManager_UpdatePostgresqlConf_Handler,
		},
		{
			MethodName: "UpdatePGHBAConf",
			Handler:    _PGManager_UpdatePGHBAConf_Handler,
		},
		{
			MethodName: "CreateRecoveryConf",
			Handler:    _PGManager_CreateRecoveryConf_Handler,
		},
		{
			MethodName: "DeleteRecoveryConf",
			Handler:    _PGManager_DeleteRecoveryConf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pg.proto",
}

func init() { proto.RegisterFile("pg.proto", fileDescriptor_pg_00c96af96ecf53ee) }

var fileDescriptor_pg_00c96af96ecf53ee = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x4d, 0xf0, 0x4f, 0x32, 0x20, 0xe2, 0x20, 0x92, 0x1a, 0x0a, 0x65, 0x0f, 0xe2, 0xa9,
	0x07, 0x7d, 0x02, 0xad, 0x12, 0xa5, 0xb4, 0x86, 0x05, 0xcf, 0x65, 0x4d, 0xc7, 0x18, 0xa8, 0xd9,
	0xed, 0xee, 0x46, 0xf0, 0xe1, 0x7c, 0x37, 0x61, 0x8d, 0xad, 0x4d, 0x93, 0x9e, 0x76, 0x66, 0x98,
	0xdf, 0xc7, 0xf0, 0x2d, 0x04, 0x2a, 0x1f, 0x2a, 0x2d, 0xad, 0xc4, 0x03, 0xf7, 0xb0, 0x4b, 0x80,
	0x34, 0xe1, 0x64, 0x94, 0x2c, 0x0d, 0x61, 0x04, 0x47, 0xa6, 0xca, 0x32, 0xa2, 0x79, 0xe4, 0x0d,
	0xbc, 0xab, 0x80, 0xff, 0xb5, 0x6c, 0x0a, 0xf1, 0x8b, 0x9a, 0x0b, 0x4b, 0xa9, 0x34, 0x36, 0xd7,
	0x64, 0x96, 0x8b, 0x91, 0x2c, 0xdf, 0x38, 0x2d, 0x2b, 0x32, 0x16, 0x11, 0xf6, 0x95, 0xd4, 0xd6,
	0xa5, 0x8e, 0xb9, 0xab, 0x31, 0x86, 0xb0, 0x30, 0xb3, 0x0f, 0x61, 0x2c, 0xe9, 0xc8, 0x77, 0xb8,
	0xa0, 0x30, 0x13, 0xd7, 0xb3, 0x31, 0x9c, 0xd7, 0xbc, 0xe4, 0xf1, 0xee, 0xb6, 0x81, 0xaa, 0x0c,
	0x69, 0x87, 0x0a, 0xb9, 0xab, 0xb1, 0x0f, 0x20, 0x4a, 0x69, 0xdf, 0x49, 0xcf, 0x0a, 0xe5, 0x58,
	0x21, 0x0f, 0xeb, 0xc9, 0x93, 0x62, 0x53, 0xe8, 0x8d, 0x34, 0x09, 0x4b, 0x9c, 0x32, 0xf9, 0x49,
	0xfa, 0xeb, 0x3f, 0x6f, 0x33, 0xeb, 0x35, 0xb2, 0xab, 0xcb, 0xfd, 0xf5, 0xe5, 0x2c, 0x86, 0xde,
	0x3d, 0x2d, 0xa8, 0x95, 0x77, 0xfd, 0xed, 0x43, 0x98, 0x26, 0x13, 0x51, 0x8a, 0x9c, 0x34, 0x3e,
	0xc3, 0x59, 0x9b, 0x17, 0x64, 0xbf, 0x9a, 0x87, 0x3b, 0xa4, 0x5d, 0x9c, 0xd6, 0x3b, 0xeb, 0x0f,
	0x60, 0x7b, 0xf8, 0x00, 0x27, 0x0d, 0x31, 0xd8, 0xdf, 0x64, 0x35, 0x84, 0xb5, 0x63, 0xc6, 0x80,
	0xdb, 0x4a, 0x70, 0x50, 0xaf, 0x76, 0xda, 0xea, 0x84, 0x6d, 0xfb, 0x58, 0xc1, 0x3a, 0x55, 0xb5,
	0xc2, 0x5e, 0x0f, 0xdd, 0xec, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x41, 0xa9, 0x3a, 0x8b,
	0x02, 0x00, 0x00,
}
