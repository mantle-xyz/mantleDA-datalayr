// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: common/interfaces/interfaceRetrieverServer/server.proto

package interfaceRetrieverServer

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

// DataRetrievalClient is the client API for DataRetrieval service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataRetrievalClient interface {
	RetrieveFramesAndData(ctx context.Context, in *FramesAndDataRequest, opts ...grpc.CallOption) (*FramesAndDataReply, error)
}

type dataRetrievalClient struct {
	cc grpc.ClientConnInterface
}

func NewDataRetrievalClient(cc grpc.ClientConnInterface) DataRetrievalClient {
	return &dataRetrievalClient{cc}
}

func (c *dataRetrievalClient) RetrieveFramesAndData(ctx context.Context, in *FramesAndDataRequest, opts ...grpc.CallOption) (*FramesAndDataReply, error) {
	out := new(FramesAndDataReply)
	err := c.cc.Invoke(ctx, "/interfaceRetrieverServer.DataRetrieval/RetrieveFramesAndData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataRetrievalServer is the server API for DataRetrieval service.
// All implementations must embed UnimplementedDataRetrievalServer
// for forward compatibility
type DataRetrievalServer interface {
	RetrieveFramesAndData(context.Context, *FramesAndDataRequest) (*FramesAndDataReply, error)
	mustEmbedUnimplementedDataRetrievalServer()
}

// UnimplementedDataRetrievalServer must be embedded to have forward compatible implementations.
type UnimplementedDataRetrievalServer struct {
}

func (UnimplementedDataRetrievalServer) RetrieveFramesAndData(context.Context, *FramesAndDataRequest) (*FramesAndDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveFramesAndData not implemented")
}
func (UnimplementedDataRetrievalServer) mustEmbedUnimplementedDataRetrievalServer() {}

// UnsafeDataRetrievalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataRetrievalServer will
// result in compilation errors.
type UnsafeDataRetrievalServer interface {
	mustEmbedUnimplementedDataRetrievalServer()
}

func RegisterDataRetrievalServer(s grpc.ServiceRegistrar, srv DataRetrievalServer) {
	s.RegisterService(&DataRetrieval_ServiceDesc, srv)
}

func _DataRetrieval_RetrieveFramesAndData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FramesAndDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataRetrievalServer).RetrieveFramesAndData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interfaceRetrieverServer.DataRetrieval/RetrieveFramesAndData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataRetrievalServer).RetrieveFramesAndData(ctx, req.(*FramesAndDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataRetrieval_ServiceDesc is the grpc.ServiceDesc for DataRetrieval service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataRetrieval_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interfaceRetrieverServer.DataRetrieval",
	HandlerType: (*DataRetrievalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveFramesAndData",
			Handler:    _DataRetrieval_RetrieveFramesAndData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/interfaces/interfaceRetrieverServer/server.proto",
}
