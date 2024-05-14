// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: common/interfaces/interfaceDL/disperse.proto

package interfaceDL

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DataDispersalClient is the client API for DataDispersal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataDispersalClient interface {
	EncodeStore(ctx context.Context, in *EncodeStoreRequest, opts ...grpc.CallOption) (*EncodeStoreReply, error)
	DisperseStore(ctx context.Context, in *DisperseStoreRequest, opts ...grpc.CallOption) (*DisperseStoreReply, error)
	EncodeAndDisperseStore(ctx context.Context, in *EncodeStoreRequest, opts ...grpc.CallOption) (*EncodeAndDisperseStoreReply, error)
	StoreFrame(ctx context.Context, in *StoreFrameRequest, opts ...grpc.CallOption) (*StoreFrameReply, error)
	StoreFrames(ctx context.Context, in *StoreFramesRequest, opts ...grpc.CallOption) (*StoreFrameReply, error)
	RetrieveFrame(ctx context.Context, in *RetrieveFrameRequest, opts ...grpc.CallOption) (DataDispersal_RetrieveFrameClient, error)
	ChallengeCoding(ctx context.Context, in *CodingChallengeRequest, opts ...grpc.CallOption) (*ChallengeReply, error)
	ChallengePayment(ctx context.Context, in *PaymentChallengeRequest, opts ...grpc.CallOption) (*ChallengeReply, error)
}

type dataDispersalClient struct {
	cc grpc.ClientConnInterface
}

func NewDataDispersalClient(cc grpc.ClientConnInterface) DataDispersalClient {
	return &dataDispersalClient{cc}
}

func (c *dataDispersalClient) EncodeStore(ctx context.Context, in *EncodeStoreRequest, opts ...grpc.CallOption) (*EncodeStoreReply, error) {
	out := new(EncodeStoreReply)
	err := c.cc.Invoke(ctx, "/interfaceDL.DataDispersal/EncodeStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataDispersalClient) DisperseStore(ctx context.Context, in *DisperseStoreRequest, opts ...grpc.CallOption) (*DisperseStoreReply, error) {
	out := new(DisperseStoreReply)
	err := c.cc.Invoke(ctx, "/interfaceDL.DataDispersal/DisperseStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataDispersalClient) EncodeAndDisperseStore(ctx context.Context, in *EncodeStoreRequest, opts ...grpc.CallOption) (*EncodeAndDisperseStoreReply, error) {
	out := new(EncodeAndDisperseStoreReply)
	err := c.cc.Invoke(ctx, "/interfaceDL.DataDispersal/EncodeAndDisperseStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataDispersalClient) StoreFrame(ctx context.Context, in *StoreFrameRequest, opts ...grpc.CallOption) (*StoreFrameReply, error) {
	out := new(StoreFrameReply)
	err := c.cc.Invoke(ctx, "/interfaceDL.DataDispersal/StoreFrame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataDispersalClient) StoreFrames(ctx context.Context, in *StoreFramesRequest, opts ...grpc.CallOption) (*StoreFrameReply, error) {
	out := new(StoreFrameReply)
	err := c.cc.Invoke(ctx, "/interfaceDL.DataDispersal/StoreFrames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataDispersalClient) RetrieveFrame(ctx context.Context, in *RetrieveFrameRequest, opts ...grpc.CallOption) (DataDispersal_RetrieveFrameClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataDispersal_ServiceDesc.Streams[0], "/interfaceDL.DataDispersal/RetrieveFrame", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataDispersalRetrieveFrameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataDispersal_RetrieveFrameClient interface {
	Recv() (*RetrieveFrameReply, error)
	grpc.ClientStream
}

type dataDispersalRetrieveFrameClient struct {
	grpc.ClientStream
}

func (x *dataDispersalRetrieveFrameClient) Recv() (*RetrieveFrameReply, error) {
	m := new(RetrieveFrameReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataDispersalClient) ChallengeCoding(ctx context.Context, in *CodingChallengeRequest, opts ...grpc.CallOption) (*ChallengeReply, error) {
	out := new(ChallengeReply)
	err := c.cc.Invoke(ctx, "/interfaceDL.DataDispersal/ChallengeCoding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataDispersalClient) ChallengePayment(ctx context.Context, in *PaymentChallengeRequest, opts ...grpc.CallOption) (*ChallengeReply, error) {
	out := new(ChallengeReply)
	err := c.cc.Invoke(ctx, "/interfaceDL.DataDispersal/ChallengePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataDispersalServer is the server API for DataDispersal service.
// All implementations must embed UnimplementedDataDispersalServer
// for forward compatibility
type DataDispersalServer interface {
	EncodeStore(context.Context, *EncodeStoreRequest) (*EncodeStoreReply, error)
	DisperseStore(context.Context, *DisperseStoreRequest) (*DisperseStoreReply, error)
	EncodeAndDisperseStore(context.Context, *EncodeStoreRequest) (*EncodeAndDisperseStoreReply, error)
	StoreFrame(context.Context, *StoreFrameRequest) (*StoreFrameReply, error)
	StoreFrames(context.Context, *StoreFramesRequest) (*StoreFrameReply, error)
	RetrieveFrame(*RetrieveFrameRequest, DataDispersal_RetrieveFrameServer) error
	ChallengeCoding(context.Context, *CodingChallengeRequest) (*ChallengeReply, error)
	ChallengePayment(context.Context, *PaymentChallengeRequest) (*ChallengeReply, error)
	mustEmbedUnimplementedDataDispersalServer()
}

// UnimplementedDataDispersalServer must be embedded to have forward compatible implementations.
type UnimplementedDataDispersalServer struct {
}

func (UnimplementedDataDispersalServer) EncodeStore(context.Context, *EncodeStoreRequest) (*EncodeStoreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncodeStore not implemented")
}
func (UnimplementedDataDispersalServer) DisperseStore(context.Context, *DisperseStoreRequest) (*DisperseStoreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisperseStore not implemented")
}
func (UnimplementedDataDispersalServer) EncodeAndDisperseStore(context.Context, *EncodeStoreRequest) (*EncodeAndDisperseStoreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncodeAndDisperseStore not implemented")
}
func (UnimplementedDataDispersalServer) StoreFrame(context.Context, *StoreFrameRequest) (*StoreFrameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreFrame not implemented")
}
func (UnimplementedDataDispersalServer) StoreFrames(context.Context, *StoreFramesRequest) (*StoreFrameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreFrames not implemented")
}
func (UnimplementedDataDispersalServer) RetrieveFrame(*RetrieveFrameRequest, DataDispersal_RetrieveFrameServer) error {
	return status.Errorf(codes.Unimplemented, "method RetrieveFrame not implemented")
}
func (UnimplementedDataDispersalServer) ChallengeCoding(context.Context, *CodingChallengeRequest) (*ChallengeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChallengeCoding not implemented")
}
func (UnimplementedDataDispersalServer) ChallengePayment(context.Context, *PaymentChallengeRequest) (*ChallengeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChallengePayment not implemented")
}
func (UnimplementedDataDispersalServer) mustEmbedUnimplementedDataDispersalServer() {}

// UnsafeDataDispersalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataDispersalServer will
// result in compilation errors.
type UnsafeDataDispersalServer interface {
	mustEmbedUnimplementedDataDispersalServer()
}

func RegisterDataDispersalServer(s grpc.ServiceRegistrar, srv DataDispersalServer) {
	s.RegisterService(&DataDispersal_ServiceDesc, srv)
}

func _DataDispersal_EncodeStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncodeStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataDispersalServer).EncodeStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interfaceDL.DataDispersal/EncodeStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataDispersalServer).EncodeStore(ctx, req.(*EncodeStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataDispersal_DisperseStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisperseStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataDispersalServer).DisperseStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interfaceDL.DataDispersal/DisperseStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataDispersalServer).DisperseStore(ctx, req.(*DisperseStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataDispersal_EncodeAndDisperseStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncodeStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataDispersalServer).EncodeAndDisperseStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interfaceDL.DataDispersal/EncodeAndDisperseStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataDispersalServer).EncodeAndDisperseStore(ctx, req.(*EncodeStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataDispersal_StoreFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreFrameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataDispersalServer).StoreFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interfaceDL.DataDispersal/StoreFrame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataDispersalServer).StoreFrame(ctx, req.(*StoreFrameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataDispersal_StoreFrames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreFramesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataDispersalServer).StoreFrames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interfaceDL.DataDispersal/StoreFrames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataDispersalServer).StoreFrames(ctx, req.(*StoreFramesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataDispersal_RetrieveFrame_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RetrieveFrameRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataDispersalServer).RetrieveFrame(m, &dataDispersalRetrieveFrameServer{stream})
}

type DataDispersal_RetrieveFrameServer interface {
	Send(*RetrieveFrameReply) error
	grpc.ServerStream
}

type dataDispersalRetrieveFrameServer struct {
	grpc.ServerStream
}

func (x *dataDispersalRetrieveFrameServer) Send(m *RetrieveFrameReply) error {
	return x.ServerStream.SendMsg(m)
}

func _DataDispersal_ChallengeCoding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodingChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataDispersalServer).ChallengeCoding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interfaceDL.DataDispersal/ChallengeCoding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataDispersalServer).ChallengeCoding(ctx, req.(*CodingChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataDispersal_ChallengePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataDispersalServer).ChallengePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interfaceDL.DataDispersal/ChallengePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataDispersalServer).ChallengePayment(ctx, req.(*PaymentChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataDispersal_ServiceDesc is the grpc.ServiceDesc for DataDispersal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataDispersal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interfaceDL.DataDispersal",
	HandlerType: (*DataDispersalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EncodeStore",
			Handler:    _DataDispersal_EncodeStore_Handler,
		},
		{
			MethodName: "DisperseStore",
			Handler:    _DataDispersal_DisperseStore_Handler,
		},
		{
			MethodName: "EncodeAndDisperseStore",
			Handler:    _DataDispersal_EncodeAndDisperseStore_Handler,
		},
		{
			MethodName: "StoreFrame",
			Handler:    _DataDispersal_StoreFrame_Handler,
		},
		{
			MethodName: "StoreFrames",
			Handler:    _DataDispersal_StoreFrames_Handler,
		},
		{
			MethodName: "ChallengeCoding",
			Handler:    _DataDispersal_ChallengeCoding_Handler,
		},
		{
			MethodName: "ChallengePayment",
			Handler:    _DataDispersal_ChallengePayment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RetrieveFrame",
			Handler:       _DataDispersal_RetrieveFrame_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "common/interfaces/interfaceDL/disperse.proto",
}