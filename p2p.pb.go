// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p2p.proto

package p2p

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

type PingPong_Message int32

const (
	PingPong_PING PingPong_Message = 0
	PingPong_PONG PingPong_Message = 1
)

var PingPong_Message_name = map[int32]string{
	0: "PING",
	1: "PONG",
}
var PingPong_Message_value = map[string]int32{
	"PING": 0,
	"PONG": 1,
}

func (x PingPong_Message) String() string {
	return proto.EnumName(PingPong_Message_name, int32(x))
}
func (PingPong_Message) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_p2p_da16c222cfc6579f, []int{2, 0}
}

// Peer is the minimum entity that a node can communicate.
type Peer struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_da16c222cfc6579f, []int{0}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (dst *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(dst, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Peer) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// Peers is a wrapper for list of peer
type Peers struct {
	Peers                []*Peer  `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peers) Reset()         { *m = Peers{} }
func (m *Peers) String() string { return proto.CompactTextString(m) }
func (*Peers) ProtoMessage()    {}
func (*Peers) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_da16c222cfc6579f, []int{1}
}
func (m *Peers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peers.Unmarshal(m, b)
}
func (m *Peers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peers.Marshal(b, m, deterministic)
}
func (dst *Peers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peers.Merge(dst, src)
}
func (m *Peers) XXX_Size() int {
	return xxx_messageInfo_Peers.Size(m)
}
func (m *Peers) XXX_DiscardUnknown() {
	xxx_messageInfo_Peers.DiscardUnknown(m)
}

var xxx_messageInfo_Peers proto.InternalMessageInfo

func (m *Peers) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

// PingPong is the valid message of ping.
type PingPong struct {
	Message              PingPong_Message `protobuf:"varint,1,opt,name=message,proto3,enum=p2p.PingPong_Message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PingPong) Reset()         { *m = PingPong{} }
func (m *PingPong) String() string { return proto.CompactTextString(m) }
func (*PingPong) ProtoMessage()    {}
func (*PingPong) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_da16c222cfc6579f, []int{2}
}
func (m *PingPong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingPong.Unmarshal(m, b)
}
func (m *PingPong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingPong.Marshal(b, m, deterministic)
}
func (dst *PingPong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingPong.Merge(dst, src)
}
func (m *PingPong) XXX_Size() int {
	return xxx_messageInfo_PingPong.Size(m)
}
func (m *PingPong) XXX_DiscardUnknown() {
	xxx_messageInfo_PingPong.DiscardUnknown(m)
}

var xxx_messageInfo_PingPong proto.InternalMessageInfo

func (m *PingPong) GetMessage() PingPong_Message {
	if m != nil {
		return m.Message
	}
	return PingPong_PING
}

func init() {
	proto.RegisterType((*Peer)(nil), "p2p.Peer")
	proto.RegisterType((*Peers)(nil), "p2p.Peers")
	proto.RegisterType((*PingPong)(nil), "p2p.PingPong")
	proto.RegisterEnum("p2p.PingPong_Message", PingPong_Message_name, PingPong_Message_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeServiceClient interface {
	Ping(ctx context.Context, in *PingPong, opts ...grpc.CallOption) (*PingPong, error)
	GetPeers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Peers, error)
}

type nodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewNodeServiceClient(cc *grpc.ClientConn) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) Ping(ctx context.Context, in *PingPong, opts ...grpc.CallOption) (*PingPong, error) {
	out := new(PingPong)
	err := c.cc.Invoke(ctx, "/p2p.NodeService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetPeers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Peers, error) {
	out := new(Peers)
	err := c.cc.Invoke(ctx, "/p2p.NodeService/GetPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
type NodeServiceServer interface {
	Ping(context.Context, *PingPong) (*PingPong, error)
	GetPeers(context.Context, *empty.Empty) (*Peers, error)
}

func RegisterNodeServiceServer(s *grpc.Server, srv NodeServiceServer) {
	s.RegisterService(&_NodeService_serviceDesc, srv)
}

func _NodeService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingPong)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2p.NodeService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Ping(ctx, req.(*PingPong))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2p.NodeService/GetPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetPeers(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "p2p.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _NodeService_Ping_Handler,
		},
		{
			MethodName: "GetPeers",
			Handler:    _NodeService_GetPeers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "p2p.proto",
}

func init() { proto.RegisterFile("p2p.proto", fileDescriptor_p2p_da16c222cfc6579f) }

var fileDescriptor_p2p_da16c222cfc6579f = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x5d, 0x9b, 0xd8, 0x64, 0x8a, 0x52, 0x06, 0x94, 0x10, 0x11, 0x43, 0xf0, 0x90, 0xd3,
	0x06, 0xd6, 0xdf, 0x20, 0xc1, 0x83, 0x31, 0xc4, 0x9b, 0x37, 0x5b, 0xc7, 0xb5, 0x60, 0xbb, 0x43,
	0x76, 0x15, 0xfc, 0xf7, 0xb2, 0xbb, 0xad, 0xd0, 0xdb, 0xdb, 0xf7, 0x3e, 0x66, 0x76, 0x1e, 0xe4,
	0xac, 0x58, 0xf2, 0x64, 0x9c, 0xc1, 0x19, 0x2b, 0x2e, 0xaf, 0xb5, 0x31, 0xfa, 0x8b, 0xda, 0x60,
	0xad, 0xbe, 0x3f, 0x5a, 0xda, 0xb2, 0xfb, 0x8d, 0x44, 0x2d, 0x21, 0x19, 0x88, 0x26, 0x44, 0x48,
	0x3e, 0x8d, 0x75, 0x85, 0xa8, 0x44, 0x93, 0x8f, 0x41, 0x7b, 0x8f, 0xcd, 0xe4, 0x8a, 0xd3, 0x4a,
	0x34, 0xe9, 0x18, 0x74, 0xdd, 0x40, 0xea, 0x79, 0x8b, 0xb7, 0x90, 0xb2, 0x17, 0x85, 0xa8, 0x66,
	0xcd, 0x42, 0xe5, 0xd2, 0x6f, 0xf5, 0xd1, 0x18, 0xfd, 0xfa, 0x15, 0xb2, 0x61, 0xb3, 0xd3, 0x83,
	0xd9, 0x69, 0x6c, 0x61, 0xbe, 0x25, 0x6b, 0xdf, 0x34, 0x85, 0x05, 0x17, 0xea, 0x32, 0xe2, 0xfb,
	0x5c, 0x3e, 0xc5, 0x70, 0x3c, 0x50, 0xf5, 0x0d, 0xcc, 0xf7, 0x1e, 0x66, 0x90, 0x0c, 0x8f, 0x7d,
	0xb7, 0x3c, 0x09, 0xea, 0xb9, 0xef, 0x96, 0x42, 0xad, 0x61, 0xd1, 0x9b, 0x77, 0x7a, 0xa1, 0xe9,
	0x67, 0xb3, 0x26, 0xbc, 0x83, 0xc4, 0x8f, 0xc2, 0xf3, 0xa3, 0xa9, 0xe5, 0xf1, 0x13, 0x25, 0x64,
	0x1d, 0xb9, 0xf8, 0xfb, 0x2b, 0x19, 0x4b, 0x91, 0x87, 0x52, 0xe4, 0x83, 0x2f, 0xa5, 0x84, 0xff,
	0x33, 0xec, 0xea, 0x2c, 0x64, 0xf7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x10, 0x85, 0x24,
	0x50, 0x01, 0x00, 0x00,
}
