// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package assignment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ApikeyServiceClient is the client API for ApikeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApikeyServiceClient interface {
	// AddApikey adds a new apikey to the user
	AddApikey(ctx context.Context, in *AddApikeyRequest, opts ...grpc.CallOption) (*AddApikeyResponse, error)
	// ListApikeys lists all currently added apikeys for the specified user
	ListApikeys(ctx context.Context, in *ListApikeysRequest, opts ...grpc.CallOption) (*ListApikeysResponse, error)
	//GetApikey requests full plaintext access to the specified apikey
	GetApikey(ctx context.Context, in *GetApikeyRequest, opts ...grpc.CallOption) (*GetApikeyResponse, error)
}

type apikeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApikeyServiceClient(cc grpc.ClientConnInterface) ApikeyServiceClient {
	return &apikeyServiceClient{cc}
}

func (c *apikeyServiceClient) AddApikey(ctx context.Context, in *AddApikeyRequest, opts ...grpc.CallOption) (*AddApikeyResponse, error) {
	out := new(AddApikeyResponse)
	err := c.cc.Invoke(ctx, "/blocksize.assignment.ApikeyService/AddApikey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apikeyServiceClient) ListApikeys(ctx context.Context, in *ListApikeysRequest, opts ...grpc.CallOption) (*ListApikeysResponse, error) {
	out := new(ListApikeysResponse)
	err := c.cc.Invoke(ctx, "/blocksize.assignment.ApikeyService/ListApikeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apikeyServiceClient) GetApikey(ctx context.Context, in *GetApikeyRequest, opts ...grpc.CallOption) (*GetApikeyResponse, error) {
	out := new(GetApikeyResponse)
	err := c.cc.Invoke(ctx, "/blocksize.assignment.ApikeyService/GetApikey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApikeyServiceServer is the server API for ApikeyService service.
// All implementations should embed UnimplementedApikeyServiceServer
// for forward compatibility
type ApikeyServiceServer interface {
	// AddApikey adds a new apikey to the user
	AddApikey(context.Context, *AddApikeyRequest) (*AddApikeyResponse, error)
	// ListApikeys lists all currently added apikeys for the specified user
	ListApikeys(context.Context, *ListApikeysRequest) (*ListApikeysResponse, error)
	//GetApikey requests full plaintext access to the specified apikey
	GetApikey(context.Context, *GetApikeyRequest) (*GetApikeyResponse, error)
}

// UnimplementedApikeyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedApikeyServiceServer struct {
}

func (UnimplementedApikeyServiceServer) AddApikey(context.Context, *AddApikeyRequest) (*AddApikeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddApikey not implemented")
}
func (UnimplementedApikeyServiceServer) ListApikeys(context.Context, *ListApikeysRequest) (*ListApikeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApikeys not implemented")
}
func (UnimplementedApikeyServiceServer) GetApikey(context.Context, *GetApikeyRequest) (*GetApikeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApikey not implemented")
}

// UnsafeApikeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApikeyServiceServer will
// result in compilation errors.
type UnsafeApikeyServiceServer interface {
	mustEmbedUnimplementedApikeyServiceServer()
}

func RegisterApikeyServiceServer(s grpc.ServiceRegistrar, srv ApikeyServiceServer) {
	s.RegisterService(&_ApikeyService_serviceDesc, srv)
}

func _ApikeyService_AddApikey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddApikeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApikeyServiceServer).AddApikey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blocksize.assignment.ApikeyService/AddApikey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApikeyServiceServer).AddApikey(ctx, req.(*AddApikeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApikeyService_ListApikeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApikeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApikeyServiceServer).ListApikeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blocksize.assignment.ApikeyService/ListApikeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApikeyServiceServer).ListApikeys(ctx, req.(*ListApikeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApikeyService_GetApikey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApikeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApikeyServiceServer).GetApikey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blocksize.assignment.ApikeyService/GetApikey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApikeyServiceServer).GetApikey(ctx, req.(*GetApikeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApikeyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blocksize.assignment.ApikeyService",
	HandlerType: (*ApikeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddApikey",
			Handler:    _ApikeyService_AddApikey_Handler,
		},
		{
			MethodName: "ListApikeys",
			Handler:    _ApikeyService_ListApikeys_Handler,
		},
		{
			MethodName: "GetApikey",
			Handler:    _ApikeyService_GetApikey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
