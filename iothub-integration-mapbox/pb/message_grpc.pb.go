// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: message.proto

package pb

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

// GeoCodeClient is the client API for GeoCode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeoCodeClient interface {
	ReverseGeo(ctx context.Context, in *ReverseGeoReq, opts ...grpc.CallOption) (*ReverseGeoResp, error)
	ForwardGeo(ctx context.Context, in *ForwardGeoReq, opts ...grpc.CallOption) (*ForwardGeoResp, error)
	Suggestions(ctx context.Context, in *SuggestApiReq, opts ...grpc.CallOption) (*SuggestApiResp, error)
	Retrieve(ctx context.Context, in *RetrieveReq, opts ...grpc.CallOption) (*RetrieveResp, error)
}

type geoCodeClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoCodeClient(cc grpc.ClientConnInterface) GeoCodeClient {
	return &geoCodeClient{cc}
}

func (c *geoCodeClient) ReverseGeo(ctx context.Context, in *ReverseGeoReq, opts ...grpc.CallOption) (*ReverseGeoResp, error) {
	out := new(ReverseGeoResp)
	err := c.cc.Invoke(ctx, "/main.GeoCode/ReverseGeo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoCodeClient) ForwardGeo(ctx context.Context, in *ForwardGeoReq, opts ...grpc.CallOption) (*ForwardGeoResp, error) {
	out := new(ForwardGeoResp)
	err := c.cc.Invoke(ctx, "/main.GeoCode/ForwardGeo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoCodeClient) Suggestions(ctx context.Context, in *SuggestApiReq, opts ...grpc.CallOption) (*SuggestApiResp, error) {
	out := new(SuggestApiResp)
	err := c.cc.Invoke(ctx, "/main.GeoCode/Suggestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoCodeClient) Retrieve(ctx context.Context, in *RetrieveReq, opts ...grpc.CallOption) (*RetrieveResp, error) {
	out := new(RetrieveResp)
	err := c.cc.Invoke(ctx, "/main.GeoCode/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeoCodeServer is the server API for GeoCode service.
// All implementations must embed UnimplementedGeoCodeServer
// for forward compatibility
type GeoCodeServer interface {
	ReverseGeo(context.Context, *ReverseGeoReq) (*ReverseGeoResp, error)
	ForwardGeo(context.Context, *ForwardGeoReq) (*ForwardGeoResp, error)
	Suggestions(context.Context, *SuggestApiReq) (*SuggestApiResp, error)
	Retrieve(context.Context, *RetrieveReq) (*RetrieveResp, error)
	mustEmbedUnimplementedGeoCodeServer()
}

// UnimplementedGeoCodeServer must be embedded to have forward compatible implementations.
type UnimplementedGeoCodeServer struct {
}

func (UnimplementedGeoCodeServer) ReverseGeo(context.Context, *ReverseGeoReq) (*ReverseGeoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReverseGeo not implemented")
}
func (UnimplementedGeoCodeServer) ForwardGeo(context.Context, *ForwardGeoReq) (*ForwardGeoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForwardGeo not implemented")
}
func (UnimplementedGeoCodeServer) Suggestions(context.Context, *SuggestApiReq) (*SuggestApiResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Suggestions not implemented")
}
func (UnimplementedGeoCodeServer) Retrieve(context.Context, *RetrieveReq) (*RetrieveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedGeoCodeServer) mustEmbedUnimplementedGeoCodeServer() {}

// UnsafeGeoCodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeoCodeServer will
// result in compilation errors.
type UnsafeGeoCodeServer interface {
	mustEmbedUnimplementedGeoCodeServer()
}

func RegisterGeoCodeServer(s grpc.ServiceRegistrar, srv GeoCodeServer) {
	s.RegisterService(&GeoCode_ServiceDesc, srv)
}

func _GeoCode_ReverseGeo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReverseGeoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoCodeServer).ReverseGeo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.GeoCode/ReverseGeo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoCodeServer).ReverseGeo(ctx, req.(*ReverseGeoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoCode_ForwardGeo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForwardGeoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoCodeServer).ForwardGeo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.GeoCode/ForwardGeo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoCodeServer).ForwardGeo(ctx, req.(*ForwardGeoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoCode_Suggestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestApiReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoCodeServer).Suggestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.GeoCode/Suggestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoCodeServer).Suggestions(ctx, req.(*SuggestApiReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoCode_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoCodeServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.GeoCode/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoCodeServer).Retrieve(ctx, req.(*RetrieveReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GeoCode_ServiceDesc is the grpc.ServiceDesc for GeoCode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeoCode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.GeoCode",
	HandlerType: (*GeoCodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReverseGeo",
			Handler:    _GeoCode_ReverseGeo_Handler,
		},
		{
			MethodName: "ForwardGeo",
			Handler:    _GeoCode_ForwardGeo_Handler,
		},
		{
			MethodName: "Suggestions",
			Handler:    _GeoCode_Suggestions_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _GeoCode_Retrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
