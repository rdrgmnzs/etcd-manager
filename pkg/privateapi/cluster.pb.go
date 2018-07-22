// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/privateapi/cluster.proto

/*
Package privateapi is a generated protocol buffer package.

It is generated from these files:
	pkg/privateapi/cluster.proto

It has these top-level messages:
	PingRequest
	PingResponse
	PeerInfo
	View
	LeaderNotificationRequest
	LeaderNotificationResponse
*/
package privateapi

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

type PingRequest struct {
	Info *PeerInfo `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingRequest) GetInfo() *PeerInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type PingResponse struct {
	Info *PeerInfo `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingResponse) GetInfo() *PeerInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type PeerInfo struct {
	Id        string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Endpoints []string `protobuf:"bytes,2,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *PeerInfo) Reset()                    { *m = PeerInfo{} }
func (m *PeerInfo) String() string            { return proto.CompactTextString(m) }
func (*PeerInfo) ProtoMessage()               {}
func (*PeerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PeerInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PeerInfo) GetEndpoints() []string {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type View struct {
	Leader          *PeerInfo   `protobuf:"bytes,1,opt,name=leader" json:"leader,omitempty"`
	LeadershipToken string      `protobuf:"bytes,2,opt,name=leadership_token,json=leadershipToken" json:"leadership_token,omitempty"`
	Healthy         []*PeerInfo `protobuf:"bytes,3,rep,name=healthy" json:"healthy,omitempty"`
}

func (m *View) Reset()                    { *m = View{} }
func (m *View) String() string            { return proto.CompactTextString(m) }
func (*View) ProtoMessage()               {}
func (*View) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *View) GetLeader() *PeerInfo {
	if m != nil {
		return m.Leader
	}
	return nil
}

func (m *View) GetLeadershipToken() string {
	if m != nil {
		return m.LeadershipToken
	}
	return ""
}

func (m *View) GetHealthy() []*PeerInfo {
	if m != nil {
		return m.Healthy
	}
	return nil
}

type LeaderNotificationRequest struct {
	View *View `protobuf:"bytes,1,opt,name=view" json:"view,omitempty"`
}

func (m *LeaderNotificationRequest) Reset()                    { *m = LeaderNotificationRequest{} }
func (m *LeaderNotificationRequest) String() string            { return proto.CompactTextString(m) }
func (*LeaderNotificationRequest) ProtoMessage()               {}
func (*LeaderNotificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LeaderNotificationRequest) GetView() *View {
	if m != nil {
		return m.View
	}
	return nil
}

type LeaderNotificationResponse struct {
	// True if this node acknowledges the new leader
	Accepted bool `protobuf:"varint,1,opt,name=accepted" json:"accepted,omitempty"`
	// If the node has a different (bigger) view, it rejects the leadership bid and sends the view
	View *View `protobuf:"bytes,2,opt,name=view" json:"view,omitempty"`
}

func (m *LeaderNotificationResponse) Reset()                    { *m = LeaderNotificationResponse{} }
func (m *LeaderNotificationResponse) String() string            { return proto.CompactTextString(m) }
func (*LeaderNotificationResponse) ProtoMessage()               {}
func (*LeaderNotificationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LeaderNotificationResponse) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *LeaderNotificationResponse) GetView() *View {
	if m != nil {
		return m.View
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "privateapi.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "privateapi.PingResponse")
	proto.RegisterType((*PeerInfo)(nil), "privateapi.PeerInfo")
	proto.RegisterType((*View)(nil), "privateapi.View")
	proto.RegisterType((*LeaderNotificationRequest)(nil), "privateapi.LeaderNotificationRequest")
	proto.RegisterType((*LeaderNotificationResponse)(nil), "privateapi.LeaderNotificationResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ClusterService service

type ClusterServiceClient interface {
	// Ping just pings another node, part of the discovery protocol
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// LeaderNotification is sent by a node that (thinks it) is the leader
	LeaderNotification(ctx context.Context, in *LeaderNotificationRequest, opts ...grpc.CallOption) (*LeaderNotificationResponse, error)
}

type clusterServiceClient struct {
	cc *grpc.ClientConn
}

func NewClusterServiceClient(cc *grpc.ClientConn) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/privateapi.ClusterService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) LeaderNotification(ctx context.Context, in *LeaderNotificationRequest, opts ...grpc.CallOption) (*LeaderNotificationResponse, error) {
	out := new(LeaderNotificationResponse)
	err := grpc.Invoke(ctx, "/privateapi.ClusterService/LeaderNotification", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClusterService service

type ClusterServiceServer interface {
	// Ping just pings another node, part of the discovery protocol
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// LeaderNotification is sent by a node that (thinks it) is the leader
	LeaderNotification(context.Context, *LeaderNotificationRequest) (*LeaderNotificationResponse, error)
}

func RegisterClusterServiceServer(s *grpc.Server, srv ClusterServiceServer) {
	s.RegisterService(&_ClusterService_serviceDesc, srv)
}

func _ClusterService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/privateapi.ClusterService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_LeaderNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).LeaderNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/privateapi.ClusterService/LeaderNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).LeaderNotification(ctx, req.(*LeaderNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "privateapi.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ClusterService_Ping_Handler,
		},
		{
			MethodName: "LeaderNotification",
			Handler:    _ClusterService_LeaderNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/privateapi/cluster.proto",
}

func init() { proto.RegisterFile("pkg/privateapi/cluster.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x51, 0x4f, 0xe2, 0x40,
	0x10, 0x4e, 0x4b, 0xc3, 0xc1, 0x70, 0xe1, 0xc8, 0xe6, 0x92, 0xeb, 0x35, 0x3c, 0x90, 0xde, 0x61,
	0x30, 0xd1, 0x36, 0xc1, 0x07, 0xf5, 0x51, 0x7d, 0x32, 0x31, 0x86, 0x54, 0xe3, 0x83, 0x0f, 0x9a,
	0xb5, 0x1d, 0x60, 0x02, 0xee, 0xae, 0xed, 0x02, 0xf1, 0x57, 0xf8, 0x57, 0xfc, 0x89, 0x86, 0xb6,
	0x50, 0x08, 0x82, 0xf1, 0x6d, 0xf7, 0x9b, 0x6f, 0xbe, 0x99, 0xf9, 0x66, 0xa0, 0xa9, 0x46, 0x03,
	0x5f, 0xc5, 0x34, 0xe5, 0x1a, 0xb9, 0x22, 0x3f, 0x1c, 0x4f, 0x12, 0x8d, 0xb1, 0xa7, 0x62, 0xa9,
	0x25, 0x83, 0x22, 0xe2, 0x1e, 0x43, 0xad, 0x47, 0x62, 0x10, 0xe0, 0xcb, 0x04, 0x13, 0xcd, 0x3a,
	0x60, 0x91, 0xe8, 0x4b, 0xdb, 0x68, 0x19, 0x9d, 0x5a, 0xf7, 0xb7, 0x57, 0x30, 0xbd, 0x1e, 0x62,
	0x7c, 0x29, 0xfa, 0x32, 0x48, 0x19, 0xee, 0x09, 0xfc, 0xcc, 0x12, 0x13, 0x25, 0x45, 0x82, 0xdf,
	0xca, 0xac, 0x2c, 0x10, 0x56, 0x07, 0x93, 0xa2, 0x34, 0xa7, 0x1a, 0x98, 0x14, 0xb1, 0x26, 0x54,
	0x51, 0x44, 0x4a, 0x92, 0xd0, 0x89, 0x6d, 0xb6, 0x4a, 0x9d, 0x6a, 0x50, 0x00, 0xee, 0x9b, 0x01,
	0xd6, 0x1d, 0xe1, 0x8c, 0x1d, 0x40, 0x79, 0x8c, 0x3c, 0xc2, 0x78, 0x67, 0xb9, 0x9c, 0xc3, 0xf6,
	0xa1, 0x91, 0xbd, 0x92, 0x21, 0xa9, 0x47, 0x2d, 0x47, 0x28, 0x6c, 0x33, 0x2d, 0xf9, 0xab, 0xc0,
	0x6f, 0xe7, 0x30, 0xf3, 0xe0, 0xc7, 0x10, 0xf9, 0x58, 0x0f, 0x5f, 0xed, 0x52, 0xab, 0xb4, 0x55,
	0x79, 0x41, 0x72, 0xcf, 0xe0, 0xef, 0x55, 0x2a, 0x71, 0x2d, 0x35, 0xf5, 0x29, 0xe4, 0x9a, 0xa4,
	0x58, 0x98, 0xf9, 0x1f, 0xac, 0x29, 0xe1, 0x2c, 0xef, 0xb1, 0xb1, 0xaa, 0x34, 0x9f, 0x22, 0x48,
	0xa3, 0xee, 0x03, 0x38, 0x9f, 0x49, 0xe4, 0xb6, 0x3a, 0x50, 0xe1, 0x61, 0x88, 0x4a, 0x63, 0x66,
	0x53, 0x25, 0x58, 0xfe, 0x97, 0xfa, 0xe6, 0x2e, 0xfd, 0xee, 0xbb, 0x01, 0xf5, 0x8b, 0x6c, 0xff,
	0x37, 0x18, 0x4f, 0x29, 0x44, 0x76, 0x0a, 0xd6, 0x7c, 0x77, 0xec, 0xcf, 0xda, 0x70, 0xc5, 0x19,
	0x38, 0xf6, 0x66, 0x20, 0xef, 0x27, 0x04, 0xb6, 0xd9, 0x2d, 0x6b, 0xaf, 0xf2, 0xb7, 0x1a, 0xe2,
	0xec, 0x7d, 0x45, 0xcb, 0x8a, 0x9c, 0xb7, 0xef, 0xff, 0x8d, 0xa4, 0x42, 0x8f, 0xa4, 0x8f, 0x3a,
	0x8c, 0x0e, 0x9f, 0xb9, 0xe0, 0x03, 0x8c, 0xfd, 0xf5, 0xab, 0x7e, 0x2a, 0xa7, 0xe7, 0x7c, 0xf4,
	0x11, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x45, 0xb3, 0x7c, 0xee, 0x02, 0x00, 0x00,
}
