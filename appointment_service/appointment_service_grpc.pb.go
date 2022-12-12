// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: appointment_service/appointment_service.proto

package appointment_service

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

// AppointmentServiceClient is the client API for AppointmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppointmentServiceClient interface {
	// 发布活动
	PublishActivity(ctx context.Context, in *PublishActivityReq, opts ...grpc.CallOption) (*PublishActivityRsp, error)
	// 获取用户发布的活动
	GetPersonActivityList(ctx context.Context, in *GetPersonActivityListReq, opts ...grpc.CallOption) (*GetPersonActivityListRsp, error)
	// 获取活动详情
	GetPersonActivity(ctx context.Context, in *GetPersonActivityReq, opts ...grpc.CallOption) (*GetPersonActivityRsp, error)
	// 预约（注意幂等）同一个活动 同一个用户
	Reserve(ctx context.Context, in *ReserveReq, opts ...grpc.CallOption) (*ReserveRsp, error)
	// 获取用户预约列表（应该需要分页）
	GetPersonReserveList(ctx context.Context, in *GetPersonReserveListReq, opts ...grpc.CallOption) (*GetPersonReserveListRsp, error)
	// 获取用户预约详情
	GetPersonReserveDetail(ctx context.Context, in *GetPersonReserveDetailReq, opts ...grpc.CallOption) (*GetPersonReserveDetailRsp, error)
	// 活动曝光上报
	ActivityExposureReport(ctx context.Context, in *ActivityExposureReportReq, opts ...grpc.CallOption) (*ActivityExposureReportRsp, error)
}

type appointmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppointmentServiceClient(cc grpc.ClientConnInterface) AppointmentServiceClient {
	return &appointmentServiceClient{cc}
}

func (c *appointmentServiceClient) PublishActivity(ctx context.Context, in *PublishActivityReq, opts ...grpc.CallOption) (*PublishActivityRsp, error) {
	out := new(PublishActivityRsp)
	err := c.cc.Invoke(ctx, "/appointment_service.AppointmentService/PublishActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) GetPersonActivityList(ctx context.Context, in *GetPersonActivityListReq, opts ...grpc.CallOption) (*GetPersonActivityListRsp, error) {
	out := new(GetPersonActivityListRsp)
	err := c.cc.Invoke(ctx, "/appointment_service.AppointmentService/GetPersonActivityList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) GetPersonActivity(ctx context.Context, in *GetPersonActivityReq, opts ...grpc.CallOption) (*GetPersonActivityRsp, error) {
	out := new(GetPersonActivityRsp)
	err := c.cc.Invoke(ctx, "/appointment_service.AppointmentService/GetPersonActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) Reserve(ctx context.Context, in *ReserveReq, opts ...grpc.CallOption) (*ReserveRsp, error) {
	out := new(ReserveRsp)
	err := c.cc.Invoke(ctx, "/appointment_service.AppointmentService/Reserve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) GetPersonReserveList(ctx context.Context, in *GetPersonReserveListReq, opts ...grpc.CallOption) (*GetPersonReserveListRsp, error) {
	out := new(GetPersonReserveListRsp)
	err := c.cc.Invoke(ctx, "/appointment_service.AppointmentService/GetPersonReserveList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) GetPersonReserveDetail(ctx context.Context, in *GetPersonReserveDetailReq, opts ...grpc.CallOption) (*GetPersonReserveDetailRsp, error) {
	out := new(GetPersonReserveDetailRsp)
	err := c.cc.Invoke(ctx, "/appointment_service.AppointmentService/GetPersonReserveDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) ActivityExposureReport(ctx context.Context, in *ActivityExposureReportReq, opts ...grpc.CallOption) (*ActivityExposureReportRsp, error) {
	out := new(ActivityExposureReportRsp)
	err := c.cc.Invoke(ctx, "/appointment_service.AppointmentService/ActivityExposureReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppointmentServiceServer is the server API for AppointmentService service.
// All implementations must embed UnimplementedAppointmentServiceServer
// for forward compatibility
type AppointmentServiceServer interface {
	// 发布活动
	PublishActivity(context.Context, *PublishActivityReq) (*PublishActivityRsp, error)
	// 获取用户发布的活动
	GetPersonActivityList(context.Context, *GetPersonActivityListReq) (*GetPersonActivityListRsp, error)
	// 获取活动详情
	GetPersonActivity(context.Context, *GetPersonActivityReq) (*GetPersonActivityRsp, error)
	// 预约（注意幂等）同一个活动 同一个用户
	Reserve(context.Context, *ReserveReq) (*ReserveRsp, error)
	// 获取用户预约列表（应该需要分页）
	GetPersonReserveList(context.Context, *GetPersonReserveListReq) (*GetPersonReserveListRsp, error)
	// 获取用户预约详情
	GetPersonReserveDetail(context.Context, *GetPersonReserveDetailReq) (*GetPersonReserveDetailRsp, error)
	// 活动曝光上报
	ActivityExposureReport(context.Context, *ActivityExposureReportReq) (*ActivityExposureReportRsp, error)
	mustEmbedUnimplementedAppointmentServiceServer()
}

// UnimplementedAppointmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppointmentServiceServer struct {
}

func (UnimplementedAppointmentServiceServer) PublishActivity(context.Context, *PublishActivityReq) (*PublishActivityRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishActivity not implemented")
}
func (UnimplementedAppointmentServiceServer) GetPersonActivityList(context.Context, *GetPersonActivityListReq) (*GetPersonActivityListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonActivityList not implemented")
}
func (UnimplementedAppointmentServiceServer) GetPersonActivity(context.Context, *GetPersonActivityReq) (*GetPersonActivityRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonActivity not implemented")
}
func (UnimplementedAppointmentServiceServer) Reserve(context.Context, *ReserveReq) (*ReserveRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reserve not implemented")
}
func (UnimplementedAppointmentServiceServer) GetPersonReserveList(context.Context, *GetPersonReserveListReq) (*GetPersonReserveListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonReserveList not implemented")
}
func (UnimplementedAppointmentServiceServer) GetPersonReserveDetail(context.Context, *GetPersonReserveDetailReq) (*GetPersonReserveDetailRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonReserveDetail not implemented")
}
func (UnimplementedAppointmentServiceServer) ActivityExposureReport(context.Context, *ActivityExposureReportReq) (*ActivityExposureReportRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivityExposureReport not implemented")
}
func (UnimplementedAppointmentServiceServer) mustEmbedUnimplementedAppointmentServiceServer() {}

// UnsafeAppointmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppointmentServiceServer will
// result in compilation errors.
type UnsafeAppointmentServiceServer interface {
	mustEmbedUnimplementedAppointmentServiceServer()
}

func RegisterAppointmentServiceServer(s grpc.ServiceRegistrar, srv AppointmentServiceServer) {
	s.RegisterService(&AppointmentService_ServiceDesc, srv)
}

func _AppointmentService_PublishActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).PublishActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment_service.AppointmentService/PublishActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).PublishActivity(ctx, req.(*PublishActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_GetPersonActivityList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonActivityListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).GetPersonActivityList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment_service.AppointmentService/GetPersonActivityList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).GetPersonActivityList(ctx, req.(*GetPersonActivityListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_GetPersonActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).GetPersonActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment_service.AppointmentService/GetPersonActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).GetPersonActivity(ctx, req.(*GetPersonActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_Reserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).Reserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment_service.AppointmentService/Reserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).Reserve(ctx, req.(*ReserveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_GetPersonReserveList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonReserveListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).GetPersonReserveList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment_service.AppointmentService/GetPersonReserveList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).GetPersonReserveList(ctx, req.(*GetPersonReserveListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_GetPersonReserveDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonReserveDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).GetPersonReserveDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment_service.AppointmentService/GetPersonReserveDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).GetPersonReserveDetail(ctx, req.(*GetPersonReserveDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_ActivityExposureReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityExposureReportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).ActivityExposureReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment_service.AppointmentService/ActivityExposureReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).ActivityExposureReport(ctx, req.(*ActivityExposureReportReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AppointmentService_ServiceDesc is the grpc.ServiceDesc for AppointmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppointmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appointment_service.AppointmentService",
	HandlerType: (*AppointmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishActivity",
			Handler:    _AppointmentService_PublishActivity_Handler,
		},
		{
			MethodName: "GetPersonActivityList",
			Handler:    _AppointmentService_GetPersonActivityList_Handler,
		},
		{
			MethodName: "GetPersonActivity",
			Handler:    _AppointmentService_GetPersonActivity_Handler,
		},
		{
			MethodName: "Reserve",
			Handler:    _AppointmentService_Reserve_Handler,
		},
		{
			MethodName: "GetPersonReserveList",
			Handler:    _AppointmentService_GetPersonReserveList_Handler,
		},
		{
			MethodName: "GetPersonReserveDetail",
			Handler:    _AppointmentService_GetPersonReserveDetail_Handler,
		},
		{
			MethodName: "ActivityExposureReport",
			Handler:    _AppointmentService_ActivityExposureReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appointment_service/appointment_service.proto",
}