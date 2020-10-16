// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package experimentationv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ExperimentsAPIClient is the client API for ExperimentsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExperimentsAPIClient interface {
	CreateExperiment(ctx context.Context, in *CreateExperimentRequest, opts ...grpc.CallOption) (*CreateExperimentResponse, error)
	CancelExperimentRun(ctx context.Context, in *CancelExperimentRunRequest, opts ...grpc.CallOption) (*CancelExperimentRunResponse, error)
	GetExperiments(ctx context.Context, in *GetExperimentsRequest, opts ...grpc.CallOption) (*GetExperimentsResponse, error)
	GetListView(ctx context.Context, in *GetListViewRequest, opts ...grpc.CallOption) (*GetListViewResponse, error)
	GetExperimentRunDetails(ctx context.Context, in *GetExperimentRunDetailsRequest, opts ...grpc.CallOption) (*GetExperimentRunDetailsResponse, error)
}

type experimentsAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewExperimentsAPIClient(cc grpc.ClientConnInterface) ExperimentsAPIClient {
	return &experimentsAPIClient{cc}
}

func (c *experimentsAPIClient) CreateExperiment(ctx context.Context, in *CreateExperimentRequest, opts ...grpc.CallOption) (*CreateExperimentResponse, error) {
	out := new(CreateExperimentResponse)
	err := c.cc.Invoke(ctx, "/clutch.chaos.experimentation.v1.ExperimentsAPI/CreateExperiment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentsAPIClient) CancelExperimentRun(ctx context.Context, in *CancelExperimentRunRequest, opts ...grpc.CallOption) (*CancelExperimentRunResponse, error) {
	out := new(CancelExperimentRunResponse)
	err := c.cc.Invoke(ctx, "/clutch.chaos.experimentation.v1.ExperimentsAPI/CancelExperimentRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentsAPIClient) GetExperiments(ctx context.Context, in *GetExperimentsRequest, opts ...grpc.CallOption) (*GetExperimentsResponse, error) {
	out := new(GetExperimentsResponse)
	err := c.cc.Invoke(ctx, "/clutch.chaos.experimentation.v1.ExperimentsAPI/GetExperiments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentsAPIClient) GetListView(ctx context.Context, in *GetListViewRequest, opts ...grpc.CallOption) (*GetListViewResponse, error) {
	out := new(GetListViewResponse)
	err := c.cc.Invoke(ctx, "/clutch.chaos.experimentation.v1.ExperimentsAPI/GetListView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentsAPIClient) GetExperimentRunDetails(ctx context.Context, in *GetExperimentRunDetailsRequest, opts ...grpc.CallOption) (*GetExperimentRunDetailsResponse, error) {
	out := new(GetExperimentRunDetailsResponse)
	err := c.cc.Invoke(ctx, "/clutch.chaos.experimentation.v1.ExperimentsAPI/GetExperimentRunDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExperimentsAPIServer is the server API for ExperimentsAPI service.
// All implementations must embed UnimplementedExperimentsAPIServer
// for forward compatibility
type ExperimentsAPIServer interface {
	CreateExperiment(context.Context, *CreateExperimentRequest) (*CreateExperimentResponse, error)
	CancelExperimentRun(context.Context, *CancelExperimentRunRequest) (*CancelExperimentRunResponse, error)
	GetExperiments(context.Context, *GetExperimentsRequest) (*GetExperimentsResponse, error)
	GetListView(context.Context, *GetListViewRequest) (*GetListViewResponse, error)
	GetExperimentRunDetails(context.Context, *GetExperimentRunDetailsRequest) (*GetExperimentRunDetailsResponse, error)
	mustEmbedUnimplementedExperimentsAPIServer()
}

// UnimplementedExperimentsAPIServer must be embedded to have forward compatible implementations.
type UnimplementedExperimentsAPIServer struct {
}

func (UnimplementedExperimentsAPIServer) CreateExperiment(context.Context, *CreateExperimentRequest) (*CreateExperimentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExperiment not implemented")
}
func (UnimplementedExperimentsAPIServer) CancelExperimentRun(context.Context, *CancelExperimentRunRequest) (*CancelExperimentRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelExperimentRun not implemented")
}
func (UnimplementedExperimentsAPIServer) GetExperiments(context.Context, *GetExperimentsRequest) (*GetExperimentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExperiments not implemented")
}
func (UnimplementedExperimentsAPIServer) GetListView(context.Context, *GetListViewRequest) (*GetListViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListView not implemented")
}
func (UnimplementedExperimentsAPIServer) GetExperimentRunDetails(context.Context, *GetExperimentRunDetailsRequest) (*GetExperimentRunDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExperimentRunDetails not implemented")
}
func (UnimplementedExperimentsAPIServer) mustEmbedUnimplementedExperimentsAPIServer() {}

// UnsafeExperimentsAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExperimentsAPIServer will
// result in compilation errors.
type UnsafeExperimentsAPIServer interface {
	mustEmbedUnimplementedExperimentsAPIServer()
}

func RegisterExperimentsAPIServer(s *grpc.Server, srv ExperimentsAPIServer) {
	s.RegisterService(&_ExperimentsAPI_serviceDesc, srv)
}

func _ExperimentsAPI_CreateExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExperimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentsAPIServer).CreateExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.chaos.experimentation.v1.ExperimentsAPI/CreateExperiment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentsAPIServer).CreateExperiment(ctx, req.(*CreateExperimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentsAPI_CancelExperimentRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelExperimentRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentsAPIServer).CancelExperimentRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.chaos.experimentation.v1.ExperimentsAPI/CancelExperimentRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentsAPIServer).CancelExperimentRun(ctx, req.(*CancelExperimentRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentsAPI_GetExperiments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExperimentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentsAPIServer).GetExperiments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.chaos.experimentation.v1.ExperimentsAPI/GetExperiments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentsAPIServer).GetExperiments(ctx, req.(*GetExperimentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentsAPI_GetListView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentsAPIServer).GetListView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.chaos.experimentation.v1.ExperimentsAPI/GetListView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentsAPIServer).GetListView(ctx, req.(*GetListViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentsAPI_GetExperimentRunDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExperimentRunDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentsAPIServer).GetExperimentRunDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.chaos.experimentation.v1.ExperimentsAPI/GetExperimentRunDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentsAPIServer).GetExperimentRunDetails(ctx, req.(*GetExperimentRunDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExperimentsAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.chaos.experimentation.v1.ExperimentsAPI",
	HandlerType: (*ExperimentsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExperiment",
			Handler:    _ExperimentsAPI_CreateExperiment_Handler,
		},
		{
			MethodName: "CancelExperimentRun",
			Handler:    _ExperimentsAPI_CancelExperimentRun_Handler,
		},
		{
			MethodName: "GetExperiments",
			Handler:    _ExperimentsAPI_GetExperiments_Handler,
		},
		{
			MethodName: "GetListView",
			Handler:    _ExperimentsAPI_GetListView_Handler,
		},
		{
			MethodName: "GetExperimentRunDetails",
			Handler:    _ExperimentsAPI_GetExperimentRunDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chaos/experimentation/v1/experimentation.proto",
}
