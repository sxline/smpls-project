// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WriteDataServiceClient is the client API for WriteDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WriteDataServiceClient interface {
	Write(ctx context.Context, in *Data, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type writeDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWriteDataServiceClient(cc grpc.ClientConnInterface) WriteDataServiceClient {
	return &writeDataServiceClient{cc}
}

func (c *writeDataServiceClient) Write(ctx context.Context, in *Data, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/WriteDataService/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriteDataServiceServer is the server API for WriteDataService service.
// All implementations should embed UnimplementedWriteDataServiceServer
// for forward compatibility
type WriteDataServiceServer interface {
	Write(context.Context, *Data) (*emptypb.Empty, error)
}

// UnimplementedWriteDataServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWriteDataServiceServer struct {
}

func (UnimplementedWriteDataServiceServer) Write(context.Context, *Data) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}

// UnsafeWriteDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WriteDataServiceServer will
// result in compilation errors.
type UnsafeWriteDataServiceServer interface {
	mustEmbedUnimplementedWriteDataServiceServer()
}

func RegisterWriteDataServiceServer(s grpc.ServiceRegistrar, srv WriteDataServiceServer) {
	s.RegisterService(&WriteDataService_ServiceDesc, srv)
}

func _WriteDataService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriteDataServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WriteDataService/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriteDataServiceServer).Write(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

// WriteDataService_ServiceDesc is the grpc.ServiceDesc for WriteDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WriteDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WriteDataService",
	HandlerType: (*WriteDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _WriteDataService_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
