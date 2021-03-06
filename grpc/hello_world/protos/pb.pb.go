// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: grpc/hello_world/protos/pb.proto

package protos

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8672450aa83da438, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return m.Size()
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8672450aa83da438, []int{1}
}
func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return m.Size()
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "protos.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "protos.HelloReply")
}

func init() { proto.RegisterFile("grpc/hello_world/protos/pb.proto", fileDescriptor_8672450aa83da438) }

var fileDescriptor_8672450aa83da438 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0x1b, 0x11, 0xa7, 0xc1, 0x53, 0xd8, 0x61, 0xec, 0x10, 0x46, 0x0f, 0xb2, 0x53, 0x3b,
	0x14, 0x04, 0xaf, 0x53, 0xd0, 0xf3, 0xfa, 0x01, 0x46, 0xbb, 0xfe, 0x4d, 0x03, 0xe9, 0x12, 0xff,
	0x49, 0x95, 0x7e, 0x0b, 0xc1, 0x2f, 0xe5, 0x71, 0x47, 0x8f, 0xd2, 0x7e, 0x11, 0x59, 0x62, 0x61,
	0x78, 0x9b, 0xb7, 0xf7, 0xfe, 0xbc, 0xf7, 0x7b, 0x81, 0xd0, 0x99, 0x40, 0xb3, 0x49, 0x2b, 0x50,
	0x4a, 0xaf, 0xdf, 0x34, 0xaa, 0x32, 0x35, 0xa8, 0x9d, 0xb6, 0xa9, 0x29, 0x12, 0xaf, 0xd8, 0x59,
	0x38, 0x4c, 0xef, 0x5c, 0x25, 0xb1, 0x5c, 0x9b, 0x1c, 0x5d, 0x9b, 0x0a, 0xe9, 0xaa, 0xa6, 0x48,
	0x36, 0xba, 0x4e, 0x85, 0x16, 0x3a, 0x94, 0x8a, 0xe6, 0xd9, 0x3b, 0x6f, 0xbc, 0x0a, 0x88, 0x38,
	0xa6, 0x97, 0x4f, 0xfb, 0x85, 0x15, 0xbc, 0x34, 0x60, 0x1d, 0x63, 0xf4, 0x74, 0x9b, 0xd7, 0x30,
	0x21, 0x33, 0x32, 0xbf, 0x58, 0x79, 0x1d, 0x5f, 0x51, 0xfa, 0x9b, 0x31, 0xaa, 0x65, 0x13, 0x3a,
	0xaa, 0xc1, 0xda, 0x5c, 0x0c, 0xa1, 0xc1, 0x5e, 0x7f, 0x9c, 0xd0, 0xd1, 0x23, 0x02, 0x38, 0x40,
	0x76, 0x4b, 0xcf, 0xb3, 0xbc, 0xf5, 0x35, 0x36, 0x0e, 0x5b, 0x36, 0x39, 0x5c, 0x9a, 0xb2, 0x3f,
	0x57, 0xa3, 0xda, 0x38, 0x62, 0x0f, 0x74, 0x3c, 0xf4, 0xee, 0x95, 0x84, 0xad, 0xcb, 0x1c, 0x42,
	0x5e, 0x1f, 0xc3, 0x98, 0x93, 0x43, 0x4a, 0x06, 0xf8, 0x0a, 0x78, 0x3c, 0x65, 0xb1, 0xa7, 0xb0,
	0x81, 0xb2, 0x94, 0xa5, 0xfc, 0xcf, 0x4b, 0x16, 0x64, 0x39, 0xf9, 0xec, 0x38, 0xd9, 0x75, 0x9c,
	0x7c, 0x77, 0x9c, 0xbc, 0xf7, 0x3c, 0xda, 0xf5, 0x3c, 0xfa, 0xea, 0x79, 0x54, 0x84, 0xef, 0xbb,
	0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0x08, 0xbb, 0x6b, 0x57, 0xe9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloClientStreamClient, error)
	SayHelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloServerStreamClient, error)
	SayHelloBidiStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloBidiStreamClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/protos.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[0], "/protos.Greeter/SayHelloClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloClientStreamClient{stream}
	return x, nil
}

type Greeter_SayHelloClientStreamClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloClientStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloClientStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayHelloClientStreamClient) CloseAndRecv() (*HelloReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayHelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[1], "/protos.Greeter/SayHelloServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_SayHelloServerStreamClient interface {
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloServerStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloServerStreamClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayHelloBidiStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloBidiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[2], "/protos.Greeter/SayHelloBidiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloBidiStreamClient{stream}
	return x, nil
}

type Greeter_SayHelloBidiStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloBidiStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloBidiStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayHelloBidiStreamClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHelloClientStream(Greeter_SayHelloClientStreamServer) error
	SayHelloServerStream(*HelloRequest, Greeter_SayHelloServerStreamServer) error
	SayHelloBidiStream(Greeter_SayHelloBidiStreamServer) error
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedGreeterServer) SayHelloClientStream(srv Greeter_SayHelloClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloClientStream not implemented")
}
func (*UnimplementedGreeterServer) SayHelloServerStream(req *HelloRequest, srv Greeter_SayHelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStream not implemented")
}
func (*UnimplementedGreeterServer) SayHelloBidiStream(srv Greeter_SayHelloBidiStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloBidiStream not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayHelloClientStream(&greeterSayHelloClientStreamServer{stream})
}

type Greeter_SayHelloClientStreamServer interface {
	SendAndClose(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayHelloClientStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloClientStreamServer) SendAndClose(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayHelloClientStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_SayHelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).SayHelloServerStream(m, &greeterSayHelloServerStreamServer{stream})
}

type Greeter_SayHelloServerStreamServer interface {
	Send(*HelloReply) error
	grpc.ServerStream
}

type greeterSayHelloServerStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloServerStreamServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_SayHelloBidiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayHelloBidiStream(&greeterSayHelloBidiStreamServer{stream})
}

type Greeter_SayHelloBidiStreamServer interface {
	Send(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayHelloBidiStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloBidiStreamServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayHelloBidiStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloClientStream",
			Handler:       _Greeter_SayHelloClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloServerStream",
			Handler:       _Greeter_SayHelloServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloBidiStream",
			Handler:       _Greeter_SayHelloBidiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/hello_world/protos/pb.proto",
}

func (m *HelloRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HelloRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPb(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HelloReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HelloReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintPb(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPb(dAtA []byte, offset int, v uint64) int {
	offset -= sovPb(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HelloRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	return n
}

func (m *HelloReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	return n
}

func sovPb(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPb(x uint64) (n int) {
	return sovPb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HelloRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HelloRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HelloReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HelloReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPb
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPb
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPb
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPb
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPb        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPb          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPb = fmt.Errorf("proto: unexpected end of group")
)
